## 什么是Bazel

Bazel是Google在2015年开源的一款构建工具，目前在GitHub上已获得超过23k的star，明显领先于类似的工具如gradle、maven和cmake。近年来，国内大型互联网公司也逐渐采用Bazel来构建自己的软件。

Bazel之所以备受青睐，正如其口号所言：“正确与快速，二者兼得”，这不仅仅是宣传口号，用户的实际使用感受也印证了这一点。

1.  加速构建和测试: Bazel只重建必要的东西，借助本地和分布式缓存、优化的依赖关系分析和并行执行，获得快速和增量构建。
2.  一个工具，多种语言: 构建和测试Java、C++、Android、iOS、Go和各种其他语言平台。Bazel可以在Windows、macOS和Linux上运行。
3.  可扩展:Bazel帮助您扩展您的组织、代码库和持续集成解决方案。它处理任意大小的代码库，在多个仓库或一个巨大的monorepo中。
4.  可根据您的需求进行扩展:使用Bazel熟悉的扩展语言，轻松添加对新语言和平台的支持。分享和重用不断发展的Bazel社区编写的语言规则。

本文不会重点介绍Bazel，而重点讨论如何通过bazel的remote executor方式实现大规模的分布式并发编译。

## 榨干bazel的性能：remote cache和分布式编译

### remote cache

Bazel的Action由构建系统本身进行设计，更加安全，因此就不会出现多个action对同一文件的竞争问题。基于这个特性，我们可以充分利用多核CPU的能力，让Action并行执行。通常我们采用CPU逻辑核心数作为Action执行的并发度，如果开启了远端执行，则可以开启更高的并发度。

除了分布式编译，bazel的远端缓存也可以尽可能地保证我们的构建是增量构建，下图展示了remote cache生效的原理。
![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/77a2feedd27f4501a22ee9bd0be67769~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=1666\&h=936\&s=130010\&e=png\&a=1\&b=fff9f9)

每个action的元数据和编译产物artifacts都会被hash为digest，这个digest是基于内容的，也就是说，只有内容发生变化时，才会重新计算digest。一个digest对应一个actionResult，内容寻址的好处是不容易污染存储空间。因此，通过remote cache我们可以持久化存储actionResult，从而加快编译构建。

当然storage不可能是无限容量的，因此一般都有LRU算法来管理CAS中的blob，当写满时，最久没有被访问的blob会被自动淘汰。

### remote execute

Bazel在构建时，可以把Action发送给另一台服务器执行，等执行完毕后，服务器可以向CAS上传ActionResult，然后本地再下载这个结果。

这种做法减少了本地执行Action的开销，使得我们设置更高的构建并发度

远程执行的原理如下：

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/4f6ccac4861e40a5b0a7ac062345e2e4~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=1002\&h=1244\&s=94421\&e=png\&a=1\&b=bfffff)

## Buildfarm

Buildfarm是一款开源的分布式编译服务，基于spring boot实现。Buildfarm可以分为server和worker两部分。

### 源码解析

#### Server端

ExecutionService

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/919a748a5dd645d282fd5907ed0e44a9~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=ZmyK5%2BoccJoNoSQ2%2FJVYeRZhgnE%3D)

ShardInstance

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/34222b26c6584fe48713483ad9eb1dc1~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=fZlUS0auQLktF8jlMvULvH%2BnXDQ%3D)

RedisShardBackplane

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/f7d30bebd1bd451a9952a708849988de~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=P8DtAeu4be1iQAAJW2Pi5NxfWms%3D)

#### worker端

ShardWorkerInstance

```java
pipeline.add(matchStage, 4);
pipeline.add(inputFetchStage, 3);
pipeline.add(executeActionStage, 2);
pipeline.add(reportResultStage, 1);
```

matchStage

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/8605fb4193904654a9aa2c6488908845~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=sEkBlQ9fBXNzUmjSc0rz%2BB2jq0A%3D)

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/e33daa2a1d754de686b83c52dc97816c~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=D20S%2FKBw2hk1pmHeca4gC3vGNFM%3D)

inputFetchStage

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/4288d37b78bc44b4bbc152636f5164f8~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=Dun9AiqDhWP1HgDkNSIe2Lg4dnc%3D)

远程拉取Input

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/f7863f8c10f344f4942185833bf820b4~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=UDGFSFHO2RpShTYCMkU4edkv8mM%3D)

executeActionStage

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/964deb5c506d4df3b27acc2f8f4c629d~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=oIx0%2BhqhNJ5meUdjc%2F0aQBIqm%2BE%3D)

reportResultStage

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/35d3536eb2a5422bbc0c93b0302211f0~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=gYtrskyOoXlzh%2FYRlo4z9Y%2Bn6Kg%3D)

从上面的源码可以看到，worker的执行分为4个步骤，通过pipeline串联起来。分别是match, fetchInput, execute和report。

#### CAS端

ShardCASFileCache
常用的cas都是filesystem的文件cache。\
![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/d2f3157935d541aca368c3dd527c53e4~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=IbuyIo3RX%2B6oOmOjc1Q36G3eq3A%3D)

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/4d6b050f27674d4aa887f6be5c067d49~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=rRmFAwq9%2BzHYKFLBA0C3%2FN%2FbxyQ%3D)

#### 定时reindexCas或者节点发生变化时执行

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/2fbc87a13aa44b24bf41bec00c4d9a92~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=8Mg%2BJkO8i2L8qfAJOrsBgoab4Qw%3D)

removeWorkerIndexesFromCas
遍历redis集群的node, 然后执行reindexNode

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/7b1691a17bd14134879a559e0a742d09~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=qqbEP90%2BkH5BsIIiJAyh2sb4j1s%3D)

reindexNode

具体清理逻辑

主要是比较active workers和cas key map set中记录的存储改digest的worker set进行一个求交集

如果交集为空，则把cas key直接删除

如果交集不为空，则对这个caskey的worker set进行一个替换，用新找的交集替换上去

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/668939846cdd458da6ee02b7ef14bb2d~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=aN%2BxMvs5aj5qnbLQHAWqwdO9RkY%3D)

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/3b1b7342eaa94c84ac00666ccaccb628~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=M46puyoxycIsrOGWPs%2Br2rcgIjg%3D)

### 常见问题

*   Missing Digest问题

![image.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/2a4e76a973fd40b7b2ae5fe9508e0471~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgd2l0aHRpbWVzZ28=:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiNDM4ODkwNjE1MTE5MTUyOCJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1729239349&x-orig-sign=ulZbNhQ80K9amXA1kXLMHF76BJE%3D)

#### 原因

*   首先明确出现问题的CI job均使用了--remote\_download\_minimal参数。

*   BwoB是一种构建模式，启用时，Bazel会推迟远程执行(或远程缓存)操作的输出下载，直到需要它们时才下载(例如，作为本地操作的输入)。这意味着，在调用过程中，Bazel需要跟踪每个输出的最少信息，以备后面检索。然后，Bazel在必要时使用元数据下载远程blobs。

*   目前，Bazel假设它以前从远程服务器获取的元数据总是可以在以后用于检索相应的blob。然而，对于使用分布式编译服务buildfarm来说，这是不正确的，因为blobs可能会由于各种原因（远程存储空间不足，弹性节点伸缩等）而从远程缓存中被逐出。从Bazel的角度来看，blobs可以在两个不同的时间被驱逐：
    *   在调用期间：在当前调用早期远程执行的action的输出被删除，而其成为了另一个正在运行的action的输入。
    *   在调用之间：当Bazel没有构建时，blobs被逐出。

#### 解决办法

Bazel 6.2.1支持--experimental\_remote\_cache\_eviction\_retries参数，指定后在遇到remote blob eviction时会重新启动一个invocation，试验了几次是能解决Missing Digest问题的。
