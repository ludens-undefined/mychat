package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ConfigAttention struct {
	Id             int    `orm:"column:id;autoIncrement"`
	ShopId         uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Type           uint8  `orm:"column:type"` //关注公众号类型 1关注公众号 2购买后引导关注
	IsAttention    uint8  `orm:"column:is_attention"` //引导关注公众号 1引导关注 2不引导关注
	IsBuyAttention uint8  `orm:"column:is_buy_attention"` //购买后引导关注 1引导关注 2不引导关注
	Desc           string `orm:"column:desc;size:40"` //对应type为2，引导描述
	Name           string `orm:"column:name;size:30"` //公众号名称
	Qrcode         string `orm:"column:qrcode;size:255"` //二维码
}

func (t *ConfigAttention) TableName() string {
	return "goouc_xmf_config_attention"
}

func init() {
	orm.RegisterModel(new(ConfigAttention))
}

// AddConfigAttention insert a new ConfigAttention into database and returns
// last inserted Id on success.
func AddConfigAttention(m *ConfigAttention) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetConfigAttentionById retrieves ConfigAttention by Id. Returns error if
// Id doesn't exist
func GetConfigAttentionById(id int) (v *ConfigAttention, err error) {
	o := orm.NewOrm()
	v = &ConfigAttention{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllConfigAttention retrieves all ConfigAttention matches certain condition. Returns empty list if
// no records exist
func GetAllConfigAttention(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ConfigAttention))
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

	var l []ConfigAttention
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

// UpdateConfigAttention updates ConfigAttention by Id and returns error if
// the record to be updated doesn't exist
func UpdateConfigAttentionById(m *ConfigAttention) (err error) {
	o := orm.NewOrm()
	v := ConfigAttention{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteConfigAttention deletes ConfigAttention by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConfigAttention(id int) (err error) {
	o := orm.NewOrm()
	v := ConfigAttention{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ConfigAttention{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
