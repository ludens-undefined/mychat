package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjTestRes struct {
	Id          int    `orm:"column:id;autoIncrement"`
	TestId      uint   `orm:"column:test_id"` //测试id,对应goouc_xet_zxgj_test表中的id
	Type        uint8  `orm:"column:type"` //1结果 2其他结果
	LowerLimit  uint8  `orm:"column:lower_limit"` //判断条件 最小值
	UpperLimit  uint8  `orm:"column:upper_limit"` //判断条件 最大值
	ResultTitle string `orm:"column:result_title;size:50"` //评测结果
	Result      string `orm:"column:result;size:255"` //评测结果
	Class       string `orm:"column:class;size:50"` //推荐课程{type:,id:}type对应goouc_xet_cource_type表中的id
}

func (t *ZxgjTestRes) TableName() string {
	return "goouc_xmf_zxgj_test_res"
}

func init() {
	orm.RegisterModel(new(ZxgjTestRes))
}

// AddZxgjTestRes insert a new ZxgjTestRes into database and returns
// last inserted Id on success.
func AddZxgjTestRes(m *ZxgjTestRes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjTestResById retrieves ZxgjTestRes by Id. Returns error if
// Id doesn't exist
func GetZxgjTestResById(id int) (v *ZxgjTestRes, err error) {
	o := orm.NewOrm()
	v = &ZxgjTestRes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjTestRes retrieves all ZxgjTestRes matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjTestRes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjTestRes))
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

	var l []ZxgjTestRes
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

// UpdateZxgjTestRes updates ZxgjTestRes by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjTestResById(m *ZxgjTestRes) (err error) {
	o := orm.NewOrm()
	v := ZxgjTestRes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjTestRes deletes ZxgjTestRes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjTestRes(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjTestRes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjTestRes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
