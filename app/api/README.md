# API   服务
## client目录
   ### 顾客接口
   #### 生成代码命令
        goctl api go --api=./client.api --dir=./ --style=goZero
   #### 运行服务命令
        cd .\app\api\client && go run client.go
   #### 生成Dockerfile
        goctl docker --go=client.go --tz=Asia/Shanghai --version=1.21.9
   #### 编译镜像
        docker build -f app/api/client/Dockerfile -t store-api-client:0.0.1 .
## store目录
   ### 店主接口
