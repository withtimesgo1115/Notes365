## 面试复盘问题
### Go的GMP模型
G: goroutine协程 go运行时调度   
M: thread线程 实际执行goroutine的 M和P一对一绑定   
P: cpu的抽象，调度器，每个P都有一个本地队列存放要执行的g，默认数量和CPU核心数一致

P 还会从全局队列中偷取 goroutine 来均衡负载。当 P 本地任务队列为空时，它会尝试从其他 P 的队列中“偷” goroutine 执行。   


#### 调度过程：

每个 P 会维护一个本地队列来存储 goroutine，当某个 M 绑定了 P 后，M 将从 P 的队列中取出 goroutine 执行。
P 还会从全局队列中偷取 goroutine 来均衡负载。当 P 本地任务队列为空时，它会尝试从其他 P 的队列中“偷” goroutine 执行。

#### 多线程并发：
GOMAXPROCS 控制了可以并发执行的 P 的数量。如果系统有 4 个 CPU 核心，默认情况下 Go 会创建 4 个 P，最多会有 4 个 goroutine 并发执行。
由于 P 与 M 一一绑定，因此有多少个 P，就有多少个 M 被分配执行

#### goroutine 的创建与切换：
当 Go 程序创建新的 goroutine 时，新的 G（goroutine）会被放到 P 的本地任务队列中。
如果一个 goroutine 发生了阻塞（如 I/O 或系统调用），与它关联的 M 会释放 P，从而使 P 能够被其他可用的 M 绑定，并继续执行其他任务。同时，阻塞的 M 会等待系统调用完成后重新获取一个 P 来继续执行。
当 G 需要被调度或中断时，Go 运行时会进行上下文切换，将当前 G 保存到任务队列中，并从任务队列中调度另一个 G 执行。

#### 任务窃取机制：
为了避免某些 P 上的 goroutine 堆积过多，而另一些 P 上没有任务执行，Go 运行时实现了任务窃取机制（work stealing）。当一个 P 没有任务可执行时，它会尝试从其他 P 的队列中窃取一半的任务，确保所有 P 都能均衡地处理 goroutine。



### GMP 调度的主要特点
#### 高效并发：
Go 通过 Goroutine 让开发者轻松创建大量并发任务，GMP 模型的设计使得 Go 运行时可以高效地管理和调度这些任务。
Goroutine 的调度不依赖操作系统内核调度器，而是由 Go 运行时进行用户态调度，避免了操作系统线程的频繁切换开销。

#### 自动扩展与负载均衡：
Go 运行时会根据需要动态创建或销毁 M，即操作系统线程，以处理大量的 goroutine，同时通过任务窃取机制实现负载均衡。

#### 阻塞处理：
当 M 阻塞时，不会影响 P 执行其他 goroutine，P 会解绑当前阻塞的 M 并绑定到其他可用的 M 上，保证其他 goroutine 能继续被调度执行。


### Go 垃圾回收 - 三色标记法
### 初始状态：
在 GC 开始时，所有对象都是白色的，表示它们尚未被标记。根对象（如全局变量和栈变量）会被初始化为灰色。

### 标记阶段：
在标记阶段，GC 会从根对象（栈、全局变量、寄存器等）开始，将所有直接可达的对象标记为灰色。
然后，GC 会处理所有灰色对象，逐个检查它们的引用。
对于每个灰色对象，GC 会将其所有引用的对象标记为灰色，并将当前对象标记为黑色。这意味着当前对象及其引用的对象都是可达的。

### 处理灰色对象：
重复这个过程，直到所有灰色对象都被处理完。处理灰色对象的过程如下：
从灰色队列中取出一个灰色对象。
将该对象标记为黑色。
遍历该对象的所有引用，将引用的对象标记为灰色（如果它们是白色的）。

### 结束标记阶段：
当灰色队列为空时，标记阶段结束。此时，所有仍然是白色的对象都被视为不可达，可以被回收。
所有黑色和灰色对象都是可达的。

### 清除阶段：
在清除之前，需要进行STW  
在标记阶段之后，GC 会进入清除阶段，回收所有白色对象的内存。此时，白色对象会被释放。  
然后解除STW。

### Go 垃圾回收 - 屏障技术
并发回收的屏障技术归根结底就是在利用内存写屏障来保证强三色不变性和弱三色不变性。

在垃圾回收的并发标记阶段，写屏障被用来保证对象的引用关系正确被追踪。具体来说，如果一个已经被标记（黑色）的对象引用了一个尚未被标记（白色）的对象，写屏障会立即将该白色对象标记（变为灰色或者黑色），以确保垃圾回收器不会错过这个新的引用关系。

Go语言的垃圾回收器在1.8版本引入了混合写屏障。这种混合写屏障结合了删除屏障（delete barrier）和插入屏障（insert barrier）两种写屏障技术，让垃圾回收器可以在并发标记阶段运行，而无需停止整个程序。


### Java hashmap底层存储结构
桶数组+链表或者红黑树

每个Key都会经过哈希算法得到在桶数组中的索引，如果位置为空，则直接插入，否则说明发生了哈希碰撞，使用链表存储元素。如果链表长度大于8且数组的长度大于 64, 那么就会变为红黑树。

HashMap 的初始容量是 16，随着元素的不断添加，HashMap 的容量可能不足，于是就需要进行扩容，阈值是capacity * loadFactor，capacity 为容量，loadFactor 为负载因子，默认为 0.75， 也就是说达到75%就进行一次扩容。

扩容后的数组大小是原来的 2 倍，然后把原来的元素重新计算哈希值，放到新的数组中，这一步也是 HashMap 最耗时的操作。

### Java ConcurrentHashMap实现线程安全机制？是否可以插入null值？
ConcurrentHashMap 在 JDK 7 时采用的是分段锁机制（Segment Locking），整个 Map 被分为若干段，每个段都可以独立地加锁。因此，不同的线程可以同时操作不同的段，从而实现并发访问。

在 JDK 8 及以上版本中，ConcurrentHashMap 的实现进行了优化，不再使用分段锁，而是使用了一种更加精细化的锁——桶锁，以及 CAS 无锁算法。每个桶（Node 数组的每个元素）都可以独立地加锁，从而实现更高级别的并发访问。

同时，对于读操作，通常不需要加锁，可以直接读取，因为 ConcurrentHashMap 内部使用了 volatile 变量来保证内存可见性。

对于写操作，ConcurrentHashMap 使用 CAS 操作来实现无锁的更新，这是一种乐观锁的实现，因为它假设没有冲突发生，在实际更新数据时才检查是否有其他线程在尝试修改数据，如果有，采用悲观的锁策略，如 synchronized 代码块来保证数据的一致性。


HashMap允许插入null值，因为hashmap不考虑并发
ConcurrentHashmap不允许，会报空指针异常，因为考虑并发，避免引入不确定性

### 静态方法和实例方法有何不同？
静态方法：static 修饰的方法，也被称为类方法。在外部调⽤静态⽅法时，可以使⽤"类名.⽅法名"的⽅式，也可以使⽤"对象名.⽅法名"的⽅式。静态方法里不能访问类的非静态成员变量和方法。

实例⽅法：依存于类的实例，需要使用"对象名.⽅法名"的⽅式调用；可以访问类的所有成员变量和方法。

### 静态方法为什么只能访问静态变量
这些方法和变量是属于类的，不依赖于对象。而非静态方法和变量是属于具体的对象的，它们依赖于对象的存在，需要具体的对象才能被调用。


### 重载和重写有什么区别？
如果一个类有多个名字相同但参数个数不同的方法，我们通常称这些方法为方法重载（overload）。如果方法的功能是一样的，但参数不同，使用相同的名字可以提高程序的可读性。

如果子类具有和父类一样的方法（参数相同、返回类型相同、方法名相同，但方法体可能不同），我们称之为方法重写（override）。方法重写用于提供父类已经声明的方法的特殊实现，是实现多态的基础条件。  

方法重载发生在同一个类中，同名的方法如果有不同的参数（参数类型不同、参数个数不同或者二者都不同）。   

方法重写发生在子类与父类之间，要求子类与父类具有相同的返回类型，方法名和参数列表，并且不能比父类的方法声明更多的异常，遵守里氏代换原则。  

### 什么是可变长参数？遇到方法重载的情况怎么办呢？会优先匹配固定参数还是可变参数的方法呢？
可变长参数（Variable Arguments，varargs）是 Java 提供的一种特性，允许方法接受零个或多个参数。使用可变长参数可以简化方法的调用，使得同一个方法可以接受不同数量的参数。

可变长参数的语法是使用三个点（...）来声明参数类型，表示方法可以接收任意数量的该类型参数。可变长参数在方法内部被视为一个数组。

在方法重载的情况下，Java 会根据以下规则决定调用哪个方法：

优先匹配固定参数：如果调用的方法参数数量与固定参数方法完全匹配，Java 将优先选择该方法。  
匹配可变参数：如果没有找到与固定参数完全匹配的方法，则 Java 会选择可变参数的方法。

可变参数必须是参数列表中的最后一个参数，不能在其后再声明其他参数。


### == 和 equals() 的区别
==：用于比较两个对象的引用，即它们是否指向同一个对象实例。
如果两个变量引用同一个对象实例，== 返回 true，否则返回 false。

对于基本数据类型（如 int, double, char 等），== 比较的是值是否相等。

equals() 方法：用于比较两个对象的内容是否相等。默认情况下，equals() 方法的行为与 == 相同，即比较对象引用。然而，equals() 方法通常被各种类重写。例如，String 类重写了 equals() 方法，以便它可以比较两个字符串的字符内容是否完全一样。  


### 为什么重写 equals() 时必须重写 hashCode() 方法？
每个对象都有一个equals方法，默认比较的是对象的引用是否相等，但是我们通常需要比较两个对象的内容是否相等，这就是要重写equals的原因。  

在使用散列数据结构时，比如哈希表，我们希望相等的对象具有相等的哈希码。如果两个相等的对象具有不同的哈希码，那么他们会被存到不同的位置，导致无法正确查找到这些对象。

### String、StringBuffer、StringBuilder 的区别？
String、StringBuilder和StringBuffer在 Java 中都是用于处理字符串的，它们之间的区别是，String 是不可变的，平常开发用得最多，当遇到大量字符串连接时，就用 StringBuilder，它不会生成很多新的对象，StringBuffer 和 StringBuilder 类似，但每个方法上都加了 synchronized 关键字，所以是线程安全的。

- String类的对象是不可变的。也就是说，一旦一个String对象被创建，它所包含的字符串内容是不可改变的。
- 每次对String对象进行修改操作（如拼接、替换等）实际上都会生成一个新的String对象，而不是修改原有对象。这可能会导致内存和性能开销，尤其是在大量字符串操作的情况下。

- StringBuilder提供了一系列的方法来进行字符串的增删改查操作，这些操作都是直接在原有字符串对象的底层数组上进行的，而不是生成新的 String 对象。
- StringBuilder不是线程安全的。这意味着在没有外部同步的情况下，它不适用于多线程环境。
- 相比于String，在进行频繁的字符串修改操作时，StringBuilder能提供更好的性能。 Java 中的字符串连+操作其实就是通过StringBuilder实现的。

- StringBuffer和StringBuilder类似，但StringBuffer是线程安全的，方法前面都加了synchronized关键字。

String：适用于字符串内容不会改变的场景，比如说作为 HashMap 的 key。   
StringBuilder：适用于单线程环境下需要频繁修改字符串内容的场景，比如在循环中拼接或修改字符串，是 String 的完美替代品。   
StringBuffer：现在已经不怎么用了，因为一般不会在多线程场景下去频繁的修改字符串内容。  

### 缓存雪崩、穿透、击穿分别指什么情况？怎么避免？
#### 缓存击穿
原因：热点key过期  

解决方法：
1. 加锁更新，⽐如请求查询 A，发现缓存中没有，对 A 这个 key 加锁，同时去数据库查询数据，写⼊缓存，再返回给⽤户，这样后⾯的请求就可以从缓存中拿到数据了。分布式锁，setnx， get key， 拿不到set value，比较重的操作

2. 将过期时间组合写在 value 中，通过异步的⽅式不断的刷新过期时间，防⽌此类现象，实现自动续期。

3. 本地队列 + 自旋锁     
    当某个热点数据的缓存失效后，首先将请求入队列，而不是直接打到数据库。可以使用一个本地阻塞队列来存储这些请求。   

    每个热点数据维护一个队列，多个热点数据可以使用多个队列来分摊请求压力。  

    每个请求在进入队列后，通过自旋锁不断检查缓存是否已更新。    
    如果缓存已更新，则直接从缓存返回数据；如果未更新，则等待获取处理结果。

    仅允许一个线程从队列中取出请求并查询数据库，将数据重新加载到缓存。   
    其他等待的请求在缓存更新后，能直接从缓存中获取数据，避免重复查询数据库。

#### 缓存穿透
缓存穿透是指查询不存在的数据，由于缓存没有命中（因为数据根本就不存在），请求每次都会穿过缓存去查询数据库。

①、缓存空值/默认值

在数据无法命中之后，把一个空对象或者默认值保存到缓存，之后再访问这个数据，就会从缓存中获取，这样就保护了数据库。

缓存空值有两大问题：

空值做了缓存，意味着缓存层中存了更多的键，需要更多的内存空间（如果是攻击，问题更严重），比较有效的方法是针对这类数据设置一个较短的过期时间，让其自动剔除。
缓存层和存储层的数据会有一段时间窗口的不一致，可能会对业务有一定影响。
例如过期时间设置为 5 分钟，如果此时存储层添加了这个数据，那此段时间就会出现缓存层和存储层数据的不一致。

这时候可以利用消息队列或者其它异步方式清理缓存中的空对象。

②、布隆过滤器

除了缓存空对象，我们还可以在存储和缓存之前，加一个布隆过滤器，做一层过滤。

布隆过滤器里会保存数据是否存在，如果判断数据不存在，就不会访问存储。

#### 缓存雪崩
大量key同时过期

对于缓存数据，设置不同的过期时间，避免大量缓存数据同时过期。可以通过在原有过期时间的基础上添加一个随机值来实现，这样可以分散缓存过期时间，减少同一时间对数据库的访问压力。

### redis内存淘汰机制

### Redis 有序集合底层实现数据结构
跳表 

ref: https://blog.csdn.net/Zhouzi_heng/article/details/127554294

- 跳表实现了二分查找的有序链表
- 每个元素插入时随机生成它的Level
- 最底层包含所有元素
- 如果一个元素出现在level(x), 那么肯定出现在x以下的level中
- 每个索引节点包含两个指针，一个向下一个向右
- 跳表查询、插入、删除的时间复杂度都是O(logn)

### Redis 有序集合为什么不用平衡树？为什么用跳表？
- 红黑树在插入、删除、查找、有序输出元素的复杂度也是OlngN, 但是按照区间来查找数据时，红黑树的效率没有跳表高。按照区间查找数据时，跳表可以做到OlogN的时间复杂度定位区间的起点，然后在原始链表中顺序往后遍历
- 实现简单，不需要平衡树复杂的旋转来维护平衡性


### Mysql InnoDB和MylSAM存储引擎的区别？
InnoDB 和 MyISAM 之间的区别主要表现在存储结构、事务支持、最小锁粒度、索引类型、主键必需、表的具体行数、外键支持等方面。

1. 存储结构：  
MyISAM：用三种格式的文件来存储，.frm 文件存储表的定义；.MYD 存储数据；.MYI 存储索引。
InnoDB：用两种格式的文件来存储，.frm 文件存储表的定义；.ibd 存储数据和索引。
2. 事务支持：   
MyISAM：不支持事务。
InnoDB：支持事务。
3. 最小锁粒度：
MyISAM：表级锁，高并发中写操作存在性能瓶颈。
InnoDB：行级锁，并发写入性能高。
4. 索引类型：
MyISAM 为非聚簇索引，索引和数据分开存储，索引保存的是数据文件的指针。
5. 外键支持：MyISAM 不支持外键；InnoDB 支持外键。
6. 主键必需：MyISAM 表可以没有主键；InnoDB 表必须有主键。
7. 表的具体行数：MyISAM 表的具体行数存储在表的属性中，查询时直接返回；InnoDB 表的具体行数需要扫描整个表才能返回。


### Mysql InnoDB 为什么用B+树不用B树
MySQL 的默认存储引擎是 InnoDB，它采用的是 B+树索引，B+树是一种自平衡的多路查找树，和红黑树、二叉平衡树不同，B+树的每个节点可以有 m 个子节点，而红黑树和二叉平衡树都只有 2 个。

和 B 树不同，B+树的非叶子节点只存储键值，不存储数据，而叶子节点存储了所有的数据，并且构成了一个有序链表。

这样做的好处是，非叶子节点上由于没有存储数据，就可以存储更多的键值对，再加上叶子节点构成了一个有序链表，范围查询时就可以直接通过叶子节点间的指针顺序访问整个查询范围内的所有记录，而无需对树进行多次遍历。查询的效率会更高。

### Mysql 聚簇索引和非聚簇索引区别
在 MySQL 的 InnoDB 存储引擎中，主键就是聚簇索引。聚簇索引不是一种新的索引，而是一种数据存储方式。

在聚簇索引中，表中的行是按照键值（索引）的顺序存储的。这意味着表中的实际数据行和键值之间存在物理排序的关系。因此，每个表只能有一个聚簇索引。

在非聚簇索引中，索引和数据是分开存储的，索引中的键值指向数据的实际存储位置。因此，非聚簇索引也被称为二级索引或辅助索引或非主键索引。表可以有多个非聚簇索引。

这意味着，当使用非聚簇索引检索数据时，数据库首先在索引中查找，然后通过索引中的指针去访问表中实际的数据行，这个过程称为“回表”（Bookmark Lookup）。

举例来说：

- InnoDB 采用的是聚簇索引，如果没有显式定义主键，InnoDB 会选择一个唯一的非空列作为隐式的聚簇索引；如果这样的列也不存在，InnoDB 会自动生成一个隐藏的行 ID 作为聚簇索引。这意味着数据与主键是紧密绑定的，行数据直接存储在索引的叶子节点上。
- MyISAM 采用的是非聚簇索引，表数据存储在一个地方，而索引存储在另一个地方，索引指向数据行的物理位置。

### MySQL几种日志？分别有什么作用？
MySQL 的日志文件主要包括：

1. 错误日志（Error Log）：记录 MySQL 服务器启动、运行或停止时出现的问题。

2. 慢查询日志（Slow Query Log）：记录执行时间超过 long_query_time 值的所有 SQL 语句。这个时间值是可配置的，默认情况下，慢查询日志功能是关闭的。可以用来识别和优化慢 SQL。

3. 一般查询日志（General Query Log）：记录所有 MySQL 服务器的连接信息及所有的 SQL 语句，不论这些语句是否修改了数据。

4. 二进制日志（Bin Log）：记录了所有修改数据库状态的 SQL 语句，以及每个语句的执行时间，如 INSERT、UPDATE、DELETE 等，但不包括 SELECT 和 SHOW 这类的操作。

5. 重做日志（Redo Log）：记录了对于 InnoDB 表的每个写操作，不是 SQL 级别的，而是物理级别的，主要用于崩溃恢复。

6. 回滚日志（Undo Log，或者叫事务日志）：记录数据被修改前的值，用于事务的回滚。

### 如何保证MQ消息不丢失
可能会在这三个阶段发生丢失：生产阶段、存储阶段、消费阶段。

#### 生产
在生产阶段，主要通过请求确认机制，来保证消息的可靠传递。

1、同步发送的时候，要注意处理响应结果和异常。如果返回响应 OK，表示消息成功发送到了 Broker，如果响应失败，或者发生其它异常，都应该重试。
2、异步发送的时候，应该在回调方法里检查，如果发送失败或者异常，都应该进行重试。
3、如果发生超时的情况，也可以通过查询日志的 API，来检查是否在 Broker 存储成功。

#### 存储
存储阶段，可以通过配置可靠性优先的 Broker 参数来避免因为宕机丢消息，简单说就是可靠性优先的场景都应该使用同步。

1. 消息只要持久化到 CommitLog（日志文件）中，即使 Broker 宕机，未消费的消息也能重新恢复再消费。 

2. Broker 的刷盘机制：同步刷盘和异步刷盘，不管哪种刷盘都可以保证消息一定存储在 pagecache 中（内存中），但是同步刷盘更可靠，它是 Producer 发送消息后等数据持久化到磁盘之后再返回响应给 Producer。

3. Broker 通过主从模式来保证高可用，Broker 支持 Master 和 Slave 同步复制、Master 和 Slave 异步复制模式，生产者的消息都是发送给 Master，但是消费既可以从 Master 消费，也可以从 Slave 消费。同步复制模式可以保证即使 Master 宕机，消息肯定在 Slave 中有备份，保证了消息不会丢失。

#### 消费
从 Consumer 角度分析，如何保证消息被成功消费？

Consumer 保证消息成功消费的关键在于确认的时机，不要在收到消息后就立即发送消费确认，而是应该在执行完所有消费业务逻辑之后，再发送消费确认。因为消息队列维护了消费的位置，逻辑执行失败了，没有确认，再去队列拉取消息，就还是之前的一条。

### 如何处理消息重复的问题呢？
需要在业务端做好消息的幂等性处理，或者做消息去重。

幂等性是指一个操作可以执行多次而不会产生副作用，即无论执行多少次，结果都是相同的。可以在业务逻辑中加入检查逻辑，确保同一消息多次消费不会产生副作用。


消息去重，是指在消费者消费消息之前，先检查一下是否已经消费过这条消息，如果消费过了，就不再消费。

业务端可以通过一个专门的表来记录已经消费过的消息 ID，每次消费消息之前，先查询一下这个表，如果已经存在，就不再消费。

### 项目-分布式编译项目  整体架构设计介绍
### 项目-分布式编译项目  Missing Digest异常情况如何处理
### 项目-分布式编译项目  worker节点存储分布不均衡问题 解决思路

### 项目-弹性缓存构建加速项目  代码仓库如何更新

### 项目-腾讯SaaS平台  线程池配合CompletableFuture并发执行具体逻辑 & 1500ms->350ms指标如何得到的  
聚合多源数据：从不同的数据库或第三方服务中获取车辆信息、销售数据、客户数据等，并将它们汇总成一个整体的营销分析结果

原来的接口执行时需要顺序调用多个耗时的子任务，导致整体执行时间很长（约 1500 ms）。为减少接口耗时，可以通过并发执行子任务来缩短整体的响应时间。这里的改进思路是使用 线程池 配合 CompletableFuture 并行执行这些子任务。

- 创建线程池   
- 定义并发任务并使用 CompletableFuture  
对于接口中的每个子任务，将其定义为一个可以返回结果的 CompletableFuture，并使用线程池执行。
- 聚合任务结果

### 项目-腾讯SaaS平台   多级缓存具体实现逻辑
（1）L1一级缓存：本地缓存guava

（2）L2二级缓存：分布式缓存redis

请求优先打到应用本地缓存，本地缓存不存在，再去redis 集群拉取，同时缓存到本地

### 项目-腾讯SaaS平台   如何保证数据一致性和时效性
主要使用发布订阅模式，完成本地cache与Redis缓存的数据同步。

推送模式：每个频道都维护着一个客户端列表，当发送消息时，会遍历该列表并将消息推送给所有订阅者。

拉取模式：发送者将消息放入一个邮箱中，所有订阅该邮箱的客户端可以随时去收取。在确保所有客户端都成功收取完整邮件后，才会删除该邮件。

1. 运营后台保存数据，写入Redis缓存，同时利用 Redis 的发布订阅功能发布信息。
2. 业务应用集群作为消息订阅者，接收到运营数据消息后，删除本地缓存，
3. 当 C 端流量请求到达时，若本地缓存不存在，则从 Redis 中加载缓存至本地缓存。
4. 防止极端情况下，Redis缓存失效，通过定时任务，将数据重新加载到Redis缓存。

####  缓存污染问题
缓存污染问题指的是留存在缓存中的数据，实际不会再被访问了，但是又占据了缓存空间。

如果这样的数据体量很大，甚至占满了缓存，每次有新数据写入缓存时，还需要把这些数据逐步淘汰出缓存，就会增加缓存操作的时间开销。

因此，要解决缓存污染问题，最关键的技术就是能识别出这些只访问一次或是访问次数很少的数据，在淘汰数据时，优先把他们筛选出来淘汰掉。所以，解决缓存污染的核心策略，叫做

缓存中主要常用的缓存淘汰策略：

- random 随机
- lru
- lfu

在实际业务应用中，LRU和LFU两个策略都有应用。

LRU和LFU两种策略关注的数据访问特征各有侧重，LRU策略更加关注数据的时效性，而LFU策略更加关注数据的访问频次。

通常情况下，实际应用的负载具有较好的时间局部性，所以LRU策略的应用会更加广泛。

#### 使用本地缓存注意事项

- 由于本地缓存会占用 Java 进程的 JVM 内存空间，因此不适合存储大量数据，需要对缓存大小进行评估。
- 如果业务能够接受短时间内的数据不一致，那么本地缓存更适用于读取场景。
- 在缓存更新策略中，无论是主动更新还是被动更新，本地缓存都应设置有效期。
- 考虑设置定时任务来同步缓存，以防止极端情况下数据丢失。
- 在 RPC 调用中，需要避免本地缓存被污染，可以通过合理的缓存淘汰策略，来解决这个问题。
- 当应用重启时，本地缓存会失效，因此需要注意加载分布式缓存的时机。
- 通过发布/订阅解决数据一致性问题时，如果发布/订阅模式不持久化消息数据，如果消息丢失，本地缓存就会删除失败。 所以，要解决发布订阅消息的高可用问题。
- 当本地缓存失效时，需要使用 synchronized 进行加锁，确保由一个线程加载 Redis 缓存，避免并发更新。


### 项目-腾讯SaaS平台   Google Guava缓存刷新策略有哪些
1. 手工刷新
```java
String value = loadingCache.get("key");
loadingCache.refresh("key");
```
2. 自动刷新    
Guava Cache 提供了刷新（refresh）机制，可以通过 refreshAfterWrite 方法来设置刷新时间，当缓存项过期的同时可以重新加载新值。  
```java
Cache<String, String> cache = CacheBuilder.newBuilder()
    .refreshAfterWrite(5, TimeUnit.MINUTES)
     // 设置并发级别为3，并发级别是指可以同时写缓存的线程数
    .concurrencyLevel(3)
    .build(new CacheLoader<String, String>() {
        @Override
        public String load(String key) throws Exception {
            // 异步加载新值的逻辑
            return fetchDataFromDataSource(key);
        }
    });
// 在获取缓存值时，如果缓存项过期，将返回旧值 
String value = cache.get("exampleKey");
```
异步加载缓存的原理是重写 reload 方法。

```java
ExecutorService executorService = Executors.newFixedThreadPool(5);
CacheLoader<String, String> cacheLoader = CacheLoader.asyncReloading(
           new CacheLoader<String, String>() {
                  //自动写缓存数据的方法
                  @Override
                  public String load(String key) {
                      System.out.println(Thread.currentThread().getName() + " 加载 key:" + key);
                      // 从数据库加载数据
                      return "value_" + key.toUpperCase();
                  }
            } , executorService);
```

自动刷新的缺点是：当缓存项到了指定过期时间，不管是同步刷新还是异步刷新，绝大部分请求线程都会返回旧的数据值，缓存值会有一定的延迟效果

### 项目-腾讯SaaS平台   CompletableFuture底层实现原理
#### Future实现原理回顾
Java中的Future是一种异步编程的技术，它允许我们在另一个线程中执行任务，并在主线程中等待任务完成后获取结果。Future的实现原理可以通过Java中的两个接口来理解：Future和FutureTask。

Future接口是Java中用于表示异步操作结果的一个接口，它定义了获取异步操作结果的方法get()，并且可以通过isDone()方法查询操作是否已经完成。

在Java 5中引入了FutureTask类，它是一个实现了Future和Runnable接口的类，它可以将一个任务（Runnable或Callable）封装成一个异步操作，通过FutureTask的get()方法可以获取任务执行的结果。

FutureTask的get()方法是一个阻塞方法，如果任务还没有完成，则会一直阻塞当前线程，直到任务完成。这个阻塞的过程可以通过一个volatile类型的变量来实现。在任务执行完成后，会调用done()方法通知FutureTask任务已经完成，并且设置执行结果。done()方法会调用FutureTask的回调函数，完成后将执行结果设置到FutureTask中。

**FutureTask并不能保证任务的执行顺序和执行结果，因为任务的执行是由线程池来控制的。如果需要保证任务的执行顺序和结果，可以使用CompletionService和ExecutorCompletionService。**

综上所述，Future的实现原理就是通过Future和FutureTask接口，将任务封装成一个异步操作，并在主线程中等待任务完成后获取执行结果。FutureTask是Future的一个具体实现，通过阻塞方法和回调函数来实现异步操作的结果获取。

#### Future局限性分析与推荐
- 阻塞问题：Future的get()方法是一个阻塞方法，如果任务没有完成，会一直阻塞当前线程，这会导致整个应用程序的响应性下降。
- 无法取消任务：Future的cancel()方法可以用于取消任务的执行，但如果任务已经开始执行，则无法取消。此时只能等待任务执行完毕，这会导致一定的性能损失。
- 缺少异常处理：Future的get()方法会抛出异常，但是如果任务执行过程中抛出异常，Future无法处理异常，只能将异常抛给调用者处理。
- 缺少组合操作：Future只能处理单个异步操作，无法支持多个操作的组合，例如需要等待多个任务全部完成后再执行下一步操作。

#### CompletableFuture核心原理简述
CompletableFuture的核心原理是基于Java的Future接口和内部的状态机实现的。它可以通过三个步骤来实现异步操作：

- 创建CompletableFuture对象：  
    通过CompletableFuture的静态工厂方法，我们可以创建一个新的CompletableFuture对象，并指定该对象的异步操作。通常情况下，我们可以通过supplyAsync()或者runAsync()方法来创建CompletableFuture对象。
- 异步操作的执行：  
    在CompletableFuture对象创建之后，异步操作就开始执行了。这个异步操作可以是一个计算任务或者一个IO操作。CompletableFuture会在另一个线程中执行这个异步操作，这样主线程就不会被阻塞。
- 对异步操作的处理：   
    异步操作执行完成后，CompletableFuture会根据执行结果修改其内部的状态，并触发相应的回调函数。如果异步操作成功完成，则会触发CompletableFuture的完成回调函数；如果异步操作抛出异常，则会触发CompletableFuture的异常回调函数。

CompletableFuture的优势在于它支持链式调用和组合操作。通过CompletableFuture的then系列方法，我们可以创建多个CompletableFuture对象，并将它们串联起来形成一个链式的操作流。在这个操作流中，每个CompletableFuture对象都可以依赖于之前的CompletableFuture对象，以实现更加复杂的异步操作。

总的来说，CompletableFuture的原理是基于Java的Future接口和内部的状态机实现的，它可以以非阻塞的方式执行异步操作，并通过回调函数来处理异步操作完成后的结果。通过链式调用和组合操作，CompletableFuture可以方便地实现复杂的异步编程任务。

在CompletableFuture中，回调是一种重要的机制，可以在异步任务完成时自动触发回调函数。

CompletableFuture的回调机制可以分为两种类型：完成回调和异常回调。

CompletableFuture的回调机制是通过Java的函数式编程实现的。在CompletableFuture中，回调函数是一个函数接口，例如CompletableFuture的thenAccept方法需要一个Consumer类型的回调函数作为参数。在异步任务完成后，CompletableFuture会自动调用回调函数并传递异步任务的结果。

在实现上，CompletableFuture的回调机制主要依赖于Java的Future接口和CompletableFuture内部的状态机。当CompletableFuture被创建时，它的状态是未完成的。当异步任务完成后，CompletableFuture会修改内部的状态，将结果或异常保存在内部，然后触发相应的回调函数。

当用户调用CompletableFuture的then系列方法时，CompletableFuture会返回一个新的CompletableFuture对象，表示一个新的异步任务。当原始的CompletableFuture完成时，它会自动触发新的CompletableFuture的回调函数。这种链式回调的设计可以方便地实现多个异步任务之间的依赖关系。

总的来说，CompletableFuture的回调机制是通过Java的函数式编程和状态机实现的。通过灵活的回调函数设置和链式调用，CompletableFuture可以方便地实现异步编程。

从CompletableFuture的使用方法可以看出，CompletableFuture主要是通过回调的方式实现异步编程，解决Future在使用过程中需要阻塞的问题。

其结构与观察者模式类似，CompletableFuture是发布者，使用链表保存观察者Completion。

CompletableFuture的postComplete方法是通知方法，用于在CompletableFuture完成时通知观察者，发送订阅的数据。
Completion的tryFire方法用于处理CompletableFuture发布的结果。