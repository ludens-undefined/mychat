package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberClass struct {
	Id         int  `orm:"column:id;autoIncrement"`
	ShopId     int  `orm:"column:shop_id"`
	Userid     int  `orm:"column:userid"`
	CourseId   int  `orm:"column:course_id"` //课程id
	ClassId    int  `orm:"column:class_id"` //班级id
	Createtime int  `orm:"column:createtime"`
	Updatetime int  `orm:"column:updatetime"`
	Status     int8 `orm:"column:status"` //状态 1有效 2申请退款 3申请通过待退款 4退款成功
	Updown     int8 `orm:"column:updown"` //课程上下 1完整课程 2上 3下
}

func (t *MemberClass) TableName() string {
	return "goouc_xmf_member_class"
}

func init() {
	orm.RegisterModel(new(MemberClass))
}

// AddMemberClass insert a new MemberClass into database and returns
// last inserted Id on success.
func AddMemberClass(m *MemberClass) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberClassById retrieves MemberClass by Id. Returns error if
// Id doesn't exist
func GetMemberClassById(id int) (v *MemberClass, err error) {
	o := orm.NewOrm()
	v = &MemberClass{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberClass retrieves all MemberClass matches certain condition. Returns empty list if
// no records exist
func GetAllMemberClass(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberClass))
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

	var l []MemberClass
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

// UpdateMemberClass updates MemberClass by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberClassById(m *MemberClass) (err error) {
	o := orm.NewOrm()
	v := MemberClass{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberClass deletes MemberClass by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberClass(id int) (err error) {
	o := orm.NewOrm()
	v := MemberClass{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberClass{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
