package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CouponCate struct {
	Id         int       `orm:"column:id;autoIncrement"`
	ShopId     int       `orm:"column:shop_id"`
	Name       string    `orm:"column:name;size:255"` //优惠券名称
	Type       int8      `orm:"column:type"` //状态 1金额 2折扣
	Starttime  time.Time `orm:"column:starttime;type:time"` //开始时间
	Endtime    time.Time `orm:"column:endtime;type:time"` //结束时间
	Money      float64   `orm:"column:money;scale:10;precision:2"` //优惠券金额
	Discount   string    `orm:"column:discount;size:6"` //优惠券折扣
	Fullmoney  float64   `orm:"column:fullmoney;scale:10;precision:2"` //优惠券使用条件
	Term       int       `orm:"column:term"` //使用学期
	Class      int       `orm:"column:class"` //使用班型
	Grade      int       `orm:"column:grade"` //使用年级
	Subject    int       `orm:"column:subject"` //使用学科
	Year       int       `orm:"column:year"` //使用年限
	Msg        string    `orm:"column:msg;size:2000"` //优惠券使用说明
	Createtime int       `orm:"column:createtime"`
	Updatetime int       `orm:"column:updatetime"`
	Status     int8      `orm:"column:status"` //状态 1显示 2隐藏
	IsDelete   int8      `orm:"column:is_delete"`
	Source     int8      `orm:"column:source"` //1普通优惠券 2老带新
	CourseId   int       `orm:"column:course_id"`
	ClassId    int       `orm:"column:class_id"`
	IsReceive  int8      `orm:"column:is_receive"` //点击领取 1开启 2关闭
}

func (t *CouponCate) TableName() string {
	return "goouc_xmf_coupon_cate"
}

func init() {
	orm.RegisterModel(new(CouponCate))
}

// AddCouponCate insert a new CouponCate into database and returns
// last inserted Id on success.
func AddCouponCate(m *CouponCate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponCateById retrieves CouponCate by Id. Returns error if
// Id doesn't exist
func GetCouponCateById(id int) (v *CouponCate, err error) {
	o := orm.NewOrm()
	v = &CouponCate{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCouponCate retrieves all CouponCate matches certain condition. Returns empty list if
// no records exist
func GetAllCouponCate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponCate))
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

	var l []CouponCate
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

// UpdateCouponCate updates CouponCate by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponCateById(m *CouponCate) (err error) {
	o := orm.NewOrm()
	v := CouponCate{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCouponCate deletes CouponCate by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCouponCate(id int) (err error) {
	o := orm.NewOrm()
	v := CouponCate{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CouponCate{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
