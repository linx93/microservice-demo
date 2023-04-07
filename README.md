# microservice-demo
### 目录

```
├─microservice-demo
│  ├─common                     公共代码
│  │  ├─ApiResult               关于api返回的封装
│  │  └─util                    工具
│  └─service                    服务
│      ├─order                  订单服务
│      │  ├─api                 对外提供api
│      │  │  ├─etc              yaml
│      │  │  └─internal         内部代码
│      │  │      ├─config       配置
│      │  │      ├─handler      处理器+路由
│      │  │      ├─logic        业务代码，核心开发业务主要就在这里编写
│      │  │      ├─svc          服务上下文
│      │  │      └─types        api需要用的model
│      │  ├─cronjob             定时任务
│      │  ├─rmq                 消息处理系统：mq和dq，处理一些高并发和延时消息业务
│      │  ├─rpc                 rpc服务，给其他子系统提供基础数据访问
│      │  └─script              脚本，处理一些临时运营需求，临时数据修复
│      └─user                   用户服务，其他目录和order服务的一致
│          └─rpc
│              ├─etc
│              ├─internal
│              │  ├─config
│              │  ├─logic
│              │  ├─server
│              │  └─svc
│              ├─types
│              │  └─user
│              └─userclient
```