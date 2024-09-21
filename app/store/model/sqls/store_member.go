package sqls

import (
	"context"
	"gorm.io/gorm"
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

type StoreMemberApi struct {
	StoreMemberId int64 `gorm:"column:store_member_id" json:"storeMemberId,string"` // 会员ID
	StoreId       int64 `gorm:"column:store_id" json:"storeId,string"`              // 店铺ID
	UserId        int64 `gorm:"column:user_id" json:"userId,string"`                // 用户ID
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
