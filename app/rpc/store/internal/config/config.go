package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"store/pkg/types"
)

type Config struct {
	zrpc.RpcServerConf
	types.ServerInfoConf
	BizRedis     types.BizRedisConf
	CacheRedis   types.CacheRedisConf
	CacheSeconds int64
	Sql          types.SqlConf
	ImRpc        zrpc.RpcClientConf
}
