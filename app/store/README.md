# 店铺 服务

## model目录
#### 存放store相关的表模型

## rpc目录
#### RPC生成代码命令
    goctl rpc protoc store.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m
#### 运行服务命令
    cd .\app\store\rpc && go run store.go