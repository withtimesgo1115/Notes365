## RAFT算法
分布式系统需要一个中间件来存储共享的信息，来保证多个系统之间的数据一致性。

### 分布式一致性
- 数据不能存在单个节点上，否则可能出现单点故障
- 多个节点需要保证具有相同的数据

### 一致性算法
- Paxos算法 (起源，图灵奖)
- Raft算法 (etcd)
- ZAB算法（zookeeper）

### 复制状态机模型
![alt text](image/image.png)
整个工作流程可以归纳为如下几步：
- 用户输入设置指令，比如将设置y为1，然后将y更改为9。  
- 集群收到用户指令之后，会将该指令同步到集群中的多台服务器上，这里你可以认为所有的变更操作都会写入到每个服务器的Log文件中。
- 根据Log中的指令序列，集群可以计算出每个变量对应的最新状态，比如y的值为9。
- 用户可以通过算法提供的API来获取到最新的变量状态。  

算法会保证变量的状态在整个集群内部是统一的，并且当集群中的部分服务器宕机后，仍然能稳定的对外提供服务。

分布式一致性问题分解为了Leader选举、日志同步和安全性保证三大子问题。

### Raft算法基础
Raft 中使用日志来记录所有操作，所有结点都有自己的日志列表来记录所有请求。算法将机器分成三种角色：Leader、Follower 和 Candidate。正常情况下只存在一个 Leader，其他均为 Follower，所有客户端都与 Leader 进行交互。

![alt text](https://imgs.lfeng.tech/images/2023/04/tcCXcV.gif)

Leader 在收到来自客户端的请求后并不会执行，只是将其写入自己的日志列表中，然后将该操作发送给所有的 Follower。Follower 在收到请求后也只是写入自己的日志列表中然后回复 Leader，当有超过半数的结点写入后 Leader 才会提交该操作并返回给客户端，同时通知所有其他结点提交该操作。

通过这一流程保证了只要提交过后的操作一定在多数结点上留有记录（在日志列表中），从而保证了该数据不会丢失。

Raft 是一个非拜占庭的一致性算法，即所有通信是正确的而非伪造的。N个结点的情况下（N为奇数）可以最多容忍(N-1)/2个结点故障。如果更多的节点故障，后续的Leader选举和日志同步将无法进行。

### 1. Leader选举
#### 首次选举
![alt text](https://imgs.lfeng.tech/images/2023/04/Kb4kPA.gif)

如果定时器超时，说明一段时间内没有收到 Leader 的消息，那么就可以认为 Leader 已死或者不存在，那么该结点就会转变成 Candidate，意思为准备竞争成为 Leader。

成为 Candidate 后结点会向所有其他结点发送请求投票的请求（RequestVote），其他结点在收到请求后会判断是否可以投给他并返回结果。Candidate 如果收到了半数以上的投票就可以成为 Leader，成为之后会立即并在任期内定期发送一个心跳信息通知其他所有结点新的 Leader 信息，并用来重置定时器，避免其他结点再次成为 Candidate。

如果 Candidate 在一定时间内没有获得足够的投票，那么就会进行一轮新的选举，直到其成为 Leader,或者其他结点成为了新的 Leader，自己变成 Follower。

#### 再次选举
1. Leader下线，此时所有其他节点的计时器不会被重置，直到一个节点成为了 Candidate，和上述一样开始一轮新的选举选出一个新的 Leader。  
![alt text](https://imgs.lfeng.tech/images/2023/04/w95KQJ.gif)

2. 某一 Follower 结点与 Leader 间通信发生问题，导致发生了分区，这时没有 Leader 的那个分区就会进行一次选举。这种情况下，因为要求获得多数的投票才可以成为 Leader，因此只有拥有多数结点的分区可以正常工作。而对于少数结点的分区，即使仍存在 Leader，但由于写入日志的结点数量不可能超过半数因此不可能提交操作。这也是为何一开始我提到Raft算法必须要半数以上节点正常才能工作。
![alt text](https://imgs.lfeng.tech/images/2023/04/03bK8U.gif)
#### 小总结
![alt text](image/image-2.png)
Leader：处理与客户端的交互和与 follower 的日志复制等，一般只有一个 Leader；
Follower：被动学习 Leader 的日志同步，同时也会在 leader 超时后转变为 Candidate 参与竞选；
Candidate：在竞选期间参与竞选；

#### 任期Term
Raft算法将时间分为一个个的任期（term），每一个term的开始都是Leader选举。
每一个任期以一次选举作为起点，所以当一个结点成为 Candidate 并向其他结点请求投票时，会将自己的 Term 加 1，表明新一轮的开始以及旧 Leader 的任期结束。所有结点在收到比自己更新的 Term 之后就会更新自己的 Term 并转成 Follower，而收到过时的消息则拒绝该请求。
在成功选举Leader之后，Leader会在整个term内管理整个集群。如果Leader选举失败，该term就会因为没有Leader而结束。
![alt text](image/image-3.png)

#### 投票
在投票时候，所有服务器采用先来先得的原则，在一个任期内只可以投票给一个结点，得到超过半数的投票才可成为 Leader，从而保证了一个任期内只会有一个 Leader 产生（Election Safety）。

在 Raft 中日志只有从 Leader 到 Follower 这一流向，所以需要保证 Leader 的日志必须正确，即必须拥有所有已在多数节点上存在的日志，这一步骤由投票来限制。

日志格式如下：
![alt text](image/image-4.png)
如上图所示，日志由有序编号（log index）的日志条目组成。每个日志条目包含它被创建时的任期号（term），和用于状态机执行的命令。如果一个日志条目被复制到大多数服务器上，就被认为可以提交（commit）了。

由于只有日志在被多数结点复制之后才会被提交并返回，所以如果一个 Candidate 并不拥有最新的已被复制的日志，那么他不可能获得多数票，从而保证了 Leader 一定具有所有已被多数拥有的日志（Leader Completeness），在后续同步时会将其同步给所有结点。

### 2. 日志同步
#### 工作流程
Leader选出后，就开始接收客户端的请求。Leader把请求作为日志条目（Log entries）加入到它的日志中，然后并行的向其他服务器发起 AppendEntries RPC复制日志条目。当这条日志被复制到大多数服务器上，Leader将这条日志应用到它的状态机并向客户端返回执行结果。

![alt text](image/image-7.png)
某些Followers可能没有成功的复制日志，Leader会无限的重试 AppendEntries RPC直到所有的Followers最终存储了所有的日志条目。

日志由有序编号（log index）的日志条目组成。每个日志条目包含它被创建时的任期号（term），和用于状态机执行的命令。如果一个日志条目被复制到大多数服务器上，就被认为可以提交（commit）了。

#### 实际处理逻辑
Leader 会给每个 Follower 发送该 RPC 以追加日志，请求中除了当前任期 term、Leader 的 id 和已提交的日志 index，还有将要追加的日志列表（空则成为心跳包），前一个日志的 index 和 term。

![alt text](image/image-8.png)

在接到该请求后，会进行如下判断：


1. 检查term，如果请求的term比自己小，说明已经过期，直接拒绝请求。


2. 如果步骤1通过，则对比先前日志的index和term，如果一致，则就可以从此处更新日志，把所有的日志写入自己的日志列表中，否则返回false。    

这里对步骤2进行展开说明，每个Leader在开始工作时，会维护 nextIndex[] 和 matchIndex[] 两个数组，分别记录了每个 Follower 下一个将要发送的日志 index 和已经匹配上的日志 index。每次成为 Leader 都会初始化这两个数组，前者初始化为 Leader 最后一条日志的 index 加 1，后者初始化为 0，每次发送 RPC 时会发送 nextIndex[i] 及之后的日志。

在步骤2中，当Leader收到返回成功时，则更新两个数组，否则说明follower上相同位置的数据和Leader不一致，这时候Leader会减小nextIndex[i]的值重试，一直找到follower上两者一致的位置，然后从这个位置开始复制Leader的数据给follower，同时follower后续已有的数据会被清空。

在复制的过程中，Raft会保证如下几点：

- Leader 绝不会覆盖或删除自己的日志，只会追加 （Leader Append-Only），成为 Leader 的结点里的日志一定拥有所有已被多数节点拥有的日志条目，所以先前的日志条目很可能已经被提交，因此不可以删除之前的日志。

- 如果两个日志的 index 和 term 相同，那么这两个日志相同 （Log Matching），第二点主要是因为一个任期内只可能出现一个 Leader，而 Leader 只会为一个 index 创建一个日志条目，而且一旦写入就不会修改，因此保证了日志的唯一性。

- 如果两个日志相同，那么他们之前的日志均相同，因为在写入日志时会检查前一个日志是否一致，从而递归的保证了前面的所有日志都一致。从而也保证了当一个日志被提交之后，所有结点在该 index 上提交的内容是一样的（State Machine Safety）。



### 3：安全性保障（核心）
- 已经commit的消息，一定会存在于后续的Leader节点上，并且绝对不会在后续操作中被删除。
- 对于并未commit的消息，可能会丢失。

#### 多数投票规则
一个candidate必须获得集群中的多数投票，才能被选为Leader；而对于每条commit过的消息，它必须是被复制到了集群中的多数节点，也就是说**成为Leader的节点，至少有1个包含了commit消息的节点给它投了票。**

而在投票的过程中每个节点都会与candidate比较日志的最后index以及相应的term，如果要成为Leader，必须有更大的index或者更新的term，所以**Leader上肯定有commit过的消息。**

#### 提交规则
Raft额外限制了 **Leader只对自己任期内的日志条目适用该规则，先前任期的条目只能由当前任期的提交而间接被提交。** 也就是说，当前任期的Leader，不会去负责之前term的日志提交，之前term的日志提交，只会随着当前term的日志提交而间接提交。
![alt text](image/image-5.png)

- 初始状态如 (a) 所示，之后 S1 下线；
- (b) 中 S5 从 S3 和 S4 处获得了投票成为了 Leader 并收到了一条来自客户端的消息，之后 S5 下线。
- (c) 中 S1 恢复并成为了 Leader，并且将日志复制给了多数结点，之后进行了一个致命操作，将 index 为 2 的日志提交了，然后 S1 下线。
- (d) 中 S5 恢复，并从 S2、S3、S4 处获得了足够投票，然后将已提交的 index 为 2 的日志覆盖了。

这个例子中，在c状态，由于Leader直接根据日志在多数节点存在的这个规则，将之前term的日志提交了，当该Term下线后，后续的Leader S5上线，就将之前已经commit的日志清空了，导致commit过的日志丢失了。

为了避免这种已提交的日志丢失，Raft只允许提交自己任期内的日志，也就是不会允许c中这种操作。（c）中可能出现的情况有如下两类：

- （c）中S1有新的客户端消息4，然后S1作为Leader将4同步到S1、S2、S3节点，并成功提交后下线。此时在新一轮的Leader选举中，S5不可能成为新的Leader，保证了commit的消息2和4不会被覆盖。


- （c）中S1有新的消息，但是在S1将数据同步到其他节点并且commit之前下线，也就是说2和4都没commit成功，这种情况下如果S5成为了新Leader，则会出现（d）中的这种情况，2和4会被覆盖，这也是符合Raft规则的，因为2和4并未提交。

### 拓展
#### 日志压缩
![alt text](image/image-6.png)
在实际的系统中，不能让日志无限增长，否则系统重启时需要花很长的时间进行回放，从而影响可用性。Raft采用对整个系统进行snapshot来解决，snapshot之前的日志都可以丢弃。

每个副本独立的对自己的系统状态进行snapshot，并且只能对已经提交的日志记录进行snapshot。

Snapshot中包含以下内容：

- 日志元数据。最后一条已提交的 log entry的 log index和term。这两个值在snapshot之后的第一条log entry的AppendEntries RPC的完整性检查的时候会被用上。
- 系统当前状态。上面的例子中，x为0，y为9.


当Leader要发给某个日志落后太多的Follower的log entry被丢弃，Leader会将snapshot发给Follower。或者当新加进一台机器时，也会发送snapshot给它。

做snapshot既不要做的太频繁，否则消耗磁盘带宽， 也不要做的太不频繁，否则一旦节点重启需要回放大量日志，影响可用性。推荐当日志达到某个固定的大小做一次snapshot。