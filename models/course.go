package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id           int     `orm:"column:id;autoIncrement"`
	ShopId       int     `orm:"column:shop_id"`
	Title        string  `orm:"column:title;size:255"` //课程标题
	Year         int16   `orm:"column:year"` //年限
	Term         int16   `orm:"column:term"` //学期
	Sort         int     `orm:"column:sort"` //排序
	SchoolType   int8    `orm:"column:school_type"` //学校类型 1小学 2初中
	Grade        int16   `orm:"column:grade"` //年级
	Subject      int16   `orm:"column:subject"` //学科
	Price        float64 `orm:"column:price;scale:10;precision:2"` //课时单价
	DelPrice     float64 `orm:"column:del_price"`
	IsMaterial   int8    `orm:"column:is_material"` //材料费 1有 2无
	MaterialFee  float64 `orm:"column:material_fee;scale:10;precision:2"` //资料费
	Classnum     int16   `orm:"column:classnum"` //课时
	Downclassnum int16   `orm:"column:downclassnum"` //新增课时
	Synopsis     string  `orm:"column:synopsis;size:2000"`
	Content      string  `orm:"column:content"`
	Createtime   int     `orm:"column:createtime"`
	Updatetime   int     `orm:"column:updatetime"`
	Status       int8    `orm:"column:status"` //状态 1启用 2禁用
	IsDelete     int8    `orm:"column:is_delete"` //状态 1启用 2删除
	IsStage      int8    `orm:"column:is_stage"` //需要续费 1全款 2续费
	IsExam       int8    `orm:"column:is_exam"` //需要考试 1需要 2不需要
	IsLimit      int8    `orm:"column:is_limit"` //限制只能报一班 1是 2否
	IsFull       int8    `orm:"column:is_full"` //是否已排课 1已排课
	Oldid        int     `orm:"column:oldid"`
}

func (t *Course) TableName() string {
	return "goouc_xmf_course"
}

func init() {
	orm.RegisterModel(new(Course))
}

// AddCourse insert a new Course into database and returns
// last inserted Id on success.
func AddCourse(m *Course) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseById retrieves Course by Id. Returns error if
// Id doesn't exist
func GetCourseById(id int) (v *Course, err error) {
	o := orm.NewOrm()
	v = &Course{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourse retrieves all Course matches certain condition. Returns empty list if
// no records exist
func GetAllCourse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Course))
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

	var l []Course
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

// UpdateCourse updates Course by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseById(m *Course) (err error) {
	o := orm.NewOrm()
	v := Course{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourse deletes Course by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourse(id int) (err error) {
	o := orm.NewOrm()
	v := Course{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Course{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
