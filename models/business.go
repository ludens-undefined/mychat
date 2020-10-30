package models

import (
	"fmt"
)

type Business struct {
	Id         int    `gorm:"primaryKey;column:id;"`
	ShopId     string `gorm:"column:shop_id;size:255;not null"` // 商铺id
	AdminType  int8   `gorm:"column:admin_type;not null"`       // 1店长 2员工
	RoleId     int    `gorm:"column:role_id;not null"`          // 角色分组
	Phone      string `gorm:"column:phone;size:32;not null"`    // 手机号
	Nickname   string `gorm:"column:nickname;size:32"`          // 昵称
	Avatar     string `gorm:"column:avatar;size:255;not null"`  // 头像
	Password   string `gorm:"column:password;size:255"`         // 用户登录密码
	Unionid    string `gorm:"column:unionid;size:255"`
	WxBinding  int8   `gorm:"column:wx_binding"`                    // 微信绑定 1已绑定 2未绑定
	WxAvatar   string `gorm:"column:wx_avatar;size:500;not null"`   // 商户微信头像
	WxNickname string `gorm:"column:wx_nickname;size:100;not null"` // 商户微信昵称
	ShareTitle string `gorm:"column:share_title;size:255"`          // 分享标题
	ShareTips  string `gorm:"column:share_tips;size:255"`           // 分享提升
	ShareImg   string `gorm:"column:share_img;size:255"`            // 分享图片
	IsAdmin    uint8  `gorm:"column:is_admin;not null"`             // 1试用 2会员
	Createtime uint   `gorm:"column:createtime;not null"`           // 注册时间
	Expiretime uint   `gorm:"column:expiretime;not null"`           // 过期时间
	Logintime  uint   `gorm:"column:logintime;not null"`            // 登录时间
	Status     int8   `gorm:"column:status"`                        // 状态 1启用 2禁用
	Pid        int    `gorm:"column:pid;not null"`                  // 推荐人
	IsDelete   int8   `gorm:"column:is_delete"`                     // 状态 1启用 2删除
}

func (*Business) TableName() string {
	return "goouc_xmf_business"
}

// AddBusiness insert a new Business into database and returns
// last inserted Id on success.
func AddBusiness(m Business) (id int64, err error) {
	result := DB.Create(m)
	return result.RowsAffected, result.Error
}

// GetBusinessById retrieves Business by Id. Returns error if
// Id doesn't exist
func GetBusinessById(id int) (v Business, err error) {
	result := DB.Debug().First(&v, id)
	if result.Error == nil {
		return v, nil
	}
	return v, result.Error
}

//
//// GetAllBusiness retrieves all Business matches certain condition. Returns empty list if
//// no records exist
//func GetAllBusiness(query map[string]string, fields []string, sortby []string, order []string,
//	offset int64, limit int64) (ml []interface{}, err error) {
//
//	qs := o.QueryTable(new(Business))
//	// query k=v
//	for k, v := range query {
//		// rewrite dot-notation to Object__Attribute
//		k = strings.Replace(k, ".", "__", -1)
//		if strings.Contains(k, "isnull") {
//			qs = qs.Filter(k, (v == "true" || v == "1"))
//		} else {
//			qs = qs.Filter(k, v)
//		}
//	}
//	// order by:
//	var sortFields []string
//	if len(sortby) != 0 {
//		if len(sortby) == len(order) {
//			// 1) for each sort field, there is an associated order
//			for i, v := range sortby {
//				orderby := ""
//				if order[i] == "desc" {
//					orderby = "-" + v
//				} else if order[i] == "asc" {
//					orderby = v
//				} else {
//					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//			qs = qs.OrderBy(sortFields...)
//		} else if len(sortby) != len(order) && len(order) == 1 {
//			// 2) there is exactly one order, all the sorted fields will be sorted by this order
//			for _, v := range sortby {
//				orderby := ""
//				if order[0] == "desc" {
//					orderby = "-" + v
//				} else if order[0] == "asc" {
//					orderby = v
//				} else {
//					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//		} else if len(sortby) != len(order) && len(order) != 1 {
//			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
//		}
//	} else {
//		if len(order) != 0 {
//			return nil, errors.New("Error: unused 'order' fields")
//		}
//	}
//
//	var l []Business
//	qs = qs.OrderBy(sortFields...)
//	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
//		if len(fields) == 0 {
//			for _, v := range l {
//				ml = append(ml, v)
//			}
//		} else {
//			// trim unused fields
//			for _, v := range l {
//				m := make(map[string]interface{})
//				val := reflect.ValueOf(v)
//				for _, fname := range fields {
//					m[fname] = val.FieldByName(fname).Interface()
//				}
//				ml = append(ml, m)
//			}
//		}
//		return ml, nil
//	}
//	return nil, err
//}

// UpdateBusiness updates Business by Id and returns error if
// the record to be updated doesn't exist
func UpdateBusinessById(m Business) (err error) {
	result := DB.Save(&m)
	// ascertain id exists in the database
	if result.Error == nil {
		fmt.Println("Number of records updated in database:", result.RowsAffected)
	}
	return result.Error
}

// DeleteBusiness deletes Business by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBusiness(id int) (err error) {
	v := Business{}
	result := DB.Delete(&v, id)
	// ascertain id exists in the database
	if result.Error == nil {
		fmt.Println("Number of records deleted in database:", result.RowsAffected)
	}
	return result.Error
}
