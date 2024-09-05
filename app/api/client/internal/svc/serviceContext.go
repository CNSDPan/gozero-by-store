package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"store/app/api/client/internal/config"
	"store/app/api/client/internal/middleware"
)

type ServiceContext struct {
	Config            config.Config
	XHeaderMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		XHeaderMiddleware: middleware.NewXHeaderMiddleware().Handle,
	}
}
