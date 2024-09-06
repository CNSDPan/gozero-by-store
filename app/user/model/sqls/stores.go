package sqls

import (
	"time"
)

const (
	STORE_STATUS_1 int8 = 1
	STORE_STATUS_2 int8 = 2
)

var StoreStatusName = map[int8]string{
	STORE_STATUS_1: "启用",
	STORE_STATUS_2: "禁用",
}

// Stores 店铺表
type Stores struct {
	ID        uint32    `gorm:"primaryKey;column:id" json:"-"`
	StoreID   int64     `gorm:"column:store_id" json:"storeId"`     // 店铺ID
	Status    int8      `gorm:"column:status" json:"status"`        // 1-启用、2-禁用
	Name      string    `gorm:"column:name" json:"name"`            // 店铺名
	Avatar    string    `gorm:"column:avatar" json:"avatar"`        // 头像
	Desc      string    `gorm:"column:desc" json:"desc"`            // 店铺描述
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type StoresApi struct {
	StoreID int64  `gorm:"column:store_id" json:"storeId,string"` // 店铺ID
	Status  int8   `gorm:"column:status" json:"status"`           // 1-启用、2-禁用
	Name    string `gorm:"column:name" json:"name"`               // 店铺名
	Avatar  string `gorm:"column:avatar" json:"avatar"`           // 头像
	Desc    string `gorm:"column:desc" json:"desc"`               // 店铺描述
}

// TableName get sql table name.获取数据库表名
func (m *Stores) TableName() string {
	return "stores"
}

// StoresColumns get sql column name.获取数据库列名
var StoresColumns = struct {
	ID        string
	StoreID   string
	Status    string
	Name      string
	Avatar    string
	Desc      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	StoreID:   "store_id",
	Status:    "status",
	Name:      "name",
	Avatar:    "avatar",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

type StoreMgr struct {
	*_BaseMgr
}
