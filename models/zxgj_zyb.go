package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZyb struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     uint   `orm:"column:shop_id"`
	Name       string `orm:"column:name;size:50"` //作业本名称
	Xhcourseid int    `orm:"column:xhcourseid"` //关联的课程id
	Teacherid  int    `orm:"column:teacherid"` //教师id
	CreateAt   uint   `orm:"column:create_at"`
	UpdateAt   uint   `orm:"column:update_at"`
	IsDelete   int8   `orm:"column:is_delete"` //1:启用   2：删除
}

func (t *ZxgjZyb) TableName() string {
	return "goouc_xmf_zxgj_zyb"
}

func init() {
	orm.RegisterModel(new(ZxgjZyb))
}

// AddZxgjZyb insert a new ZxgjZyb into database and returns
// last inserted Id on success.
func AddZxgjZyb(m *ZxgjZyb) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZybById retrieves ZxgjZyb by Id. Returns error if
// Id doesn't exist
func GetZxgjZybById(id int) (v *ZxgjZyb, err error) {
	o := orm.NewOrm()
	v = &ZxgjZyb{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZyb retrieves all ZxgjZyb matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZyb(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZyb))
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

	var l []ZxgjZyb
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

// UpdateZxgjZyb updates ZxgjZyb by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZybById(m *ZxgjZyb) (err error) {
	o := orm.NewOrm()
	v := ZxgjZyb{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZyb deletes ZxgjZyb by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZyb(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZyb{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZyb{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
