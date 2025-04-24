# GoZero 构建企业级微服务 + WebSocket 实时通信 + Docker 全流程部署
## 🔥项目亮点
* 全栈教学：代码 + Docker + 文档，覆盖开发到部署全流程
* docker容器部署：docker-compose
* 实战场景：集成 WebSocket 实现实时消息推送（含身份鉴权）
* 最佳实践：模块化拆分、配置管理、性能优化技巧
* 断开socket 基于 context + WaitGroup 的协程优雅退出机制
## 目录架构
```
.
├── app/
│   ├── api/              # API
│       ├── client/       # 会员服务
│       ├── im/           # 即使通信服务（websocket）
│       └── store/        # 
│   └── rpc/              # GRPC
│       ├── api/          #  
│       ├── im/           # 消息广播服务
│       ├── store/        # 
│       └── user/         # 
├── db/                   # gorm & gentool
├── doc/                  # 表结构
├── docker/               # docker-compose.yaml,容器搭建
├── pkg/                  # 工具
└── docker-build-dev.sh/  # 各服务的镜像打包
```

# 演示地址
http://8.135.237.23:8081/login

# 演示

![people.png](https://raw.githubusercontent.com/CNSDPan/store/master/static/images/people.png)

![store.png](https://raw.githubusercontent.com/CNSDPan/store/master/static/images/store.png)

![message.png](https://raw.githubusercontent.com/CNSDPan/store/master/static/images/message.png)