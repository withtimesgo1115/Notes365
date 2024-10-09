## Spring 是什么？特性？有哪些模块？
一句话概括：Spring 是一个轻量级、非入侵式的控制反转 (IoC) 和面向切面 (AOP) 的框架。

### Spring 有哪些特性呢？
1. IoC 和 DI 的支持
Spring 的核心就是一个大的工厂容器，可以维护所有对象的创建和依赖关系，Spring 工厂用于生成 Bean，并且管理 Bean 的生命周期，实现高内聚低耦合的设计理念。

2. AOP 编程的支持
Spring 提供了面向切面编程，可以方便的实现对程序进行权限拦截、运行监控等切面功能。

3. 声明式事务的支持
支持通过配置就来完成对事务的管理，而不需要通过硬编码的方式，以前重复的一些事务提交、回滚的 JDBC 代码，都可以不用自己写了。

4. 快捷测试的支持
Spring 对 Junit 提供支持，可以通过注解快捷地测试 Spring 程序。

5. 快速集成功能
方便集成各种优秀框架，Spring 不排斥各种优秀的开源框架，其内部提供了对各种优秀框架（如：Struts、Hibernate、MyBatis、Quartz 等）的直接支持。

6. 复杂 API 模板封装
Spring 对 JavaEE 开发中非常难用的一些 API（JDBC、JavaMail、远程调用等）都提供了模板化的封装，这些封装 API 的提供使得应用难度大大降低。

## Spring 有哪些模块呢？
Spring 框架是分模块存在，除了最核心的Spring Core Container是必要模块之外，其他模块都是可选，大约有 20 多个模块。

最主要的七大模块：

Spring Core：Spring 核心，它是框架最基础的部分，提供 IoC 和依赖注入 DI 特性。  
Spring Context：Spring 上下文容器，它是 BeanFactory 功能加强的一个子接口。  
Spring Web：它提供 Web 应用开发的支持。  
Spring MVC：它针对 Web 应用中 MVC 思想的实现。  
Spring DAO：提供对 JDBC 抽象层，简化了 JDBC 编码，同时，编码更具有健壮性。  
Spring ORM：它支持用于流行的 ORM 框架的整合，比如：Spring + Hibernate、Spring + iBatis、Spring + JDO 的整合等。   
Spring AOP：即面向切面编程，它提供了与 AOP 联盟兼容的编程实现。    

## Spring 有哪些常用注解呢？
![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-8d0a1518-a425-4887-9735-45321095d927.png)

### Web 开发方面有哪些注解呢？     
①、@Controller：用于标注控制层组件。

②、@RestController：是@Controller 和 @ResponseBody 的结合体，返回 JSON 数据时使用。

③、@RequestMapping：用于映射请求 URL 到具体的方法上，还可以细分为：

@GetMapping：只能用于处理 GET 请求  
@PostMapping：只能用于处理 POST 请求   
@DeleteMapping：只能用于处理 DELETE 请求   

④、@ResponseBody：直接将返回的数据放入 HTTP 响应正文中，一般用于返回 JSON 数据。

⑤、@RequestBody：表示一个方法参数应该绑定到 Web 请求体。

⑥、@PathVariable：用于接收路径参数，比如 @RequestMapping(“/hello/{name}”)，这里的 name 就是路径参数。

⑦、@RequestParam：用于接收请求参数。比如 @RequestParam(name = "key") String key，这里的 key 就是请求参数。

### 容器类注解有哪些呢？
- @Component：标识一个类为 Spring 组件，使其能够被 Spring 容器自动扫描和管理。
- @Service：标识一个业务逻辑组件（服务层）。比如 @Service("userService")，这里的 userService 就是 Bean 的名称。
- @Repository：标识一个数据访问组件（持久层）。
- @Autowired：按类型自动注入依赖。
- @Configuration：用于定义配置类，可替换 XML 配置文件。
- @Value：用于将 Spring Boot 中 application.properties 配置的属性值赋值给变量。

### AOP 方面有哪些注解呢？
@Aspect 用于声明一个切面，可以配合其他注解一起使用，比如：

@After：在方法执行之后执行。
@Before：在方法执行之前执行。
@Around：方法前后均执行。
@PointCut：定义切点，指定需要拦截的方法。

### 事务注解有哪些？
主要就是 @Transactional，用于声明一个方法需要事务支持。


## Spring 中应用了哪些设计模式呢？
Spring 框架中用了蛮多设计模式的：

①、工厂模式：IoC 容器本身可以看作是一个巨大的工厂，负责创建和管理 Bean 的生命周期和依赖关系。

像 BeanFactory 和 ApplicationContext 接口都提供了工厂模式的实现，负责实例化、配置和组装 Bean。

②、代理模式：AOP 的实现就是基于代理模式的，如果配置了事务管理，Spring 会使用代理模式创建一个连接数据库的代理对象，来进行事务管理。

③、单例模式：Spring 容器中的 Bean 默认都是单例的，这样可以保证 Bean 的唯一性，减少系统开销。

④、模板模式：Spring 中的 JdbcTemplate，HibernateTemplate 等以 Template 结尾的类，都使用了模板方法模式。

比如，我们使用 JdbcTemplate，只需要提供 SQL 语句和需要的参数就可以了，至于如何创建连接、执行 SQL、处理结果集等都由 JdbcTemplate 这个模板方法来完成。

④、观察者模式：Spring 事件驱动模型就是观察者模式很经典的一个应用，Spring 中的 ApplicationListener 就是观察者，当有事件（ApplicationEvent）被发布，ApplicationListener 就能接收到信息。

⑤、适配器模式：Spring MVC 中的 HandlerAdapter 就用了适配器模式。它允许 DispatcherServlet 通过统一的适配器接口与多种类型的请求处理器进行交互。

⑥、策略模式：Spring 中有一个 Resource 接口，它的不同实现类，会根据不同的策略去访问资源。

## Spring 容器、Web 容器之间的区别？（补充）
Spring 容器是 Spring 框架的核心部分，负责管理应用程序中的对象生命周期和依赖注入。

Web 容器（也称 Servlet 容器），是用于运行 Java Web 应用程序的服务器环境，支持 Servlet、JSP 等 Web 组件。常见的 Web 容器包括 Apache Tomcat、Jetty等。

Spring MVC 是 Spring 框架的一部分，专门用于处理 Web 请求，基于 MVC（Model-View-Controller）设计模式。

## 说一说什么是 IoC？什么是 DI？
所谓的IoC（控制反转，Inversion of Control），就是由容器来控制对象的生命周期和对象之间的关系。以前是我们想要什么就自己创建什么，现在是我们需要什么容器就帮我们送来什么。

也就是说，控制对象生命周期的不再是引用它的对象，而是容器，这就叫控制反转。

婚介所就相当于一个 IoC 容器，我就是一个对象，我需要的女朋友就是另一个对象，我不用关心女朋友是怎么来的，我只需要告诉婚介所我需要什么样的女朋友，婚介所就帮我去找。

Spring 倡导的开发方式就是这样，所有的类创建都通过 Spring 容器来，不再是开发者去 new，去 = null 销毁，这些创建和销毁的工作都交给 Spring 容器来。

于是，对于某个对象来说，以前是它控制它依赖的对象，现在是所有对象都被 Spring 控制，这就是控制反转。

## 说说什么是 DI？
DI（依赖注入，Dependency Injection）：有人说 IoC 和 DI 是一回事，有人说 IoC 是思想，DI 是 IoC 的实现。2004 年，Martin Fowler 在他的文章《控制反转容器&依赖注入模式》首次提出了依赖注入这个名词。

控制反转这个词太宽泛，并不能很好地解释这个框架的具体实现，于是就想到了一个新名词：依赖注入。

打个比方，你现在想吃韭菜馅的饺子，这时候就有人用针管往你吃的饺子里注入韭菜鸡蛋馅。就好像 A 类需要 B 类，以前是 A 类自己 new 一个 B 类，现在是有人把 B 类注入到 A 类里。

### 为什么要使用 IoC 呢？
在平时的 Java 开发中，如果我们要实现某一个功能，可能至少需要两个以上的对象来协助完成，在没有 Spring 之前，每个对象在需要它的合作对象时，需要自己 new 一个，比如说 A 要使用 B，A 就对 B 产生了依赖，也就是 A 和 B 之间存在了一种耦合关系。

有了 Spring 之后，就不一样了，创建 B 的工作交给了 Spring 来完成，Spring 创建好了 B 对象后就放到容器中，A 告诉 Spring 我需要 B，Spring 就从容器中取出 B 交给 A 来使用。

至于 B 是怎么来的，A 就不再关心了，Spring 容器想通过 newnew 创建 B 还是 new 创建 B，无所谓。

这就是 IoC 的好处，它降低了对象之间的耦合度，使得程序更加灵活，更加易于维护。

### 能简单说一下 Spring IoC 的实现机制吗？
Spring 的 IoC 本质就是一个大工厂，我们想想一个工厂是怎么运行的呢？

- 生产产品：一个工厂最核心的功能就是生产产品。在 Spring 里，不用 Bean 自己来实例化，而是交给 Spring，应该怎么实现呢？——答案毫无疑问，反射。 那么这个厂子的生产管理是怎么做的？你应该也知道——工厂模式。

- 库存产品：工厂一般都是有库房的，用来库存产品，毕竟生产的产品不能立马就拉走。Spring 我们都知道是一个容器，这个容器里存的就是对象，不能每次来取对象，都得现场来反射创建对象，得把创建出的对象存起来。

- 订单处理：还有最重要的一点，工厂根据什么来提供产品呢？订单。这些订单可能五花八门，有线上签签的、有到工厂签的、还有工厂销售上门签的……最后经过处理，指导工厂的出货。

    ![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-1d55c63d-2d12-43b1-9f43-428f5f4a1413.png)

### Bean 定义：
Bean 通过一个配置文件定义，把它解析成一个类型。
```
userDao:cn.fighter3.bean.UserDao
```
BeanDefinition.java

bean 定义类，配置文件中 bean 定义对应的实体

```java
public class BeanDefinition {

    private String beanName;

    private Class beanClass;
     //省略getter、setter
 }
```

ResourceLoader.java  
资源加载器，用来完成配置文件中配置的加载  
```java
public class ResourceLoader {

    public static Map<String, BeanDefinition> getResource() {
        Map<String, BeanDefinition> beanDefinitionMap = new HashMap<>(16);
        Properties properties = new Properties();
        try {
            InputStream inputStream = ResourceLoader.class.getResourceAsStream("/beans.properties");
            properties.load(inputStream);
            Iterator<String> it = properties.stringPropertyNames().iterator();
            while (it.hasNext()) {
                String key = it.next();
                String className = properties.getProperty(key);
                BeanDefinition beanDefinition = new BeanDefinition();
                beanDefinition.setBeanName(key);
                Class clazz = Class.forName(className);
                beanDefinition.setBeanClass(clazz);
                beanDefinitionMap.put(key, beanDefinition);
            }
            inputStream.close();
        } catch (IOException | ClassNotFoundException e) {
            e.printStackTrace();
        }
        return beanDefinitionMap;
    }

}
```

BeanRegister.java

对象注册器，这里用于单例 bean 的缓存，我们大幅简化，默认所有 bean 都是单例的。可以看到所谓单例注册，也很简单，不过是往 HashMap 里存对象。
```java
public class BeanRegister {

    //单例Bean缓存
    private Map<String, Object> singletonMap = new HashMap<>(32);

    /**
     * 获取单例Bean
     *
     * @param beanName bean名称
     * @return
     */
    public Object getSingletonBean(String beanName) {
        return singletonMap.get(beanName);
    }

    /**
     * 注册单例bean
     *
     * @param beanName
     * @param bean
     */
    public void registerSingletonBean(String beanName, Object bean) {
        if (singletonMap.containsKey(beanName)) {
            return;
        }
        singletonMap.put(beanName, bean);
    }

}
```

BeanFactory.java   
对象工厂，我们最核心的一个类，在它初始化的时候，创建了 bean 注册器，完成了资源的加载。

获取 bean 的时候，先从单例缓存中取，如果没有取到，就创建并注册一个 bean
```java
public class BeanFactory {

    private Map<String, BeanDefinition> beanDefinitionMap = new HashMap<>();

    private BeanRegister beanRegister;

    public BeanFactory() {
        //创建bean注册器
        beanRegister = new BeanRegister();
        //加载资源
        this.beanDefinitionMap = new ResourceLoader().getResource();
    }

    /**
     * 获取bean
     *
     * @param beanName bean名称
     * @return
     */
    public Object getBean(String beanName) {
        //从bean缓存中取
        Object bean = beanRegister.getSingletonBean(beanName);
        if (bean != null) {
            return bean;
        }
        //根据bean定义，创建bean
        return createBean(beanDefinitionMap.get(beanName));
    }

    /**
     * 创建Bean
     *
     * @param beanDefinition bean定义
     * @return
     */
    private Object createBean(BeanDefinition beanDefinition) {
        try {
            Object bean = beanDefinition.getBeanClass().newInstance();
            //缓存bean
            beanRegister.registerSingletonBean(beanDefinition.getBeanName(), bean);
            return bean;
        } catch (InstantiationException | IllegalAccessException e) {
            e.printStackTrace();
        }
        return null;
    }
}
```

- 测试
UserDao.java

我们的 Bean 类，很简单
```java
public class UserDao {
    public void queryUserInfo(){
        System.out.println("A good man.");
    }
}
```

```java
public class ApiTest {
    @Test
    public void test_BeanFactory() {
        //1.创建bean工厂(同时完成了加载资源、创建注册单例bean注册器的操作)
        BeanFactory beanFactory = new BeanFactory();

        //2.第一次获取bean（通过反射创建bean，缓存bean）
        UserDao userDao1 = (UserDao) beanFactory.getBean("userDao");
        userDao1.queryUserInfo();

        //3.第二次获取bean（从缓存中获取bean）
        UserDao userDao2 = (UserDao) beanFactory.getBean("userDao");
        userDao2.queryUserInfo();
    }
}
```
##  BeanFactory 和 ApplicantContext?
可以这么比喻，BeanFactory 是 Spring 的“心脏”，而 ApplicantContext 是 Spring 的完整“身躯”。

BeanFactory 主要负责配置、创建和管理 bean，为 Spring 提供了基本的依赖注入（DI）支持。
ApplicationContext 是 BeanFactory 的子接口，在 BeanFactory 的基础上添加了企业级的功能支持。

### 详细说说 BeanFactory
BeanFactory 位于整个 Spring IoC 容器的顶端，ApplicationContext 算是 BeanFactory 的子接口。

它最主要的方法就是 getBean()，这个方法负责从容器中返回特定名称或者类型的 Bean 实例。

```java
class HelloWorldApp{
   public static void main(String[] args) {
      BeanFactory factory = new XmlBeanFactory (new ClassPathResource("beans.xml"));
      HelloWorld obj = (HelloWorld) factory.getBean("itwanger");
      obj.getMessage();
   }
}
```

### 请详细说说 ApplicationContext
ApplicationContext 继承了 HierachicalBeanFactory 和 ListableBeanFactory 接口，算是 BeanFactory 的自动挡版本，是 Spring 应用的默认方式。  

ApplicationContext 会在启动时预先创建和配置所有的单例 bean，并支持如 JDBC、ORM 框架的集成，内置面向切面编程（AOP）的支持，可以配置声明式事务管理等。

```java
class MainApp {
    public static void main(String[] args) {
        // 使用 AppConfig 配置类初始化 ApplicationContext
        ApplicationContext context = new AnnotationConfigApplicationContext(AppConfig.class);

        // 从 ApplicationContext 获取 messageService 的 bean
        MessageService service = context.getBean(MessageService.class);

        // 使用 bean
        service.printMessage();
    }
}
```

通过 AnnotationConfigApplicationContext 类，我们可以使用 Java 配置类来初始化 ApplicationContext，这样就可以使用 Java 代码来配置 Spring 容器。
```java
@Configuration
@ComponentScan(basePackages = "com.github.paicoding.forum.test.javabetter.spring1") // 替换为你的包名
public class AppConfig {
}
```

##  Spring 容器启动阶段会干什么吗？
Spring 的 IoC 容器工作的过程，其实可以划分为两个阶段：容器启动阶段和Bean 实例化阶段。

其中容器启动阶段主要做的工作是加载和解析配置文件，保存到对应的 Bean 定义中。

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-8f8103f7-2a51-4858-856e-96a4ac400d76.png)

容器启动开始，首先会通过某种途径加载 Congiguration MetaData，在大部分情况下，容器需要依赖某些工具类（BeanDefinitionReader）对加载的 Congiguration MetaData 进行解析和分析，并将分析后的信息组为相应的 BeanDefinition。

最后把这些保存了 Bean 定义必要信息的 BeanDefinition，注册到相应的 BeanDefinitionRegistry，这样容器启动就完成了。

### 说说 Spring 的 Bean 实例化方式
Spring 提供了 4 种不同的方式来实例化 Bean，以满足不同场景下的需求。

### 构造方法的方式
在类上使用@Component（或@Service、@Repository 等特定于场景的注解）标注类，然后通过构造方法注入依赖。

```java
@Component
public class ExampleBean {
    private DependencyBean dependency;

    @Autowired
    public ExampleBean(DependencyBean dependency) {
        this.dependency = dependency;
    }
}
```

### 静态工厂的方式
在这种方式中，Bean 是由一个静态方法创建的，而不是直接通过构造方法。
```java
public class ClientService {
    private static ClientService clientService = new ClientService();

    private ClientService() {}

    public static ClientService createInstance() {
        return clientService;
    }
}
```

### 实例工厂方法实例化的方式
与静态工厂方法相比，实例工厂方法依赖于某个类的实例来创建 Bean。这通常用在需要通过工厂对象的非静态方法来创建 Bean 的场景。

```java
public class ServiceLocator {
    public ClientService createClientServiceInstance() {
        return new ClientService();
    }
}
```

### FactoryBean 接口实例化方式
FactoryBean 是一个特殊的 Bean 类型，可以在 Spring 容器中返回其他对象的实例。通过实现 FactoryBean 接口，可以自定义实例化逻辑，这对于构建复杂的初始化逻辑非常有用。

## Spring Bean 生命周期吗？
Spring 中 Bean 的生命周期大致分为四个阶段：实例化（Instantiation）、属性赋值（Populate）、初始化（Initialization）、销毁（Destruction）。

- 实例化：Spring 容器根据 Bean 的定义创建 Bean 的实例，相当于执行构造方法，也就是 new 一个对象。
- 属性赋值：相当于执行 setter 方法为字段赋值。
- 初始化：初始化阶段允许执行自定义的逻辑，比如设置某些必要的属性值、开启资源、执行预加载操作等，以确保 Bean 在使用之前是完全配置好的。
- 销毁：相当于执行 = null，释放资源。

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-942a927a-86e4-4a01-8f52-9addd89642ff.png)

## Bean 定义和依赖定义有哪些方式？
有三种方式：直接编码方式、配置文件方式、注解方式。

- 直接编码方式：我们一般接触不到直接编码的方式，但其实其它的方式最终都要通过直接编码来实现。
- 配置文件方式：通过 xml、propreties 类型的配置文件，配置相应的依赖关系，Spring 读取配置文件，完成依赖关系的注入。
- 注解方式：注解方式应该是我们用的最多的一种方式了，在相应的地方使用注解修饰，Spring 会扫描注解，完成依赖关系的注入。

## 有哪些依赖注入的方法？
Spring 支持构造方法注入、属性注入、工厂方法注入,其中工厂方法注入，又可以分为静态工厂方法注入和非静态工厂方法注入。

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-491f8444-54ba-4628-b8eb-8418a2197096.png)

- 构造方法注入   
通过调用类的构造方法，将接口实现类通过构造方法变量传入
```java
public CatDaoImpl(String message){
   this. message = message;
 }

<bean id="CatDaoImpl" class="com.CatDaoImpl">
    <constructor-arg value=" message "></constructor-arg>
</bean>
```
- 属性注入   
通过 Setter 方法完成调用类所需依赖的注入
```java
 public class Id {
    private int id;

    public int getId() { return id; }

    public void setId(int id) { this.id = id; }
}

<bean id="id" class="com.id ">
  <property name="id" value="123"></property>
</bean>
```

- 静态工厂注入   
静态工厂顾名思义，就是通过调用静态工厂的方法来获取自己需要的对象，为了让 Spring 管理所有对象，我们不能直接通过"工程类.静态方法()"来获取对象，而是依然通过 Spring 注入的形式获取：  

```java
public class DaoFactory { //静态工厂

   public static final FactoryDao getStaticFactoryDaoImpl(){
      return new StaticFacotryDaoImpl();
   }
}

public class SpringAction {

 //注入对象
 private FactoryDao staticFactoryDao;

 //注入对象的 set 方法
 public void setStaticFactoryDao(FactoryDao staticFactoryDao) {
     this.staticFactoryDao = staticFactoryDao;
 }

}

//factory-method="getStaticFactoryDaoImpl"指定调用哪个工厂方法
 <bean name="springAction" class=" SpringAction" >
   <!--使用静态工厂的方法注入对象,对应下面的配置文件-->
   <property name="staticFactoryDao" ref="staticFactoryDao"></property>
 </bean>

 <!--此处获取对象的方式是从工厂类中获取静态方法-->
<bean name="staticFactoryDao" class="DaoFactory"
  factory-method="getStaticFactoryDaoImpl"></bean>
```

- 非静态工厂注入  
非静态工厂，也叫实例工厂，意思是工厂方法不是静态的，所以我们需要首先 new 一个工厂实例，再调用普通的实例方法。  
```java
//非静态工厂
public class DaoFactory {
   public FactoryDao getFactoryDaoImpl(){
     return new FactoryDaoImpl();
   }
 }

public class SpringAction {
  //注入对象
  private FactoryDao factoryDao;

  public void setFactoryDao(FactoryDao factoryDao) {
    this.factoryDao = factoryDao;
  }
}

 <bean name="springAction" class="SpringAction">
   <!--使用非静态工厂的方法注入对象,对应下面的配置文件-->
   <property name="factoryDao" ref="factoryDao"></property>
 </bean>

 <!--此处获取对象的方式是从工厂类中获取实例方法-->
 <bean name="daoFactory" class="com.DaoFactory"></bean>

<bean name="factoryDao" factory-bean="daoFactory" factory-method="getFactoryDaoImpl"></bean>
```

## Spring 有哪些自动装配的方式？
Spring IoC 容器知道所有 Bean 的配置信息，此外，通过 Java 反射机制还可以获知实现类的结构信息，如构造方法的结构、属性等信息。掌握所有 Bean 的这些信息后，Spring IoC 容器就可以按照某种规则对容器中的 Bean 进行自动装配，而无须通过显式的方式进行依赖配置。

Spring 提供的这种方式，可以按照某些规则进行 Bean 的自动装配，<bean>元素提供了一个指定自动装配类型的属性：autowire="<自动装配类型>"

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-034120d9-88c7-490b-af07-7d48f3b6b7bc.png)

### Spring 提供了 4 种自动装配类型
- byName：根据名称进行自动匹配，假设 Boss 又一个名为 car 的属性，如果容器中刚好有一个名为 car 的 bean，Spring 就会自动将其装配给 Boss 的 car 属性
- byType：根据类型进行自动匹配，假设 Boss 有一个 Car 类型的属性，如果容器中刚好有一个 Car 类型的 Bean，Spring 就会自动将其装配给 Boss 这个属性
- constructor：与 byType 类似， 只不过它是针对构造函数注入而言的。如果 Boss 有一个构造函数，构造函数包含一个 Car 类型的入参，如果容器中有一个 Car 类型的 Bean，则 Spring 将自动把这个 Bean 作为 Boss 构造函数的入参；如果容器中没有找到和构造函数入参匹配类型的 Bean，则 Spring 将抛出异常。
- autodetect：根据 Bean 的自省机制决定采用 byType 还是 constructor 进行自动装配，如果 Bean 提供了默认的构造函数，则采用 byType，否则采用 constructor。

## Spring 中的 Bean 的作用域有哪些?
- singleton : 在 Spring 容器仅存在一个 Bean 实例，Bean 以单实例的方式存在，是 Bean 默认的作用域。
- prototype : 每次从容器重调用 Bean 时，都会返回一个新的实例。
以下三个作用域于只在 Web 应用中适用：    
- request : 每一次 HTTP 请求都会产生一个新的 Bean，该 Bean 仅在当前 HTTP Request 内有效。
- session : 同一个 HTTP Session 共享一个 Bean，不同的 HTTP Session 使用不同的 Bean。
- globalSession：同一个全局 Session 共享一个 Bean，只用于基于 Protlet 的 Web 应用，Spring5 中已经不存在了。

## Spring 中的单例 Bean 会存在线程安全问题吗？
Spring Bean 的默认作用域是单例（Singleton），这意味着 Spring 容器中只会存在一个 Bean 实例，并且该实例会被多个线程共享。

如果单例 Bean 是无状态的，也就是没有成员变量，那么这个单例 Bean 是线程安全的。比如 Spring MVC 中的 Controller、Service、Dao 等，基本上都是无状态的。

但如果 Bean 的内部状态是可变的，且没有进行适当的同步处理，就可能出现线程安全问题。  

### 单例 Bean 线程安全问题怎么解决呢？
1. 使用局部变量。局部变量是线程安全的，因为每个线程都有自己的局部变量副本。尽量使用局部变量而不是共享的成员变量。
2. 尽量使用无状态的 Bean，即不在 Bean 中保存任何可变的状态信息。
3. 同步访问。如果 Bean 中确实需要保存可变状态，可以通过 synchronized 关键字或者 Lock 接口来保证线程安全。 或者将 Bean 中的成员变量保存到 ThreadLocal 中，ThreadLocal 可以保证多线程环境下变量的隔离。再或者使用线程安全的工具类，比如说 AtomicInteger、ConcurrentHashMap、CopyOnWriteArrayList 等。
5. 将 Bean 定义为原型作用域（Prototype）。原型作用域的 Bean 每次请求都会创建一个新的实例，因此不存在线程安全问题。

## 说说循环依赖?
A 依赖 B，B 依赖 A，或者 C 依赖 C，就成了循环依赖。
![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-f8fea53f-56fa-4cca-9199-ec7f648da625.png)

当然了，循环依赖只发生在 Singleton 作用域的 Bean 之间，因为如果是 Prototype 作用域的 Bean，Spring 会直接抛出异常。

原因很简单，AB 循环依赖，A 实例化的时候，发现依赖 B，创建 B 实例，创建 B 的时候发现需要 A，创建 A1 实例……无限套娃。。。。

```java
@Component
@Scope("prototype")
public class PrototypeBeanA {
    private final PrototypeBeanB prototypeBeanB;

    @Autowired
    public PrototypeBeanA(PrototypeBeanB prototypeBeanB) {
        this.prototypeBeanB = prototypeBeanB;
    }
}
```

```java
@Component
@Scope("prototype")
public class PrototypeBeanB {
    private final PrototypeBeanA prototypeBeanA;

    @Autowired
    public PrototypeBeanB(PrototypeBeanA prototypeBeanA) {
        this.prototypeBeanA = prototypeBeanA;
    }
}
```
```java
@SpringBootApplication
public class DemoApplication {

    public static void main(String[] args) {
        SpringApplication.run(DemoApplication.class, args);
    }

    @Bean
    CommandLineRunner commandLineRunner(ApplicationContext ctx) {
        return args -> {
            // 尝试获取PrototypeBeanA的实例
            PrototypeBeanA beanA = ctx.getBean(PrototypeBeanA.class);
        };
    }
}
```

## Spring 可以解决哪些情况的循环依赖？
- AB 均采用构造器注入，不支持
- AB 均采用 setter 注入，支持
- AB 均采用属性自动注入，支持
- A 中注入的 B 为 setter 注入，B 中注入的 A 为构造器注入，支持
- B 中注入的 A 为 setter 注入，A 中注入的 B 为构造器注入，不支持   

第四种可以，第五种不可以的原因是 Spring 在创建 Bean 时默认会根据自然排序进行创建，所以 A 会先于 B 进行创建。

简单总结下，当循环依赖的实例都采用 setter 方法注入时，Spring 支持，都采用构造器注入的时候，不支持；构造器注入和 setter 注入同时存在的时候，看天（😂）。

## 那 Spring 怎么解决循环依赖的呢？  
Spring 通过三级缓存（Three-Level Cache）机制来解决循环依赖。

- 一级缓存：用于存放完全初始化好的单例 Bean。
- 二级缓存：用于存放正在创建但未完全初始化的 Bean 实例。
- 三级缓存：用于存放 Bean 工厂对象，用于提前暴露 Bean。  

### 三级缓存解决循环依赖的过程是什么样的？
假如 A、B 两个类发生循环依赖：

A 实例的初始化过程：

①、创建 A 实例，实例化的时候把 A 的对象⼯⼚放⼊三级缓存，表示 A 开始实例化了，虽然这个对象还不完整，但是先曝光出来让大家知道。

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-1a8bdc29-ff43-4ff4-9b61-3eedd9da59b3.png)

②、A 注⼊属性时，发现依赖 B，此时 B 还没有被创建出来，所以去实例化 B。

③、同样，B 注⼊属性时发现依赖 A，它就从缓存里找 A 对象。依次从⼀级到三级缓存查询 A。

发现可以从三级缓存中通过对象⼯⼚拿到 A，虽然 A 不太完善，但是存在，就把 A 放⼊⼆级缓存，同时删除三级缓存中的 A，此时，B 已经实例化并且初始化完成了，把 B 放入⼀级缓存。
![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-bf2507bf-96aa-4b88-a58b-7ec41d11bc70.png)

④、接着 A 继续属性赋值，顺利从⼀级缓存拿到实例化且初始化完成的 B 对象，A 对象创建也完成，删除⼆级缓存中的 A，同时把 A 放⼊⼀级缓存

⑤、最后，⼀级缓存中保存着实例化、初始化都完成的 A、B 对象。
![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-022f7cb9-2c83-4fe9-b252-b02bd0fb2435.png)

## 为什么要三级缓存？⼆级不⾏吗？
不行，主要是为了**⽣成代理对象**。如果是没有代理的情况下，使用二级缓存解决循环依赖也是 OK 的。但是如果存在代理，三级没有问题，二级就不行了。

因为三级缓存中放的是⽣成具体对象的匿名内部类，获取 Object 的时候，它可以⽣成代理对象，也可以返回普通对象。使⽤三级缓存主要是为了保证不管什么时候使⽤的都是⼀个对象。

假设只有⼆级缓存的情况，往⼆级缓存中放的显示⼀个普通的 Bean 对象，Bean 初始化过程中，通过 BeanPostProcessor 去⽣成代理对象之后，覆盖掉⼆级缓存中的普通 Bean 对象，那么可能就导致取到的 Bean 对象不一致了。

## @Autowired 的实现原理？
AutowiredAnnotationBeanPostProcessor

在 Bean 的初始化阶段，会通过 Bean 后置处理器来进行一些前置和后置的处理。

实现@Autowired 的功能，也是通过后置处理器来完成的。这个后置处理器就是 AutowiredAnnotationBeanPostProcessor。

- Spring 在创建 bean 的过程中，最终会调用到 doCreateBean()方法，在 doCreateBean()方法中会调用 populateBean()方法，来为 bean 进行属性填充，完成自动装配等工作。

- 在 populateBean()方法中一共调用了两次后置处理器，第一次是为了判断是否需要属性填充，如果不需要进行属性填充，那么就会直接进行 return，如果需要进行属性填充，那么方法就会继续向下执行，后面会进行第二次后置处理器的调用，这个时候，就会调用到 AutowiredAnnotationBeanPostProcessor 的 postProcessPropertyValues()方法，在该方法中就会进行@Autowired 注解的解析，然后实现自动装配。

## 说说什么是 AOP？
AOP，也就是 Aspect-oriented Programming，译为面向切面编程，是 Spring 中最重要的核心概念之一。

简单点说，AOP 就是把一些业务逻辑中的相同代码抽取到一个独立的模块中，让业务逻辑更加清爽。

### AOP 有哪些核心概念？
- 切面（Aspect）：类是对物体特征的抽象，切面就是对横切关注点的抽象
- 连接点（Join Point）：被拦截到的点，因为 Spring 只支持方法类型的连接点，所以在 Spring 中，连接点指的是被拦截到的方法，实际上连接点还可以是字段或者构造方法
- 切点（Pointcut）：对连接点进行拦截的定位
- 通知（Advice）：指拦截到连接点之后要执行的代码，也可以称作增强
- 目标对象 （Target）：代理的目标对象
- 引介（introduction）：一种特殊的增强，可以动态地为类添加一些属性和方法
- 织入（Weabing）：织入是将增强添加到目标类的具体连接点上的过程。

### 织入有哪几种方式？
①、编译期织入：切面在目标类编译时被织入。

②、类加载期织入：切面在目标类加载到 JVM 时被织入。需要特殊的类加载器，它可以在目标类被引入应用之前增强该目标类的字节码。

③、运行期织入：切面在应用运行的某个时刻被织入。一般情况下，在织入切面时，AOP 容器会为目标对象动态地创建一个代理对象。

Spring AOP 采用运行期织入

### AOP 有哪些环绕方式？
AOP 一般有 5 种环绕方式：
1. 前置通知 (@Before)
2. 返回通知 (@AfterReturning)
3. 异常通知 (@AfterThrowing)
4. 后置通知 (@After)
5. 环绕通知 (@Around)

### Spring AOP 发生在什么时候？
Spring AOP 基于运行时代理机制，这意味着 Spring AOP 是在运行时通过动态代理生成的，而不是在编译时或类加载时生成的。

在 Spring 容器初始化 Bean 的过程中，Spring AOP 会检查 Bean 是否需要应用切面。如果需要，Spring 会为该 Bean 创建一个代理对象，并在代理对象中织入切面逻辑。这一过程发生在 Spring 容器的后处理器（BeanPostProcessor）阶段。

### 简单总结一下 AOP
AOP，也就是面向切面编程，是一种编程范式，旨在提高代码的模块化。比如说可以将日志记录、事务管理等分离出来，来提高代码的可重用性。

AOP 的核心概念包括切面（Aspect）、连接点（Join Point）、通知（Advice）、切点（Pointcut）和织入（Weaving）等。

① 像日志打印、事务管理等都可以抽离为切面，可以声明在类的方法上。像 @Transactional 注解，就是一个典型的 AOP 应用，它就是通过 AOP 来实现事务管理的。我们只需要在方法上添加 @Transactional 注解，Spring 就会在方法执行前后添加事务管理的逻辑。

② Spring AOP 是基于代理的，它默认使用 JDK 动态代理和 CGLIB 代理来实现 AOP。

③ Spring AOP 的织入方式是运行时织入，而 AspectJ 支持编译时织入、类加载时织入。


## 你平时有用到 AOP 吗？
我利用 AOP 打印了接口的入参和出参日志，以及执行时间。

1. 自定义一个注解作为切点
2. 配置 AOP 切面
3. 在使用的地方加上自定义注解
4. 当接口被调用时，就可以看到对应的执行日志。

## 说说 JDK 动态代理和 CGLIB 代理？
Spring 的 AOP 是通过动态代理来实现的，动态代理主要有两种方式：JDK 动态代理和 CGLIB 代理。

①、JDK 动态代理是基于接口的代理，只能代理实现了接口的类。使用 JDK 动态代理时，Spring AOP 会创建一个代理对象，该代理对象实现了目标对象所实现的接口，并在方法调用前后插入横切逻辑。

优点：只需依赖 JDK 自带的 java.lang.reflect.Proxy 类，不需要额外的库；缺点：只能代理接口，不能代理类本身。

②、CGLIB 动态代理是基于继承的代理，可以代理没有实现接口的类。使用 CGLIB 动态代理时，Spring AOP 会生成目标类的子类，并在方法调用前后插入横切逻辑。

优点：可以代理没有实现接口的类，灵活性更高；缺点：需要依赖 CGLIB 库，创建代理对象的开销相对较大。

### 选择 CGLIB 还是 JDK 动态代理？
如果目标对象没有实现任何接口，则只能使用 CGLIB 代理。如果目标对象实现了接口，通常首选 JDK 动态代理。    
 
虽然 CGLIB 在代理类的生成过程中可能消耗更多资源，但在运行时具有较高的性能。对于性能敏感且代理对象创建频率不高的场景，可以考虑使用 CGLIB。

JDK 动态代理是 Java 原生支持的，不需要额外引入库。而 CGLIB 需要将 CGLIB 库作为依赖加入项目中。


## 说说 Spring AOP 和 AspectJ AOP 区别?
Spring AOP 属于运行时增强，主要具有如下特点：

基于动态代理来实现，默认如果使用接口的，用 JDK 提供的动态代理实现，如果是方法则使用 CGLIB 实现

Spring AOP 需要依赖 IoC 容器来管理，并且只能作用于 Spring 容器，使用纯 Java 代码实现

在性能上，由于 Spring AOP 是基于动态代理来实现的，在容器启动时需要生成代理实例，在方法调用上也会增加栈的深度，使得 Spring AOP 的性能不如 AspectJ 的那么好。

Spring AOP 致力于解决企业级开发中最普遍的 AOP(方法织入)。

AspectJ 是一个易用的功能强大的 AOP 框架，属于编译时增强， 可以单独使用，也可以整合到其它框架中，是 AOP 编程的完全解决方案。AspectJ 需要用到单独的编译器 ajc。

AspectJ 属于静态织入，通过修改代码来实现，在实际运行之前就完成了织入，所以说它生成的类是没有额外运行时开销的，一般有如下几个织入的时机：

编译期织入（Compile-time weaving）：如类 A 使用 AspectJ 添加了一个属性，类 B 引用了它，这个场景就需要编译期的时候就进行织入，否则没法编译类 B。

编译后织入（Post-compile weaving）：也就是已经生成了 .class 文件，或已经打成 jar 包了，这种情况我们需要增强处理的话，就要用到编译后织入。

类加载后织入（Load-time weaving）：指的是在加载类的时候进行织入，要实现这个时期的织入，有几种常见的方法

## 说说 AOP 和反射的区别？（补充）
反射：用于检查和操作类的方法和字段，动态调用方法或访问字段。反射是 Java 提供的内置机制，直接操作类对象。  

动态代理：通过生成代理类来拦截方法调用，通常用于 AOP 实现。动态代理使用反射来调用被代理的方法。 



## 事务  
Spring 事务的本质其实就是数据库对事务的支持，没有数据库的事务支持，Spring 是无法提供事务功能的。Spring 只提供统一事务管理接口，具体实现都是由各数据库自己实现，数据库事务的提交和回滚是通过数据库自己的事务机制实现。

## Spring 事务的种类？
在 Spring 中，事务管理可以分为两大类：声明式事务管理和编程式事务管理。 
### 介绍一下编程式事务管理？
编程式事务可以使用 TransactionTemplate 和 PlatformTransactionManager 来实现，需要显式执行事务。允许我们在代码中直接控制事务的边界，通过编程方式明确指定事务的开始、提交和回滚。

### 介绍一下声明式事务管理？
声明式事务是建立在 AOP 之上的。其本质是通过 AOP 功能，对方法前后进行拦截，将事务处理的功能编织到拦截的方法中，也就是在目标方法开始之前启动一个事务，在目标方法执行完之后根据执行情况提交或者回滚事务。

相比较编程式事务，优点是不需要在业务逻辑代码中掺杂事务管理的代码， Spring 推荐通过 @Transactional 注解的方式来实现声明式事务管理，也是日常开发中最常用的。

不足的地方是，声明式事务管理最细粒度只能作用到方法级别，无法像编程式事务那样可以作用到代码块级别。

### 说说两者的区别？
编程式事务管理：需要在代码中显式调用事务管理的 API 来控制事务的边界，比较灵活，但是代码侵入性较强，不够优雅。    
声明式事务管理：这种方式使用 Spring 的 AOP 来声明事务，将事务管理代码从业务代码中分离出来。优点是代码简洁，易于维护。但缺点是不够灵活，只能在预定义的方法上使用事务。


## 说说 Spring 的事务隔离级别？
事务的隔离级别定义了一个事务可能受其他并发事务影响的程度。SQL 标准定义了四个隔离级别，Spring 都支持，并且提供了对应的机制来配置它们，定义在 TransactionDefinition 接口中。

①、ISOLATION_DEFAULT：使用数据库默认的隔离级别（你们爱咋咋滴 😁），MySQL 默认的是可重复读，Oracle 默认的读已提交。

②、ISOLATION_READ_UNCOMMITTED：读未提交，允许事务读取未被其他事务提交的更改。这是隔离级别最低的设置，可能会导致“脏读”问题。

③、ISOLATION_READ_COMMITTED：读已提交，确保事务只能读取已经被其他事务提交的更改。这可以防止“脏读”，但仍然可能发生“不可重复读”和“幻读”问题。

④、ISOLATION_REPEATABLE_READ：可重复读，确保事务可以多次从一个字段中读取相同的值，即在这个事务内，其他事务无法更改这个字段，从而避免了“不可重复读”，但仍可能发生“幻读”问题。

⑤、ISOLATION_SERIALIZABLE：串行化，这是最高的隔离级别，它完全隔离了事务，确保事务序列化执行，以此来避免“脏读”、“不可重复读”和“幻读”问题，但性能影响也最大。


## Spring 的事务传播机制？
事务的传播机制定义了在方法被另一个事务方法调用时，这个方法的事务行为应该如何。

Spring 提供了一系列事务传播行为，这些传播行为定义了事务的边界和事务上下文如何在方法调用链中传播。

REQUIRED：如果当前存在事务，则加入该事务；如果当前没有事务，则创建一个新的事务。Spring 的默认传播行为。   
SUPPORTS：如果当前存在事务，则加入该事务；如果当前没有事务，则以非事务方式执行。   
MANDATORY：如果当前存在事务，则加入该事务；如果当前没有事务，则抛出异常。  
REQUIRES_NEW：总是启动一个新的事务，如果当前存在事务，则将当前事务挂起。  
NOT_SUPPORTED：总是以非事务方式执行，如果当前存在事务，则将当前事务挂起。   
NESTED：如果当前存在事务，则在嵌套事务内执行。如果当前事务不存在，则行为与 REQUIRED 一样。嵌套事务是一个子事务，它依赖于父事务。父事务失败时，会回滚子事务所做的所有操作。但子事务异常不一定会导致父事务的回滚。

事务传播机制是使用 ThreadLocal 实现的，所以，如果调用的方法是在新线程中，事务传播会失效。

### protected 和 private 加事务会生效吗
在 Spring 中，只有通过 Spring 容器的 AOP 代理调用的公开方法（public method）上的@Transactional注解才会生效。

如果在 protected、private 方法上使用@Transactional，这些事务注解将不会生效，原因：Spring 默认使用基于 JDK 的动态代理（当接口存在时）或基于 CGLIB 的代理（当只有类时）来实现事务。这两种代理机制都只能代理公开的方法。


## 声明式事务实现原理了解吗？
Spring 的声明式事务管理是通过 AOP（面向切面编程）和代理机制实现的。

1. 在 Bean 初始化阶段创建代理对象：

Spring 容器在初始化单例 Bean 的时候，会遍历所有的 BeanPostProcessor 实现类，并执行其 postProcessAfterInitialization 方法。

在执行 postProcessAfterInitialization 方法时会遍历容器中所有的切面，查找与当前 Bean 匹配的切面，这里会获取事务的属性切面，也就是 @Transactional 注解及其属性值。

然后根据得到的切面创建一个代理对象，默认使用 JDK 动态代理创建代理，如果目标类是接口，则使用 JDK 动态代理，否则使用 Cglib。

2. 在执行目标方法时进行事务增强操作：

当通过代理对象调用 Bean 方法的时候，会触发对应的 AOP 增强拦截器，声明式事务是一种环绕增强，对应接口为MethodInterceptor，事务增强对该接口的实现为TransactionInterceptor

事务拦截器TransactionInterceptor在invoke方法中，通过调用父类TransactionAspectSupport的invokeWithinTransaction方法进行事务处理，包括开启事务、事务提交、异常回滚等。

## 声明式事务在哪些情况下会失效？
1. @Transactional 应用在非 public 修饰的方法上
2. @Transactional 注解属性 propagation 设置错误
3. @Transactional 注解属性 rollbackFor 设置错误
4. 同一个类中方法调用，导致@Transactional 失效



## Spring MVC 的核心组件？
1. DispatcherServlet：前置控制器，是整个流程控制的核心，控制其他组件的执行，进行统一调度，降低组件之间的耦合性，相当于总指挥。
2. Handler：处理器，完成具体的业务逻辑，相当于 Servlet 或 Action。
3. HandlerMapping：DispatcherServlet 接收到请求之后，通过 HandlerMapping 将不同的请求映射到不同的 Handler。
4. HandlerInterceptor：处理器拦截器，是一个接口，如果需要完成一些拦截处理，可以实现该接口。
5. HandlerExecutionChain：处理器执行链，包括两部分内容：Handler 和 HandlerInterceptor（系统会有一个默认的 HandlerInterceptor，如果需要额外设置拦截，可以添加拦截器）。
6. HandlerAdapter：处理器适配器，Handler 执行业务方法之前，需要进行一系列的操作，包括表单数据的验证、数据类型的转换、将表单数据封装到 JavaBean 等，这些操作都是由 HandlerApater 来完成，开发者只需将注意力集中业务逻辑的处理上，DispatcherServlet 通过 HandlerAdapter 执行不同的 Handler。
7. ModelAndView：装载了模型数据和视图信息，作为 Handler 的处理结果，返回给 DispatcherServlet。
8. ViewResolver：视图解析器，DispatcheServlet 通过它将逻辑视图解析为物理视图，最终将渲染结果响应给客户端。

## Spring MVC 的工作流程？
Spring MVC 是基于模型-视图-控制器的 Web 框架，它的工作流程也主要是围绕着 Model、View、Controller 这三个组件展开的。

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-e29a122b-db07-48b8-8289-7251032e87a1.png)

①、发起请求：客户端通过 HTTP 协议向服务器发起请求。

②、前端控制器：这个请求会先到前端控制器 DispatcherServlet，它是整个流程的入口点，负责接收请求并将其分发给相应的处理器。

③、处理器映射：DispatcherServlet 调用 HandlerMapping 来确定哪个 Controller 应该处理这个请求。通常会根据请求的 URL 来确定。

④、处理器适配器：一旦找到目标 Controller，DispatcherServlet 会使用 HandlerAdapter 来调用 Controller 的处理方法。

⑤、执行处理器：Controller 处理请求，处理完后返回一个 ModelAndView 对象，其中包含模型数据和逻辑视图名。

⑥、视图解析器：DispatcherServlet 接收到 ModelAndView 后，会使用 ViewResolver 来解析视图名称，找到具体的视图页面。

⑦、渲染视图：视图使用模型数据渲染页面，生成最终的页面内容。

⑧、响应结果：DispatcherServlet 将视图结果返回给客户端。

Spring MVC 虽然整体流程复杂，但是实际开发中很简单，大部分的组件不需要我们开发人员创建和管理，真正需要处理的只有 Controller 、View 、Model。

在前后端分离的情况下，步骤 ⑥、⑦、⑧ 会略有不同，后端通常只需要处理数据，并将 JSON 格式的数据返回给前端就可以了，而不是返回完整的视图页面。

## 这个 Handler 是什么东西啊？为什么还需要 HandlerAdapter
Handler 一般就是指 Controller，Controller 是 Spring MVC 的核心组件，负责处理请求，返回响应。

Spring MVC 允许使用多种类型的处理器。不仅仅是标准的@Controller注解的类，还可以是实现了特定接口的其他类（如 HttpRequestHandler 或 SimpleControllerHandlerAdapter 等）。这些处理器可能有不同的方法签名和交互方式。

HandlerAdapter 的主要职责就是调用 Handler 的方法来处理请求，并且适配不同类型的处理器。HandlerAdapter 确保 DispatcherServlet 可以以统一的方式调用不同类型的处理器，无需关心具体的执行细节。

## SpringMVC Restful 风格的接口的流程是什么样的呢？
我们都知道 Restful 接口，响应格式是 json，这就用到了一个常用注解：@ResponseBody
```java
@GetMapping("/user")
@ResponseBody
public User user(){
    return new User(1,"张三");
}
```
加入了这个注解后，整体的流程上和使用 ModelAndView 大体上相同，但是细节上有一些不同

![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/spring-2da963a0-5da9-4b3a-aafd-fd8dbc7e1807.png)

1. 客户端向服务端发送一次请求，这个请求会先到前端控制器 DispatcherServlet

2. DispatcherServlet 接收到请求后会调用 HandlerMapping 处理器映射器。由此得知，该请求该由哪个 Controller 来处理

3. DispatcherServlet 调用 HandlerAdapter 处理器适配器，告诉处理器适配器应该要去执行哪个 Controller

4. Controller 被封装成了 ServletInvocableHandlerMethod，HandlerAdapter 处理器适配器去执行 invokeAndHandle 方法，完成对 Controller 的请求处理

5. HandlerAdapter 执行完对 Controller 的请求，会调用 HandlerMethodReturnValueHandler 去处理返回值，主要的过程：
    - 调用 RequestResponseBodyMethodProcessor，创建 ServletServerHttpResponse（Spring 对原生 ServerHttpResponse 的封装）实例

    - 使用 HttpMessageConverter 的 write 方法，将返回值写入 ServletServerHttpResponse 的 OutputStream 输出流中

    - 在写入的过程中，会使用 JsonGenerator（默认使用 Jackson 框架）对返回值进行 Json 序列化

6. 执行完请求后，返回的 ModealAndView 为 null，ServletServerHttpResponse 里也已经写入了响应，所以不用关心 View 的处理


