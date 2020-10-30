package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Discount struct {
	Id         int       `orm:"column:id;autoIncrement"`
	ShopId     int       `orm:"column:shop_id"`
	Name       string    `orm:"column:name;size:255"` //优惠名称
	Userid     int       `orm:"column:userid"` //所属用户 0未发放
	Starttime  time.Time `orm:"column:starttime;type:time"` //开始时间
	Endtime    time.Time `orm:"column:endtime;type:time"` //结束时间
	Createtime int       `orm:"column:createtime"`
	Updatetime int       `orm:"column:updatetime"`
	Grade      int       `orm:"column:grade"` //使用年级
	Subject    int       `orm:"column:subject"` //使用学科
	Year       int       `orm:"column:year"` //使用年限
	Term       int       `orm:"column:term"` //使用学期
	Discount   float64   `orm:"column:discount;scale:2;precision:1"` //优惠券折扣
	Msg        string    `orm:"column:msg;size:2000"` //说明
	IsDelete   int8      `orm:"column:is_delete"`
}

func (t *Discount) TableName() string {
	return "goouc_xmf_discount"
}

func init() {
	orm.RegisterModel(new(Discount))
}

// AddDiscount insert a new Discount into database and returns
// last inserted Id on success.
func AddDiscount(m *Discount) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDiscountById retrieves Discount by Id. Returns error if
// Id doesn't exist
func GetDiscountById(id int) (v *Discount, err error) {
	o := orm.NewOrm()
	v = &Discount{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDiscount retrieves all Discount matches certain condition. Returns empty list if
// no records exist
func GetAllDiscount(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Discount))
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

	var l []Discount
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

// UpdateDiscount updates Discount by Id and returns error if
// the record to be updated doesn't exist
func UpdateDiscountById(m *Discount) (err error) {
	o := orm.NewOrm()
	v := Discount{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDiscount deletes Discount by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDiscount(id int) (err error) {
	o := orm.NewOrm()
	v := Discount{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Discount{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
