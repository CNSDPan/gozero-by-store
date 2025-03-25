# RPC-API   服务
## rpc目录
#### RPC生成代码命令
    goctl rpc protoc api.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m
    每次重新生成代码后，需要编辑 api.pb.go 里的结构体，因为json的tag需要接收string
    type StoreItem struct {
        StoreId  int64  `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty,string"`
    }
    type UserItem struct{
        UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty,string"`
    } 
#### 运行服务命令
    cd .\app\rpc\api && go run api.go
#### 生成Dockerfile
    goctl docker --go=api.go --tz=Asia/Shanghai --version=1.21.9
#### 编译镜像
    docker build -f app/rpc/api/Dockerfile -t store-api:0.0.1 .