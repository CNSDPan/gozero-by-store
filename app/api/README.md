# API   服务
## client目录
   #### 顾客接口
   #### 生成代码命令
      goctl api go --api=./client.api --dir=./ --style=goZero
   #### 运行服务命令
      cd .\app\api\client && go run client.go
      
## store目录
   #### 店主接口

## rpc目录
   #### 业务处理服务
   #### 生成代码命令
      goctl rpc protoc api.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m
   #### 运行服务命令
      cd .\app\api\rpc && go run api.go