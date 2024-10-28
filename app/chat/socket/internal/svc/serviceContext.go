package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/api/rpc/api/apistore"
	"store/app/api/rpc/api/apitoken"
	"store/app/api/rpc/api/apiuser"
	"store/app/chat/rpc/chat/socket"
	"store/app/chat/socket/internal/config"
	"store/app/chat/socket/internal/middleware"
	"store/app/chat/socket/server"
	"store/pkg/inital"
	"strconv"
)

type ServiceContext struct {
	Config            config.Config
	AuthMiddleware    rest.Middleware
	XHeaderMiddleware rest.Middleware
	Node              *snowflake.Node
	ApiRpcCl          ApiRpc
	SocketRpcCl       SocketRpc
	WsServer          *server.Server
	BizConn           *redis.Client
}

// ApiRpc API的RPC服务
type ApiRpc struct {
	Store apistore.ApiStore
	User  apiuser.ApiUser
	Auth  apitoken.ApiToken
}
type SocketRpc struct {
	Socket socket.Socket
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 服务uuid节点池
	nodeId, err := strconv.ParseInt(c.ServiceId, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("%s nodeId节点初始化失败 fail:%s", c.ServiceName, err.Error()))
	}
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic(fmt.Sprintf("%s node节点池初始化失败 fail:%s", c.ServiceName, err.Error()))
	}

	apiRPC := zrpc.MustNewClient(c.ApiRPC)
	socketRPC := zrpc.MustNewClient(c.SocketRPC)
	ApiRpcCl := ApiRpc{
		Store: apistore.NewApiStore(apiRPC),
		User:  apiuser.NewApiUser(apiRPC),
		Auth:  apitoken.NewApiToken(apiRPC),
	}
	SocketRpcCl := SocketRpc{
		Socket: socket.NewSocket(socketRPC),
	}

	wsServer := server.NewServer(c.ServiceId, c.ServiceName, c.ServiceIp, c.SocketOptions, node, logx.WithContext(context.Background()))
	wsServer.SetSocketRpc(SocketRpcCl.Socket)

	bizConn := inital.NewBizRedisConn(c.BizRedis, c.Name)
	err = server.NewRedisMq(bizConn)
	if err != nil {
		panic(fmt.Sprintf("%s MQ消息订阅出事异常 fail:%s", c.ServiceName, err.Error()))
	}
	return &ServiceContext{
		Config:            c,
		Node:              node,
		AuthMiddleware:    middleware.NewAuthMiddleware(ApiRpcCl.Auth).Handle,
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
		ApiRpcCl:          ApiRpcCl,
		SocketRpcCl:       SocketRpcCl,
		WsServer:          wsServer,
		BizConn:           bizConn,
	}
}
