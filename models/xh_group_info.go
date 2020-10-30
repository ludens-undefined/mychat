package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhGroupInfo struct {
	Id         int    `orm:"column:id;autoIncrement"`
	Groupid    int    `orm:"column:groupid"` //团ID  course_group
	Memberid   int    `orm:"column:memberid"` //学员id
	Isgrouper  int8   `orm:"column:isgrouper"` //默认是团长 1：团长   2 ：成员
	Xhcourseid int    `orm:"column:xhcourseid"` //课程id
	Createtime int    `orm:"column:createtime"` //参团时间
	OutTradeNo string `orm:"column:out_trade_no;size:255"` //订单id
	ShopId     int    `orm:"column:shop_id"`
	Status     int8   `orm:"column:status"` //团状态 默认1 拼团成功 2 拼团未完成 3团已过期4:拼团退款
	TuiStatus  int8   `orm:"column:tui_status"` //退费：1 已退费   2：未退费
	Xoldid     int    `orm:"column:xoldid"`
}

func (t *XhGroupInfo) TableName() string {
	return "goouc_xmf_xh_group_info"
}

func init() {
	orm.RegisterModel(new(XhGroupInfo))
}

// AddXhGroupInfo insert a new XhGroupInfo into database and returns
// last inserted Id on success.
func AddXhGroupInfo(m *XhGroupInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhGroupInfoById retrieves XhGroupInfo by Id. Returns error if
// Id doesn't exist
func GetXhGroupInfoById(id int) (v *XhGroupInfo, err error) {
	o := orm.NewOrm()
	v = &XhGroupInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhGroupInfo retrieves all XhGroupInfo matches certain condition. Returns empty list if
// no records exist
func GetAllXhGroupInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhGroupInfo))
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

	var l []XhGroupInfo
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

// UpdateXhGroupInfo updates XhGroupInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhGroupInfoById(m *XhGroupInfo) (err error) {
	o := orm.NewOrm()
	v := XhGroupInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhGroupInfo deletes XhGroupInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhGroupInfo(id int) (err error) {
	o := orm.NewOrm()
	v := XhGroupInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhGroupInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
