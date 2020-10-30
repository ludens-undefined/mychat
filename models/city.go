package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type City struct {
	Id       int    `orm:"column:id;autoIncrement"`
	ShopId   int    `orm:"column:shop_id"`
	Name     string `orm:"column:name;size:32"`
	Qq       string `orm:"column:qq;size:16"` //QQ
	Contact  string `orm:"column:contact;size:255"` //联系方式
	Address  string `orm:"column:address;size:255"` //报名地址
	Status   int8   `orm:"column:status"` //状态 1显示 2隐藏
	Sort     int    `orm:"column:sort"` //排序 由小到大
	IsDelete int8   `orm:"column:is_delete"` //状态 1启用 2删除
	Qrcode   string `orm:"column:qrcode;size:255"` //二维码
	Desc     string `orm:"column:desc;size:500"` //详细描述
}

func (t *City) TableName() string {
	return "goouc_xmf_city"
}

func init() {
	orm.RegisterModel(new(City))
}

// AddCity insert a new City into database and returns
// last inserted Id on success.
func AddCity(m *City) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCityById retrieves City by Id. Returns error if
// Id doesn't exist
func GetCityById(id int) (v *City, err error) {
	o := orm.NewOrm()
	v = &City{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCity retrieves all City matches certain condition. Returns empty list if
// no records exist
func GetAllCity(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(City))
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

	var l []City
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

// UpdateCity updates City by Id and returns error if
// the record to be updated doesn't exist
func UpdateCityById(m *City) (err error) {
	o := orm.NewOrm()
	v := City{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCity deletes City by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCity(id int) (err error) {
	o := orm.NewOrm()
	v := City{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&City{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
