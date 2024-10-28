package svc

import (
	"github.com/redis/go-redis/v9"
	"store/app/chat/rpc/internal/config"
	"store/pkg/inital"
)

type ServiceContext struct {
	Config  config.Config
	BizConn *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizConn := inital.NewBizRedisConn(c.BizRedis, c.Name)

	return &ServiceContext{
		Config:  c,
		BizConn: bizConn,
	}
}
