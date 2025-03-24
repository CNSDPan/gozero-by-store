package query

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"store/db/dao/model"
	"time"
)

type _ChatLogMgr struct {
	*_BaseMgr
}

// ChatLogMgr open func
func ChatLogMgr(db *gorm.DB) *_ChatLogMgr {
	if db == nil {
		panic(fmt.Errorf("ChatLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ChatLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("chat_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_ChatLogMgr) Debug() *_ChatLogMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ChatLogMgr) GetTableName() string {
	return "chat_log"
}

// Reset 重置gorm会话
func (obj *_ChatLogMgr) Reset() *_ChatLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ChatLogMgr) Get() (result model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ChatLogMgr) Gets() (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ChatLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 用户ID
func (obj *_ChatLogMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStoreID store_id获取 店铺ID
func (obj *_ChatLogMgr) WithStoreID(storeID int64) Option {
	return optionFunc(func(o *options) { o.query["store_id"] = storeID })
}

// WithMessage message获取 消息
func (obj *_ChatLogMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithTimestamp timestamp获取 记录时间;微秒
func (obj *_ChatLogMgr) WithTimestamp(timestamp int64) Option {
	return optionFunc(func(o *options) { o.query["timestamp"] = timestamp })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ChatLogMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ChatLogMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ChatLogMgr) GetByOption(opts ...Option) (result model.ChatLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ChatLogMgr) GetByOptions(opts ...Option) (results []*model.ChatLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ChatLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.ChatLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where(options.query)
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

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_ChatLogMgr) GetFromUserID(userID int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_ChatLogMgr) GetBatchFromUserID(userIDs []int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromStoreID 通过store_id获取内容 店铺ID
func (obj *_ChatLogMgr) GetFromStoreID(storeID int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`store_id` = ?", storeID).Find(&results).Error

	return
}

// GetBatchFromStoreID 批量查找 店铺ID
func (obj *_ChatLogMgr) GetBatchFromStoreID(storeIDs []int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`store_id` IN (?)", storeIDs).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容 消息
func (obj *_ChatLogMgr) GetFromMessage(message string) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`message` = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量查找 消息
func (obj *_ChatLogMgr) GetBatchFromMessage(messages []string) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`message` IN (?)", messages).Find(&results).Error

	return
}

// GetFromTimestamp 通过timestamp获取内容 记录时间;微秒
func (obj *_ChatLogMgr) GetFromTimestamp(timestamp int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`timestamp` = ?", timestamp).Find(&results).Error

	return
}

// GetBatchFromTimestamp 批量查找 记录时间;微秒
func (obj *_ChatLogMgr) GetBatchFromTimestamp(timestamps []int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`timestamp` IN (?)", timestamps).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ChatLogMgr) GetFromCreatedAt(createdAt time.Time) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_ChatLogMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ChatLogMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_ChatLogMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByIDxUID  获取多个内容
func (obj *_ChatLogMgr) FetchIndexByIDxUID(userID int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByIDxStoreid  获取多个内容
func (obj *_ChatLogMgr) FetchIndexByIDxStoreid(storeID int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`store_id` = ?", storeID).Find(&results).Error

	return
}

// FetchIndexByIDxTime  获取多个内容
func (obj *_ChatLogMgr) FetchIndexByIDxTime(timestamp int64) (results []*model.ChatLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where("`timestamp` = ?", timestamp).Find(&results).Error

	return
}

// GetStoreChatPage
// @Desc：
// @param：page
// @param：where 构建查询条件
// @param：opts
// @return：resultPage
// @return：err
func (obj *_ChatLogMgr) GetStoreChatPage(page IPage, where interface{}, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]model.ChatLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Where(options.query).Where(where)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

func (obj *_ChatLogMgr) CreateChatLog(storeId int64, userId int64, message string, createAt time.Time) {
	chatLog := model.ChatLog{
		StoreId:   storeId,
		UserId:    userId,
		Message:   message,
		CreatedAt: createAt,
	}
	obj.DB.WithContext(obj.ctx).Model(model.StoreUsers{}).Create(&chatLog)
}

// InsertChatLogs
// @Desc：批量插入
// @param：chatLogs
func (obj *_ChatLogMgr) InsertChatLogs(chatLogs []model.ChatLog) error {
	return obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Create(&chatLogs).Error
}

// InitChatLog
// @Desc：获取每个店铺群的10条最新聊天记录,每次最多获取10个店铺
// @param：page
// @param：userId
// @return：resultPage
// @return：err
func (obj *_ChatLogMgr) InitChatLog(page IPage, userId int64) (resultPage IPage, err error) {
	resultPage = page
	results := make([]model.ChatLogApi, 0)
	// 子查询
	subQuery := obj.DB.WithContext(obj.ctx).Model(model.ChatLog{}).Select("store_id, MAX(timestamp) AS max_timestamp").Group("store_id")
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.ChatLog{})
	query.Select("chat_log.*,stores.name as store_name")
	query.Joins("join stores on stores.store_id = chat_log.store_id")
	query.Joins("JOIN (?) AS last ON chat_log.store_id = last.store_id AND chat_log.timestamp = last.max_timestamp", subQuery)
	query.Where("user_id = ?", userId)
	query.Count(&count)
	resultPage.SetTotal(count)

	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Order("chat_log.timestamp desc").Find(&results).Error

	resultPage.SetRecords(results)
	return
}

func (obj *_ChatLogMgr) SelectPageChatLog(page IPage, storeId int64, timestamp int64) (resultPage IPage, err error) {
	resultPage = page
	results := make([]model.ChatLogApi, 0)

	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(model.ChatLog{})
	query.Select("chat_log.*,stores.name as store_name")
	query.Joins("join stores on stores.store_id = chat_log.store_id")
	query.Where("chat_log.store_id = ?", storeId).Where("timestamp < ?", timestamp)
	query.Count(&count)
	resultPage.SetTotal(count)

	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Order("chat_log.timestamp desc").Find(&results).Error

	resultPage.SetRecords(results)
	return
}
