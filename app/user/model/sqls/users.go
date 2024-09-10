package sqls

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const (
	USER_STATUS_1 int8 = 1
	USER_STATUS_2 int8 = 2
)

var UserStatusName = map[int8]string{
	USER_STATUS_1: "启用",
	USER_STATUS_2: "禁用",
}

// Users 用户表
type Users struct {
	ID        uint32    `gorm:"primaryKey;column:id" json:"-"`
	UserID    int64     `gorm:"column:user_id" json:"userId"`          // 用户IID
	Status    int8      `gorm:"column:status;default:1" json:"status"` // 1-启用、2-禁用
	Mobile    int64     `gorm:"column:mobile" json:"mobile"`           // 手机号
	Password  string    `gorm:"column:password" json:"password"`       // 密码
	Name      string    `gorm:"column:name" json:"name"`               // 昵称
	Avatar    string    `gorm:"column:avatar" json:"avatar"`           // 头像
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`    // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`    // 更新时间
}

type UsersApi struct {
	UserID int64  `gorm:"column:user_id" json:"userId,string"` // 用户IID
	Status int8   `gorm:"column:status" json:"status"`         // 1-启用、2-禁用
	Mobile int64  `gorm:"column:mobile" json:"mobile,string"`  // 手机号
	Name   string `gorm:"column:name" json:"name"`             // 昵称
	Avatar string `gorm:"column:avatar" json:"avatar"`         // 头像
}

// UsersColumns get sql column name.获取数据库列名
var UsersColumns = struct {
	ID        string
	UserID    string
	Status    string
	Mobile    string
	Password  string
	Name      string
	Avatar    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Status:    "status",
	Mobile:    "mobile",
	Password:  "Password",
	Name:      "name",
	Avatar:    "avatar",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

type UsersMgr struct {
	*_BaseMgr
}

func UsersTableName() string {
	return "users"
}

func NewUserMgr(db *gorm.DB) *UsersMgr {
	ctx, cancel := context.WithCancel(context.Background())
	return &UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table(UsersTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *UsersMgr) Reset() *UsersMgr {
	obj.New()
	return obj
}

// CreatUser
// @Desc：创建用户
// @param：users
// @return：err
func (obj *UsersMgr) CreatUser(users Users) (err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Create(&users).Error
	return
}

// GetFromUserID
// @Desc：通过UUID获取
// @param：userID
// @return：result
// @return：err
func (obj *UsersMgr) GetFromUserID(userID int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`user_id` = ?", userID).Find(&result).Error
	return
}

// GetBatchFromUserID 批量查找 用户IID
func (obj *UsersMgr) GetBatchFromUserID(userIDs []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	return
}

// GetFromMobile 通过mobil获取内容
func (obj *UsersMgr) GetFromMobile(mobile int32) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`mobile` = ?", mobile).Find(&result).Error
	return
}

// GetUserApi
// @Desc：获取基本用户信息
// @param：users
// @return：result
// @return：err
func (obj *UsersMgr) GetUserApi(users Users) (result UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(&users).Find(&result).Error
	return
}

// SelectPageApi
// @Desc：分页获取基本用户信息
// @param：page
// @param：opts
// @return：resultPage
// @return：err
func (obj *UsersMgr) SelectPageApi(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UsersApi, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query)
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
func (obj *UsersMgr) GetFromUserIDApi(userID int64) (result UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`user_id` = ?", userID).Find(&result).Error
	return
}

// GetBatchFromUserIDApi 批量查找 用户IID
func (obj *UsersMgr) GetBatchFromUserIDApi(userIDs []int64) (results []*UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	return
}

// GetFromMobileApi 通过mobil获取内容
func (obj *UsersMgr) GetFromMobileApi(mobile int32) (result UsersApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`mobile` = ?", mobile).Find(&result).Error
	return
}
