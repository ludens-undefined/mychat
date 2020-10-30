package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjActOrder struct {
	Id         int     `orm:"column:id;pk"`
	OrderSn    string  `orm:"column:order_sn;size:30"` //订单号
	ActId      uint    `orm:"column:act_id"` //活动管理id，对应goouc_xet_zxgj_act表中的id
	UserId     uint    `orm:"column:user_id"` //对应goouc_xet_user表中的id
	Type       uint8   `orm:"column:type"` //1免费票 2收费票
	Num        uint16  `orm:"column:num"` //买票数量
	ActUserId  uint    `orm:"column:act_user_id"` //购买票的用户id,对应goouc_xet_zxgj_act_order_user表中的id
	Price      float64 `orm:"column:price;scale:10;precision:2"` //支付价格
	OrderState uint8   `orm:"column:order_state"` //1未支付 2已支付
	PayTime    uint    `orm:"column:pay_time"` //支付时间
	OrderTime  uint    `orm:"column:order_time"` //下订单时间
}

func (t *ZxgjActOrder) TableName() string {
	return "goouc_xmf_zxgj_act_order"
}

func init() {
	orm.RegisterModel(new(ZxgjActOrder))
}

// AddZxgjActOrder insert a new ZxgjActOrder into database and returns
// last inserted Id on success.
func AddZxgjActOrder(m *ZxgjActOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjActOrderById retrieves ZxgjActOrder by Id. Returns error if
// Id doesn't exist
func GetZxgjActOrderById(id int) (v *ZxgjActOrder, err error) {
	o := orm.NewOrm()
	v = &ZxgjActOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjActOrder retrieves all ZxgjActOrder matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjActOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjActOrder))
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

	var l []ZxgjActOrder
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

// UpdateZxgjActOrder updates ZxgjActOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjActOrderById(m *ZxgjActOrder) (err error) {
	o := orm.NewOrm()
	v := ZxgjActOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjActOrder deletes ZxgjActOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjActOrder(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjActOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjActOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
