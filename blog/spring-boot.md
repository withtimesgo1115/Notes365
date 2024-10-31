## Spring Boot自动配置原理？
结论：springboot 所有自动配置都是在启动的时候扫描并加载：`spring.factories`所有的自动配置类都在这里面，但是不一定生效，要判断条件是否成立，只要导入了对应的starter, 就有对应的启动器了，有了启动器，我们自动装配就会生效，然后就配置成功了。

- @SpringBootApplication
    - @SpringBootConfiguration
        - @Configuration
            - @Component
    - @ComponentScan 扫描当前主启动类同级的包
    - @EnableConfiguration启动配置
        - @AutoConfigurationPackage自动导入包
            - @Import(AutoConfigurationPackages.Register.class)自动注册包
        - @Import(AutoConfigurationImportSelector.class)自动导入包的核心
            - AutoConfigurationImportSelector
                - 获得自动配置实体
                - 获取候选配置
                    - 标注enableAutoConfiguration的类
                - 获取所有的加载配置
                - loadSpringFactories
                    - 项目资源 （META-INF/spring.factories）
                    - 系统资源



## 