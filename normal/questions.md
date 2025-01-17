### 最有挑战的事是什么？怎么解决的？
线上问题的排查，影响了平台可用性。
大部分线上问题都是更新导致的，必要时回滚版本，先把影响降低到最低然后再排查问题根因。

missing digest问题排查，细读bazel源码发现加的一个参数开启了bwob，远端cache由于存储打满可能会被剔除，导致了数据不一致。

增加了日志埋点打印报错的Host主机，分布式trace能力

### 怎么跨部门沟通？
OKR不一致的话，其他团队的人确实很难愿意配合。如果事情高优，则上升到团队leader，让leader间明确人力投入，这样的话大家目标一致，更能把事情做好。

### 不同事项冲突，如何安排？
先跟leader反馈，明确哪个优先级更高。1. 要不就先把手头的事完善好 2. 要不就先去做更紧迫的事情。都很重要的话，请求leader协调人力支持。

### 你有什么想问我的？
请问XX部门/XX团队当下的重点工作是什么呢？
请问对于XXX这个岗位，你对我当下的努力方向与重点有什么建议吗？
期望我的加入带来哪些改变?
对于领导您，您觉得最喜欢公司的地方是哪里？
这个部门在公司是什么定位？

### 你的预期薪资是多少？
我想先了解下Pony的薪资构成，base/年终奖/股票都是怎么算的？以及我的职级是怎样的？

我目前是34K*15+2w的期权，总包是53w。涨幅30%是比较理想的，当然这些对我来说都是可以谈的。

### 最低能接受的薪水是多少？
我知道这个工作的薪水的大概范围是30K-45K。我希望公司能根据我的情况和市场标准的水平，给我合理的薪水

### 离职原因？
离职原因主要是个人家庭原因，我跟我老婆都想到广州生活和发展，她现在已经过去了，所以我在看看机会。

### 面试体验如何？
面试体验非常好，反馈迅速，流程很快。面试官水平很高，聊得问题也很深入。


### 两段工作你觉得有什么提升？
1. 在腾讯让我接触到了规范的开发流程，业务划分等，和世界范围内的客户有了交流，熟练使用后端开发用到的组件。
2. 在qcraft，我深入研究了分布式编译平台和泊车自动标注pipeline，尤其是对高性能高可用平台的开发和维护和有了深刻的理解，同时参与自动标注项目也让我对算法优化积累了行业经验。

### 为什么选择小马？
其实一直都有关注小马的新闻，前段时间我在官网看了下机会没有找到合适的，正好前段时间您邀请我面试，我才发现这个岗位，对这个岗位非常感兴趣。

小马作为自动驾驶行业的No.1, 几乎是所有从业工程师中都想加入的企业。企业知名度非常高，行业影响力很大，做的产品也很酷，前段时间也顺利上市，证明了公司的实力。

### 问手上offer，如何考虑？
如果pony给我发offer，薪资合适，肯定就接了。岗位的工作内容我更有兴趣，机器人相对来说可能不是特别稳。

### 最大的优点和缺点（学习能力，自制力，适应能力，乐于助人，缺点是会对追求细节，可能会花费过多的时间）

### 最困难的问题是什么，如何应对的
突出为了团队可以付出自己的牺牲

### 你朋友对你的评价？（我的朋友都说我是一个可以信赖的人。因为，我一旦答应别人的事情，就一定会做到。如果我做不到，我就不会轻易许诺。我觉的我是一个比较随和的人，与不同的人都可以友好相处。在我与人相处时，我总是能站在别人的角度考虑问题。）

### 领导交给你一个很重要但又很艰难的工作，你怎么去处理？
与领导进行详细的沟通，确保我完全理解任务的目标、优先级和期望的成果，以及截止日期，
我会再进一步拆解任务，我会考虑最坏的情况，并预先规划解决方案。
每天汇报工作中的进展和卡点，必要时寻求资源和帮助。

### 你的未来工作规划？
对AI infra很感兴趣，且我有平台开发和深度学习算法开发背景，迫切希望在这个领域深耕

### 怎么设计一个L4打车系统？

### AI infra了解什么？
#### DeepSpeed
DeepSpeed的核心就在于，GPU显存不够，CPU内存来凑。具体点说，DeepSpeed将当前时刻，训练模型用不到的参数，缓存到CPU中，等到要用到了，再从CPU挪到GPU。这里的“参数”，不仅指的是模型参数，还指optimizer、梯度等。

- 分布式计算环境中，主节点负责协调其他节点和进程的工作
- deepspeed支持更大规模的模型训练
- 混合精度训练
- ZeRO可以减少内存占用，优化大模型训练，将模型参数分成了三个部分：Optimizer States、Gradient 和 Model Parameter。在使用 ZeRO 进行分布式训练时，可以选择 ZeRO-Offload 和 ZeRO-Stage3 等不同的优化技术。

大规模深度学习模型训练中有个主要范式：

1. 数据并行
2. 模型并行

- 用 3D 并行化实现万亿参数模型训练： DeepSpeed 实现了三种并行方法的灵活组合：ZeRO 支持的数据并行，流水线并行和张量切片模型并行。
- ZeRO-Offload 使 GPU 单卡能够训练 10 倍大的模型
- 通过 DeepSpeed Sparse Attention 用6倍速度执行10倍长的序列
- 1 比特 Adam：减少 5 倍通信量

ZeRO-Offload和ZeRO-Stage3是DeepSpeed中的不同的Zero-Redundancy Optimization技术，用于加速分布式训练，主要区别在资源占用和通信开销方面。

- ZeRO-Stage3将模型参数分片到不同的GPU上，通过交换节点间通信来降低显存占用，但需要进行额外的通信操作，因此可能会导致训练速度的下降。
- ZeRO-Offload将模型参数分布在CPU和GPU上，通过CPU去计算一部分梯度，从而减少显存占用，但也会带来一定的计算开销。

ZeRO-Offload的做法是：

forward和backward计算量高，因此和它们相关的部分，例如 参数W（fp16），activation，就全放入GPU。

update的部分计算量低，因此和它相关的部分，全部放入CPU中。例如 optimizer states（fp32）和gradients(fp16)等。






