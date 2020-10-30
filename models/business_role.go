package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type BusinessRole struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	BusinessId uint   `orm:"column:business_id"` //对应goouc_xet_business表中的id，后台用户id
	RoleId     string `orm:"column:role_id;size:255"` //角色id,对应goouc_xet_role表中的id，{1,2}
	Phone      string `orm:"column:phone;size:11"` //小鹅通账号，对应goouc_xet_business表中的phone
	Name       string `orm:"column:name;size:30"` //姓名
	Img        string `orm:"column:img;size:255"` //头像
	Decription string `orm:"column:decription;size:50"` //描述
	Brief      string `orm:"column:brief;size:255"` //简介
	State      uint8  `orm:"column:state"` //状态设置 1启用 2停用
}

func (t *BusinessRole) TableName() string {
	return "goouc_xmf_business_role"
}

func init() {
	orm.RegisterModel(new(BusinessRole))
}

// AddBusinessRole insert a new BusinessRole into database and returns
// last inserted Id on success.
func AddBusinessRole(m *BusinessRole) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBusinessRoleById retrieves BusinessRole by Id. Returns error if
// Id doesn't exist
func GetBusinessRoleById(id int) (v *BusinessRole, err error) {
	o := orm.NewOrm()
	v = &BusinessRole{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBusinessRole retrieves all BusinessRole matches certain condition. Returns empty list if
// no records exist
func GetAllBusinessRole(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BusinessRole))
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

	var l []BusinessRole
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

// UpdateBusinessRole updates BusinessRole by Id and returns error if
// the record to be updated doesn't exist
func UpdateBusinessRoleById(m *BusinessRole) (err error) {
	o := orm.NewOrm()
	v := BusinessRole{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBusinessRole deletes BusinessRole by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBusinessRole(id int) (err error) {
	o := orm.NewOrm()
	v := BusinessRole{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BusinessRole{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
