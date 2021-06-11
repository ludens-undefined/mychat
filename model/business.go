package model

import (
	"context"
	"fmt"
)

type Business struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	ShopID     string `gorm:"index:idx_shop_id;column:shop_id;type:varchar(255);not null" json:"shop_id"` // 商铺id
	AdminType  bool   `gorm:"column:admin_type;type:tinyint(1);not null" json:"admin_type"`               // 1店长 2员工
	RoleID     int    `gorm:"column:role_id;type:int(11);not null" json:"role_id"`                        // 角色分组
	Phone      string `gorm:"column:phone;type:varchar(32);not null" json:"phone"`                        // 手机号
	Nickname   string `gorm:"column:nickname;type:varchar(32)" json:"nickname"`                           // 昵称
	Avatar     string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`                     // 头像
	Password   string `gorm:"column:password;type:varchar(255)" json:"password"`                          // 用户登录密码
	Unionid    string `gorm:"column:unionid;type:varchar(255)" json:"unionid"`
	WxBinding  int    `gorm:"column:wx_binding;type:tinyint(1)" json:"wx_binding"`                // 微信绑定 1已绑定 2未绑定
	WxAvatar   string `gorm:"column:wx_avatar;type:varchar(500);not null" json:"wx_avatar"`       // 商户微信头像
	WxNickname string `gorm:"column:wx_nickname;type:varchar(100);not null" json:"wx_nickname"`   // 商户微信昵称
	ShareTitle string `gorm:"column:share_title;type:varchar(255)" json:"share_title"`            // 分享标题
	ShareTips  string `gorm:"column:share_tips;type:varchar(255)" json:"share_tips"`              // 分享提升
	ShareImg   string `gorm:"column:share_img;type:varchar(255)" json:"share_img"`                // 分享图片
	IsAdmin    uint8  `gorm:"column:is_admin;type:tinyint(3) unsigned;not null" json:"is_admin"`  // 1试用 2会员
	Createtime int    `gorm:"column:createtime;type:int(11) unsigned;not null" json:"createtime"` // 注册时间
	Expiretime int    `gorm:"column:expiretime;type:int(11) unsigned;not null" json:"expiretime"` // 过期时间
	Logintime  uint32 `gorm:"column:logintime;type:int(1) unsigned;not null" json:"logintime"`    // 登录时间
	Status     bool   `gorm:"column:status;type:tinyint(1)" json:"status"`                        // 状态 1启用 2禁用
	Pid        int    `gorm:"column:pid;type:int(11);not null" json:"pid"`                        // 推荐人
	IsDelete   bool   `gorm:"column:is_delete;type:tinyint(1)" json:"is_delete"`                  // 状态 1启用 2删除
}

func (t *Business) TableName() string {
	return "goouc_xmf_business"
}

type _BusinessModel struct {
	*_BaseMgr
}

// BusinessModel 获取表对象
func BusinessModel() *_BusinessModel {
	if db == nil {
		panic(fmt.Errorf("GooucXmfActivityFormMgr need init by db"))
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &_BusinessModel{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetBusinessById retrieves Business by Id. Returns error if
// Id doesn't exist
func (obj *_BusinessModel) GetBusinessById(id int) (result Business, err error) {
	err = obj.DB.Debug().First(&result, id).Error
	return
}

// GetsBusiness GetBusinessById retrieves Business by Id. Returns error if
// Id doesn't exist
func (obj *_BusinessModel) GetsBusiness() (result Business, err error) {
	err = obj.DB.Debug().Find(&result).Error
	return
}

// AddBusiness insert a new Business into database and returns
// last inserted Id on success.
func (obj *_BusinessModel) AddBusiness(m Business) (id int64, err error) {
	result := obj.DB.Create(m)
	return result.RowsAffected, result.Error
}

// UpdateBusiness updates Business by Id and returns error if
// the record to be updated doesn't exist
func (obj *_BusinessModel) UpdateBusinessById(m Business) (err error) {
	err = obj.DB.Save(&m).Error
	return
}

// DeleteBusiness deletes Business by Id and returns error if
// the record to be deleted doesn't exist
func (obj *_BusinessModel) DeleteBusiness(id int) (err error) {
	err = obj.DB.Delete(&Business{}, id).Error
	return
}
