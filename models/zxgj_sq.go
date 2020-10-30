package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjSq struct {
	Id           int     `orm:"column:id;autoIncrement"`
	ShopId       uint    `orm:"column:shop_id"` //店铺id，对应goouc_xet_business_shop表中的id
	Name         string  `orm:"column:name;size:50"` //社群名称
	Brief        string  `orm:"column:brief;size:50"` //社群简介
	Img          string  `orm:"column:img;size:255"` //社群封面
	JoinPrice    float64 `orm:"column:join_price;scale:10;precision:2"` //入群价格->付费入群，0免费
	JoinRelative string  `orm:"column:join_relative;size:50"` //入群价格->关联专栏或会员,格式：{type:1,id:1}，type对应goouc_xet_cource_type中的id
	IsShow       uint8   `orm:"column:is_show"` //在店铺内显示 1立即显示 2隐藏
	QzId         uint    `orm:"column:qz_id"` //群主，对应goouc_xet_user表中的id
	CreateAt     uint    `orm:"column:create_at"` //创建时间
	UpdateAt     uint    `orm:"column:update_at"` //修改时间
}

func (t *ZxgjSq) TableName() string {
	return "goouc_xmf_zxgj_sq"
}

func init() {
	orm.RegisterModel(new(ZxgjSq))
}

// AddZxgjSq insert a new ZxgjSq into database and returns
// last inserted Id on success.
func AddZxgjSq(m *ZxgjSq) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjSqById retrieves ZxgjSq by Id. Returns error if
// Id doesn't exist
func GetZxgjSqById(id int) (v *ZxgjSq, err error) {
	o := orm.NewOrm()
	v = &ZxgjSq{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjSq retrieves all ZxgjSq matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjSq(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjSq))
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

	var l []ZxgjSq
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

// UpdateZxgjSq updates ZxgjSq by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjSqById(m *ZxgjSq) (err error) {
	o := orm.NewOrm()
	v := ZxgjSq{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjSq deletes ZxgjSq by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjSq(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjSq{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjSq{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
