package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhTicket struct {
	Id         int     `orm:"column:id;autoIncrement"`
	Desc       string  `orm:"column:desc;size:50"` //兑课券使用描述
	Money      float64 `orm:"column:money;scale:11;precision:2"` //兑课券抵扣金额
	Amount     float64 `orm:"column:amount;scale:11;precision:2"` //订单总额超出设定值才可使用
	Icon       int     `orm:"column:icon"` //兑换兑课券需要消耗的晓禾币的数量
	ValidTime  int     `orm:"column:valid_time"` //有效期，单位为天，默认为30天
	Status     int8    `orm:"column:status"` //状态默认不显示      1：显示  2：不显示
	Sort       int     `orm:"column:sort"` //排序值
	ShopId     int     `orm:"column:shop_id"` //店铺id
	Createtime int     `orm:"column:createtime"` //创建时间
	Updatetime int     `orm:"column:updatetime"` //更新时间
	IsDelete   int8    `orm:"column:is_delete"` //1：启用  2：删除
	Xoldid     int     `orm:"column:xoldid"`
}

func (t *XhTicket) TableName() string {
	return "goouc_xmf_xh_ticket"
}

func init() {
	orm.RegisterModel(new(XhTicket))
}

// AddXhTicket insert a new XhTicket into database and returns
// last inserted Id on success.
func AddXhTicket(m *XhTicket) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhTicketById retrieves XhTicket by Id. Returns error if
// Id doesn't exist
func GetXhTicketById(id int) (v *XhTicket, err error) {
	o := orm.NewOrm()
	v = &XhTicket{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhTicket retrieves all XhTicket matches certain condition. Returns empty list if
// no records exist
func GetAllXhTicket(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhTicket))
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

	var l []XhTicket
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

// UpdateXhTicket updates XhTicket by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhTicketById(m *XhTicket) (err error) {
	o := orm.NewOrm()
	v := XhTicket{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhTicket deletes XhTicket by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhTicket(id int) (err error) {
	o := orm.NewOrm()
	v := XhTicket{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhTicket{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
