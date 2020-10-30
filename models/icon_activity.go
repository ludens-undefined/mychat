package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type IconActivity struct {
	Id         int    `orm:"column:id;autoIncrement"` //ID
	ShopId     int    `orm:"column:shop_id"`
	Name       string `orm:"column:name;size:500"` //活动名称
	Num        int    `orm:"column:num"` //发放数量
	Remark     string `orm:"column:remark"`
	Status     int8   `orm:"column:status"` //1:开启  2：关闭
	IsDelete   int8   `orm:"column:is_delete"` //1开启  2;删除
	Createtime int    `orm:"column:createtime"`
	Starttime  int    `orm:"column:starttime"` //活动开始时间
	Endtime    int    `orm:"column:endtime"` //活动结束时间
	NumStatus  int8   `orm:"column:num_status"` //1:数量统一   2：数量自定义
	Sort       int8   `orm:"column:sort"`
}

func (t *IconActivity) TableName() string {
	return "goouc_xmf_icon_activity"
}

func init() {
	orm.RegisterModel(new(IconActivity))
}

// AddIconActivity insert a new IconActivity into database and returns
// last inserted Id on success.
func AddIconActivity(m *IconActivity) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetIconActivityById retrieves IconActivity by Id. Returns error if
// Id doesn't exist
func GetIconActivityById(id int) (v *IconActivity, err error) {
	o := orm.NewOrm()
	v = &IconActivity{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllIconActivity retrieves all IconActivity matches certain condition. Returns empty list if
// no records exist
func GetAllIconActivity(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(IconActivity))
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

	var l []IconActivity
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

// UpdateIconActivity updates IconActivity by Id and returns error if
// the record to be updated doesn't exist
func UpdateIconActivityById(m *IconActivity) (err error) {
	o := orm.NewOrm()
	v := IconActivity{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIconActivity deletes IconActivity by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIconActivity(id int) (err error) {
	o := orm.NewOrm()
	v := IconActivity{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&IconActivity{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
