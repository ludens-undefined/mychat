package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhIconRecord struct {
	Id             int    `orm:"column:id;autoIncrement"` //ID
	ShopId         uint   `orm:"column:shop_id"` //店铺id
	Memberid       uint   `orm:"column:memberid"` //学员id
	Operator       int    `orm:"column:operator"` //后台导入者id   -1:系统发放 
	Lastnum        int    `orm:"column:lastnum"` //更改之前的兑换币数量
	Nownum         int    `orm:"column:nownum"` //更新之后兑换币数量
	Num            int    `orm:"column:num"` //操作数量（增加正号，减少负号）
	Xhcourseid     int    `orm:"column:xhcourseid"` //课程id（也存储礼品id）
	Chapterid      int    `orm:"column:chapterid"` //章节id
	IconActivityid int    `orm:"column:icon_activityid"` //活动id
	NumType        uint8  `orm:"column:num_type"` //1:增加  2：减少
	GetType        int8   `orm:"column:get_type"` //1：注册，2：签到，3：分享，4：设置头像，5：完善信息，6：购买课程赠送，7：按时上课，8：查看回放，9：兑换兑换券,10：兑换课程，11：后台导入 ，12：后台单个操作，13：礼品兑换 ，14：活动导入，15：活动扣除
	Remark         string `orm:"column:remark;size:1000"` //备注（备注记录哪个活动，哪个课程，哪个兑换券等）
	Createtime     int    `orm:"column:createtime"` //记录时间
	IsDelete       int8   `orm:"column:is_delete"` //1:显示 2：不显示
}

func (t *XhIconRecord) TableName() string {
	return "goouc_xmf_xh_icon_record"
}

func init() {
	orm.RegisterModel(new(XhIconRecord))
}

// AddXhIconRecord insert a new XhIconRecord into database and returns
// last inserted Id on success.
func AddXhIconRecord(m *XhIconRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhIconRecordById retrieves XhIconRecord by Id. Returns error if
// Id doesn't exist
func GetXhIconRecordById(id int) (v *XhIconRecord, err error) {
	o := orm.NewOrm()
	v = &XhIconRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhIconRecord retrieves all XhIconRecord matches certain condition. Returns empty list if
// no records exist
func GetAllXhIconRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhIconRecord))
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

	var l []XhIconRecord
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

// UpdateXhIconRecord updates XhIconRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhIconRecordById(m *XhIconRecord) (err error) {
	o := orm.NewOrm()
	v := XhIconRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhIconRecord deletes XhIconRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhIconRecord(id int) (err error) {
	o := orm.NewOrm()
	v := XhIconRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhIconRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
