package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjRldkDz struct {
	Id         int    `orm:"column:id;autoIncrement"`
	UserId     uint   `orm:"column:user_id"` //点赞者id，对应goouc_xet_user表中的id
	WxNickname string `orm:"column:wx_nickname;size:100"` //微信昵称，对应goouc_xet_user表中的wx_nickname
	WxAvatar   string `orm:"column:wx_avatar;size:255"` //微信头像，对应goouc_xet_user表中的wx_avatar
	DiaryId    uint   `orm:"column:diary_id"` //日记id，对应goouc_xet_zxgj_rldk_diary表中的id
	IsValid    uint8  `orm:"column:is_valid"` //1点赞 2取消点赞
	ZDate      uint   `orm:"column:z_date"` //点赞时间
	QDate      uint   `orm:"column:q_date"` //取消时间
}

func (t *ZxgjRldkDz) TableName() string {
	return "goouc_xmf_zxgj_rldk_dz"
}

func init() {
	orm.RegisterModel(new(ZxgjRldkDz))
}

// AddZxgjRldkDz insert a new ZxgjRldkDz into database and returns
// last inserted Id on success.
func AddZxgjRldkDz(m *ZxgjRldkDz) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjRldkDzById retrieves ZxgjRldkDz by Id. Returns error if
// Id doesn't exist
func GetZxgjRldkDzById(id int) (v *ZxgjRldkDz, err error) {
	o := orm.NewOrm()
	v = &ZxgjRldkDz{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjRldkDz retrieves all ZxgjRldkDz matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjRldkDz(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjRldkDz))
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

	var l []ZxgjRldkDz
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

// UpdateZxgjRldkDz updates ZxgjRldkDz by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjRldkDzById(m *ZxgjRldkDz) (err error) {
	o := orm.NewOrm()
	v := ZxgjRldkDz{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjRldkDz deletes ZxgjRldkDz by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjRldkDz(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjRldkDz{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjRldkDz{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
