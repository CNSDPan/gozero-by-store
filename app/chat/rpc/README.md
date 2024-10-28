# 生成rpc代码命令(多服务)
goctl rpc protoc chat.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=goZero -m