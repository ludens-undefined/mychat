package model

import (
	"context"
	"fmt"
	"time"
)

type Users struct {
	ID        int       `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`                // 姓名
	Phone     string    `gorm:"column:phone;type:varchar(32);not null" json:"phone"`               // 手机号
	Email     string    `gorm:"column:email;type:varchar(32);default:'';not null" json:"email"`    // 邮箱
	Avatar    string    `gorm:"column:avatar;type:varchar(255);default:'';not null" json:"avatar"` // 头像
	Password  string    `gorm:"column:password;type:varchar(255)" json:"password"`                 // 用户登录密码
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (t *Users) TableName() string {
	return "users"
}

type _UsersModel struct {
	*_BaseMgr
}

// UsersModel 获取表对象
func UsersModel() *_UsersModel {
	if db == nil {
		panic(fmt.Errorf("GooucXmfActivityFormMgr need init by db"))
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersModel{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetUsersById retrieves Users by ID. Returns error if
// ID doesn't exist
func (obj *_UsersModel) GetUsersById(id int) (result Users, err error) {
	err = obj.DB.Debug().First(&result, id).Error
	return
}

// GetsUsers GetUsersById retrieves Users by ID. Returns error if
// ID doesn't exist
func (obj *_UsersModel) GetsUsers() (result Users, err error) {
	err = obj.DB.Debug().Find(&result).Error
	return
}

// AddUsers insert a new Users into database and returns
// last inserted ID on success.
func (obj *_UsersModel) AddUsers(m Users) (id int64, err error) {
	result := obj.DB.Create(m)
	return result.RowsAffected, result.Error
}

// UpdateUsersById UpdateBusiness updates Users by ID and returns error if
// the record to be updated doesn't exist
func (obj *_UsersModel) UpdateUsersById(m Users) (err error) {
	err = obj.DB.Save(&m).Error
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func (obj *_UsersModel) DeleteUsers(id int) (err error) {
	err = obj.DB.Delete(&Users{}, id).Error
	return
}
