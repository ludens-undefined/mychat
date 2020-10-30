package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id           int    `orm:"column:id;autoIncrement"`
	ShopId       int    `orm:"column:shop_id"`
	Name         string `orm:"column:name;size:32"` //姓名
	Tag          string `orm:"column:tag;size:32"` //姓名首字母
	Spell        string `orm:"column:spell;size:32"` //教学定位
	Sex          int8   `orm:"column:sex"` //状态 1男生 2女生
	Mobile       string `orm:"column:mobile;size:32"` //电话
	Subject      int    `orm:"column:subject"` //学科
	SchoolType   int8   `orm:"column:school_type"` //学校类型 1小学 2初中
	Grade        int    `orm:"column:grade"` //年级
	Motto        string `orm:"column:motto;size:255"` //座右铭
	Avatarurl    string `orm:"column:avatarurl;size:255"`
	Content      string `orm:"column:content"` //详细介绍
	Graduate     string `orm:"column:graduate;size:255"` //毕业学校
	Userdata     string `orm:"column:userdata;size:5000"` //个人数据
	Experience   string `orm:"column:experience;size:5000"` //教学经历
	Teacherstyle string `orm:"column:teacherstyle;size:5000"` //教学风格
	Status       int8   `orm:"column:status"` //状态 1显示 2隐藏
	IsHot        int8   `orm:"column:is_hot"` //状态 1普通 2推荐
	Sort         int    `orm:"column:sort"` //排序
	Createtime   int    `orm:"column:createtime"`
	Updatetime   int    `orm:"column:updatetime"`
	IsDelete     int8   `orm:"column:is_delete"` //状态 1启用 2删除
	Cityid       int    `orm:"column:cityid"` //讲师城市一对一
	Label        string `orm:"column:label"` //标签
	Oldid        int    `orm:"column:oldid"`
	Xoldid       int    `orm:"column:xoldid"`
	Password     string `orm:"column:password;size:255"`
	Logintime    int    `orm:"column:logintime"`
}

func (t *Teacher) TableName() string {
	return "goouc_xmf_teacher"
}

func init() {
	orm.RegisterModel(new(Teacher))
}

// AddTeacher insert a new Teacher into database and returns
// last inserted Id on success.
func AddTeacher(m *Teacher) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTeacherById retrieves Teacher by Id. Returns error if
// Id doesn't exist
func GetTeacherById(id int) (v *Teacher, err error) {
	o := orm.NewOrm()
	v = &Teacher{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTeacher retrieves all Teacher matches certain condition. Returns empty list if
// no records exist
func GetAllTeacher(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Teacher))
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

	var l []Teacher
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

// UpdateTeacher updates Teacher by Id and returns error if
// the record to be updated doesn't exist
func UpdateTeacherById(m *Teacher) (err error) {
	o := orm.NewOrm()
	v := Teacher{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTeacher deletes Teacher by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTeacher(id int) (err error) {
	o := orm.NewOrm()
	v := Teacher{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Teacher{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
