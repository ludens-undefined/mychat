package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjExamFalseRecord struct {
	Id         int    `orm:"column:id;autoIncrement"` //ID
	ShopId     int    `orm:"column:shop_id"` //店铺id
	ZxgjType   int8   `orm:"column:zxgj_type"` //助学工具分类：考试，作业本，练习
	Memberid   int    `orm:"column:memberid"` //学员id
	ExamId     int    `orm:"column:exam_id"` //试卷id  goouc_xmf_zxgj_exam
	ExamUserId int    `orm:"column:exam_user_id"` //学员本次考试记录主表id   goouc_xmf_zxgj_exam_user
	TkTestId   int    `orm:"column:tk_test_id"` //题库中该题id
	ExamTkId   int    `orm:"column:exam_tk_id"` //试卷中该题id
	MAnswer    string `orm:"column:m_answer;size:1000"` //会员答案
	TrueAnswer string `orm:"column:true_answer;size:1000"` //真实答案
	Comment    string `orm:"column:comment;size:2000"` //评语
	Createtime int    `orm:"column:createtime"` //创建时间
	IsDelete   int8   `orm:"column:is_delete"` //1:启用   2：删除
}

func (t *ZxgjExamFalseRecord) TableName() string {
	return "goouc_xmf_zxgj_exam_false_record"
}

func init() {
	orm.RegisterModel(new(ZxgjExamFalseRecord))
}

// AddZxgjExamFalseRecord insert a new ZxgjExamFalseRecord into database and returns
// last inserted Id on success.
func AddZxgjExamFalseRecord(m *ZxgjExamFalseRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjExamFalseRecordById retrieves ZxgjExamFalseRecord by Id. Returns error if
// Id doesn't exist
func GetZxgjExamFalseRecordById(id int) (v *ZxgjExamFalseRecord, err error) {
	o := orm.NewOrm()
	v = &ZxgjExamFalseRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjExamFalseRecord retrieves all ZxgjExamFalseRecord matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjExamFalseRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjExamFalseRecord))
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

	var l []ZxgjExamFalseRecord
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

// UpdateZxgjExamFalseRecord updates ZxgjExamFalseRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjExamFalseRecordById(m *ZxgjExamFalseRecord) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamFalseRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjExamFalseRecord deletes ZxgjExamFalseRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjExamFalseRecord(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamFalseRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjExamFalseRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
