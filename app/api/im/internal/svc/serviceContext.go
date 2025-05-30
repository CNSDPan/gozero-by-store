package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/api/im/internal/config"
	"store/app/api/im/internal/middleware"
	"store/app/api/im/server"
	"store/app/rpc/api/client/apistore"
	"store/app/rpc/api/client/apitoken"
	"store/app/rpc/api/client/apiuser"
	"store/app/rpc/im/client/socket"
	"store/pkg/inital"
	"strconv"
)

type ServiceContext struct {
	Config            config.Config
	XHeaderMiddleware rest.Middleware
	AuthMiddleware    rest.Middleware
	Node              *snowflake.Node
	ApiRpcCl          ApiRpc
	SocketRpcCl       SocketRpc
	WsServer          *server.Server
	BizConn           *redis.Client
}
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
	socketRPC := zrpc.MustNewClient(c.ImRPC)
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
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
		AuthMiddleware:    middleware.NewAuthMiddleware(ApiRpcCl.Auth).Handle,
		ApiRpcCl:          ApiRpcCl,
		SocketRpcCl:       SocketRpcCl,
		WsServer:          wsServer,
		BizConn:           bizConn,
	}
}
