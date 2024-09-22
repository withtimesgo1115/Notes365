## 并行跟并发有什么区别？
- 并行是指多个处理器同时执行多个任务，每个核心实际上可以在同一时间独立地执行不同的任务。
- 并发是指系统有处理多个任务的能力，但是任意时刻只有一个任务在执行。在单核处理器上，多个任务是通过时间片轮转的方式实现的。但这种切换非常快，给人感觉是在同时执行。

## 你对线程安全的理解是什么？
线程安全是并发编程中一个重要的概念，如果一段代码块或者一个方法在多线程环境中被多个线程同时执行时能够正确地处理共享数据，那么这段代码块或者方法就是线程安全的。

可以从三个要素来确保线程安全：

①、原子性：确保当某个线程修改共享变量时，没有其他线程可以同时修改这个变量，即这个操作是不可分割的。
原子性可以通过互斥锁（如 synchronized）或原子操作（如 AtomicInteger 类中的方法）来保证。

②、可见性：确保一个线程对共享变量的修改可以立即被其他线程看到。
volatile 关键字可以保证了变量的修改对所有线程立即可见，并防止编译器优化导致的可见性问题。

③、活跃性问题：要确保线程不会因为死锁、饥饿、活锁等问题导致无法继续执行。

## 说说什么是进程和线程？
进程说简单点就是我们在电脑上启动的一个个应用，比如我们启动一个浏览器，就会启动了一个浏览器进程。进程是操作系统资源分配的最小单位，它包括了程序、数据和进程控制块等。

线程说简单点就是我们在 Java 程序中启动的一个 main 线程，一个进程至少会有一个线程。当然了，我们也可以启动多个线程，比如说一个线程进行 IO 读写，一个线程进行加减乘除计算，这样就可以充分发挥多核 CPU 的优势，因为 IO 读写相对 CPU 计算来说慢得多。线程是 CPU 分配资源的基本单位。


一个进程中可以有多个线程，多个线程共用进程的堆和方法区（Java 虚拟机规范中的一个定义，JDK 8 以后的实现为元空间）资源，但是每个线程都会有自己的程序计数器和栈。

## 如何理解协程？
协程通常被视为比线程更轻量级的并发单元，它们主要在一些支持异步编程模型的语言中得到了原生支持，如 Kotlin、Go 等。

```java
class CompletableFutureExample {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        // 异步执行任务1
        CompletableFuture<Integer> future1 = CompletableFuture.supplyAsync(() -> {
            try {
                Thread.sleep(1000); // 模拟耗时操作
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            return 10;
        });

        // 异步执行任务2
        CompletableFuture<Integer> future2 = CompletableFuture.supplyAsync(() -> {
            try {
                Thread.sleep(1000); // 模拟耗时操作
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            return 20;
        });

        // 合并两个任务的结果并计算
        CompletableFuture<Integer> resultFuture = future1.thenCombine(future2, Integer::sum);

        // 等待最终结果并打印
        System.out.println("结果: " + resultFuture.get());
    }
}
```

## 说说线程的共享内存？
线程之间想要进行通信，可以通过消息传递和共享内存两种方法来完成。那 Java 采用的是共享内存的并发模型。

这个模型被称为 Java 内存模型，也就是 JMM，JMM 决定了一个线程对共享变量的写入何时对另外一个线程可见。

线程之间的共享变量存储在主内存（main memory）中，每个线程都有一个私有的本地内存（local memory），本地内存中存储了共享变量的副本。当然了，本地内存是 JMM 的一个抽象概念，并不真实存在。

线程 A 与线程 B 之间如要通信的话，必须要经历下面 2 个步骤：
- 线程 A 把本地内存 A 中的共享变量副本刷新到主内存中。
- 线程 B 到主内存中读取线程 A 刷新过的共享变量，再同步到自己的共享变量副本中。


## 线程有几种创建方式？
Java 中创建线程主要有三种方式，分别为继承 Thread 类、实现 Runnable 接口、实现 Callable 接口。

1. 继承 Thread 类，重写 run()方法，调用 start()方法启动线程。
```java
class ThreadTask extends Thread {
    public void run() {
        System.out.println("看完二哥的 Java 进阶之路，上岸了!");
    }

    public static void main(String[] args) {
        ThreadTask task = new ThreadTask();
        task.start();
    }
}
```
这种方法的缺点是，由于 Java 不支持多重继承，所以如果类已经继承了另一个类，就不能使用这种方法了。    

2. 实现 Runnable 接口，重写 run() 方法，然后创建 Thread 对象，将 Runnable 对象作为参数传递给 Thread 对象，调用 start() 方法启动线程。
```java
class RunnableTask implements Runnable {
    public void run() {
        System.out.println("看完二哥的 Java 进阶之路，上岸了!");
    }

    public static void main(String[] args) {
        RunnableTask task = new RunnableTask();
        Thread thread = new Thread(task);
        thread.start();
    }
}
```
这种方法的优点是可以避免 Java 的单继承限制，并且更符合面向对象的编程思想，因为 Runnable 接口将任务代码和线程控制的代码解耦了。

3. 实现 Callable 接口，重写 call() 方法，然后创建 FutureTask 对象，参数为 Callable 对象；紧接着创建 Thread 对象，参数为 FutureTask 对象，调用 start() 方法启动线程。
```java
class CallableTask implements Callable<String> {
    public String call() {
        return "看完二哥的 Java 进阶之路，上岸了!";
    }

    public static void main(String[] args) throws ExecutionException, InterruptedException {
        CallableTask task = new CallableTask();
        FutureTask<String> futureTask = new FutureTask<>(task);
        Thread thread = new Thread(futureTask);
        thread.start();
        System.out.println(futureTask.get());
    }
}
```
这种方法的优点是可以获取线程的执行结果。

## 一个 8G 内存的系统最多能创建多少线程?
在确定一个系统最多可以创建多个线程时，除了需要考虑系统的内存大小外，Java 虚拟机栈的大小也是值得考虑的因素。

线程在创建的时候会被分配一个虚拟机栈，在 64 位操作系统中，默认大小为 1M。

通过 java -XX:+PrintFlagsFinal -version | grep ThreadStackSize 这个命令可以查看 JVM 栈的默认大小。

换句话说，8GB = 8 _ 1024 MB = 8 _ 1024 _ 1024 KB，所以一个 8G 内存的系统可以创建的线程数为 8 _ 1024 = 8192 个。

但操作系统本身的运行也需要消耗一定的内存，所以实际上可以创建的线程数肯定会比 8192 少一些。

## 启动一个 Java 程序，你能说说里面有哪些线程吗？
首先是 main 线程，这是程序开始执行的入口。

然后是垃圾回收线程，它是一个后台线程，负责回收不再使用的对象。

还有编译器线程，在及时编译中（JIT），负责把一部分热点代码编译后放到 codeCache 中，以提升程序的执行效率。

- Thread: main (ID=1) - 主线程，Java 程序启动时由 JVM 创建。
- Thread: Reference Handler (ID=2) - 这个线程是用来处理引用对象的，如软引用（SoftReference）、弱引用（WeakReference）和虚引用（PhantomReference）。负责清理被 JVM 回收的对象。
- Thread: Finalizer (ID=3) - 终结器线程，负责调用对象的 finalize 方法。对象在垃圾回收器标记为可回收之前，由该线程执行其 finalize 方法，用于执行特定的资源释放操作。
- Thread: Signal Dispatcher (ID=4) - 信号调度线程，处理来自操作系统的信号，将它们转发给 JVM 进行进一步处理，例如响应中断、停止等信号。
- Thread: Monitor Ctrl-Break (ID=5) - 监视器线程，通常由一些特定的 IDE 创建，用于在开发过程中监控和管理程序执行或者处理中断。


## 调用 start()方法时会执行 run()方法，那怎么不直接调用 run()方法？

在 Java 中，启动一个新的线程应该调用其start()方法，而不是直接调用run()方法。

当调用start()方法时，会启动一个新的线程，并让这个新线程调用run()方法。这样，run()方法就在新的线程中运行，从而实现多线程并发。

如果直接调用run()方法，那么run()方法就在当前线程中运行，没有新的线程被创建，也就没有实现多线程的效果。

## 线程有哪些常用的调度方法？
![alt text](images/java-concurrent/image.png)

### 线程等待
在 Object 类中有一些方法可以用于线程的等待与通知。

①、wait()：当一个线程 A 调用一个共享变量的 wait() 方法时，线程 A 会被阻塞挂起，直到发生下面几种情况才会返回：

线程 B 调用了共享对象 notify()或者 notifyAll() 方法；
其他线程调用了线程 A 的 interrupt() 方法，线程 A 抛出 InterruptedException 异常返回。

②、wait(long timeout) ：这个方法相比 wait() 方法多了一个超时参数，它的不同之处在于，如果线程 A 调用共享对象的 wait(long timeout)方法后，没有在指定的 timeout 时间内被其它线程唤醒，那么这个方法还是会因为超时而返回。

③、wait(long timeout, int nanos)，其内部调用的是 wait(long timout) 方法。

### 线程唤醒
唤醒线程主要有下面两个方法：

①、notify()：一个线程 A 调用共享对象的 notify() 方法后，会唤醒一个在这个共享变量上调用 wait 系列方法后被挂起的线程。

一个共享变量上可能会有多个线程在等待，具体唤醒哪个等待的线程是随机的。

②、notifyAll()：不同于在共享变量上调用 notify() 方法会唤醒被阻塞到该共享变量上的一个线程，notifyAll 方法会唤醒所有在该共享变量上调用 wait 系列方法而被挂起的线程。

Thread 类还提供了一个 join() 方法，意思是如果一个线程 A 执行了 thread.join()，当前线程 A 会等待 thread 线程终止之后才从 thread.join() 返回。

### 线程休眠
sleep(long millis)：Thread 类中的静态方法，当一个执行中的线程 A 调用了 Thread 的 sleep 方法后，线程 A 会暂时让出指定时间的执行权。

但是线程 A 所拥有的监视器资源，比如锁，还是持有不让出的。指定的睡眠时间到了后该方法会正常返回，接着参与 CPU 的调度，获取到 CPU 资源后就可以继续运行。

### 让出优先权
yield()：Thread 类中的静态方法，当一个线程调用 yield 方法时，实际是在暗示线程调度器，当前线程请求让出自己的 CPU，但是线程调度器可能会“装看不见”忽略这个暗示。

### 线程中断
Java 中的线程中断是一种线程间的协作模式，通过设置线程的中断标志并不能直接终止该线程的执行。被中断的线程会根据中断状态自行处理。

### 




