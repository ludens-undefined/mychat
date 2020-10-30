package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CourseClasstable struct {
	Id          int       `orm:"column:id;autoIncrement"`
	ShopId      int       `orm:"column:shop_id"`
	CourseId    int       `orm:"column:course_id"` //课程id
	ClassId     int       `orm:"column:class_id"` //班级id
	CtableId    int       `orm:"column:ctable_id"` //课表id
	Classdate   time.Time `orm:"column:classdate;type(date)"` //上课日期
	Starttime   time.Time `orm:"column:starttime;type(time)"` //上课时间
	Endtime     time.Time `orm:"column:endtime;type(time)"` //下课时间
	Createtime  int       `orm:"column:createtime"`
	Updatetime  int       `orm:"column:updatetime"`
	Have        int8      `orm:"column:have"` //状态 1未上课 2已上课
	TeacherHave int8      `orm:"column:teacher_have"` //教师考勤 1空白 2正常出勤 3课前迟到 4课后迟到
	Status      int8      `orm:"column:status"` //状态 1启用 2禁用
	IsDelete    int8      `orm:"column:is_delete"` //状态 1启用 2删除
	Oldid       int       `orm:"column:oldid"`
}

func (t *CourseClasstable) TableName() string {
	return "goouc_xmf_course_classtable"
}

func init() {
	orm.RegisterModel(new(CourseClasstable))
}

// AddCourseClasstable insert a new CourseClasstable into database and returns
// last inserted Id on success.
func AddCourseClasstable(m *CourseClasstable) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseClasstableById retrieves CourseClasstable by Id. Returns error if
// Id doesn't exist
func GetCourseClasstableById(id int) (v *CourseClasstable, err error) {
	o := orm.NewOrm()
	v = &CourseClasstable{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseClasstable retrieves all CourseClasstable matches certain condition. Returns empty list if
// no records exist
func GetAllCourseClasstable(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseClasstable))
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

	var l []CourseClasstable
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

// UpdateCourseClasstable updates CourseClasstable by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseClasstableById(m *CourseClasstable) (err error) {
	o := orm.NewOrm()
	v := CourseClasstable{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseClasstable deletes CourseClasstable by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseClasstable(id int) (err error) {
	o := orm.NewOrm()
	v := CourseClasstable{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseClasstable{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
