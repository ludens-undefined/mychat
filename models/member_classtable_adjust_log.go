package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberClasstableAdjustLog struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     int    `orm:"column:shop_id"`
	Userid     int    `orm:"column:userid"`
	Mcid       int    `orm:"column:mcid"` //用户课程编号
	Mctid      int    `orm:"column:mctid"` //用户课程表编号
	FromTid    int    `orm:"column:from_tid"` //当前课程表id
	ToTid      int    `orm:"column:to_tid"` //目标课程表id
	Createtime int    `orm:"column:createtime"`
	Msg        string `orm:"column:msg;size:2000"` //调课理由
	Status     int8   `orm:"column:status"` //状态 1有效 2无效
	Operator   string `orm:"column:operator;size:32"` //操作人
	Platform   string `orm:"column:platform;size:32"` //操作平台
}

func (t *MemberClasstableAdjustLog) TableName() string {
	return "goouc_xmf_member_classtable_adjust_log"
}

func init() {
	orm.RegisterModel(new(MemberClasstableAdjustLog))
}

// AddMemberClasstableAdjustLog insert a new MemberClasstableAdjustLog into database and returns
// last inserted Id on success.
func AddMemberClasstableAdjustLog(m *MemberClasstableAdjustLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberClasstableAdjustLogById retrieves MemberClasstableAdjustLog by Id. Returns error if
// Id doesn't exist
func GetMemberClasstableAdjustLogById(id int) (v *MemberClasstableAdjustLog, err error) {
	o := orm.NewOrm()
	v = &MemberClasstableAdjustLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberClasstableAdjustLog retrieves all MemberClasstableAdjustLog matches certain condition. Returns empty list if
// no records exist
func GetAllMemberClasstableAdjustLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberClasstableAdjustLog))
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

	var l []MemberClasstableAdjustLog
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

// UpdateMemberClasstableAdjustLog updates MemberClasstableAdjustLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberClasstableAdjustLogById(m *MemberClasstableAdjustLog) (err error) {
	o := orm.NewOrm()
	v := MemberClasstableAdjustLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberClasstableAdjustLog deletes MemberClasstableAdjustLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberClasstableAdjustLog(id int) (err error) {
	o := orm.NewOrm()
	v := MemberClasstableAdjustLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberClasstableAdjustLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
