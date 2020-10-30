package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type IconOrderRecord struct {
	Id              int    `orm:"column:id;autoIncrement"`
	ShopId          int    `orm:"column:shop_id"`
	Memberid        int    `orm:"column:memberid"` //会员id
	Xhcourseid      int    `orm:"column:xhcourseid"` //礼品id
	Iconnum         int    `orm:"column:iconnum"` //兑换币数量
	Address         string `orm:"column:address"` //收货地址
	Consignee       string `orm:"column:consignee;size:50"` //收货人
	ConsigneeMobile string `orm:"column:consignee_mobile;size:30"` //收货手机号
	Createtime      int    `orm:"column:createtime"` //兑换时间
	Remarks         string `orm:"column:remarks"` //备注
	LogisticsNo     string `orm:"column:logistics_no;size:50"` //物流单号
}

func (t *IconOrderRecord) TableName() string {
	return "goouc_xmf_icon_order_record"
}

func init() {
	orm.RegisterModel(new(IconOrderRecord))
}

// AddIconOrderRecord insert a new IconOrderRecord into database and returns
// last inserted Id on success.
func AddIconOrderRecord(m *IconOrderRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetIconOrderRecordById retrieves IconOrderRecord by Id. Returns error if
// Id doesn't exist
func GetIconOrderRecordById(id int) (v *IconOrderRecord, err error) {
	o := orm.NewOrm()
	v = &IconOrderRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllIconOrderRecord retrieves all IconOrderRecord matches certain condition. Returns empty list if
// no records exist
func GetAllIconOrderRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(IconOrderRecord))
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

	var l []IconOrderRecord
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

// UpdateIconOrderRecord updates IconOrderRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateIconOrderRecordById(m *IconOrderRecord) (err error) {
	o := orm.NewOrm()
	v := IconOrderRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIconOrderRecord deletes IconOrderRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIconOrderRecord(id int) (err error) {
	o := orm.NewOrm()
	v := IconOrderRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&IconOrderRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
