package sqls

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// StoreUsers 店主表
type StoreUsers struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"-"`
	StoreUserID int64     `gorm:"column:store_user_id" json:"storeUserId"` // 店主ID
	StoreID     int64     `gorm:"column:store_id" json:"storeId"`          // 店铺ID
	UserID      int64     `gorm:"column:user_id" json:"userId"`            // 用户ID
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`      // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`      // 更新时间
}
type StoreUsersApi struct {
	StoreUserID int64  `gorm:"column:store_user_id" json:"storeUserId,string"` // 店主ID
	StoreID     int64  `gorm:"column:store_id" json:"storeId,string"`          // 店铺ID
	UserID      int64  `gorm:"column:user_id" json:"userId,string"`            // 用户ID
	StoreInfo   Stores `gorm:"foreignkey:StoreID"`
}

// StoreUsersColumns get sql column name.获取数据库列名
var StoreUsersColumns = struct {
	ID          string
	StoreUserID string
	StoreID     string
	UserID      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	StoreUserID: "store_user_id",
	StoreID:     "store_id",
	UserID:      "user_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

type StoreUsersMgr struct {
	*_BaseMgr
}

func StoreUsersTableName() string {
	return "store_users"
}

func NewStoreUsersMgr(db *gorm.DB) *StoreUsersMgr {
	ctx, cancel := context.WithCancel(context.Background())
	return &StoreUsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table(StoreUsersTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *StoreUsersMgr) Reset() *StoreUsersMgr {
	obj.New()
	return obj
}

// SelectPageApi
// @Desc：分页获取基本店铺信息
// @param：page
// @param：opts
// @return：resultPage
// @return：err
func (obj *StoreUsersMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]StoreUsersApi, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(StoreUsers{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

// GetFromUserIdApi 通过user_id获取内容 用户IID
func (obj *StoreUsersMgr) GetFromUserIdApi(userId int64) (result StoreUsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StoreUsers{}).Where("`user_id` = ?", userId).Find(&result).Error
	return
}

// GetFromStoreIdApi 通过store_id获取内容 店铺ID
func (obj *StoreUsersMgr) GetFromStoreIdApi(storeId int64) (result StoreUsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StoreUsers{}).Where("`store_id` = ?", storeId).Find(&result).Error
	return
}

// GetFromStoreUserIdApi 通过store_user_id获取内容 店主ID
func (obj *StoreUsersMgr) GetFromStoreUserIdApi(storeUserId int64) (result StoreUsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(StoreUsers{}).Where("`store_user_id` = ?", storeUserId).Find(&result).Error
	return
}

// CreateStoreUser
// @Desc：创建门店和店长
// @param：storeUsers
// @param：store
// @return：err
func (obj *StoreUsersMgr) CreateStoreUser(storeUsers StoreUsers, store Stores) (err error) {
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
	if err = tx.Table(StoresTableName()).Create(&store).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Table(StoreUsersTableName()).Create(&storeUsers).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
