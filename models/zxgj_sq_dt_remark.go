package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjSqDtRemark struct {
	Id         int    `orm:"column:id;autoIncrement"`
	DtId       uint   `orm:"column:dt_id"` //动态id，对应goouc_xet_sq_dt表中的id
	UserId     uint   `orm:"column:user_id"` //评论者id，对应goouc_xet_user表中的id
	WxNickname string `orm:"column:wx_nickname;size:100"` //微信昵称
	WxAvatar   string `orm:"column:wx_avatar;size:255"` //微信头像
	Remark     string `orm:"column:remark;size:255"` //评论内容
	CreateAt   uint   `orm:"column:create_at"` //评论时间
}

func (t *ZxgjSqDtRemark) TableName() string {
	return "goouc_xmf_zxgj_sq_dt_remark"
}

func init() {
	orm.RegisterModel(new(ZxgjSqDtRemark))
}

// AddZxgjSqDtRemark insert a new ZxgjSqDtRemark into database and returns
// last inserted Id on success.
func AddZxgjSqDtRemark(m *ZxgjSqDtRemark) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjSqDtRemarkById retrieves ZxgjSqDtRemark by Id. Returns error if
// Id doesn't exist
func GetZxgjSqDtRemarkById(id int) (v *ZxgjSqDtRemark, err error) {
	o := orm.NewOrm()
	v = &ZxgjSqDtRemark{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjSqDtRemark retrieves all ZxgjSqDtRemark matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjSqDtRemark(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjSqDtRemark))
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

	var l []ZxgjSqDtRemark
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

// UpdateZxgjSqDtRemark updates ZxgjSqDtRemark by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjSqDtRemarkById(m *ZxgjSqDtRemark) (err error) {
	o := orm.NewOrm()
	v := ZxgjSqDtRemark{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjSqDtRemark deletes ZxgjSqDtRemark by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjSqDtRemark(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjSqDtRemark{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjSqDtRemark{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
