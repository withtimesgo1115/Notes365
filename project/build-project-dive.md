## Part1. QBot机器人服务
### 平台如何限流？
1. 在平台编译资源紧张时（比如PV打满），优先把PV分配给优先级高的Job，其他job则阻塞
2. Gitlab Runner Concurrent 参数配置
3. PV资源池QUATA限制
4. trigger调用API触发的job由令牌桶算法控制

### 平台如何监控？
配置普罗米修斯和Grafana监控告警

### 平台如何可视化？
1. 实现失败pipeline，job等统计分析接口
2. 实现失败job的查询接口
3. Runner健康情况的监控告警

### 平台如何自动化管理？
1. 轮询机制对已经存在失败作业的pipeline，cancel其他running的job
2. 事件机制webhook钩子来处理MR，pipeline，消息通知等
3. Marge Bot机器人服务实现自动合入，单线程轮询，检查是否符合合入要求，逐个处理
4. Bot机器人结合GPT API实现code review

## Part2. Buildfarm分布式构建服务

    
## Part3. PVC Runner缓存加速部署