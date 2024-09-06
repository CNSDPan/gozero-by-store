package svc

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"store/app/user/model/sqls"
	"store/app/user/rpc/internal/config"
	"store/pkg/inital"
	"strconv"
)

type ServiceContext struct {
	Config    config.Config
	Node      *snowflake.Node
	UserModel *sqls.UsersMgr
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 服务uuid节点池
	nodeId, err := strconv.ParseInt(c.ServiceId, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("%s nodeId节点初始化失败 fail:%s", c.ServiceName, err.Error()))
	}
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic(fmt.Sprintf("%s node节点池初始化失败 fail:%s", c.ServiceName, err.Error()))
	}

	return &ServiceContext{
		Config:    c,
		Node:      node,
		UserModel: sqls.NewUserMgr(inital.NewSqlDB(c.Sql, "userModel")),
	}
}
