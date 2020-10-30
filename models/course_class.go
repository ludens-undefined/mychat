package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CourseClass struct {
	Id            int       `orm:"column:id;autoIncrement"`
	ShopId        int       `orm:"column:shop_id"`
	CourseId      int       `orm:"column:course_id"` //课程id
	Title         string    `orm:"column:title;size:32"` //班级名称
	Class         int16     `orm:"column:class"` //班型
	School        int16     `orm:"column:school"` //校区
	Classroom     int16     `orm:"column:classroom"` //教室
	Teacher       int16     `orm:"column:teacher"` //教师
	Week          int       `orm:"column:week"` //星期
	BeginDate     time.Time `orm:"column:begin_date;type(date)"` //开班日期
	EndDate       time.Time `orm:"column:end_date;type(date)"` //结课日期
	Number        int       `orm:"column:number"` //限额
	BeginTime     int       `orm:"column:begin_time"` //报名开始时间
	EndTime       int       `orm:"column:end_time"` //报名截至时间
	DownBeginTime int       `orm:"column:down_begin_time"` //报名开始时间
	DownEndTime   int       `orm:"column:down_end_time"` //报名截至时间
	Starttime     string    `orm:"column:starttime;size:32"` //上课时间
	Endtime       string    `orm:"column:endtime;size:32"` //下课时间
	Reserve       int       `orm:"column:reserve"` //预留时间 单位s
	Sort          int       `orm:"column:sort"` //排序
	Status        int8      `orm:"column:status"` //状态 1启用 2禁用
	Createtime    int       `orm:"column:createtime"`
	Updatetime    int       `orm:"column:updatetime"`
	IsDelete      int8      `orm:"column:is_delete"` //状态 1启用 2删除
	IsChange      int8      `orm:"column:is_change"`
	IsRefund      int8      `orm:"column:is_refund"`
	IsAdjust      int8      `orm:"column:is_adjust"`
	IsOnlinepay   int8      `orm:"column:is_onlinepay"`
	Pid           int       `orm:"column:pid"`
	IsShow        int8      `orm:"column:is_show"` //系统状态 1显示 2隐藏
}

func (t *CourseClass) TableName() string {
	return "goouc_xmf_course_class"
}

func init() {
	orm.RegisterModel(new(CourseClass))
}

// AddCourseClass insert a new CourseClass into database and returns
// last inserted Id on success.
func AddCourseClass(m *CourseClass) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseClassById retrieves CourseClass by Id. Returns error if
// Id doesn't exist
func GetCourseClassById(id int) (v *CourseClass, err error) {
	o := orm.NewOrm()
	v = &CourseClass{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseClass retrieves all CourseClass matches certain condition. Returns empty list if
// no records exist
func GetAllCourseClass(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseClass))
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

	var l []CourseClass
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

// UpdateCourseClass updates CourseClass by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseClassById(m *CourseClass) (err error) {
	o := orm.NewOrm()
	v := CourseClass{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseClass deletes CourseClass by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseClass(id int) (err error) {
	o := orm.NewOrm()
	v := CourseClass{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseClass{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
