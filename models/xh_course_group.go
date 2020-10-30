package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhCourseGroup struct {
	Id         int  `orm:"column:id;autoIncrement"`
	Memberid   int  `orm:"column:memberid"` //学员id
	Xhcourseid int  `orm:"column:xhcourseid"` //课程id
	Groupnum   int  `orm:"column:groupnum"` //需要人数
	Status     int8 `orm:"column:status"` //团状态 默认1 拼团成功 2 拼团未完成 3团已过期4:拼团退款
	Ctime      int  `orm:"column:ctime"` //开团时间
	OverTime   int  `orm:"column:over_time"` //过期时间
	ShopId     int  `orm:"column:shop_id"` //店铺id
	MemberNum  int  `orm:"column:member_num"` //已存在成员数量
	Xoldid     int  `orm:"column:xoldid"`
}

func (t *XhCourseGroup) TableName() string {
	return "goouc_xmf_xh_course_group"
}

func init() {
	orm.RegisterModel(new(XhCourseGroup))
}

// AddXhCourseGroup insert a new XhCourseGroup into database and returns
// last inserted Id on success.
func AddXhCourseGroup(m *XhCourseGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhCourseGroupById retrieves XhCourseGroup by Id. Returns error if
// Id doesn't exist
func GetXhCourseGroupById(id int) (v *XhCourseGroup, err error) {
	o := orm.NewOrm()
	v = &XhCourseGroup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhCourseGroup retrieves all XhCourseGroup matches certain condition. Returns empty list if
// no records exist
func GetAllXhCourseGroup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhCourseGroup))
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

	var l []XhCourseGroup
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

// UpdateXhCourseGroup updates XhCourseGroup by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhCourseGroupById(m *XhCourseGroup) (err error) {
	o := orm.NewOrm()
	v := XhCourseGroup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhCourseGroup deletes XhCourseGroup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhCourseGroup(id int) (err error) {
	o := orm.NewOrm()
	v := XhCourseGroup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhCourseGroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
