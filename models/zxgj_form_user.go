package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjFormUser struct {
	Id         int    `orm:"column:id;autoIncrement"`
	UserId     uint   `orm:"column:user_id"` //对应goouc_xet_user表中的id
	WxNickname string `orm:"column:wx_nickname;size:100"` //微信昵称
	WxAvatar   string `orm:"column:wx_avatar;size:255"` //微信头像
	WxGender   uint8  `orm:"column:wx_gender"` //1男 2女
	FormId     uint   `orm:"column:form_id"` //对应goouc_xet_zxgj_form表中的id
	Content    string `orm:"column:content;size:2000"` //提交内容
	CreateAt   uint   `orm:"column:create_at"` //提交时间
}

func (t *ZxgjFormUser) TableName() string {
	return "goouc_xmf_zxgj_form_user"
}

func init() {
	orm.RegisterModel(new(ZxgjFormUser))
}

// AddZxgjFormUser insert a new ZxgjFormUser into database and returns
// last inserted Id on success.
func AddZxgjFormUser(m *ZxgjFormUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjFormUserById retrieves ZxgjFormUser by Id. Returns error if
// Id doesn't exist
func GetZxgjFormUserById(id int) (v *ZxgjFormUser, err error) {
	o := orm.NewOrm()
	v = &ZxgjFormUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjFormUser retrieves all ZxgjFormUser matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjFormUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjFormUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ZxgjFormUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateZxgjFormUser updates ZxgjFormUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjFormUserById(m *ZxgjFormUser) (err error) {
	o := orm.NewOrm()
	v := ZxgjFormUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjFormUser deletes ZxgjFormUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjFormUser(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjFormUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjFormUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
