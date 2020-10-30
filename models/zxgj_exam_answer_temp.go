package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjExamAnswerTemp struct {
	Id         int    `orm:"column:id;autoIncrement"` //ID
	ShopId     int    `orm:"column:shop_id"`
	ExamUserId int    `orm:"column:exam_user_id"` //学员考试记录表id    对应goouc_xmf_exam_user的id
	Memberid   int    `orm:"column:memberid"` //学员id
	ExamId     int    `orm:"column:exam_id"` //考试id，对应goouc_xet_zxgj_exam表中的id
	ExamAnswer string `orm:"column:examAnswer"` //学员答案
	Usetime    int8   `orm:"column:usetime"` //考试用时
	Createtime int    `orm:"column:createtime"`
	IsDelete   int8   `orm:"column:is_delete"` //1:启用    2:删除
}

func (t *ZxgjExamAnswerTemp) TableName() string {
	return "goouc_xmf_zxgj_exam_answer_temp"
}

func init() {
	orm.RegisterModel(new(ZxgjExamAnswerTemp))
}

// AddZxgjExamAnswerTemp insert a new ZxgjExamAnswerTemp into database and returns
// last inserted Id on success.
func AddZxgjExamAnswerTemp(m *ZxgjExamAnswerTemp) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjExamAnswerTempById retrieves ZxgjExamAnswerTemp by Id. Returns error if
// Id doesn't exist
func GetZxgjExamAnswerTempById(id int) (v *ZxgjExamAnswerTemp, err error) {
	o := orm.NewOrm()
	v = &ZxgjExamAnswerTemp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjExamAnswerTemp retrieves all ZxgjExamAnswerTemp matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjExamAnswerTemp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjExamAnswerTemp))
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

	var l []ZxgjExamAnswerTemp
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

// UpdateZxgjExamAnswerTemp updates ZxgjExamAnswerTemp by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjExamAnswerTempById(m *ZxgjExamAnswerTemp) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamAnswerTemp{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjExamAnswerTemp deletes ZxgjExamAnswerTemp by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjExamAnswerTemp(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamAnswerTemp{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjExamAnswerTemp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
