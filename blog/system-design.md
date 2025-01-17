## 如何设计本地缓存
本地缓存不需要跨网络传输，应用和cache都在同一个进程内部，快速请求，适用于首页这种数据更新频率较低的业务场景。本地缓存虽然带来性能优化，不过也是有一些弊端的，缓存与应用程序耦合，多个应用程序无法直接的共享缓存，各应用或集群的各节点都需要维护自己的单独缓存，对内存是一种浪费。  

### 数据结构
因为不同的业务场景使用的数据类型不同，为了通用，在java中我们可以使用泛型，Go语言中暂时没有泛型，我们可以使用interface类型来代替   
- 数据结构：哈希表；    
- key：string类型；   
- value：interface类型；   

### 并发安全
本地缓存的应用肯定会面对并发读写的场景，这是就要考虑并发安全的问题。因为我们选择的是哈希结构，Go语言中主要提供了两种哈希，一种是非线程安全的map，一种是线程安全的sync.map。这里肯定要选择线程安全的map也可以选择更加灵活的map+ sync.RWMutex

### 高性能并发访问
对锁竞争要进行优化。我们可以根据key进行分桶处理，减少锁竞争。我们的key都是string类型，可以使用djb2哈希算法把Key打散进行分桶，然后对每一个桶进行加锁，也就是锁细化，减少竞争。

### 对象上限
因为本地缓存是在内存中存储的，内存都是有限制的，我们不可能无限存储，所以我们可以指定缓存对象的数量，根据我们具体的应用场景去预估这个上限值，默认我们选择缓存的数量为1024。

### 淘汰策略
#### LFU
最近不常用算法，根据数据的历史访问频率来淘汰数据，这种算法核心思想认为最近使用频率低的数据,很大概率不会再使用，把使用频率最小的数据置换出去。

存在的问题：

某些数据在短时间内被高频访问，在之后的很长一段时间不再被访问，因为之前的访问频率急剧增加，那么在之后不会在短时间内被淘汰，占据着队列前头的位置，会导致更频繁使用的块更容易被清除掉，刚进入的缓存新数据也可能会很快的被删除。

#### LRU
最近最少使用算法，根据数据的历史访问记录来淘汰数据，这种算法核心思想认为最近使用的数据很大概率会再次使用，最近一段时间没有使用的数据，很大概率不会再次使用，把最长时间未被访问的数据置换出去

存在问题：

当某个客户端访问了大量的历史数据时，可能会使缓存中的数据被历史数据替换，降低缓存命中率。

#### FIFO
即先进先出算法，这种算法的核心思想是最近刚访问的，将来访问的可能性比较大，先进入缓存的数据最先被淘汰掉。

这种算法采用绝对公平的方式进行数据置换，很容易发生缺页中断问题。

### 过期清除
除了使用缓存淘汰策略清除数据外，还可以添加一个过期时间做双重保证，避免不经常访问的数据一直占用内存。可以有两种做法：

- 数据过期了直接删除
- 数据过期了不删除，异步更新数据

两种做法各有利弊，异步更新数据需要具体业务场景选择，为了迎合大多数业务，我们采用数据过期了直接删除这种方法更友好，这里我们采用懒加载的方式，在获取数据的时候判断数据是否过期，同时设置一个定时任务，每天定时删除过期的数据。

### 缓存监控
很多人对于缓存的监控也比较忽略，基本写完后不报错就默认他已经生效了，这就无法感知这个缓存是否起作用了，所以对于缓存各种指标的监控，也比较重要，通过其不同的指标数据，我们可以对缓存的参数进行优化，从而让缓存达到最优化。如果是企业应用，我们可以使用Prometheus进行监控上报，我们自测可以简单写一个小组件，定时打印缓存数、缓存命中率等指标

### GC调优 
对于大量使用本地缓存的应用，由于涉及到缓存淘汰，那么GC问题必定是常事。如果出现GC较多，STW时间较长，那么必定会影响服务可用性；对于这个事项一般是具体case具体分析，本地缓存上线后记得经常查看GC监控

## 分布式集群中如何保证线程安全？
### 串行化
可以通过串行化可能产生并发问题操作，牺牲性能和扩展性，来满足对数据一致性的要求

### 分布式锁
- 互斥性。在任意时刻，只有一个客户端能持有锁。

- 不会发生死锁。即使有一个客户端在持有锁的期间崩溃而没有主动解锁，也能保证后续其他客户端能加锁。

- 解铃还须系铃人。加锁和解锁必须是同一个客户端，客户端自己不能把别人加的锁给解了。

- 加锁和解锁必须具有原子性。

### Redis如何实现呢？
相当于把Redis当做一个锁的中心，所有的服务器如果有加锁的需求，都需要通过这个“中心”实现。
```bash
set lock uuid nx ex 12
```
设置lock 的值为uuid

nx 表示为不允许再设置

ex 表示过期时间 12s

![](https://i-blog.csdnimg.cn/blog_migrate/8469fbde933f496b9baf4aab2d3e5559.png)

### setnx刚好获取到锁，业务逻辑出现异常，导致锁无法释放
设置过期时间，自动释放锁。

### 可能会释放其他服务器的锁
一台服务器做完需要同步的操作后就可以 del 锁了 让别的服务器去用

在删除锁之前先进行判断看是不是自己的锁，通过 uuid进行锁的标识

如果是自己的锁 那么删除 别人在set lock 为自己的uuid，如果不是自己的说明 已经过了过期时间 则自动释放。

### 删除操作缺乏原子性。
在判断是自己的uuid 后准备删除 这时候正好过期，被另一个线程拿到了这把锁。

也会误删 所以需要使得判断和删除时原子性的，可以使用Lua脚本实现判断和删除操作的原子性。

## 分布式多个机器生成id，如何保证不重复?
### snowflake
snowflake是Twitter开源的分布式ID生成算法，结果是一个long型的ID（64位）。其核心思想是：使用41bit作为毫秒数，10bit作为机器的ID（5个bit是数据中心，5个bit的机器ID），12bit作为毫秒内的流水号（意味着每个节点在每毫秒可以产生 4096 个 ID），最后还有一个符号位，永远是0。

优点：   
1. 毫秒数在高位，自增序列在低位，整个ID都是趋势递增的。

2. 不依赖数据库等第三方系统，以服务的方式部署，稳定性更高，生成ID的性能也是非常高的。

3. 可以根据自身业务特性分配bit位，非常灵活。

缺点：     
1. 强依赖机器时钟，如果机器上时钟回拨，会导致发号重复或者服务会处于不可用状态。


### 用Redis生成ID
因为Redis是单线程的，也可以用来生成全局唯一ID。可以用Redis的原子操作INCR和INCRBY来实现。

此外，可以使用Redis集群来获取更高的吞吐量。假如一个集群中有5台Redis，可以初始化每台Redis的值分别是1,2,3,4,5，步长都是5，各Redis生成的ID如下：

优点：  
1. 不依赖于数据库，灵活方便，且性能优于数据库。
2. 数字ID天然排序，对分页或需要排序的结果很有帮助。

缺点：  
1. 如果系统中没有Redis，需要引入新的组件，增加系统复杂度。
2. 需要编码和配置的工作量较大。

### UUID
常见的方式。可以利用数据库也可以利用程序生成，一般来说全球唯一。UUID是由32个的16进制数字组成，所以每个UUID的长度是128位（16^32 = 2^128）。UUID作为一种广泛使用标准，有多个实现版本，影响它的因素包括时间、网卡MAC地址、自定义Namesapce等等。

优点：   
1. 简单，代码方便。
2. 生成ID性能非常好，基本不会有性能问题。
3. 全球唯一，在遇见数据迁移，系统数据合并，或者数据库变更等情况下，可以从容应对。

缺点：   
1. 没有排序，无法保证趋势递增。
2. UUID往往是使用字符串存储，查询的效率比较低。
3. 存储空间比较大，如果是海量数据库，就需要考虑存储量的问题。
4. 传输数据量大
5. 不可读。

## 为了实现微信发红包的API，并确保红包金额精确到分且不能有人领到的红包里面没钱

1. 确定红包发放的总金额和红包个数
2. 确定每个红包的最小和最大金额范围
3. 计算每个红包的平均金额
4. 以分为单位操作，避免浮点数误差
5. 生成每个红包的金额
6. 转为元
7. 返回红包列表

```py
import random

def distribute_red_packet(total_amount, total_count):
    # 以分为单位操作
    total_amount = int(total_amount * 100) # 扩大100倍 转为分
    min_amount = 1  # 每个红包至少 1 分
    result = []

    for i in range(1, total_count):
        # 计算当前最大金额，确保剩余金额足够剩下的红包分配
        max_amount = (total_amount - (total_count - i) * min_amount)
        
        # 随机生成红包金额在[min_amount, max_amount]之间
        amount = random.randint(min_amount, max_amount)
        result.append(amount)
        total_amount -= amount

    # 最后一个红包直接等于剩余的所有金额
    result.append(total_amount)
    
    # 将分转化为元
    result = [x / 100.0 for x in result]
    
    return result

# 示例使用
red_packets = distribute_red_packet(100, 10)
print("红包金额列表：", red_packets)
print("总金额校验：", sum(red_packets))  # 应该为 100 元
```

## 扫码登录是如何实现的
### 同一产品中的扫码登录
![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d54068a8c7f44589bd48a6c6e405e11e~tplv-k3u1fbpfcp-zoom-in-crop-mark:1512:0:0:0.awebp)

    1、用户发起二维码登录      
    此时网站会先生成一个二维码，同时把这个二维码对应的标识保存起来，以便跟踪二维码的扫码状态，然后将二维码页面返回到浏览器中；浏览器先展示这个二维码，再按照Javascript脚本的指示发起扫码状态的轮询。所谓轮询就是浏览器每隔几秒调用网站的API查询二维码的扫码登录结果，查询时携带二维码的标识。有的文章说这里可以使用WebSocket，虽然WebSocket响应比较及时，但是从兼容性和复杂度考虑，大部分方案还是会选择轮询或者长轮询，毕竟此时通信稍微延迟下也没多大关系。      

    2、用户扫码确认登录：
    用户打开手机App，使用App自带的扫码功能，扫描浏览器中展现的二维码，然后App提取出二维码中的登录信息，显示登录确认的页面，这个页面可以是App的Native页面，也可以是远程H5页面，这里采用Native页面，用户点击确认或者同意按钮后，App将二维码信息和当前用户的Token一起提交到网站API，网站API确认用户Token有效后，更新在步骤1中创建的二维码标识的状态为“确认登录”，同时绑定当前用户。  
    
    3、网站验证登录成功：  
    在步骤1中，二维码登录页面启动了一个扫码状态的轮询，如果用户已经“确认登录”，则轮询访问网站API时，网站会生成二维码绑定用户的登录Session，然后向前端返回登录成功消息。这里登录状态维护是采用的Session机制，也可以换成其它的机制，比如JWT。

为了保证登录的安全，有必要采取一些安全措施，可能包括以下若干方法：
1. 对二维码承载的信息按照某种规则进行处理，App可以在扫码时进行验证，避免任何扫码都去请求登录；
2. 对二维码设置一个过期时间，过期就自动删除，这样使其占用的资源保持在合理范围之内；
3. 限制二维码只能使用一次，防止重放攻击；
4. 二维码使用足够长的随机性字符串，防止被恶意穷举占用；
5. 使用HTTPS传输，保护登录数据不被窃听和篡改。

### 第三方应用的扫码登录
![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/b2de8595cf004b8abd5f67856d16c732~tplv-k3u1fbpfcp-zoom-in-crop-mark:1512:0:0:0.awebp)

1、步骤3 生成微信登录请求记录：当用户扫码并同意登录之后，步骤25中浏览器会重定向到第三方应用，如果之前没有创建一条登录请求记录，网站并不能确定这次登录就是自己发起的，这可能导致跨站请求伪造攻击。比如使用某个应用的微信登录二维码，骗取用户的授权，然后最终回调跳转到其它站点，被回调站点只能被动接受，虽然下一步验证授权Code通不过，微信也可能会认为第三方应用出了某种问题，搞不好被封掉。因此第三方应用创建一条登录请求记录之后，还要把记录的标识拼接到访问微信登录二维码的url中，微信会在用户同意登录后原样返回这个标识，步骤26中第三方应用可以验证这个标识是不是有效的。    
2、步骤17 显示应用名称和请求授权信息：因为微信支持很多的第三方应用，需要明确告知用户正在登录哪个应用，应用可以访问自己的什么信息，这都是用户做出登录决定的必要信息。因此扫码之后，微信手机端就需要去微信开放平台查询下二维码对应的第三方应用信息。    
3、步骤24 登录临时授权Code：微信开放平台没有直接向浏览器返回登录用户的信息，这是因为第三方应用还需要对用户进行授权并保持会话的状态，这适合在应用的服务端来处理；而且直接返回用户信息到浏览器也是不安全的，并不能保证二维码登录请求就是通过指定的第三方应用发起的。第三方应用会在步骤27中携带这个授权Code，加上应用的AppId和AppSecret，再去向微信开放平台发起登录请求，临时授权Code只能使用1次，存下来也不能再用，且只能用在指定的应用（即绑定了AppId），AppSecret是应用从服务端提取的，用来验证应用的身份，这些措施保证了微信授权登录的安全性。不过验证通过后还是没有直接返回用户信息，而是返回了一个access token，应用可以使用这个token再去请求获取用户信息的接口，这是因为开放平台提供了很多接口，访问这些接口都需要有授权才行，所以发放了一个access token给第三方应用，这种授权登录方式叫做OAuth 2.0。基于安全方面的考虑，access token的有效期比较短，开放平台一般还会发放一个refresh token，access token过期之后，第三方应用可以拿着这个refresh token再去换一个新的access token，如果refresh token也过期了或者用户取消了授权，则不能获取到新的access token，第三方应用此时应该注销用户的登录。这些token都不能泄漏，所以需要保存在第三方应用的服务端。


## 一个外卖平台上有一个外卖单子，现在有多名骑手想接这一单，如何保证只有一个骑手可以接到单子？
1. 在发布外卖配送单时，生成一个唯一的标识符（比如订单ID或随机UUID），作为这个配送单的唯一标识。           
2. 当骑手想要接单时，首先通过Redis的分布式锁机制尝试获取锁。只有一个骑手能够成功获取到锁，表示该骑手接到了单子。 
    ```bash
    SETNX order_lock:{orderId} A
    ``` 
    该命令会在 order_lock:{orderId} 不存在时将其设置为骑手 A。如果此操作成功，说明骑手 A 抢到了该单。   

4. 如果骑手成功获取到锁，即成功接到单子，将配送单的信息存储在Redis中，例如使用Hash结构保存配送单的详细信息。 

    设定过期时间：设置锁的过期时间，例如 30 秒，以防止某个骑手抢单后长时间不处理订单    
5. 成功抢单后，骑手在一定时间内确认接单，同时删除该锁 DEL order_lock:{orderId}，表示抢单完成。 如果骑手没有成功获取到锁，表示已经有其他骑手接到了单子，可以给骑手返回一个提示或者重新获取其他的配送单。  

## 如何把一个文件快速下发到100w个服务器？
1. 使用分发工具：使用专门的分发工具，如BitTorrent、rsync等，可以帮助在多个服务器之间并行地进行文件传输，提高传输效率。

2. 利用分布式文件系统：使用分布式文件系统，如Hadoop  HDFS、GlusterFS等，将文件存储在分布式节点上，可以更快地将文件分发到多个服务器上。

3. 使用多线程或并行传输：在传输文件时，采用多线程或并行传输的方式，可以同时向多个服务器传输文件，提高传输速度。

4. 使用多个传输通道：尝试使用多个传输通道同时传输文件，可以增加传输的带宽，加快传输速度。

5. 利用CDN（内容分发网络）：如果服务器部署在不同的地理位置，可以考虑使用CDN服务，将文件缓存到离用户较近的CDN节点，通过CDN节点将文件分发给多个服务器，提高传输速度和可靠性。

## 典型TOPk系列的问题：10亿个数，找出最大的10个。等(10万个数，输出从小到大？有十万个单词，找出重复次数最高十个？)
在10亿个数中找出最大的10个数：    
-  解法1：使用堆数据结构。维护一个大小为10的最小堆，遍历10亿个数，如果当前数比堆顶元素大，则将堆顶元素替换为当前数并对堆进行调整，保持堆的大小为10。最终堆中的数就是最大的10个数。
-  解法2：使用快速选择算法。类似于快速排序算法，每次选择一个枢纽元素将数据分为两部分，左边的部分均小于枢纽元素，右边的部分均大于枢纽元素。如果枢纽元素的位置大于k，则在左边部分继续查找，否则在右边部分继续查找。最终得到的子数组中的前k个元素就是最大的k个数。

在10万个数中输出从小到大的排序结果：
-  解法1：使用快速排序算法。对于给定的数组，选择一个枢纽元素将数组分为两部分，左边的部分都小于枢纽元素，右边的部分都大于枢纽元素。然后递归地对左右两个部分进行快速排序，最终得到的数组就是从小到大的排序结果。   
-  解法2：使用归并排序算法。将数组分成两个部分，分别对两个部分进行排序，然后将排好序的两个部分合并成一个有序的数组。通过递归地进行这个操作，最终得到的数组就是从小到大的排序结果。    

另外，对于十万个单词中找出重复次数最高的十个单词的问题，可以使用哈希表来统计每个单词出现的次数，并维护一个大小为10的最小堆，遍历哈希表，对于每个单词的出现次数，如果大于堆顶元素，则将堆顶元素替换为当前单词，并对堆进行调整。最终堆中的元素就是重复次数最高的十个单词。   

## 某网站/app首页每天会从10000个商家里面推荐50个商家置顶，每个商家有一个权值，你如何来推荐？第二天怎么更新推荐的商家？
初始推荐：根据商家的权值进行排序，选取排名前50的商家作为推荐商家，并将其置顶在网站/app首页上。权值高的商家有更大的概率被选中。   

更新推荐的商家：  

更新商家权值：根据前一天推荐商家的点击量、购买量、评价等指标，对商家的权值进行更新。点击量高、购买量多、评价好的商家权值会提高。   

选择推荐商家：根据更新后的商家权值重新排序，选取排名前50的商家作为第二天的推荐商家，并进行置顶。     

通过这种方式，每天根据商家的表现和用户的反馈实时更新商家的权值，从而不断优化推荐的商家列表，提供更为精准和个性化的推荐服务。

## 给每个组分配不同的IP段，怎么设计一种结构使的快速得知IP是哪个组的?
可以使用CIDR（无类域间路由）来设计一种结构

## 商城秒杀
1. 使用CDN做页面静态化   
    会对活动页面做静态化处理。用户浏览商品等常规操作，并不会请求到服务端。只有到了秒杀时间点，并且用户主动点了秒杀按钮才允许访问服务端。
2. 秒杀按钮  
    大部分用户怕错过秒杀时间点，一般会提前进入活动页面。此时看到的秒杀按钮是置灰，不可点击的。只有到了秒杀时间点那一时刻，秒杀按钮才会自动点亮，变成可点击的。  

    当秒杀开始的时候系统会生成一个新的js文件，此时标志为true，并且随机参数生成一个新值，然后同步给CDN。由于有了这个随机参数，CDN不会缓存数据，每次都能从CDN中获取最新的js代码。
3. 读多写少
    通过缓存来检查库存
4. 缓存击穿
    在高并发下，同一时刻会有大量的请求，都在秒杀同一件商品，这些请求同时去查缓存中没有数据，然后又同时访问数据库。结果悲剧了，数据库可能扛不住压力，直接挂掉。

    使用分布式锁

    当然，针对这种情况，最好在项目启动之前，先把缓存进行预热。即事先把所有的商品，同步到缓存中，这样商品基本都能直接从缓存中获取到，就不会出现缓存击穿的问题了。

5. 缓存穿透   
    布隆过滤器 ，布隆过滤器绝大部分使用在缓存数据更新很少的场景中。
    如果缓存数据更新非常频繁，又该如何处理呢？ 这时，就需要把不存在的商品id也缓存起来。

6. 库存问题    
    真正的秒杀商品的场景，不是说扣完库存，就完事了，如果用户在一段时间内，还没完成支付，扣减的库存是要加回去的。

    所以，在这里引出了一个预扣库存的概念

7. 数据库扣减库存   
    基于数据库的乐观锁，这样会少一次数据库查询，而且能够天然的保证数据操作的原子性。

    ```sql
    update product set stock=stock-1 where id=product and stock > 0;
    ```

8. redis扣减库存   
    redis的incr方法是原子性的，可以用该方法扣减库存。

9. lua脚本扣减库存  
    lua脚本，是能够保证原子性的，它跟redis一起配合使用，能够完美解决上面的问题

10. 分布式锁   
    - set加锁 NX PX
    - 自旋锁，保障性能
    - redission

11. mq异步处理   
真正并发量大的是秒杀功能，下单和支付功能实际并发量很小。所以，我们在设计秒杀系统时，有必要把下单和支付功能从秒杀的主流程中拆分出来，特别是下单功能要做成mq异步处理的。而支付功能，比如支付宝支付，是业务场景本身保证的异步。   

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ab1ac5478e2c40008c3d64b3ac4c17ee~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=746&h=334&s=58244&e=png&b=fefdfd)

增加消息发送表避免消息丢失

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/3552c19d69c447fb92d8648a2300ea1c~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp#?w=818&h=470&s=110904&e=png&b=fefcfc)

如果生产者把消息写入消息发送表之后，再发送mq消息到mq服务端的过程中失败了，造成了消息丢失。

这时候，要如何处理呢？

答：使用job，增加重试机制。

#### 消息重复消费问题
答：加一张消息处理表。

消费者读到消息之后，先判断一下消息处理表，是否存在该消息，如果存在，表示是重复消费，则直接返回。如果不存在，则进行下单操作，接着将该消息写入消息处理表中，再返回。

#### 延迟消费问题
答：使用延迟队列。

我们都知道rocketmq，自带了延迟队列的功能。

12. 如何限流？
- 对同一个用户限流
- 对同一ip限流
- 对接口限流
- 加验证码



## 排行榜（微信步数等）
### 基于redis
Zset的API可以实现全服排行榜

每个人看见的数据排行数据来源自己的微信好友，而微信好友各不相同，所以看到的排行榜也各不相同。

我们当前的 key 是 sport:ranking:20210227，里面只包含了某一天的信息。
只要我们在 key 里面加上用户的属性就可以了，假设我的微信号是 why。那么 key 可以设计为这样 sport:ranking:why:20210227。

如果每个用户都有在redis有一个自己的排行榜，一个用户的分数更新的时候就需要对所有好友的zset更新，这多大的代价啊，对吧？

当以用户为纬度做排行榜的时候，就会出现排行榜巨多的情况，导致维护成本升高。

每个用户看到的排行榜不一样，我们其实不用时时刻刻帮用户维护好排行榜。维护好了，用户还不一定来看，出力不讨好的节奏。所以还不如延迟到用户请求的阶段。当用户请求查看排行榜的时候，再去根据用户的好友关系，循环获取好友的步数，生成排行榜。

![](https://why-image-1300252878.cos.ap-chengdu.myqcloud.com/img/83/20210225213708.png)

```bash
hmset sport:ranking:why:20210227:jay nickName jay headPhoto xxx likeNum 520 walkNum 66079
```

有些排行榜可以考虑在前端/客户端做，比如：排序数据量不大/排序场景很固定，面试时提到这点很加分

## 类微博的feed流系统
什么是 Feeds 流？ 从用户层面来说， 各种手机 APP 里面， 特别是社交类的， 我们可以看到关注的内容、好友的动态聚合成一个列表(最典型的就是微信朋友圈)都是 feeds 流的一种形式。

Feeds 流的核心功能就是: 信息聚合
它可以根据你的行为去聚合你想要的信息，然后再将它们以轻松易得的方式提供给你。

### feeds流分类
从信息源聚合来看， Feeds 的信息源聚合有三种场景:  
- 无依赖关系: 如抖音推荐页可以从你的操作行为中生成你的用户画像，再去匹配聚合信息
- 单向依赖关系: 譬如微博我关注了某个大v，就可以获取他发布的信息。这里的信息聚合依据是单向的关注关系
- 双向依赖关系: 如微信朋友圈，需要两个人互相通过好友，才会聚合对方的信息到自己的朋友圈中

从展示逻辑上来看， 又分为两种:    
权重推荐: 如抖音， 依据隐含兴趣推荐信息,按权重排序展示的feeds流    
timeline 展示: 如微博和朋友圈， 依据用户关系拉取信息,按时间顺序展示的feeds流 


|名称	|  说明 |	备注 |
| ---- | ---- | ---- |
| Feed	 | Feed流中的每一条状态或者消息都是Feed，比如朋友圈中的一个状态就是一个Feed，微博中的一条微博就是一个Feed	| 无
| Feed流	| Feed流本质上是数据流，核心逻辑是服务端系统将 “多个发布者的信息内容” 通过 “关注收藏屏蔽等关系” 推送给 “多个接收者”，如公众号订阅消息 |	三大特点：少部分人发布；基于订阅行为关联关系；大多数人读取信息
| Timeline	| Timeline其实是一种Feed流的类型，微博，朋友圈都是Timeline类型的Feed流，但是由于Timeline类型出现最早，使用最广泛，最为人熟知，有时候也用Timeline来表示Feed流 | 	也叫时间线
| 关注页timeline	| 展示其他人Feed消息的页面，比如朋友圈，微博的首页等。	| 又叫做收件箱，每个用户能看到的消息都会被存储到收件箱中
| 个人页Timeline	| 展示自己发送过的Feed消息的页面，比如微信中的相册，微博的个人页等	| 又叫做发件箱，自己发布的消息都会被记录到自己的发件箱中。别人的收件箱内的消息，也是从他的各个关注人的发件箱内同步过来的

### timeline Feeds 设计功能点
一个 timeline Feeds 模型需要开发的功能包括:

1. 用户发布/删除 Feed  
2. 用户关注/取消关注其他用户
3. 用户查看订阅的消息流（Feeds流）：用户可以以timeline的形式查看所有订阅的消息源发布的消息。消息的删除和更新，都会实时被用户感知到。Feeds流的翻页问题：用户翻页Feeds流的时候，不管Feeds流更新了多少内容，此时都是沿着最后一次看到的信息往下看。Feeds流前面的信息被删改不予理会
4. 用户可以查看某个用户的主页， 看其他用户曾经发布的 feed
5. 用户对某条 feed 阅读/点赞/评论/转发等
6. 额外功能: 发布内容安全合规审核/黑白名单配置等


### 设计面临的问题
feeds 流系统通用的特点(挑战):

1. 实时性: 消息是实时产生，实时消费，实时推送的。 整体性能要求较高
2. 海量数据: 消息来自不同的数据源， 产生的消息是海量的
3. 读多写少: 一般读写比为 100:1 , 一个用户发布 feed 有 100 个用户会阅读此 feed

### 读扩散模式
1. 订阅者去拉取 feeds 时，订阅者主动去查询关注列表，逐一请求出所有关注人的发件箱中未阅读过的 feed(通过上一次拉取的时间戳)
2. 拿到多个 feed ID 后通过时间戳对其排序， 得到一个 list， 然后进行聚合展示返回

#### 读扩散分页问题:
由于读扩散下，用户的收件箱是实时计算出来的，翻页的时候，需要去所有关注人的发件箱中拉取一定量的数据。拉取后，需要记录当前拉取到了写信箱的 write_last_id，多少个关注就要记录了多少个 write_last_id。而后翻页的时候，需要用这些write_last_id往后拉取新的一定量（比如page_size个）的数据。再用这些数据组成的新收件箱列表，筛选 page_size 条返回前端。同时，还需要更新他实际拉取了消息的写信箱中的 write_last_id，并且存储。当下一次翻页的时候，这批 write_last_id 将作为下次的翻页时定位的依据

总结: 读扩散模式，写 feed 逻辑简单， 节约存储， 但是读性能差， 分页功能实现复杂

### 写扩散模式
1. 当发布 feed 时， 查询发布者的粉丝列表， 并将发布的 feed ID 写入粉丝的收件箱
2. 读取时， 直接读取自身的收件箱， 然后打包成 feeds list 进行聚合展示

#### 写扩散下分页:
由于用户收件箱都是写好的， 直接用 last_id 往下翻即可

总结: 写扩散模式读性能较好，但是浪费存储， 并且大V用户写扩散太慢会出现时效性问题

### 改进方案-推拉结合
我们上面提到过 feeds 流系统是一个读多写少的系统, 所以选择写扩散会更好， 不过针对上面提到的大V用户问题对写的放大太严重了， 性能受到较大影响。

所以我们采取推拉结合模式:

- 针对大V用户， 读扩散， 生成 feed 列表
- 针对普通用户， 写扩散， 生成 feed 列表

具体操作:
1. 发布 feed 时， 如果是大V则仅写入自己的发件箱中
2. 发布 feed 时， 如果是普通用户则进行写扩散推出去
3. 读 feed 时， 读取关注列表判断哪些是大V用户， 拉取大V的发件箱(同样按照上面的 write_last_id 拉取)， 并行读取自己的收件箱， 拿到两个 feedID list 进行合并

### 继续改进-用户分级策略
当我们解决了大V的写扩散问题后， 又面临着新的问题:

1. 如何识别大V用户才能避免边界问题导致性能抖动(用户的粉丝量是一个动态的值， 如何标记一个用户是大V？)
2. app 注册用户很多， 但是活跃用户很少， 如果为某个用户都存储收件箱是否会占据太多的存储成本(存储浪费)

针对大V用户进行打标:

1. 通过粉丝数/离线热度计算/机器学习模型打标等手段进行标识用户是否是大V， 并且将大V作为一种用户标签进行存储
2. 通过 flink 等流式计算， 来标识是否是大V发文
3. 大V用户只能升级不能降级， 一旦降级需要回溯所有粉丝的收件箱(重新写入所有粉丝的收件箱)

针对活跃用户进行用户分级:

1. 基于日活/月活来判断一个用户是否是活跃用户， 甚至可以维护一个活跃级别
2. 譬如月内活跃为一级，收件箱长度保留100条。周活跃为二级，收件箱长度保留300条。日活跃为3级，收件箱长度保留1000条(节约存储成本)

### 冷热分离+预拉取-收件箱过大问题
如果用户关注的列表过多，会导致这个用户的收件箱列表成为一个大 key， 这类用户的性能上会有影响

1. 为了避免用户的收件箱在 redis 中无限增长， 可以对活跃用户做一个限制， 默认最多刷新1000条
2. 如果用户持续拉取内容， 超过1000条， 可以退化为拉模式， 去关注者的发件箱拉取(每次拉取100条来更新用户的收件箱)
3. 在写扩散的过程中， 只添加新的 feed 到列表， 删除超过限制的 feed(写入新的 100条， 删除最老的 100条)

### 软删除+懒删除-写扩散下删除问题
写扩散模式下，用户发布消息可以慢慢扩散出去，但是删除，修改都要扩散出去，速度过慢会出现时效性问题。而且，如果真的是删除了数据，可能会影响Feeds流的分页功能

这种情况， 我们可以采用软删除+懒删除机制:   
软删除是指消息内容不进行实际删除，而是将消息置为删除状态即可，不扩散出去。如此一来，用户在自己的读取收件箱中消息的时候，是先获取了消息 Id 后，再去数据库查出消息内容，而后判断状态进行过滤，把已经删除的状态剔除，不返回给前端。此时也需要重新进行捞数据，填充分页内容。  
懒删除是指如果过滤了某个消息，此时才把消息从用户收件箱中真正删除。（redis的zset中的对应id进行剔除，完成Feeds流表的刷新）

软删除+懒删除的机制具体的实现方案较: 读扩散回查:
我们在写扩散时，只写了一个消息id到用户的收件箱中，所以，用户查询收件箱信息的时候，要进行一个回查将信息丰富（该方案相比直接把内容一起写入收件箱内会更加节约内存，减少冗余数据，同时消息删除无需扩散）。

## 点赞系统设计
### 点赞服务应该支持下面的能力：
- 对某个稿件点赞（取消点赞）、点踩（取消点踩）
- 点赞状态查询
- 查询某个稿件的点赞数
- 查询某个用户的点赞列表
- 查询某个稿件的点赞人列表
- 查询用户收到的总点赞数

### 平台能力：
- 提供业务快速接入的能力
- 数据存储上，具备数据隔离存储能力

### 容灾能力：
- 存储不可用
- 消息队列不可用
- 数据同步延迟
- 点赞消息堆积

### 流量压力
全局流量压力：    
- 针对写流量，为了保证数据写入性能，我们在写入【点赞数】数据的时候，在内存中做了部分聚合写入，比如聚合10s内的点赞数，一次性写入。如此可大量减少数据库的IO次数。
- 数据库的写入我们也做了全面的异步化处理，保证了数据库能以合理的速率处理写入请求。
- 为了保证点赞状态的正确性，且同时让程序拥有自我纠错的能力，我们在每一次更新点赞状态之前，都会取出老的点赞状态作为更新依据。例如：如果当前是取消点赞，那么需要数据库中已有点赞状态。

热点流量压力：   
- 当一个稿件成为超级热门的时候，大量的流量就会涌入到存储的单个分片上，造成读写热点问题。此时需要有热点识别机制来识别该热点，并将数据缓存至本地，并设置合理的TTL。   

数据存储压力：
- KV化存储

面对未知灾难：  
- DB宕机、Redis集群抖动、机房故障、网络故障等   

### 系统架构
![](https://i1.hdslb.com/bfs/article/758b2b4bef2f3dd719ef82ccf3bf077f9331d7e4.png@1192w.avif)

- 流量路由层
- 业务网关层
- 点赞服务
- 点赞异步任务
- 数据层

#### 基本数据类型
- 点赞记录表：记录用户在什么时间对什么实体进行了什么类型的操作(是赞还是踩，是取消点赞还是取消点踩)等
- 点赞计数表：记录被点赞实体的累计点赞（踩）数量  

第一层存储 DB层  
- 点赞记录表，每一次点赞记录，在用户ID+消息ID建立联合索引
- 点赞计数表，业务ID+消息ID作为主键
- 由于DB采用的是分布式数据库，业务无序分库分表

第二层存储缓存层 CacheAside模式
- 点赞数   
key-value = count:patten:{business_id}:{message_id} - {likes},{disLikes}
用业务ID和该业务下的实体ID作为缓存的Key,并将点赞数与点踩数拼接起来存储以及更新
- 用户点赞列表      
key-value = user:likes:patten:{mid}:{business_id} - member(messageID)-score(likeTimestamp)   
用mid与业务ID作为key，value则是一个ZSet,member为被点赞的实体ID，score为点赞的时间。当改业务下某用户有新的点赞操作的时候，被点赞的实体则会通过 zadd的方式把最新的点赞记录加入到该ZSet里面来
为了维持用户点赞列表的长度（不至于无限扩张），需要在每一次加入新的点赞记录的时候，按照固定长度裁剪用户的点赞记录缓存。该设计也就代表用户的点赞记录在缓存中是有限制长度的，超过该长度的数据请求需要回源DB查询  

第三层存储 本地缓存
LocalCache - 本地缓存
- 本地缓存的建立，目的是为了应对缓存热点问题。
- 利用最小堆算法，在可配置的时间窗口范围内，统计出访问最频繁的缓存Key,并将热Key（Value）按照业务可接受的TTL存储在本地内存中。
- 其中热点的发现之前也有同步过：https://mp.weixin.qq.com/s/C8CI-1DDiQ4BC_LaMaeDBg


  

