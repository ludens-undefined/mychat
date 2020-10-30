package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PayModelOrder struct {
	Id         int    `orm:"column:id;autoIncrement"`
	OrderSn    string `orm:"column:order_sn;size:30"` //订单编号
	ShopId     uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	ModelId    uint   `orm:"column:model_id"` //对应goouc_xet_pay_model表中的id
	IsPay      uint8  `orm:"column:is_pay"` //是否已经立即订购 1已订购 2未订购
	IsUse      uint8  `orm:"column:is_use"` //1立即订购 2免费试用
	OrderState uint8  `orm:"column:order_state"` //1未支付 2已支付
	OrderTime  uint   `orm:"column:order_time"` //下订单时间
	PayTime    uint   `orm:"column:pay_time"` //支付时间
	EndTime    uint   `orm:"column:end_time"` //到期时间
}

func (t *PayModelOrder) TableName() string {
	return "goouc_xmf_pay_model_order"
}

func init() {
	orm.RegisterModel(new(PayModelOrder))
}

// AddPayModelOrder insert a new PayModelOrder into database and returns
// last inserted Id on success.
func AddPayModelOrder(m *PayModelOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPayModelOrderById retrieves PayModelOrder by Id. Returns error if
// Id doesn't exist
func GetPayModelOrderById(id int) (v *PayModelOrder, err error) {
	o := orm.NewOrm()
	v = &PayModelOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPayModelOrder retrieves all PayModelOrder matches certain condition. Returns empty list if
// no records exist
func GetAllPayModelOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PayModelOrder))
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

	var l []PayModelOrder
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

// UpdatePayModelOrder updates PayModelOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdatePayModelOrderById(m *PayModelOrder) (err error) {
	o := orm.NewOrm()
	v := PayModelOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePayModelOrder deletes PayModelOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeletePayModelOrder(id int) (err error) {
	o := orm.NewOrm()
	v := PayModelOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PayModelOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
