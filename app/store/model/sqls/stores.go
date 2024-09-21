package sqls

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const (
	STORE_STATUS_1 int8 = 1
	STORE_STATUS_2 int8 = 2
)

var StoresStatusName = map[int8]string{
	STORE_STATUS_1: "启用",
	STORE_STATUS_2: "禁用",
}

// Stores 店铺表
type Stores struct {
	Id        uint32    `gorm:"primaryKey;column:id" json:"-"`
	StoreId   int64     `gorm:"column:store_id" json:"storeId"`        // 店铺ID
	Status    int8      `gorm:"column:status;default:1" json:"status"` // 1-启用、2-禁用
	Name      string    `gorm:"column:name" json:"name"`               // 昵称
	Avatar    string    `gorm:"column:avatar" json:"avatar"`           // 头像
	Desc      string    `gorm:"column:desc" json:"desc"`               // 店铺描述
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`    // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`    // 更新时间
}
type StoresApi struct {
	StoreId int64  `gorm:"column:store_id" json:"storeId,string"` // 店铺ID
	Status  int8   `gorm:"column:status" json:"status"`           // 1-启用、2-禁用
	Name    string `gorm:"column:name" json:"name"`               // 昵称
	Avatar  string `gorm:"column:avatar" json:"avatar"`           // 头像
	Desc    string `gorm:"column:desc" json:"desc"`               // 店铺描述
	// 虚拟字段
	Contacts    int64 `gorm:"-;default:0" json:"contacts"`           // 虚拟字段,店铺会员数
	StoreUserId int64 `gorm:"-;default:0" json:"storeUserId,string"` // 虚拟字段,店铺会员数
	UserId      int64 `gorm:"-;default:0" json:"userId,string"`      // 虚拟字段,店铺会员数
}

type StoresInfoApi struct {
	StoreId int64  `gorm:"column:store_id" json:"storeId,string"` // 店铺ID
	Status  int8   `gorm:"column:status" json:"status"`           // 1-启用、2-禁用
	Name    string `gorm:"column:name" json:"name"`               // 昵称
	Avatar  string `gorm:"column:avatar" json:"avatar"`           // 头像
	Desc    string `gorm:"column:desc" json:"desc"`               // 店铺描述
	// 虚拟字段
	Contacts    int64 `gorm:"contacts;default:0" json:"contacts"`                // 虚拟字段,店铺会员数
	StoreUserId int64 `gorm:"store_user_id;default:0" json:"storeUserId,string"` // 虚拟字段,店铺会员数
	UserId      int64 `gorm:"user_id;default:0" json:"userId,string"`            // 虚拟字段,店铺会员数
}

// StoresColumns get sql column name.获取数据库列名
var StoresColumns = struct {
	Id        string
	StoreId   string
	Status    string
	Name      string
	Avatar    string
	Desc      string
	CreatedAt string
	UpdatedAt string
}{
	Id:        "id",
	StoreId:   "store_id",
	Status:    "status",
	Name:      "name",
	Avatar:    "avatar",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

type StoresMgr struct {
	*_BaseMgr
}

func StoresTableName() string {
	return "stores"
}
func StoresTableJoinName() string {
	return "stores as s"
}

func NewStoresMgr(db *gorm.DB) *StoresMgr {
	ctx, cancel := context.WithCancel(context.Background())
	return &StoresMgr{_BaseMgr: &_BaseMgr{DB: db.Table(StoresTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *StoresMgr) Reset() *StoresMgr {
	obj.New()
	return obj
}

// CreatStore
// @Desc：创建店铺
// @param：stores
// @return：err
func (obj *StoresMgr) CreatStore(stores Stores) (err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Stores{}).Create(&stores).Error
	return
}

// SelectPageApi
// @Desc：分页获取基本店铺信息
// @param：page
// @param：opts
// @return：resultPage
// @return：err
func (obj *StoresMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]StoresApi, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Stores{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

// GetFromStoreIDApi 通过store_id获取内容 店铺IID
func (obj *StoresMgr) GetFromStoreIDApi(storeId int64) (result StoresInfoApi, err error) {
	query := obj.DB.WithContext(obj.ctx).Table(StoresTableJoinName())
	query.Joins("left join store_users as su on s.store_id = su.store_id")
	query.Where("`s`.`store_id` = ?", storeId)
	query.Select("s.*,su.store_user_id,su.user_id,(SELECT COUNT(*) FROM store_member WHERE store_id = ?) as contacts", storeId)
	err = query.Find(&result).Error
	return
}

// GetFromNameApi 通过name获取内容
func (obj *StoresMgr) GetFromNameApi(name string) (result StoresApi) {
	obj.DB.WithContext(obj.ctx).Model(Stores{}).Where("`name` = ?", name).Find(&result)
	return
}

// 分页条件
func (obj *StoresMgr) WithStatusEnable() Option {
	return optionFunc(func(o *options) { o.query["status"] = STORE_STATUS_1 })
}
