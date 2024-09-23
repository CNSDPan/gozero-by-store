package mysqltest

import (
	sqlsStore "store/app/store/model/sqls"
	"store/pkg/inital"
	"store/pkg/types"
)

type StoreModel struct {
	StoresMgr       *sqlsStore.StoresMgr
	StoreUsersMgr   *sqlsStore.StoreUsersMgr
	StoresMemberMgr *sqlsStore.StoresMemberMgr
}

var storeModel StoreModel

func init() {
	sqlConf := types.SqlConf{
		Separation:  2,
		MasterSlave: types.MasterSlaveConf{},
		SqlSource: types.SqlSourceConf{
			Addr: "root:root@tcp(192.168.33.10:3307)/store2?loc=Local&parseTime=True&charset=utf8mb4",
		},
	}
	storeModel = StoreModel{
		StoresMgr:       sqlsStore.NewStoresMgr(inital.NewSqlDB(sqlConf, "source.storeModel")),
		StoreUsersMgr:   sqlsStore.NewStoreUsersMgr(inital.NewSqlDB(sqlConf, "source.storeUsersModel")),
		StoresMemberMgr: sqlsStore.NewStoresMemberMgr(inital.NewSqlDB(sqlConf, "source.storesMemberModel")),
	}
}
