## 构建平台亮点
### 背景
- 平台用户是公司的研发、测试、版本经理等
- 公司代码以C++为主，对编译有较高的要求
- 平台的意义在于提供稳定、易用、快速的持续集成和持续部署能力

### 高可用性
高可用性是这个平台最关键的能力，为此，我做了：   
1. 排查和解决Missing Digest, Create Exec Dir Error等线上问题
2. bazel remote cache从单节点迁移到云上的分布式编译服务, 避免单点故障
3. 容错机制的设计，对于部分失败job进行自动重试，超过3次则报警通知。对于API触发的job进行限流处理，避免资源都被API调用所占用
4. 增加隔离性，不同环境使用不同的cache目录，灰度更新
5. 容量规划，定期进行性能测试和容量规划，评估系统负载能力，确保能应对未来的业务增长
6. 完善监控告警，监控operation队列排队情况来了解平台当前压力；出现报错则触发告警

### 高性能
#### 开发分布式编译服务
旧：在机房部署单个节点从而提供一个单机的remote cache地址，只使用了remote cache的能力同时不支持高可用和弹性
新：同时使用remote cache和remote executor, 分布式服务server端通过实现remote execution api来做远程执行。

分布式服务支持多个server节点和多个worker节点且可以单独部署，增加这二者的节点数量，可以大大提高编译job action的并发量，从而显著提升编译性能。

#### Runner优化
1. 量产项目发包对于性能的要求很高，排查发现很大一部分耗时在拉取代码库和lfs文件上
2. bazel在本地使用时可以通过内存缓存构建图，因此增量编译的性能非常炸裂。但是如果放在pod容器内使用的话，由于pod存在生命周期，那么这些构建图就本地cache就无法被复用了。

PVC runner是基于gitlab runner k8s executor进行的一个优化，使用ephemeral pvc的方式，去挂载固定资源池里的pv，临时pvc和pod一对一绑定，pod销毁，临时卷也就被销毁了。但是pvc所绑定的pv并不会被销毁，而是retain住，从而保证repo和bazel cache可以被持久化以便后续使用这个PV的pod进行复用, pv资源需要进行一个监听，每次pv状态变化的时候对released状态的pv改为available。每次使用前需要判断pv是否存储空间合法，如果存储超过阈值，需要先进行清理。

这个优化已经被我提交到gitlab runner社区了。

### 可视化编译
编译是黑盒子，并不确定编译targets时每个action的执行情况，通过bazel提供的build event protocol可以拿到编译每个action时的详细信息，从而进行一个可视化。方便排查问题和了解编译每个targets时的具体细节。

### 项目技术架构
Bazel + Spring Boot + Redis

当前技术选型可以满足项目场景的需求，引入过多的模块会增加项目复杂度

Bazel负责构建图的创建，优化等
Spring Boot负责开发Server, Worker服务等
Redis负责记录Cache的位置, 当前活跃workers等

### 具体实现
实现一个server服务，一个worker服务。server中维护一个operation队列，和redis通信记录CAS的数据，同时和bazel client通信接收action，并把action派发给worker。worker则是维护了一个workers队列, 做具体的执行。

关键就在于实现了Remote Execution API。  
客户端需要提供：
- Command
- Action
- ContentAddressableStorage   

这些也被叫做action_digest，发给server，server安排对action_digest的执行。

- ContentAddressableStorage
- cache ActionResults by a key ActionCache
- execute actions asynchronously with Execution

#### workers 
worker主要承担两个角色：执行构建和cas分片。  
执行通过一个pipeline来进行，同时需要保证同一个执行不会被其他管道执行   
cas分片，每个从server端的blob写入操作是随机选择worker的，作为inputs的blob会被复制到其他worker上，但是作为outputs的不会被复制。

考虑到背压，每个阶段都需要下一阶段允许接收才会释放。
1. 匹配
2. 输入获取
    从CAS下载，会有IO写入带来的性能问题
3. 执行
    编译action，超时停止
4. 报告结果
    将输出注入CAS，完成后销毁执行目录

#### operaion队列
分为CPU队列和GPU队列，可以分别用于监控当前编译的压力

#### ActionCache
ActionCache是一个服务，用来查询operation是否已经执行，如果已经执行，那么就下载结果，ActionCache需要contentAddressableStorage服务来存储文件数据。

Action封装了执行操作所需的所有信息。这些信息包括命令、包含子目录/文件树的输入树、环境变量、平台信息。所有这些信息都将有助于进行digest计算，以便多次执行Action产生相同的输出。这样， action的hash值可以作为ActionResult:XXXX键。ActionResult可以被上传到ActionCache。

#### CAS
- 读取内容    
    通过BatchReadBlobs 还是 ByteStream Read 方法访问。沿着内容的长度推进偏移量，直到完成连续的请求。
- 写入内容    
    将内容写入 CAS 需要事先使用所选的digest方法(SHA256)计算所有内容的地址。可以使用 BatchUpdateBlobs 或 ByteStream Write 方法发起写入。   
    在收到内容之前失败的写入应该是渐进的，通过 ByteStream QueryWriteStatus 检查偏移量以恢复上传
- 分片
    分片在 redis 上注册，redis中维护了enrty和worker映射。FindMissingBlobs 请求在分片中循环，减少缺失条目列表。
    写入随机选择目标节点并且目前不会复制，读取时尝试在每个公布的分片上请求条目，并在出现 NOT_FOUND 或暂时性 grpc 错误时进行故障转移。读取是乐观的，因为不会请求预期找不到的 blob，分片 CAS 将在 blob 完全缺失时故障转移到整个集群搜索条目。如果还是找不到则反馈missing digest。

#### Redis
为了平衡不同redis节点CPU使用，维护一个平衡的operaion队列，底层使用redis List，通过redis hashtags来做负载均衡。

### 改进点
#### cache目前存在节点本地
从成本方面考虑，我们的节点由固定节点和弹性节点组成。当前cache是保存在固定节点的/tmp/worker目录下的，这种方式实现起来相对简单，但是不如将cache存储在S3存储上好，这样的话cache不会因为节点存储达到阈值被剔除，影响编译的可用性，同时S3存储也更可靠。不过这种方式也存在问题，如果cache一直都不被清除，对象存储cache目录会越来越大，保存了很多失效的编译产物，需要有对应的清理机制。

#### worker的产物复制
当前一个worker的outputs不会被复制到其他worker上，如果这个节点或者outputs被剔除的话，那么在bazel bwob时就会出错。

#### 如何限流
对不同的编译Job进行优先级划分：  
1. 发包作业 > 编译作业 > 测试作业 > 代码风格检查等
2. 可以考虑在MR上通过用户指定标签来声明优先级。  

对于优先级高的Job，平台优先将资源分配给它，保证高优任务能够快速执行。在业务压力大的时候，低优先度的job会被放到阻塞队列中等待执行，从而使得平台的编译流量保持平稳。快速完成当前任务后可以从阻塞队列中获取任务来执行。

#### 定期reindex（已完成）
由于我们使用了一部分弹性节点，所以reindex非常重要，不然redis中的记录可能不准确了。获取活跃workers，然后对于redis中的CAS记录取交集，进行一个更新，保证redis中的记录准确。

## BevSlots泊车自动标注
### 背景
为泊车模型训练提供低成本、高质量的自动标注数据。目前方案存在的问题是基于模型的推理结果优化后反用于训练模型，导致模型接收的数据质量并不高，结果不是特别理想。后面会训练出的大模型，会用于泊车标注。

### 三个步骤
1. 局部建图
2. 预标
3. 标注生产

局部建图阶段主要是生成点云底图，点云底图既可以用于自动标注也可以用于人工标注。

预标阶段主要是通过BevSlots模型根据底图，帧信息等对进行车位推理。推理后的结果通过后处理算法进行融合，输出更准确的结果。同时这一步还会生成AVM模型的推理结果，后续在第三步中用于过滤筛选。

标注阶段主要是对预测出的结果进行投影，转换为标准的输出格式，统计车位和检测信息，最后质检过滤并记录需要筛掉的帧。

每个步骤都会记录执行的状态，方便debug。

### 点云底图

### 时序融合算法
- 轨迹更新和管理
对于每一个BEVSlots车位i，我们在时序上会有多帧的观测。由于是时间序列，我们通过引入Tracker机制来跟踪这些观测的轨迹，如果匹配上属于同一个轨迹就加入，否则就新建轨迹。如果某条轨迹在连续N帧中未能匹配任何观测，则认为该轨迹消失。

- 平滑轨迹    
加权平均：根据观测的置信度（如检测得分）对轨迹中的信息加权。   
异常点剔除：在融合过程中剔除异常观测点
轨迹补全（TODO）：在观测帧中因漏检导致轨迹中断时，可以通过插值方法对轨迹进行补全。基于运动模型预测：使用卡尔曼滤波或其他预测方法估计丢失帧的车位信息。


### 基于AVM的过滤算法
匈牙利算法匹配，匹配AVM车位点和BEVSlots模型车位点，找到之后判断距离最小的2个点的平均值是否小于给定阈值。如果小于则说明匹配上了。没有匹配上的AVM车位大于一定阈值，则过滤掉这一帧。如果AVM没结果也过滤掉，如果可见范围内没有车位也过滤掉。

### 优化点
- 模型算法优化
- AVM模型优化
- 多传感器数据
- 引入时间加权机制，对轨迹中的近期帧给予更高权重。

### 问题
1. 每个pipeline的耗时
2. 对于30s的分片，泊车pipeline耗时
3. 对于失败的pipeline重刷有版本的概念吗？