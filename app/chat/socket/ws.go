package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"store/app/chat/socket/internal/config"
	"store/app/chat/socket/internal/handler"
	"store/app/chat/socket/internal/svc"
	"store/pkg/util"
)

var configFile = flag.String("socket-file", "etc/ws.yaml", "the config file")

func main() {
	flag.Parse()
	var (
		err error
		c   config.Config
	)
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	c.ServiceIp, err = util.GetServerIP()
	if err != nil {
		panic(c.ServiceName + " 获取服务信息 fail:" + err.Error())
	}
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
