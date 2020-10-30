package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjExamTk struct {
	Id         int     `orm:"column:id;autoIncrement"`
	ShopId     int     `orm:"column:shop_id"`
	ExamId     uint    `orm:"column:exam_id"` //考试id，对应goouc_xet_zxgj_exam表中的id
	TkTestId   uint    `orm:"column:tk_test_id"` //试题id，对应goouc_xmf_zxgj_tk_test表中的id
	Score      float64 `orm:"column:score;scale:10;precision:1"` //该题总分值
	LxScore    float64 `orm:"column:lx_score;scale:10;precision:1"` //漏选分值
	BlankScore float64 `orm:"column:blank_score;scale:10;precision:1"` //填空题，每个空分值
	Order      uint16  `orm:"column:order"` //排序
	IsDelete   uint8   `orm:"column:is_delete"` //状态 1启用 2删除
}

func (t *ZxgjExamTk) TableName() string {
	return "goouc_xmf_zxgj_exam_tk"
}

func init() {
	orm.RegisterModel(new(ZxgjExamTk))
}

// AddZxgjExamTk insert a new ZxgjExamTk into database and returns
// last inserted Id on success.
func AddZxgjExamTk(m *ZxgjExamTk) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjExamTkById retrieves ZxgjExamTk by Id. Returns error if
// Id doesn't exist
func GetZxgjExamTkById(id int) (v *ZxgjExamTk, err error) {
	o := orm.NewOrm()
	v = &ZxgjExamTk{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjExamTk retrieves all ZxgjExamTk matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjExamTk(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjExamTk))
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

	var l []ZxgjExamTk
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

// UpdateZxgjExamTk updates ZxgjExamTk by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjExamTkById(m *ZxgjExamTk) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamTk{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjExamTk deletes ZxgjExamTk by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjExamTk(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamTk{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjExamTk{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
