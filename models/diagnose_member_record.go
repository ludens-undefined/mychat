package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DiagnoseMemberRecord struct {
	Id        int    `orm:"column:id;autoIncrement"`
	ShopId    int    `orm:"column:shop_id"`
	ExamId    int    `orm:"column:exam_id"` //报告id
	Userid    int    `orm:"column:userid"` //用户id
	Score     int    `orm:"column:score"` //总分
	ClassName string `orm:"column:class_name;size:255"` //推荐班型
	Status    int8   `orm:"column:status"` //状态 1显示 2隐藏
	IsDelete  int8   `orm:"column:is_delete"`
}

func (t *DiagnoseMemberRecord) TableName() string {
	return "goouc_xmf_diagnose_member_record"
}

func init() {
	orm.RegisterModel(new(DiagnoseMemberRecord))
}

// AddDiagnoseMemberRecord insert a new DiagnoseMemberRecord into database and returns
// last inserted Id on success.
func AddDiagnoseMemberRecord(m *DiagnoseMemberRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDiagnoseMemberRecordById retrieves DiagnoseMemberRecord by Id. Returns error if
// Id doesn't exist
func GetDiagnoseMemberRecordById(id int) (v *DiagnoseMemberRecord, err error) {
	o := orm.NewOrm()
	v = &DiagnoseMemberRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDiagnoseMemberRecord retrieves all DiagnoseMemberRecord matches certain condition. Returns empty list if
// no records exist
func GetAllDiagnoseMemberRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DiagnoseMemberRecord))
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

	var l []DiagnoseMemberRecord
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

// UpdateDiagnoseMemberRecord updates DiagnoseMemberRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateDiagnoseMemberRecordById(m *DiagnoseMemberRecord) (err error) {
	o := orm.NewOrm()
	v := DiagnoseMemberRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDiagnoseMemberRecord deletes DiagnoseMemberRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDiagnoseMemberRecord(id int) (err error) {
	o := orm.NewOrm()
	v := DiagnoseMemberRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DiagnoseMemberRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
