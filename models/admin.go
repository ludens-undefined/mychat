package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int    `orm:"column:id;autoIncrement"`
	Type       int8   `orm:"column:type"` //1店长 2员工
	Phone      string `orm:"column:phone;size:32"` //手机号
	Username   string `orm:"column:username;size:32"` //昵称
	Avatar     string `orm:"column:avatar;size:255"` //头像
	Password   string `orm:"column:password;size:255"` //用户登录密码
	Createtime uint   `orm:"column:createtime"` //注册时间
	Logintime  uint   `orm:"column:logintime"` //登录时间
	Status     int8   `orm:"column:status"` //状态 1启用 2禁用
	RoleId     int    `orm:"column:role_id"` //角色分组
	Pid        int    `orm:"column:pid"` //父级
	IsDelete   int8   `orm:"column:is_delete"` //状态 1启用 2删除
}

func (t *Admin) TableName() string {
	return "goouc_xmf_admin"
}

func init() {
	orm.RegisterModel(new(Admin))
}

// AddAdmin insert a new Admin into database and returns
// last inserted Id on success.
func AddAdmin(m *Admin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdminById retrieves Admin by Id. Returns error if
// Id doesn't exist
func GetAdminById(id int) (v *Admin, err error) {
	o := orm.NewOrm()
	v = &Admin{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdmin retrieves all Admin matches certain condition. Returns empty list if
// no records exist
func GetAllAdmin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Admin))
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

	var l []Admin
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

// UpdateAdmin updates Admin by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminById(m *Admin) (err error) {
	o := orm.NewOrm()
	v := Admin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdmin deletes Admin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdmin(id int) (err error) {
	o := orm.NewOrm()
	v := Admin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Admin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
