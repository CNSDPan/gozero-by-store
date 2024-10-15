package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/api/client/internal/config"
	"store/app/api/client/internal/middleware"
	"store/app/api/rpc/api/apistore"
	"store/app/api/rpc/api/apitoken"
	"store/app/api/rpc/api/apiuser"
	"store/app/store/rpc/store/storebecome"
	"store/app/user/rpc/user/userRegister"
)

type ServiceContext struct {
	Config            config.Config
	XHeaderMiddleware rest.Middleware
	AuthMiddleware    rest.Middleware
	ApiRpcCl          ApiRpc
	UserRpcCl         UserRpc
	StoreRpcCl        StoreRpc
}
type ApiRpc struct {
	Store apistore.ApiStore
	User  apiuser.ApiUser
	Auth  apitoken.ApiToken
}
type UserRpc struct {
	Register userregister.UserRegister
}
type StoreRpc struct {
	Become storebecome.StoreBecome
}

func NewServiceContext(c config.Config) *ServiceContext {
	apiRPC := zrpc.MustNewClient(c.ApiRPC)
	userRPC := zrpc.MustNewClient(c.UserRPC)
	storeRPC := zrpc.MustNewClient(c.StoreRPC)

	ApiRpcCl := ApiRpc{
		Store: apistore.NewApiStore(apiRPC),
		User:  apiuser.NewApiUser(apiRPC),
		Auth:  apitoken.NewApiToken(apiRPC),
	}
	return &ServiceContext{
		Config:            c,
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
		AuthMiddleware:    middleware.NewAuthMiddleware(ApiRpcCl.User, ApiRpcCl.Auth).Handle,
		ApiRpcCl:          ApiRpcCl,
		UserRpcCl: UserRpc{
			Register: userregister.NewUserRegister(userRPC),
		},
		StoreRpcCl: StoreRpc{
			Become: storebecome.NewStoreBecome(storeRPC),
		},
	}
}
