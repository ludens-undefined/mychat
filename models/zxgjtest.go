package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjTest struct {
	Id        int    `orm:"column:id;autoIncrement"`
	ShopId    uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name      string `orm:"column:name;size:50"` //测试名称
	Brief     string `orm:"column:brief;size:255"` //测试简介
	Img       string `orm:"column:img;size:255"` //封面配图
	TkId      uint   `orm:"column:tk_id"` //题库，对应goouc_xet_zxgj_tk表中的id
	StartTime uint   `orm:"column:start_time"` //有效开始时间
	EndTime   uint   `orm:"column:end_time"` //有效结束时间
	IsDelete  uint8  `orm:"column:is_delete"` //1启用 2删除
	CreateAt  uint   `orm:"column:create_at"` //创建时间
	UpdateAt  uint   `orm:"column:update_at"` //修改时间
	DeleteAt  uint   `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjTest) TableName() string {
	return "goouc_xmf_zxgj_test"
}

func init() {
	orm.RegisterModel(new(ZxgjTest))
}

// AddZxgjTest insert a new ZxgjTest into database and returns
// last inserted Id on success.
func AddZxgjTest(m *ZxgjTest) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjTestById retrieves ZxgjTest by Id. Returns error if
// Id doesn't exist
func GetZxgjTestById(id int) (v *ZxgjTest, err error) {
	o := orm.NewOrm()
	v = &ZxgjTest{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjTest retrieves all ZxgjTest matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjTest(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjTest))
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

	var l []ZxgjTest
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

// UpdateZxgjTest updates ZxgjTest by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjTestById(m *ZxgjTest) (err error) {
	o := orm.NewOrm()
	v := ZxgjTest{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjTest deletes ZxgjTest by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjTest(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjTest{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjTest{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
