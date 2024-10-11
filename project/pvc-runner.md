## Gitlab Runner

Gitlab Runner是gitlab CI中一个非常核心的组件，每个gitlab CI job都需要指定一个runner来执行。Runner的类型有很多，比如常见的shell, docker, kubernetes等。shell和docker的启动速度很快，适合轻量化的Job, kubernetes类型具有弹性，并发量大的优点，应用范围更加广泛。

## 什么是PVC Runner

PVC runner是我们提出的一种通过使用持久化存储PVC来复用代码库和bazel cache，同时能够兼容kubernetes executor所具备的高可用，高并发，弹性等特性的runner的自定义runner。

之所以设计这个runner，是我们在使用kubernetes runner时，在针对复杂的CI job时，job耗时大部分都是源自于git，git lfs操作和bazel编译过程，如何加速这两部分成为优化CI平台编译性能的关键。

对于第一个问题，gitlab CI job默认会浅clone代码仓库然后获取指定个数的commits，这个过程会花费一定的时间。因此我们的想法是持久化存储项目代码库，这需要用到一些提前申请的云盘。而对于第二个问题，每个job都会在编译时产生bazel cache，而这些中间产物可能会被后续的action所复用，因此如果能够持久化存储bazel cache的目录，那么就可以对编译进行大幅加速。请注意，我们讨论这个问题的前提是我们在生产环境下所使用的集群节点是弹性伸缩的，而不是固定的，如果是固定的，可以考虑持久化到节点本地而非PVC中。

## 初探

最初的尝试是使用docker runner的docker in docker方法，通过挂载单独的云盘实现对代码库（包括LFS文件）和bazel cache的缓存，以及我们把每个docker runner并发度设为1，此举是为了保证每个Job的代码不会互相影响。

尽管这个方案是有效的，但是存在很大的问题，那就是每个docker runner实例的并发数只有1，为了能够多支持一些job并发执行，我们需要手动部署多个实例，实际上我们部署了8、9个实例，但是有时业务繁忙，后续job不得不排队等待前面的Job完成任务。同时docker runner并不弹性，docker runner作为执行的pod需要分配到高配的节点上长时间占用，也会带来较高的成本。

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/6c2430d8987f4cce9c0fe64d4de0e7a9~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239323&x-orig-sign=vdmUdPiqI2BEbB1r9K6terEyKBY%3D)

PVC runner的方案如图所示，考虑到高可用，我们会部署3副本，副本名称由前缀决定，这个前缀同时也作为PVC和并发值的前缀来使用。

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/1d54ac13ff4e4e85a5c40c252c22a235~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239323&x-orig-sign=kNqaeFISdg3Xqongnm%2BiWH5YqTM%3D)

### 存在的问题

#### pod长时间起不起来导致超时

*   高负载\
    排查发现由于负载高导致kubelet的操作有较大的时延，这会导致后续挂载冲突。
    什么是高负载？\
    根据核心数，系统负载值超过CPU核心个数的2倍可以算作高负载了，需要重点关注和处理。2倍说明有一半的线程在排队状态。

如何降低节点负载？

1.  降低节点负载可以考虑降低每个节点上的pod个数，这可以通过调大pod资源申请值来规划。
2.  也可以单独申请一个节点池，将pvc runner的job分配到新的资源池中，避免受到其他高负载pod的影响
3.  排查哪些job属于高负载任务，找到后把他们排除出去，比如使用单独的云盘以加强IO读写性能，这也是一种思路

*   fsGroup权限配置导致长耗时
    在解决了高负载导致kubelet操作延时问题后，问题得到一定的解决，于是我们再次进行job迁移，当job数量上来后，再次出现pod长时间初始化失败的问题。通过查看grafana系统负载监控，可以看到系统负载并不高。那么问题出在哪里呢？

通过分析/var/log/messages文件，我们注意到applyFSGroup存在超时异常，最终导致云盘被强制unmount

于是我们删除了fsGroup的配置项，避免每次pod启动时都做一遍深层遍历权限设置，实验表明这次修改非常关键，pod的启动速度变得飞快。

#### runner重启后出现挂载问题

问题原因是存在长耗时的job一直占用着某些云盘，修改参数重新部署runner后依然会用到原来的名字和PVC名字，后续再起pod可能导致分配已经被占用的PVC，从而导致挂载冲突。

#### LFS文件每次都要拉取，无法利用缓存，导致长耗时

在发版任务中，需要用到lfs文件，但是我们发现pvc runner中git lfs pull过程很慢，因为无法利用缓存对象。经过排查后发现，git checkout过程会发生指针转换，将LFS二进制大文件转换为了指针，从而后续执行git lfs pull相当于重新拉一遍。

解决办法：通过修改源码，禁用掉强制SKIP\_SMUDGE以及对应的git lfs pull操作，我们只允许通过环境变量LFS\_SKIP\_SMUDGE来控制LFS指针变换的行为，从而确保速度和LFS文件的准确性。

经过实验验证，我们的PVC方案能够大幅度提高作业的执行速度，提速可达**40%-85%！**

## 继续优化

### 当前现状

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/7b4c158d614c41eca3536225fecbc80a~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239323&x-orig-sign=4oS2CvqLHp3EWc7cozBpYV388%2BA%3D)

### 我们的目标

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/e51fdaab494e48a787ac7b0de1689723~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239323&x-orig-sign=x8f2uOnYG0xu0VD4DJ1%2B1Qc3MfI%3D)

目标1： Pod和PVC动态生成，且一对一绑定\
目标2： PV的复用\
目标3： PVC前缀固定的模式

针对目标1，我们的解决办法是通过ephemeral pvc实现Pod和pvc动态绑定，临时PVC可以随着pod的生命周期自动创建和释放，非常灵活，符合我们这里的需求。

针对目标2，我们先提前创建storage class资源池，然后在某后台服务中监听k8s pv的变化，一旦pv从bound变为released状态时，这个服务会自动patch，将pv的状态改为available，从而实现复用。

针对目标3，删除helm chart中pvc前缀相关的配置即可

通过这一系列方法，PVC runner真正具备了多副本，弹性，高并发的能力，相当于最初的架构方案更加优雅。

## 潜在优化点

*   如果针对不同分支部署不同的runner实例，pv却是可以共用的，此时就无法保证云盘只给特定分支复用，怎么解决？
    *   可以通过新建Storage class的方式，区分不同的分支
*   在给PV修改状态时，想办法对PV的disk usage, inode等情况进行检查和清理
    *   可以放到后天服务中进行
