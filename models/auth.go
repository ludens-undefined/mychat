package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Auth struct {
	Id         int     `orm:"column:id;autoIncrement"`
	Name       string  `orm:"column:name;size:30"` //菜单名称
	Pid        uint    `orm:"column:pid"` //父id
	Type       uint8   `orm:"column:type"` //菜单类型;1:有界面可访问菜单(详情),2:无界面可访问菜单（按钮）,0:只作为菜单(左侧栏大标题)
	Status     uint8   `orm:"column:status"` //在左侧栏是否显示，1显示 0不显示
	Order      float32 `orm:"column:order"` //排序
	Icon       string  `orm:"column:icon;size:20"` //图标
	App        string  `orm:"column:app;size:40"` //应用名
	Controller string  `orm:"column:controller;size:50"` //控制器
	Action     string  `orm:"column:action;size:50"` //方法名
	Route      string  `orm:"column:route;size:255"` //前端路由
	Level      uint8   `orm:"column:level"` //级别
	Remark     string  `orm:"column:remark;size:255"` //备注
}

func (t *Auth) TableName() string {
	return "goouc_xmf_auth"
}

func init() {
	orm.RegisterModel(new(Auth))
}

// AddAuth insert a new Auth into database and returns
// last inserted Id on success.
func AddAuth(m *Auth) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAuthById retrieves Auth by Id. Returns error if
// Id doesn't exist
func GetAuthById(id int) (v *Auth, err error) {
	o := orm.NewOrm()
	v = &Auth{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAuth retrieves all Auth matches certain condition. Returns empty list if
// no records exist
func GetAllAuth(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Auth))
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

	var l []Auth
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

// UpdateAuth updates Auth by Id and returns error if
// the record to be updated doesn't exist
func UpdateAuthById(m *Auth) (err error) {
	o := orm.NewOrm()
	v := Auth{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAuth deletes Auth by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAuth(id int) (err error) {
	o := orm.NewOrm()
	v := Auth{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Auth{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
