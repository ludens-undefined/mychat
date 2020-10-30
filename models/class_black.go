package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ClassBlack struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     int    `orm:"column:shop_id"`
	Userid     int    `orm:"column:userid"`
	ClassId    int    `orm:"column:class_id"` //班级id
	Createtime int    `orm:"column:createtime"`
	Updatetime int    `orm:"column:updatetime"`
	IsDelete   int8   `orm:"column:is_delete"` //状态 1启用 2删除
	Operator   string `orm:"column:operator;size:32"` //操作人
}

func (t *ClassBlack) TableName() string {
	return "goouc_xmf_class_black"
}

func init() {
	orm.RegisterModel(new(ClassBlack))
}

// AddClassBlack insert a new ClassBlack into database and returns
// last inserted Id on success.
func AddClassBlack(m *ClassBlack) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetClassBlackById retrieves ClassBlack by Id. Returns error if
// Id doesn't exist
func GetClassBlackById(id int) (v *ClassBlack, err error) {
	o := orm.NewOrm()
	v = &ClassBlack{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClassBlack retrieves all ClassBlack matches certain condition. Returns empty list if
// no records exist
func GetAllClassBlack(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ClassBlack))
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

	var l []ClassBlack
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

// UpdateClassBlack updates ClassBlack by Id and returns error if
// the record to be updated doesn't exist
func UpdateClassBlackById(m *ClassBlack) (err error) {
	o := orm.NewOrm()
	v := ClassBlack{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClassBlack deletes ClassBlack by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClassBlack(id int) (err error) {
	o := orm.NewOrm()
	v := ClassBlack{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ClassBlack{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
