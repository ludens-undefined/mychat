package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjCgdkKs struct {
	Id       int    `orm:"column:id;autoIncrement"`
	CgdkId   int    `orm:"column:cgdk_id"` //闯关打卡id，对应goouc_xet_zxgj_cgdk表中的id
	Img      string `orm:"column:img;size:255"` //课程封面
	DkCount  uint   `orm:"column:dk_count"` //打卡次数
	Zan      uint   `orm:"column:zan"` //点赞数
	Remark   uint   `orm:"column:remark"` //评论数
	Dp       uint   `orm:"column:dp"` //点评数
	Jx       uint   `orm:"column:jx"` //精选数
	CreateAt uint   `orm:"column:create_at"` //创建时间
	UpdateAt uint   `orm:"column:update_at"` //修改时间
}

func (t *ZxgjCgdkKs) TableName() string {
	return "goouc_xmf_zxgj_cgdk_ks"
}

func init() {
	orm.RegisterModel(new(ZxgjCgdkKs))
}

// AddZxgjCgdkKs insert a new ZxgjCgdkKs into database and returns
// last inserted Id on success.
func AddZxgjCgdkKs(m *ZxgjCgdkKs) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjCgdkKsById retrieves ZxgjCgdkKs by Id. Returns error if
// Id doesn't exist
func GetZxgjCgdkKsById(id int) (v *ZxgjCgdkKs, err error) {
	o := orm.NewOrm()
	v = &ZxgjCgdkKs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjCgdkKs retrieves all ZxgjCgdkKs matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjCgdkKs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjCgdkKs))
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

	var l []ZxgjCgdkKs
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

// UpdateZxgjCgdkKs updates ZxgjCgdkKs by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjCgdkKsById(m *ZxgjCgdkKs) (err error) {
	o := orm.NewOrm()
	v := ZxgjCgdkKs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjCgdkKs deletes ZxgjCgdkKs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjCgdkKs(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjCgdkKs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjCgdkKs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
