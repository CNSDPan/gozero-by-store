package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"store/app/api/im/internal/config"
	"store/app/api/im/internal/middleware"
)

type ServiceContext struct {
	Config            config.Config
	XHeaderMiddleware rest.Middleware
	AuthMiddleware    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
		AuthMiddleware:    middleware.NewAuthMiddleware().Handle,
	}
}
