package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DiagnoseExamAppraise struct {
	Id          int    `orm:"column:id;autoIncrement"`
	ShopId      int    `orm:"column:shop_id"`
	ExamId      int    `orm:"column:exam_id"` //报告id
	KnowledgeId int    `orm:"column:knowledge_id"` //一级知识点id
	MinScore    int    `orm:"column:min_score"` //起点分数
	MaxScore    int    `orm:"column:max_score"` //终点分数
	Content     string `orm:"column:content"` //评价
	Status      int8   `orm:"column:status"` //状态 1显示 2隐藏
	IsDelete    int8   `orm:"column:is_delete"`
}

func (t *DiagnoseExamAppraise) TableName() string {
	return "goouc_xmf_diagnose_exam_appraise"
}

func init() {
	orm.RegisterModel(new(DiagnoseExamAppraise))
}

// AddDiagnoseExamAppraise insert a new DiagnoseExamAppraise into database and returns
// last inserted Id on success.
func AddDiagnoseExamAppraise(m *DiagnoseExamAppraise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDiagnoseExamAppraiseById retrieves DiagnoseExamAppraise by Id. Returns error if
// Id doesn't exist
func GetDiagnoseExamAppraiseById(id int) (v *DiagnoseExamAppraise, err error) {
	o := orm.NewOrm()
	v = &DiagnoseExamAppraise{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDiagnoseExamAppraise retrieves all DiagnoseExamAppraise matches certain condition. Returns empty list if
// no records exist
func GetAllDiagnoseExamAppraise(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DiagnoseExamAppraise))
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

	var l []DiagnoseExamAppraise
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

// UpdateDiagnoseExamAppraise updates DiagnoseExamAppraise by Id and returns error if
// the record to be updated doesn't exist
func UpdateDiagnoseExamAppraiseById(m *DiagnoseExamAppraise) (err error) {
	o := orm.NewOrm()
	v := DiagnoseExamAppraise{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDiagnoseExamAppraise deletes DiagnoseExamAppraise by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDiagnoseExamAppraise(id int) (err error) {
	o := orm.NewOrm()
	v := DiagnoseExamAppraise{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DiagnoseExamAppraise{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
