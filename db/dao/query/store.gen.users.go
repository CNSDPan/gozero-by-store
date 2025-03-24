package query

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"store/db/dao/model"
	"time"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UsersMgr) Debug() *_UsersMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Reset 重置gorm会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsersMgr) Get() (result model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(model.Users{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UsersMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户ID
func (obj *_UsersMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatus status获取 1-启用、2-禁用
func (obj *_UsersMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMobile mobile获取 手机号
func (obj *_UsersMgr) WithMobile(mobile int64) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithPassword password获取 密码
func (obj *_UsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithName name获取 昵称
func (obj *_UsersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAvatar avatar获取 头像
func (obj *_UsersMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UsersMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_UsersMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result model.Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*model.Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UsersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.Users, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where(options.query)
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
func (obj *_UsersMgr) GetFromID(id uint32) (result model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UsersMgr) GetBatchFromID(ids []uint32) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 1-启用、2-禁用
func (obj *_UsersMgr) GetFromStatus(status bool) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 1-启用、2-禁用
func (obj *_UsersMgr) GetBatchFromStatus(statuss []bool) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号
func (obj *_UsersMgr) GetBatchFromMobile(mobiles []int64) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_UsersMgr) GetFromPassword(password string) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_UsersMgr) GetBatchFromPassword(passwords []string) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 昵称
func (obj *_UsersMgr) GetFromName(name string) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 昵称
func (obj *_UsersMgr) GetBatchFromName(names []string) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_UsersMgr) GetFromAvatar(avatar string) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_UsersMgr) GetBatchFromAvatar(avatars []string) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UsersMgr) GetFromCreatedAt(createdAt time.Time) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_UsersMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_UsersMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_UsersMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(id uint32) (result model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUnxUID primary or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByUnxUID(userID int64) (result model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`user_id` = ?", userID).First(&result).Error

	return
}

// FetchUniqueByUnxMobile primary or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByUnxMobile(mobile int64) (result model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`mobile` = ?", mobile).First(&result).Error

	return
}

// CreatUser
// @Desc：创建用户
// @param：users
// @return：err
func (obj *_UsersMgr) CreatUser(users model.Users) (err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Create(&users).Error
	return
}

// GetFromMobileApi 通过mobil获取内容
func (obj *_UsersMgr) GetFromMobileApi(mobile int64) (result model.UsersApi, err error) {
	err = obj.WithContext(obj.ctx).Model(model.Users{}).Where("`mobile` = ?", mobile).Find(&result).Error
	return
}

// MyStoreIds 获取我的店铺
func (obj *_UsersMgr) MyStoreIds(userId int64) []int64 {
	var (
		storeIds    = make([]int64, 0)
		storeUsers  = make([]int64, 0)
		storeMember = make([]int64, 0)
	)

	obj.DB.WithContext(obj.ctx).Table(model.StoreMemberTableName()).Where("`user_id` = ?", userId).Select("store_id").Find(&storeMember)
	storeIds = append(storeIds, storeUsers...)
	storeIds = append(storeIds, storeMember...)
	return storeIds
}

// GetFromUserID
// @Desc：通过UUID获取
// @param：userID
// @return：result
// @return：err
func (obj *_UsersMgr) GetFromUserID(userID int64) (result model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`user_id` = ?", userID).Find(&result).Error
	return
}

// GetBatchFromUserID 批量查找 用户IID
func (obj *_UsersMgr) GetBatchFromUserID(userIDs []int64) (results []*model.Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	return
}

// GetFromMobile 通过mobil获取内容
func (obj *_UsersMgr) GetFromMobile(mobile int64) (result model.Users, err error) {

	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`mobile` = ?", mobile).Find(&result).Error
	return
}

// GetUserApi
// @Desc：获取基本用户信息
// @param：users
// @return：result
// @return：err
func (obj *_UsersMgr) GetUserApi(users model.Users) (result model.UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where(&users).Find(&result).Error
	return
}

// SelectPageApi
// @Desc：分页获取基本用户信息
// @param：page
// @param：opts
// @return：resultPage
// @return：err
func (obj *_UsersMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.UsersApi, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

// GetFromUserIDApi 通过user_id获取内容 用户IID
func (obj *_UsersMgr) GetFromUserIDApi(userID int64) (result model.UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`user_id` = ?", userID).Find(&result).Error
	return
}

// GetBatchFromUserIDApi 批量查找 用户IID
func (obj *_UsersMgr) GetBatchFromUserIDApi(userIDs []int64) (results []*model.UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.Users{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	return
}
