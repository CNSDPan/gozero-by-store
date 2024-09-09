package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/api/client/internal/config"
	"store/app/api/client/internal/middleware"
	"store/app/api/rpc/api/apistore"
	"store/app/api/rpc/api/apiuser"
	"store/app/user/rpc/user/userRegister"
)

type ServiceContext struct {
	Config            config.Config
	XHeaderMiddleware rest.Middleware
	ApiRpcCl          ApiRpc
	UserRpcCl         UserRpc
}
type ApiRpc struct {
	Store apistore.ApiStore
	User  apiuser.ApiUser
}
type UserRpc struct {
	Register userregister.UserRegister
}

func NewServiceContext(c config.Config) *ServiceContext {
	apiRPC := zrpc.MustNewClient(c.ApiRPC)
	userRPC := zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:            c,
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
		ApiRpcCl: ApiRpc{
			Store: apistore.NewApiStore(apiRPC),
			User:  apiuser.NewApiUser(apiRPC),
		},
		UserRpcCl: UserRpc{
			Register: userregister.NewUserRegister(userRPC),
		},
	}
}
