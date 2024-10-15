package sqls

import (
	"context"
	"gorm.io/gorm"
	sqlsUser "store/app/user/model/sqls"
	"time"
)

// StoreMember 店铺会员表
type StoreMember struct {
	Id            uint32    `gorm:"primaryKey;column:id" json:"-"`
	StoreMemberId int64     `gorm:"column:store_member_id" json:"storeMemberId"` // 会员ID
	StoreId       int64     `gorm:"column:store_id" json:"storeId"`              // 店铺ID
	UserId        int64     `gorm:"column:user_id" json:"userId"`                // 用户ID
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt"`          // 创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updatedAt"`          // 更新时间
}

type MemberUserItem struct {
	StoreMemberId int64             `gorm:"column:store_member_id" json:"storeMemberId,string"` // 会员ID
	StoreId       int64             `gorm:"column:store_id" json:"storeId,string"`              // 店铺ID
	UserId        int64             `gorm:"column:user_id" json:"userId,string"`                // 用户ID
	User          sqlsUser.UsersApi `gorm:"foreignkey:user_id;references:user_id"`
}

type StoresMemberMgr struct {
	*_BaseMgr
}

func StoreMemberTableName() string {
	return "store_member"
}

func NewStoresMemberMgr(db *gorm.DB) *StoresMemberMgr {
	ctx, cancel := context.WithCancel(context.Background())
	return &StoresMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table(StoreMemberTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *StoresMemberMgr) Reset() *StoresMemberMgr {
	obj.New()
	return obj
}

func (obj *StoresMemberMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MemberUserItem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(StoreMember{}).Where(options.query)
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

func (obj *StoresMemberMgr) WithStoreId(storeId int64) Option {
	return optionFunc(func(o *options) { o.query["store_id"] = storeId })
}

func (obj *StoresMemberMgr) GetStoreIdsByUserId(userId int64) []int64 {
	storeIds := make([]int64, 0)
	obj.DB.WithContext(obj.ctx).Model(StoreMember{}).Where("user_id = ?", userId).Select("store_id").Find(&storeIds)
	return storeIds
}

func (obj *StoresMemberMgr) MemberJoin(storeId int64, userId int64, storeMemberId int64) (row int64, err error) {
	obj.DB.WithContext(obj.ctx).Model(StoreMember{}).Where("user_id = ?", userId).Where("store_id = ?", storeId).Count(&row)
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
	tx.Table(StoreMemberTableName()).Create(&StoreMember{
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
func (obj *StoresMemberMgr) GetMemberContacts(storeId int64) int64 {
	var contacts int64
	obj.DB.WithContext(obj.ctx).Model(StoreMember{}).Where("store_id = ?", storeId).Count(&contacts)
	return contacts
}
