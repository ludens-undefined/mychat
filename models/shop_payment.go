package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ShopPayment struct {
	Id            int    `orm:"column:id;autoIncrement"`
	ShopId        int    `orm:"column:shop_id"`
	Name          string `orm:"column:name;size:32"` //第三方平台名称
	Code          string `orm:"column:code;size:32"` //第三方平台编码
	Config        string `orm:"column:config"` //参数
	ApiclientCert string `orm:"column:apiclient_cert;size:255"` //证书
	ApiclientKey  string `orm:"column:apiclient_key;size:255"` //证书密钥
	Enabled       int8   `orm:"column:enabled"` //是否安装 1安装 2卸载
	Position      int8   `orm:"column:position"`
}

func (t *ShopPayment) TableName() string {
	return "goouc_xmf_shop_payment"
}

func init() {
	orm.RegisterModel(new(ShopPayment))
}

// AddShopPayment insert a new ShopPayment into database and returns
// last inserted Id on success.
func AddShopPayment(m *ShopPayment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetShopPaymentById retrieves ShopPayment by Id. Returns error if
// Id doesn't exist
func GetShopPaymentById(id int) (v *ShopPayment, err error) {
	o := orm.NewOrm()
	v = &ShopPayment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShopPayment retrieves all ShopPayment matches certain condition. Returns empty list if
// no records exist
func GetAllShopPayment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShopPayment))
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

	var l []ShopPayment
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

// UpdateShopPayment updates ShopPayment by Id and returns error if
// the record to be updated doesn't exist
func UpdateShopPaymentById(m *ShopPayment) (err error) {
	o := orm.NewOrm()
	v := ShopPayment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShopPayment deletes ShopPayment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShopPayment(id int) (err error) {
	o := orm.NewOrm()
	v := ShopPayment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ShopPayment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
