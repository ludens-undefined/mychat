package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DiagnoseExamQuestion struct {
	Id          int    `orm:"column:id;autoIncrement"`
	ShopId      int    `orm:"column:shop_id"`
	ExamId      int    `orm:"column:exam_id"` //报告id
	Subject     int    `orm:"column:subject"` //科目
	KnowledgeId int    `orm:"column:knowledge_id"` //一级知识点id
	Knowledge   string `orm:"column:knowledge;size:255"` //题目知识点
	Score       int    `orm:"column:score"` //分数
	Sort        int    `orm:"column:sort"` //题号
	Status      int8   `orm:"column:status"` //状态 1显示 2隐藏
	IsDelete    int8   `orm:"column:is_delete"`
}

func (t *DiagnoseExamQuestion) TableName() string {
	return "goouc_xmf_diagnose_exam_question"
}

func init() {
	orm.RegisterModel(new(DiagnoseExamQuestion))
}

// AddDiagnoseExamQuestion insert a new DiagnoseExamQuestion into database and returns
// last inserted Id on success.
func AddDiagnoseExamQuestion(m *DiagnoseExamQuestion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDiagnoseExamQuestionById retrieves DiagnoseExamQuestion by Id. Returns error if
// Id doesn't exist
func GetDiagnoseExamQuestionById(id int) (v *DiagnoseExamQuestion, err error) {
	o := orm.NewOrm()
	v = &DiagnoseExamQuestion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDiagnoseExamQuestion retrieves all DiagnoseExamQuestion matches certain condition. Returns empty list if
// no records exist
func GetAllDiagnoseExamQuestion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DiagnoseExamQuestion))
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

	var l []DiagnoseExamQuestion
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

// UpdateDiagnoseExamQuestion updates DiagnoseExamQuestion by Id and returns error if
// the record to be updated doesn't exist
func UpdateDiagnoseExamQuestionById(m *DiagnoseExamQuestion) (err error) {
	o := orm.NewOrm()
	v := DiagnoseExamQuestion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDiagnoseExamQuestion deletes DiagnoseExamQuestion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDiagnoseExamQuestion(id int) (err error) {
	o := orm.NewOrm()
	v := DiagnoseExamQuestion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DiagnoseExamQuestion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
