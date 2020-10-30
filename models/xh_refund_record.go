package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhRefundRecord struct {
	Id         int     `orm:"column:id;autoIncrement"`
	ShopId     int     `orm:"column:shop_id"`
	Memberid   int     `orm:"column:memberid"`
	Orderid    int     `orm:"column:orderid"`
	OutTradeNo string  `orm:"column:out_trade_no;size:255"`
	RefundNo   string  `orm:"column:refund_no;size:255"`
	Money      float64 `orm:"column:money;scale:10;precision:2"` //退款金额
	RefundType int8    `orm:"column:refund_type"` //1:手动  2：自动
	Operator   int     `orm:"column:operator"` //操作者
	Platform   string  `orm:"column:platform;size:255"` //退款平台
	Status     int8    `orm:"column:status"` //1:成功 2：失败
	Createtime int     `orm:"column:createtime"`
}

func (t *XhRefundRecord) TableName() string {
	return "goouc_xmf_xh_refund_record"
}

func init() {
	orm.RegisterModel(new(XhRefundRecord))
}

// AddXhRefundRecord insert a new XhRefundRecord into database and returns
// last inserted Id on success.
func AddXhRefundRecord(m *XhRefundRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhRefundRecordById retrieves XhRefundRecord by Id. Returns error if
// Id doesn't exist
func GetXhRefundRecordById(id int) (v *XhRefundRecord, err error) {
	o := orm.NewOrm()
	v = &XhRefundRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhRefundRecord retrieves all XhRefundRecord matches certain condition. Returns empty list if
// no records exist
func GetAllXhRefundRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhRefundRecord))
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

	var l []XhRefundRecord
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

// UpdateXhRefundRecord updates XhRefundRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhRefundRecordById(m *XhRefundRecord) (err error) {
	o := orm.NewOrm()
	v := XhRefundRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhRefundRecord deletes XhRefundRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhRefundRecord(id int) (err error) {
	o := orm.NewOrm()
	v := XhRefundRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhRefundRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
