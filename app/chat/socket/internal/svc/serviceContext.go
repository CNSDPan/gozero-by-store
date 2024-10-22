package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/api/rpc/api/apistore"
	"store/app/api/rpc/api/apitoken"
	"store/app/api/rpc/api/apiuser"
	"store/app/chat/socket/internal/config"
	"store/app/chat/socket/internal/middleware"
	"store/app/chat/socket/server"
	"strconv"
)

type ServiceContext struct {
	Config            config.Config
	AuthMiddleware    rest.Middleware
	XHeaderMiddleware rest.Middleware
	Node              *snowflake.Node
	ApiRpcCl          ApiRpc
	WsServer          *server.Server
}

// ApiRpc API的RPC服务
type ApiRpc struct {
	Store apistore.ApiStore
	User  apiuser.ApiUser
	Auth  apitoken.ApiToken
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
	ApiRpcCl := ApiRpc{
		Store: apistore.NewApiStore(apiRPC),
		User:  apiuser.NewApiUser(apiRPC),
		Auth:  apitoken.NewApiToken(apiRPC),
	}
	return &ServiceContext{
		Config:            c,
		Node:              node,
		AuthMiddleware:    middleware.NewAuthMiddleware(ApiRpcCl.Auth).Handle,
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
		ApiRpcCl:          ApiRpcCl,
		WsServer:          server.NewServer(c.ServiceId, c.ServiceName, c.ServiceIp, c.SocketOptions, node, logx.WithContext(context.Background())),
	}
}
