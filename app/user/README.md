# 用户 服务

## model目录
#### 存放user相关的表模型

## rpc目录
#### RPC生成代码命令
    goctl rpc protoc user.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m
#### 运行服务命令
    cd .\app\user\rpc && go run user.go