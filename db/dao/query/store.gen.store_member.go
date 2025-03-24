package query

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"store/db/dao/model"
	"strconv"
	"time"
)

type _StoreMemberMgr struct {
	*_BaseMgr
}

// StoreMemberMgr open func
func StoreMemberMgr(db *gorm.DB) *_StoreMemberMgr {
	if db == nil {
		panic(fmt.Errorf("StoreMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("store_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_StoreMemberMgr) Debug() *_StoreMemberMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StoreMemberMgr) GetTableName() string {
	return "store_member"
}

// Reset 重置gorm会话
func (obj *_StoreMemberMgr) Reset() *_StoreMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_StoreMemberMgr) Get() (result model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StoreMemberMgr) Gets() (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_StoreMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_StoreMemberMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStoreMemberID store_member_id获取 会员ID
func (obj *_StoreMemberMgr) WithStoreMemberID(storeMemberID int64) Option {
	return optionFunc(func(o *options) { o.query["store_member_id"] = storeMemberID })
}

// WithStoreID store_id获取 店铺ID
func (obj *_StoreMemberMgr) WithStoreID(storeID int64) Option {
	return optionFunc(func(o *options) { o.query["store_id"] = storeID })
}

// WithUserID user_id获取 用户ID
func (obj *_StoreMemberMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_StoreMemberMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_StoreMemberMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_StoreMemberMgr) GetByOption(opts ...Option) (result model.StoreMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StoreMemberMgr) GetByOptions(opts ...Option) (results []*model.StoreMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_StoreMemberMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.StoreMember, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where(options.query)
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
func (obj *_StoreMemberMgr) GetFromID(id uint32) (result model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_StoreMemberMgr) GetBatchFromID(ids []uint32) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStoreMemberID 通过store_member_id获取内容 会员ID
func (obj *_StoreMemberMgr) GetFromStoreMemberID(storeMemberID int64) (result model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`store_member_id` = ?", storeMemberID).First(&result).Error

	return
}

// GetBatchFromStoreMemberID 批量查找 会员ID
func (obj *_StoreMemberMgr) GetBatchFromStoreMemberID(storeMemberIDs []int64) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`store_member_id` IN (?)", storeMemberIDs).Find(&results).Error

	return
}

// GetFromStoreID 通过store_id获取内容 店铺ID
func (obj *_StoreMemberMgr) GetFromStoreID(storeID int64) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`store_id` = ?", storeID).Find(&results).Error

	return
}

// GetBatchFromStoreID 批量查找 店铺ID
func (obj *_StoreMemberMgr) GetBatchFromStoreID(storeIDs []int64) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`store_id` IN (?)", storeIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_StoreMemberMgr) GetFromUserID(userID int64) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_StoreMemberMgr) GetBatchFromUserID(userIDs []int64) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_StoreMemberMgr) GetFromCreatedAt(createdAt time.Time) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_StoreMemberMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_StoreMemberMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_StoreMemberMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_StoreMemberMgr) FetchByPrimaryKey(id uint32) (result model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUnxSmid primary or index 获取唯一内容
func (obj *_StoreMemberMgr) FetchUniqueByUnxSmid(storeMemberID int64) (result model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`store_member_id` = ?", storeMemberID).First(&result).Error

	return
}

// FetchIndexByIDxSUID  获取多个内容
func (obj *_StoreMemberMgr) FetchIndexByIDxSUID(storeID int64, userID int64) (results []*model.StoreMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("`store_id` = ? AND `user_id` = ?", storeID, userID).Find(&results).Error

	return
}

func (obj *_StoreMemberMgr) WithStoreId(storeId int64) Option {
	return optionFunc(func(o *options) { o.query["store_id"] = storeId })
}

func (obj *_StoreMemberMgr) GetStoreIdsByUserId(userId int64) []int64 {
	storeIds := make([]int64, 0)
	obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("user_id = ?", userId).Select("store_id").Find(&storeIds)
	return storeIds
}

func (obj *_StoreMemberMgr) MemberJoin(storeId int64, userId int64, storeMemberId int64) (row int64, err error) {
	obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("user_id = ?", userId).Where("store_id = ?", storeId).Count(&row)
	if row > 0 {
		return row, nil
	}
	tx := obj.DB.WithContext(obj.ctx).Begin()
	defer func() {
		// 防止panic
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return 0, tx.Error
	}
	tx.Table(model.StoreMemberTableName()).Create(&model.StoreMember{
		StoreMemberId: storeMemberId,
		StoreId:       storeId,
		UserId:        userId,
	})
	return 0, tx.Commit().Error
}

// GetMemberContacts
// @Desc：获取店铺会员总数
// @param：storeId
// @return：int64
func (obj *_StoreMemberMgr) GetMemberContacts(storeId int64) int64 {
	var contacts int64
	obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("store_id = ?", storeId).Count(&contacts)
	return contacts
}

// MapKeyUserId
// @Desc： 获取店铺所有会员，并且已map的形式放回
// @param：storeId
// @return：map[string]MemberUserItem
func (obj *_StoreMemberMgr) MapKeyUserId(storeId int64) map[string]model.MemberUserItem {
	result := make(map[string]model.MemberUserItem)
	results := make([]model.MemberUserItem, 0)
	query := obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where("store_id = ?", storeId)
	query.Preload("User").Find(&results)

	for _, item := range results {
		key := strconv.FormatInt(item.UserId, 10)
		result[key] = item
	}
	return result
}

// InitChatLog
// @Desc：获取每个店铺群的10条最新聊天记录,每次最多获取10个店铺
// @param：page
// @param：userId
// @return：resultPage
// @return：err
func (obj *_StoreMemberMgr) InitChatLog(page IPage, userId int64) (resultPage IPage, err error) {
	resultPage = page
	results := make([]model.MemberChatLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.MemberChatLog{})
	query.Select("store_member.store_id,store_member.user_id,stores.name as store_name")
	query.Joins("join stores on stores.store_id = store_member.store_id")
	query.Joins("left join chat_log on chat_log.store_id = store_member.store_id")
	query.Preload("ChatLog")
	query.Where("store_member.user_id = ?", userId).Group("store_member.store_id")
	query.Count(&count)
	resultPage.SetTotal(count)

	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Order("timestamp desc").Find(&results).Error

	resultPage.SetRecords(results)
	return
}

func (obj *_StoreMemberMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.MemberUserItem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.StoreMember{}).Where(options.query)
	query.Preload("User")
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}
