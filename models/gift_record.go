package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type GiftRecord struct {
	Id              int    `orm:"column:id;autoIncrement"`
	ShopId          int    `orm:"column:shop_id"`
	Memberid        int    `orm:"column:memberid"` //会员id
	Giftid          int    `orm:"column:giftid"` //礼品id
	Iconnum         int    `orm:"column:iconnum"` //兑换币数量
	Address         string `orm:"column:address"` //收货地址
	Consignee       string `orm:"column:consignee;size:50"` //收货人
	ConsigneeMobile string `orm:"column:consignee_mobile;size:30"` //收货手机号
	Createtime      int    `orm:"column:createtime"` //兑换时间
	Remarks         string `orm:"column:remarks"` //备注
	Status          int8   `orm:"column:status"` //2:待收货  1：未发货 3：已收货
	Logistics       string `orm:"column:logistics;size:100"` //物流单号
	Updatetime      int    `orm:"column:updatetime"`
}

func (t *GiftRecord) TableName() string {
	return "goouc_xmf_gift_record"
}

func init() {
	orm.RegisterModel(new(GiftRecord))
}

// AddGiftRecord insert a new GiftRecord into database and returns
// last inserted Id on success.
func AddGiftRecord(m *GiftRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGiftRecordById retrieves GiftRecord by Id. Returns error if
// Id doesn't exist
func GetGiftRecordById(id int) (v *GiftRecord, err error) {
	o := orm.NewOrm()
	v = &GiftRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGiftRecord retrieves all GiftRecord matches certain condition. Returns empty list if
// no records exist
func GetAllGiftRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GiftRecord))
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

	var l []GiftRecord
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

// UpdateGiftRecord updates GiftRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateGiftRecordById(m *GiftRecord) (err error) {
	o := orm.NewOrm()
	v := GiftRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGiftRecord deletes GiftRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGiftRecord(id int) (err error) {
	o := orm.NewOrm()
	v := GiftRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GiftRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
