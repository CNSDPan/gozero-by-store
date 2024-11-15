package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/chat/rpc/internal/config"
	"store/app/chat/rpc/internal/writer"
	"store/app/store/rpc/store/storebecome"
	"store/pkg/inital"
)

type ServiceContext struct {
	Config     config.Config
	BizConn    *redis.Client
	StoreRpcCl StoreRpc

	WriterHandle *writer.BroadcastWriter
}

type StoreRpc struct {
	Become storebecome.StoreBecome
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizConn := inital.NewBizRedisConn(c.BizRedis, c.Name)
	storeRPC := zrpc.MustNewClient(c.StoreRPC)
	storeRpcCl := StoreRpc{
		Become: storebecome.NewStoreBecome(storeRPC),
	}

	// 写聊天记录协程
	writerHandle := writer.NewBroadcastWriter()
	go writerHandle.WriteChat(storeRpcCl.Become)
	return &ServiceContext{
		Config:       c,
		BizConn:      bizConn,
		StoreRpcCl:   storeRpcCl,
		WriterHandle: writerHandle,
	}
}
