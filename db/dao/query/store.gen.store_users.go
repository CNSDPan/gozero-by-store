package query

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"store/db/dao/model"
	"time"
)

type _StoreUsersMgr struct {
	*_BaseMgr
}

// StoreUsersMgr open func
func StoreUsersMgr(db *gorm.DB) *_StoreUsersMgr {
	if db == nil {
		panic(fmt.Errorf("StoreUsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreUsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("store_users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_StoreUsersMgr) Debug() *_StoreUsersMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StoreUsersMgr) GetTableName() string {
	return "store_users"
}

// Reset 重置gorm会话
func (obj *_StoreUsersMgr) Reset() *_StoreUsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_StoreUsersMgr) Get() (result model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StoreUsersMgr) Gets() (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_StoreUsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_StoreUsersMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStoreUserID store_user_id获取 店主ID
func (obj *_StoreUsersMgr) WithStoreUserID(storeUserID int64) Option {
	return optionFunc(func(o *options) { o.query["store_user_id"] = storeUserID })
}

// WithStoreID store_id获取 店铺ID
func (obj *_StoreUsersMgr) WithStoreID(storeID int64) Option {
	return optionFunc(func(o *options) { o.query["store_id"] = storeID })
}

// WithUserID user_id获取 用户ID
func (obj *_StoreUsersMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_StoreUsersMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_StoreUsersMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_StoreUsersMgr) GetByOption(opts ...Option) (result model.StoreUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StoreUsersMgr) GetByOptions(opts ...Option) (results []*model.StoreUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_StoreUsersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.StoreUsers, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where(options.query)
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
func (obj *_StoreUsersMgr) GetFromID(id uint32) (result model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_StoreUsersMgr) GetBatchFromID(ids []uint32) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStoreUserID 通过store_user_id获取内容 店主ID
func (obj *_StoreUsersMgr) GetFromStoreUserID(storeUserID int64) (result model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_user_id` = ?", storeUserID).First(&result).Error

	return
}

// GetBatchFromStoreUserID 批量查找 店主ID
func (obj *_StoreUsersMgr) GetBatchFromStoreUserID(storeUserIDs []int64) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_user_id` IN (?)", storeUserIDs).Find(&results).Error

	return
}

// GetFromStoreID 通过store_id获取内容 店铺ID
func (obj *_StoreUsersMgr) GetFromStoreID(storeID int64) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_id` = ?", storeID).Find(&results).Error

	return
}

// GetBatchFromStoreID 批量查找 店铺ID
func (obj *_StoreUsersMgr) GetBatchFromStoreID(storeIDs []int64) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_id` IN (?)", storeIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_StoreUsersMgr) GetFromUserID(userID int64) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_StoreUsersMgr) GetBatchFromUserID(userIDs []int64) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_StoreUsersMgr) GetFromCreatedAt(createdAt time.Time) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_StoreUsersMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_StoreUsersMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_StoreUsersMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_StoreUsersMgr) FetchByPrimaryKey(id uint32) (result model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUnxSuid primary or index 获取唯一内容
func (obj *_StoreUsersMgr) FetchUniqueByUnxSuid(storeUserID int64) (result model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_user_id` = ?", storeUserID).First(&result).Error

	return
}

// FetchUniqueIndexByUnxSUID primary or index 获取唯一内容
func (obj *_StoreUsersMgr) FetchUniqueIndexByUnxSUID(storeID int64, userID int64) (result model.StoreUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_id` = ? AND `user_id` = ?", storeID, userID).First(&result).Error

	return
}

// SelectPageApi
// @Desc：分页获取基本店铺信息
// @param：page
// @param：opts
// @return：resultPage
// @return：err
func (obj *_StoreUsersMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.StoreUsersApi, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

// CreateStoreUser
// @Desc：创建门店和店长
// @param：storeUsers
// @param：store
// @return：err
func (obj *_StoreUsersMgr) CreateStoreUser(storeMember model.StoreMember, storeUsers model.StoreUsers, store model.Stores) (err error) {
	tx := obj.DB.WithContext(obj.ctx).Begin()
	defer func() {
		// 防止panic
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return err
	}
	if err = tx.Table(model.StoresTableName()).Create(&store).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Table(model.StoreUsersTableName()).Create(&storeUsers).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Table(model.StoreMemberTableName()).Create(&storeMember).Error; err != nil {
	}
	return tx.Commit().Error
}

// GetStoreIdByUserId 通过用户ID获取店铺ID
func (obj *_StoreUsersMgr) GetStoreIdByUserId(userId int64) (storeId int64) {
	obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`user_id` = ?", userId).Select("store_id").Find(&storeId)
	return
}

// GetFromUserIdApi 通过user_id获取内容 用户IID
func (obj *_StoreUsersMgr) GetFromUserIdApi(userId int64) (result model.StoreUsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`user_id` = ?", userId).Find(&result).Error
	return
}

// GetFromStoreIdApi 通过store_id获取内容 店铺ID
func (obj *_StoreUsersMgr) GetFromStoreIdApi(storeId int64) (result model.StoreUsersApi) {
	obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_id` = ?", storeId).Find(&result)
	return
}

// GetFromStoreUserIdApi 通过store_user_id获取内容 店主ID
func (obj *_StoreUsersMgr) GetFromStoreUserIdApi(storeUserId int64) (result model.StoreUsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Where("`store_user_id` = ?", storeUserId).Find(&result).Error
	return
}
