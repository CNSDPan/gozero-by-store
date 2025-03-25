package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/pkg/types"
)

type Config struct {
	rest.RestConf
	types.ServerInfoConf
	ServiceIp string `json:",omitempty"`
	Auth      struct {
		AccessSecret string
		AccessExpire int64
	}
	BizRedis      types.BizRedisConf
	SocketOptions types.SocketOptionsConf
	ImRPC         zrpc.RpcClientConf
	ApiRPC        zrpc.RpcClientConf
}
