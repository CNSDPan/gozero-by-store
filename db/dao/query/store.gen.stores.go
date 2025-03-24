package query

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"store/db/dao/model"
	"time"
)

type _StoresMgr struct {
	*_BaseMgr
}

// StoresMgr open func
func StoresMgr(db *gorm.DB) *_StoresMgr {
	if db == nil {
		panic(fmt.Errorf("StoresMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoresMgr{_BaseMgr: &_BaseMgr{DB: db.Table("stores"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_StoresMgr) Debug() *_StoresMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StoresMgr) GetTableName() string {
	return "stores"
}

// Reset 重置gorm会话
func (obj *_StoresMgr) Reset() *_StoresMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_StoresMgr) Get() (result model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StoresMgr) Gets() (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_StoresMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_StoresMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStoreID store_id获取 店铺ID
func (obj *_StoresMgr) WithStoreID(storeID int64) Option {
	return optionFunc(func(o *options) { o.query["store_id"] = storeID })
}

// WithStatus status获取 1-启用、2-禁用
func (obj *_StoresMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithName name获取 昵称
func (obj *_StoresMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAvatar avatar获取 头像
func (obj *_StoresMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithDesc desc获取 店铺描述
func (obj *_StoresMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_StoresMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_StoresMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_StoresMgr) GetByOption(opts ...Option) (result model.Stores, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StoresMgr) GetByOptions(opts ...Option) (results []*model.Stores, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_StoresMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.Stores, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_StoresMgr) GetFromID(id uint32) (result model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_StoresMgr) GetBatchFromID(ids []uint32) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStoreID 通过store_id获取内容 店铺ID
func (obj *_StoresMgr) GetFromStoreID(storeID int64) (result model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`store_id` = ?", storeID).First(&result).Error

	return
}

// GetBatchFromStoreID 批量查找 店铺ID
func (obj *_StoresMgr) GetBatchFromStoreID(storeIDs []int64) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`store_id` IN (?)", storeIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 1-启用、2-禁用
func (obj *_StoresMgr) GetFromStatus(status bool) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 1-启用、2-禁用
func (obj *_StoresMgr) GetBatchFromStatus(statuss []bool) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 昵称
func (obj *_StoresMgr) GetFromName(name string) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 昵称
func (obj *_StoresMgr) GetBatchFromName(names []string) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_StoresMgr) GetFromAvatar(avatar string) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_StoresMgr) GetBatchFromAvatar(avatars []string) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 店铺描述
func (obj *_StoresMgr) GetFromDesc(desc string) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 店铺描述
func (obj *_StoresMgr) GetBatchFromDesc(descs []string) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_StoresMgr) GetFromCreatedAt(createdAt time.Time) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_StoresMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_StoresMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_StoresMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_StoresMgr) FetchByPrimaryKey(id uint32) (result model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUnxSid primary or index 获取唯一内容
func (obj *_StoresMgr) FetchUniqueByUnxSid(storeID int64) (result model.Stores, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`store_id` = ?", storeID).First(&result).Error

	return
}

func (obj *_StoresMgr) StoresTableJoinName() string {
	return "stores as s"
}

// CreatStore
// @Desc：创建店铺
// @param：stores
// @return：err
func (obj *_StoresMgr) CreatStore(stores model.Stores) (err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Create(&stores).Error
	return
}

// GetFromNameApi 通过name获取内容
func (obj *_StoresMgr) GetFromStoreIDApi(storeId int64) (result model.StoresInfoApi, err error) {
	query := obj.WithContext(obj.ctx).Table(obj.StoresTableJoinName())
	query.Joins("left join store_users as su on s.store_id = su.store_id")
	query.Where("`s`.`store_id` = ?", storeId)
	query.Select("s.*,su.store_user_id,su.user_id,(SELECT COUNT(*) FROM store_member WHERE store_id = ?) as contacts", storeId)
	err = query.Find(&result).Error
	return
}

// GetFromNameApi 通过name获取内容
func (obj *_StoresMgr) GetFromNameApi(name string) (result model.StoresApi) {
	obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where("`name` = ?", name).Find(&result)
	return
}

// 分页条件
func (obj *_StoresMgr) WithStatusEnable() Option {
	return optionFunc(func(o *options) { o.query["status"] = model.STORE_STATUS_1 })
}

// SelectPageApi
// @Desc：分页获取基本店铺信息
// @param：page
// @param：opts
// @return：resultPage
// @return：err
func (obj *_StoresMgr) SelectPageApi(page IPage, where interface{}, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.StoresApi, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.Stores{}).Where(options.query).Where(where)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}
