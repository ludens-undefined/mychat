package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjRelative struct {
	Id          int   `orm:"column:id;autoIncrement"`
	ShopId      uint  `orm:"column:shop_id"`
	TypeId      uint  `orm:"column:type_id"` //助学工具type,对应goouc_xmf_course_type表中的id
	RelativeId  uint  `orm:"column:relative_id"` //关联者id,对应课程管理中的标签，例如gooux_xmf_course_imgtxt表中的id
	TypedId     uint  `orm:"column:typed_id"` //关联课程类型,对应goouc_xmf_course_type表中的id，如图文等
	RelativedId uint  `orm:"column:relatived_id"` //被关联者id
	IsDelete    uint8 `orm:"column:is_delete"` //1启用 2禁用
}

func (t *ZxgjRelative) TableName() string {
	return "goouc_xmf_zxgj_relative"
}

func init() {
	orm.RegisterModel(new(ZxgjRelative))
}

// AddZxgjRelative insert a new ZxgjRelative into database and returns
// last inserted Id on success.
func AddZxgjRelative(m *ZxgjRelative) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjRelativeById retrieves ZxgjRelative by Id. Returns error if
// Id doesn't exist
func GetZxgjRelativeById(id int) (v *ZxgjRelative, err error) {
	o := orm.NewOrm()
	v = &ZxgjRelative{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjRelative retrieves all ZxgjRelative matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjRelative(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjRelative))
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

	var l []ZxgjRelative
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

// UpdateZxgjRelative updates ZxgjRelative by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjRelativeById(m *ZxgjRelative) (err error) {
	o := orm.NewOrm()
	v := ZxgjRelative{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjRelative deletes ZxgjRelative by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjRelative(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjRelative{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjRelative{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
