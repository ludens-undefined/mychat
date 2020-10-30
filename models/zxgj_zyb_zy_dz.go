package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZybZyDz struct {
	Id         int    `orm:"column:id;autoIncrement"`
	UserId     uint   `orm:"column:user_id"` //点赞者id，对应goouc_xet_user表中的id
	WxNickname string `orm:"column:wx_nickname;size:100"` //微信昵称，对应goouc_xet_user表中的wx_nickname
	WxAvatar   string `orm:"column:wx_avatar;size:255"` //微信头像，对应goouc_xet_user表中的wx_avatar
	ZyId       uint   `orm:"column:zy_id"` //作业id，对应goouc_xet_zxgj_zyb_zy表中的id
	IsValid    uint8  `orm:"column:is_valid"` //1点赞 2取消点赞
	ZDate      uint   `orm:"column:z_date"` //点赞时间
	QDate      uint   `orm:"column:q_date"` //取消时间
}

func (t *ZxgjZybZyDz) TableName() string {
	return "goouc_xmf_zxgj_zyb_zy_dz"
}

func init() {
	orm.RegisterModel(new(ZxgjZybZyDz))
}

// AddZxgjZybZyDz insert a new ZxgjZybZyDz into database and returns
// last inserted Id on success.
func AddZxgjZybZyDz(m *ZxgjZybZyDz) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZybZyDzById retrieves ZxgjZybZyDz by Id. Returns error if
// Id doesn't exist
func GetZxgjZybZyDzById(id int) (v *ZxgjZybZyDz, err error) {
	o := orm.NewOrm()
	v = &ZxgjZybZyDz{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZybZyDz retrieves all ZxgjZybZyDz matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZybZyDz(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZybZyDz))
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

	var l []ZxgjZybZyDz
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

// UpdateZxgjZybZyDz updates ZxgjZybZyDz by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZybZyDzById(m *ZxgjZybZyDz) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZyDz{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZybZyDz deletes ZxgjZybZyDz by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZybZyDz(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZyDz{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZybZyDz{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
