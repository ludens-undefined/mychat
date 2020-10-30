package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Relative struct {
	Id          int     `orm:"column:id;autoIncrement"`
	ShopId      uint    `orm:"column:shop_id"`
	TypeId      uint8   `orm:"column:type_id"` //关联者类型,对应goouc_xmf_course_type表中的id
	RelativeId  uint    `orm:"column:relative_id"` //关联者id,对应课程管理中的标签，例如gooux_xmf_course_imgtxt表中的id
	Name        string  `orm:"column:name;size:50"` //对应type_id,relative_id表中的name
	Img         string  `orm:"column:img;size:255"` //对应type_id,relative_id中的img
	Price       float64 `orm:"column:price;scale:10;precision:2"` //价格
	TypedId     uint8   `orm:"column:typed_id"` //被关联者类型,对应goouc_xmf_course_type表中的id
	RelativedId uint    `orm:"column:relatived_id"` //被关联者id，对应课程管理中的标签，例如gooux_xmf_course_imgtxt表中的id
	Named       string  `orm:"column:named;size:50"` //对应typed_id,relatived_id中的name
	Imged       string  `orm:"column:imged;size:255"` //对应typed_id,relatived_id表中的img
	Priced      float64 `orm:"column:priced;scale:10;precision:2"` //价格
	IsDelete    uint8   `orm:"column:is_delete"` //状态 1启用 2删除
}

func (t *Relative) TableName() string {
	return "goouc_xmf_relative"
}

func init() {
	orm.RegisterModel(new(Relative))
}

// AddRelative insert a new Relative into database and returns
// last inserted Id on success.
func AddRelative(m *Relative) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRelativeById retrieves Relative by Id. Returns error if
// Id doesn't exist
func GetRelativeById(id int) (v *Relative, err error) {
	o := orm.NewOrm()
	v = &Relative{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRelative retrieves all Relative matches certain condition. Returns empty list if
// no records exist
func GetAllRelative(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Relative))
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

	var l []Relative
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

// UpdateRelative updates Relative by Id and returns error if
// the record to be updated doesn't exist
func UpdateRelativeById(m *Relative) (err error) {
	o := orm.NewOrm()
	v := Relative{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRelative deletes Relative by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRelative(id int) (err error) {
	o := orm.NewOrm()
	v := Relative{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Relative{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
