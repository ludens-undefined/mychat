package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Signin struct {
	Id       int    `orm:"column:id;autoIncrement"`
	ShopId   int    `orm:"column:shop_id"`
	Memberid int    `orm:"column:memberid"` //用户签到标识
	Num      int    `orm:"column:num"` //记录连续签到天数
	Date     string `orm:"column:date;size:200"` //当天签到日期
	Month    string `orm:"column:month;size:200"` //签到月份第一天日期
	Days     string `orm:"column:days"` //当月签到汇总， 数组形式
	Status   int8   `orm:"column:status"` //领取每月奖励状态    1:没领取  2：已经领取
	Ctime    int    `orm:"column:ctime"` //签到时间
	Xoldid   int    `orm:"column:xoldid"`
}

func (t *Signin) TableName() string {
	return "goouc_xmf_signin"
}

func init() {
	orm.RegisterModel(new(Signin))
}

// AddSignin insert a new Signin into database and returns
// last inserted Id on success.
func AddSignin(m *Signin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSigninById retrieves Signin by Id. Returns error if
// Id doesn't exist
func GetSigninById(id int) (v *Signin, err error) {
	o := orm.NewOrm()
	v = &Signin{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSignin retrieves all Signin matches certain condition. Returns empty list if
// no records exist
func GetAllSignin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Signin))
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

	var l []Signin
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

// UpdateSignin updates Signin by Id and returns error if
// the record to be updated doesn't exist
func UpdateSigninById(m *Signin) (err error) {
	o := orm.NewOrm()
	v := Signin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSignin deletes Signin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSignin(id int) (err error) {
	o := orm.NewOrm()
	v := Signin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Signin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
