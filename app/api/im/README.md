# 聊天服务

## socket 服务
#### ws握手服务
#### 生成代码命令
    goctl api go --api=./socket.api --dir=./ --style=goZero
#### 运行服务命令
    cd .\app\api\im\socket && go run ws.go

## 聊天RPC 服务
#### 业务处理服务
#### 生成代码命令
    goctl rpc protoc chat.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m
#### 运行服务命令
    cd .\app\chat\rpc && go run chat.go