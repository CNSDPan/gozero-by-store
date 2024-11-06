package sqls

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// ChatLog 聊天记录表
type ChatLog struct {
	UserId    int64     `gorm:"column:user_id" json:"userId"`       // 用户ID
	StoreId   int64     `gorm:"column:store_id" json:"storeId"`     // 店铺ID
	Message   string    `gorm:"column:message" json:"message"`      // 消息
	Timestamp int64     `gorm:"column:timestamp" json:"timestamp"`  // 记录时间;微秒
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type ChatLogApi struct {
	UserId    int64  `gorm:"column:user_id" json:"userId,string"`      // 用户ID
	StoreId   int64  `gorm:"column:store_id" json:"storeId,string"`    // 店铺ID
	StoreName string `gorm:"column:store_name" json:"storeName"`       // 店铺
	Message   string `gorm:"column:message" json:"message"`            // 消息
	Timestamp int64  `gorm:"column:timestamp" json:"timestamp,string"` // 记录时间;微秒
	CreatedAt string `gorm:"column:created_at" json:"createdAt"`       // 创建时间
}

type WithChatLog struct {
	StoreId   int64  `gorm:"column:store_id" json:"storeId,string"`    // 店铺ID
	Message   string `gorm:"column:message" json:"message"`            // 消息
	Timestamp int64  `gorm:"column:timestamp" json:"timestamp,string"` // 记录时间;微秒
	CreatedAt string `gorm:"column:created_at" json:"createdAt"`       // 创建时间
}

type ChatLogMgr struct {
	*_BaseMgr
}

// TableName get sql table name.获取数据库表名
func (m *ChatLog) TableName() string {
	return "chat_log"
}

func (m *WithChatLog) TableName() string {
	return "chat_log"
}

func ChatLogTableName() string {
	return "chat_log"
}

// ChatLogColumns get sql column name.获取数据库列名
var ChatLogColumns = struct {
	UserID    string
	StoreID   string
	Message   string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_id",
	StoreID:   "store_id",
	Message:   "message",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func NewChatLogMgr(db *gorm.DB) *ChatLogMgr {
	ctx, cancel := context.WithCancel(context.Background())
	return &ChatLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table(ChatLogTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *ChatLogMgr) Reset() *ChatLogMgr {
	obj.New()
	return obj
}

// GetStoreChatPage
// @Desc：
// @param：page
// @param：where 构建查询条件
// @param：opts
// @return：resultPage
// @return：err
func (obj *ChatLogMgr) GetStoreChatPage(page IPage, where interface{}, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ChatLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ChatLog{}).Where(options.query).Where(where)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

func (obj *ChatLogMgr) CreateChatLog(storeId int64, userId int64, message string, createAt time.Time) {
	chatLog := ChatLog{
		StoreId:   storeId,
		UserId:    userId,
		Message:   message,
		CreatedAt: createAt,
	}
	obj.DB.WithContext(obj.ctx).Model(StoreUsers{}).Create(&chatLog)
}

// InitChatLog
// @Desc：获取每个店铺群的10条最新聊天记录,每次最多获取10个店铺
// @param：page
// @param：userId
// @return：resultPage
// @return：err
func (obj *ChatLogMgr) InitChatLog(page IPage, userId int64) (resultPage IPage, err error) {
	resultPage = page
	results := make([]ChatLogApi, 0)
	// 子查询
	subQuery := obj.DB.WithContext(obj.ctx).Model(ChatLog{}).Select("store_id, MAX(created_at) AS max_created_at").Group("store_id")
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ChatLog{})
	query.Select("chat_log.*,stores.name as store_name")
	query.Joins("join stores on stores.store_id = chat_log.store_id")
	query.Joins("JOIN (?) AS last ON chat_log.store_id = last.store_id AND chat_log.created_at = last.max_created_at", subQuery)
	query.Where("user_id = ?", userId)
	query.Count(&count)
	resultPage.SetTotal(count)

	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Order("chat_log.created_at desc").Find(&results).Error

	resultPage.SetRecords(results)
	return
}

func (obj *ChatLogMgr) SelectPageChatLog(page IPage, storeId int64, timestamp int64) (resultPage IPage, err error) {
	resultPage = page
	results := make([]ChatLogApi, 0)

	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ChatLog{})
	query.Select("chat_log.*,stores.name as store_name")
	query.Joins("join stores on stores.store_id = chat_log.store_id")
	query.Where("chat_log.store_id = ?", storeId).Where("timestamp < ?", timestamp)
	query.Count(&count)
	resultPage.SetTotal(count)

	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Order("chat_log.created_at desc").Find(&results).Error

	resultPage.SetRecords(results)
	return
}
