package inital

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"store/pkg/types"
	"time"
)

var DBMap = make(map[string]*gorm.DB)

// NewSqlDB
// @Desc：公共初始化GORM连接，若各模块设置不一样，通过传参配置
// @param：sqlConf
// @param：model		模块;例如：user、store
// @return：*gorm.DB
func NewSqlDB(sqlConf types.SqlConf, model string) *gorm.DB {
	if dbConn, ok := DBMap[model]; ok {
		return dbConn
	}
	if sqlConf.Separation == 1 {
		dbConn, err := gorm.Open(mysql.New(mysql.Config{
			DSN: sqlConf.MasterSlave.MasterAddr,
		}), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("gorm 初始化master失败:%s", err.Error()))
		}
		// replicate 从库,只拥有读权限
		replicates := []gorm.Dialector{}
		for _, dsn := range sqlConf.MasterSlave.SlaveAddr.Connect {
			sConf := mysql.Config{
				DSN: dsn,
			}
			replicates = append(replicates, mysql.New(sConf))
		}
		err = dbConn.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.New(mysql.Config{DSN: sqlConf.MasterSlave.MasterAddr})},
			Replicas: replicates,
			Policy:   dbresolver.RandomPolicy{},
		}).SetMaxIdleConns(10).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))
		if err != nil {
			panic(fmt.Sprintf("gorm 初始化master、replicate失败:%s", err.Error()))
		}
		DBMap[model] = dbConn
		return dbConn
	} else {
		dbConn, err := gorm.Open(mysql.New(mysql.Config{
			DSN: sqlConf.SqlSource.Addr,
		}), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("gorm 初始化数据库实例失败:%s", err.Error()))
		}
		db, e := dbConn.DB()
		if e != nil {
			panic(fmt.Sprintf("gorm 获取数据库实例失败:%s", e.Error()))
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		db.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		db.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		db.SetConnMaxLifetime(time.Hour)
		DBMap[model] = dbConn
		return dbConn
	}
}
