## 分布式系统
### 为什么需要分布式系统？
1. 增大系统容量
2. 加强系统可用

### 分布式系统弊端
- 架构设计变得复杂（尤其是其中的分布式事务）。
- 部署单个服务会比较快，但是如果一次部署需要多个服务，流程会变得复杂。
- 系统的吞吐量会变大，但是响应时间会变长。
- 运维复杂度会因为服务变多而变得很复杂。
- 架构复杂导致学习曲线变大。
- 测试和查错的复杂度增大。
- 技术多元化，这会带来维护和运维的复杂度。
- 管理分布式系统中的服务和调度变得困难和复杂。

### 分布式系统难点
- 异构系统的不标准问题 
- 服务依赖链中，出现“木桶短板效应”——整个 SLA 由最差的那个服务所决定。
- 故障发生的概率更大。一方面，信息太多等于没有信息，另一方面，SLA 要求我们定义出“Key Metrics”，也就是所谓的关键指标。

### 提高架构的性能
- 缓存系统。加入缓存系统，可以有效地提高系统的访问能力。从前端的浏览器，到网络，再到后端的服务。
- 负载均衡系统。负载均衡系统是水平扩展的关键技术，它可以使用多台机器来共同分担一部分流量请求。
- 异步调用。异步系统主要通过消息队列来对请求做排队处理，这样可以把前端的请求的峰值给“削平”了，而后端通过自己能够处理的速度来处理请求。这样可以增加系统的吞吐量，但是实时性就差很多了。同时，还会引入消息丢失的问题，所以要对消息做持久化，这会造成“有状态”的结点，从而增加了服务调度的难度。
- 数据分区和数据镜像。数据分区是把数据按一定的方式分成多个区（比如通过地理位置），不同的数据区来分担不同区的流量。这需要一个数据路由的中间件，会导致跨库的 Join 和跨库的事务非常复杂。而数据镜像是把一个数据库镜像成多份一样的数据，这样就不需要数据路由的中间件了。你可以在任意结点上进行读写，内部会自行同步数据。然而，数据镜像中最大的问题就是数据的一致性问题。

### 提高架构的稳定性
- 服务拆分。服务拆分主要有两个目的：一是为了隔离故障，二是为了重用服务模块。但服务拆分完之后，会引入服务调用间的依赖问题。
- 服务冗余。服务冗余是为了去除单点故障，并可以支持服务的弹性伸缩，以及故障迁移。然而，对于一些有状态的服务来说，冗余这些有状态的服务带来了更高的复杂性。其中一个是弹性伸缩时，需要考虑数据的复制或是重新分片，迁移的时候还要迁移数据到其它机器上。
- 限流降级。当系统实在扛不住压力时，只能通过限流或者功能降级的方式来停掉一部分服务，或是拒绝一部分用户，以确保整个架构不会挂掉。这些技术属于保护措施。
- 高可用架构。通常来说高可用架构是从冗余架构的角度来保障可用性。比如，多租户隔离，灾备多活，或是数据可以在其中复制保持一致性的集群。总之，就是为了不出单点故障。
- 高可用运维。高可用运维指的是 DevOps 中的 CI/CD（持续集成 / 持续部署）。一个良好的运维应该是一条很流畅的软件发布管线，其中做了足够的自动化测试，还可以做相应的灰度发布，以及对线上系统的自动化控制。这样，可以做到“计划内”或是“非计划内”的宕机事件的时长最短。

### 分布式系统的关键技术
- 服务治理。服务拆分、服务调用、服务发现、服务依赖、服务的关键度定义……服务治理的最大意义是需要把服务间的依赖关系、服务调用链，以及关键的服务给梳理出来，并对这些服务进行性能和可用性方面的管理。
- 架构软件管理。服务之间有依赖，而且有兼容性问题，所以，整体服务所形成的架构需要有架构版本管理、整体架构的生命周期管理，以及对服务的编排、聚合、事务处理等服务调度功能。
- DevOps。分布式系统可以更为快速地更新服务，但是对于服务的测试和部署都会是挑战。所以，还需要 DevOps 的全流程，其中包括环境构建、持续集成、持续部署等。自动化运维。
- 有了 DevOps 后，我们就可以对服务进行自动伸缩、故障迁移、配置管理、状态管理等一系列的自动化运维技术了。
- 资源调度管理。应用层的自动化运维需要基础层的调度支持，也就是云计算 IaaS 层的计算、存储、网络等资源调度、隔离和管理。
- 整体架构监控。如果没有一个好的监控系统，那么自动化运维和资源调度管理只可能成为一个泡影，因为监控系统是你的眼睛。没有眼睛，没有数据，就无法进行高效运维。所以说，监控是非常重要的部分。这里的监控需要对三层系统（应用层、中间件层、基础层）进行监控。
- 流量控制。最后是我们的流量控制，负载均衡、服务路由、熔断、降级、限流等和流量相关的调度都会在这里，包括灰度发布之类的功能也在这里。

### 分布式系统的“纲”
全栈系统监控；服务 / 资源调度；流量调度；状态 / 数据调度；开发和运维的自动化。

## 监控
### 多层体系的监控
所谓全栈监控，其实就是三层监控。
1. 基础层：监控主机和底层资源。比如：CPU、内存、网络吞吐、硬盘 I/O、硬盘使用等。
2. 中间层：就是中间件层的监控。比如：Nginx、Redis、ActiveMQ、Kafka、MySQL、Tomcat 等。
3. 应用层：监控应用层的使用。比如：HTTP 访问的吞吐量、响应时间、返回码、调用链路分析、性能瓶颈，还包括用户端的监控。

### 什么才是好的监控系统
1. 关注于整体应用的 SLA。
2. 关联指标聚合。
3. 快速故障定位。

#### “体检” 
- 容量管理。  
提供一个全局的系统运行时数据的展示，可以让工程师团队知道是否需要增加机器或者其它资源。
- 性能管理。   
可以通过查看大盘，找到系统瓶颈，并有针对性地优化系统和相应代码。

#### “急诊”
- 定位问题。   
可以快速地暴露并找到问题的发生点，帮助技术人员诊断问题。
- 性能分析。   
当出现非预期的流量提升时，可以快速地找到系统的瓶颈，并帮助开发人员深入代码。

只有做到了上述的这些关键点才能是一个好的监控系统。

一个分布式系统，或是一个自动化运维系统，或是一个 Cloud Native 的云化系统，最重要的事就是把监控系统做好。在把数据收集好的同时，更重要的是把数据关联好。这样，我们才可能很快地定位故障，进而才能进行自动化调度。

## 服务调度
资源 / 服务调度

### 服务状态的维持和拟合
类似K8S

### 服务的弹性伸缩和故障迁移
而对于故障迁移，也就是服务的某个实例出现问题时，我们需要自动地恢复它。对于服务来说，有两种模式，一种是宠物模式，一种是奶牛模式。


### 服务编排
API Gateway 或一个简单的消息队列来做相应的编排工作。在 Spring Cloud 中，所有的请求都统一通过 API Gateway（Zuul）来访问内部的服务。这个和 Kubernetes 中的 Ingress 相似。