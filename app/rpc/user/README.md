# RPC-USER   服务
## rpc目录
#### RPC生成代码命令
    goctl rpc protoc user.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m
#### 运行服务命令
    cd .\app\rpc\user && go run user.go
#### 生成Dockerfile
    goctl docker --go=user.go --tz=Asia/Shanghai --version=1.21.9
#### 编译镜像
    docker build -f app/rpc/user/Dockerfile -t store-user:0.0.1 .