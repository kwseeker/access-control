# Casbin

## 基本概念

+ 模型：
  
+ subject(访问实体), object(访问资源), action(访问方法)
  
+ Model：
  + model.conf

    + PERM

      + [request_definition]

      + [policy_definition]

      + [policy_effect]

        对policy生效范围的定义

        | Policy effect定义                                            | 意义             | 示例                                                         |
        | ------------------------------------------------------------ | ---------------- | ------------------------------------------------------------ |
        | some(where (p.eft == allow))                                 | allow-override   | [ACL, RBAC, etc.](https://casbin.org/zh/docs/supported-models#examples) |
        | !some(where (p.eft == deny))                                 | deny-override    | [拒绝改写](https://casbin.org/zh/docs/supported-models#examples) |
        | some(where (p.eft == allow)) && !some(where (p.eft == deny)) | allow-and-deny   | [同意与拒绝](https://casbin.org/zh/docs/supported-models#examples) |
        | priority(p.eft) \|\| deny                                    | priority         | [优先级](https://casbin.org/zh/docs/supported-models#examples) |
        | subjectPriority(p.eft)                                       | 基于角色的优先级 | [主题优先级](https://casbin.org/zh/docs/supported-models#examples) |

        + 优先级模型

      + [matchers]

        + Matcher 函数
        + 自定义Matcher函数
        + 超级管理员 root

  + policy.csv

    + 定义sub, obj, act 策略
    + 定义资源-角色映射关系

  + ACL

  + RBAC
    + Role
      + [role_definition]
    + Pattern 模式匹配
    + 域租户角色（类似组的概念）
    + Casbin RBAC vs RBAC96
  + ABAC

  + RESTful

  + 优先级

  + 存储

    + Model存储

    + Policy存储

      + 策略存储适配器

        支持 File、SQL、ORM、NoSQL、云、KVstore、Stream、String、HTTP 等存储形式。

      + 策略子集加载

+ 执行器 Enforcer

+ 鉴权结果 Effector

+ 监视器

  提供分布式多节点服务，保持多个Casbin执行器实例之间的一致性。

+ 调度器

+ 基于Casbin的服务
  
  + [casbin-server](https://github.com/casbin/casbin-server)



## 工作原理

惊叹Casbin代码量如此小居然能实现这么强大而灵活的功能，代码量很小就直接看源码吧。

TODO:

+ 多Casbin实例+监视器原理



## 使用

和业界其他框架一样都是在拦截器、过滤器、AOP这种地方加授权；

### Gin

### Beego



## 其他

+ 在线配置编辑器 https://casbin.org/editor