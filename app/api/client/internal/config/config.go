package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"store/pkg/types"
)

type Config struct {
	rest.RestConf
	types.ServerInfoConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	ApiRPC   zrpc.RpcClientConf
	StoreRPC zrpc.RpcClientConf
	UserRPC  zrpc.RpcClientConf
}
