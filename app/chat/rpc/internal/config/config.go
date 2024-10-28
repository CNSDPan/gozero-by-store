package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"store/pkg/types"
)

type Config struct {
	zrpc.RpcServerConf
	types.ServerInfoConf
	BizRedis types.BizRedisConf
	Sql      types.SqlConf
}
