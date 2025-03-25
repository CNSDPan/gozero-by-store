#!/bin/bash
image=$1
version=$2
if [ ! -n "$image" ]; then
echo "请输入镜像名!"
exit
fi
if [ ! -n "$version" ]; then
echo "请输入版本号!"
exit
fi
echo "开始构建镜像..."
if [ "$image" = "store-client" ]; then
 docker build -f app/api/client/Dockerfile -t $image:$version .
elif [ "$image" = "store-ws" ]; then
 docker build -f app/api/im/Dockerfile -t $image:$version .
elif [ "$image" = "store-api" ]; then
 docker build -f app/rpc/api/Dockerfile -t $image:$version .
elif [ "$image" = "store-im" ]; then
 docker build -f app/rpc/im/Dockerfile -t $image:$version .
elif [ "$image" = "store-store" ]; then
 docker build -f app/rpc/store/Dockerfile -t $image:$version .
elif [ "$image" = "store-user" ]; then
 docker build -f app/rpc/user/Dockerfile -t $image:$version .
fi
echo "请提前登录aliyun镜像仓库..."&&
echo "开始给镜像打标签..."&&
docker tag $image:$version registry.cn-hangzhou.aliyuncs.com/XXX/$image:$version&&
echo "开始推送镜像..."&&
docker push registry.cn-hangzhou.aliyuncs.com/XXX/$image:$version&&
echo "执行结束"


#window
# ./docker-build.sh store-client 0.0.1
# ./docker-build.sh store-ws 0.0.1
# ./docker-build.sh store-api 0.0.1
# ./docker-build.sh store-im 0.0.1
# ./docker-build.sh store-store 0.0.1
# ./docker-build.sh store-user 0.0.1