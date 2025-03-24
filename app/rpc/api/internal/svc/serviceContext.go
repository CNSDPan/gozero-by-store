package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"store/app/rpc/api/internal/config"
	"store/db/dao/query"
	"store/pkg/cache"
	"store/pkg/inital"
	"strconv"
)

type ServiceContext struct {
	Config       config.Config
	Node         *snowflake.Node
	CacheConn    *redis.Client
	CacheConnApi *cache.CacheItem
	BizConn      *redis.Client
	//UserModel    *sqls.UsersMgr
	//StoreModel   StoreModel
	Mysql      *gorm.DB
	MysqlQuery *query.Query
}

//type StoreModel struct {
//	StoresMgr       *model.Store
//	StoreUsersMgr   *sqlsStore.StoreUsersMgr
//	StoresMemberMgr *sqlsStore.StoresMemberMgr
//	ChatLogMgr      *sqlsStore.ChatLogMgr
//}

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

	cacheConn := inital.NewCacheRedisConn(c.CacheRedis, c.Name)
	bizConn := inital.NewBizRedisConn(c.BizRedis, c.Name)

	mysqlDb := inital.NewSqlDB(c.Sql, "model")
	return &ServiceContext{
		Config:       c,
		Node:         node,
		CacheConn:    cacheConn,
		CacheConnApi: cache.NewCache(cache.NewCacheUser(context.Background(), cacheConn), cache.NewCacheStore(context.Background(), cacheConn)),
		BizConn:      bizConn,
		//UserModel:    sqls.NewUserMgr(inital.NewSqlDB(c.Sql, "source.userModel")),
		//StoreModel: StoreModel{
		//	StoresMgr:       model.NewStore(),
		//	StoreUsersMgr:   sqlsStore.NewStoreUsersMgr(inital.NewSqlDB(c.Sql, "source.storeUsersModel")),
		//	StoresMemberMgr: sqlsStore.NewStoresMemberMgr(inital.NewSqlDB(c.Sql, "source.storesMemberModel")),
		//	ChatLogMgr:      sqlsStore.NewChatLogMgr(inital.NewSqlDB(c.Sql, "source.chatLogModel")),
		//},
		Mysql:      mysqlDb,
		MysqlQuery: query.Use(mysqlDb),
	}
}
