package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/rpc/im/internal/config"
	"store/app/rpc/store/client/storebecome"
	"store/pkg/inital"
)

type ServiceContext struct {
	Config       config.Config
	BizConn      *redis.Client
	StoreRpcCl   StoreRpc
	WriterHandle *BroadcastWriter
}

type StoreRpc struct {
	Become storebecome.StoreBecome
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizConn := inital.NewBizRedisConn(c.BizRedis, c.Name)
	storeRPC := zrpc.MustNewClient(c.StoreRPC)

	sCtx := &ServiceContext{
		Config:  c,
		BizConn: bizConn,
	}
	go func() {
		sCtx.StoreRpcCl = StoreRpc{
			Become: storebecome.NewStoreBecome(storeRPC),
		}
		// 写聊天记录协程
		writerHandle := NewBroadcastWriter()
		go writerHandle.WriteChat(sCtx.StoreRpcCl.Become)
	}()

	return sCtx
}
