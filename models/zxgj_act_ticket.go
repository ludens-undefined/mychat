package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjActTicket struct {
	Id           int     `orm:"column:id;autoIncrement"`
	ActId        uint    `orm:"column:act_id"` //活动管理id，对应goouc_xet_zxgj_act表中的id
	Type         uint8   `orm:"column:type"` //1免费票 2收费票
	Name         string  `orm:"column:name;size:50"` //票种名称
	Price        float64 `orm:"column:price;scale:10;precision:2"` //价格
	Num          uint16  `orm:"column:num"` //票数
	IsSh         uint8   `orm:"column:is_sh"` //1审核 2不审核
	Illustration string  `orm:"column:illustration;size:255"` //票种说明
	BuyNum       uint16  `orm:"column:buy_num"` //单次购买最多数 0不限制
	Taopiao      uint8   `orm:"column:taopiao"` //套票设置
}

func (t *ZxgjActTicket) TableName() string {
	return "goouc_xmf_zxgj_act_ticket"
}

func init() {
	orm.RegisterModel(new(ZxgjActTicket))
}

// AddZxgjActTicket insert a new ZxgjActTicket into database and returns
// last inserted Id on success.
func AddZxgjActTicket(m *ZxgjActTicket) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjActTicketById retrieves ZxgjActTicket by Id. Returns error if
// Id doesn't exist
func GetZxgjActTicketById(id int) (v *ZxgjActTicket, err error) {
	o := orm.NewOrm()
	v = &ZxgjActTicket{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjActTicket retrieves all ZxgjActTicket matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjActTicket(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjActTicket))
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

	var l []ZxgjActTicket
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

// UpdateZxgjActTicket updates ZxgjActTicket by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjActTicketById(m *ZxgjActTicket) (err error) {
	o := orm.NewOrm()
	v := ZxgjActTicket{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjActTicket deletes ZxgjActTicket by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjActTicket(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjActTicket{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjActTicket{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
