package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ActivityFormField struct {
	Id             int    `orm:"column:id;autoIncrement"`
	ShopId         int    `orm:"column:shop_id"`
	FormId         int    `orm:"column:form_id"` //活动id
	Code           string `orm:"column:code;size:32"` //字段类型
	Name           string `orm:"column:name;size:32"` //名称
	Customizedtype string `orm:"column:customizedtype"` //自定义类型
	Sort           int    `orm:"column:sort"` //排序 由小到大
	Choices        string `orm:"column:choices"` //选项
	Type           int8   `orm:"column:type"`
	Required       int8   `orm:"column:required"`
	IsDelete       int8   `orm:"column:is_delete"`
}

func (t *ActivityFormField) TableName() string {
	return "goouc_xmf_activity_form_field"
}

func init() {
	orm.RegisterModel(new(ActivityFormField))
}

// AddActivityFormField insert a new ActivityFormField into database and returns
// last inserted Id on success.
func AddActivityFormField(m *ActivityFormField) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetActivityFormFieldById retrieves ActivityFormField by Id. Returns error if
// Id doesn't exist
func GetActivityFormFieldById(id int) (v *ActivityFormField, err error) {
	o := orm.NewOrm()
	v = &ActivityFormField{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllActivityFormField retrieves all ActivityFormField matches certain condition. Returns empty list if
// no records exist
func GetAllActivityFormField(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ActivityFormField))
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

	var l []ActivityFormField
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

// UpdateActivityFormField updates ActivityFormField by Id and returns error if
// the record to be updated doesn't exist
func UpdateActivityFormFieldById(m *ActivityFormField) (err error) {
	o := orm.NewOrm()
	v := ActivityFormField{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteActivityFormField deletes ActivityFormField by Id and returns error if
// the record to be deleted doesn't exist
func DeleteActivityFormField(id int) (err error) {
	o := orm.NewOrm()
	v := ActivityFormField{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ActivityFormField{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
