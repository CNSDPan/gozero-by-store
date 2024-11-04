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
        每次重新生成代码后，需要编辑 api.pb.go 里的结构体，因为json的tag需要接收string
        type StoreItem struct {
            StoreId  int64  `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty,string"`
        }
        type StoreItem UserItem{
            UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty,string"`
        } 
   #### 运行服务命令
      cd .\app\api\rpc && go run api.go