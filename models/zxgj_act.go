package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjAct struct {
	Id            int    `orm:"column:id;autoIncrement"`
	ShopId        uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name          string `orm:"column:name;size:50"` //活动名称
	Addr          string `orm:"column:addr;size:255"` //活动地点
	ActStartTime  uint   `orm:"column:act_start_time"` //活动开始时间
	ActEndTime    uint   `orm:"column:act_end_time"` //活动结束时间
	BmType        uint8  `orm:"column:bm_type"` //报名类型 1默认 2自定义
	BmStartTime   uint   `orm:"column:bm_start_time"` //对应bm_type=2，报名开始时间
	BmEndTime     uint   `orm:"column:bm_end_time"` //报名结束时间
	Img           string `orm:"column:img;size:255"` //活动海报
	Detail        string `orm:"column:detail;size:100"` //活动详情路径
	Class         string `orm:"column:class;size:50"` //所属课程{type:,id:}type为goouc_xet_cource_type表中的id
	IsShow        uint8  `orm:"column:is_show"` //报名人数显示 1开启 2关闭
	IsSend        uint8  `orm:"column:is_send"` //提醒通知设置 1开启 2关闭
	IsCollectUser uint8  `orm:"column:is_collect_user"` //报名信息收集 1开启 2关闭
	Content       string `orm:"column:content;size:2000"` //报名信息
	TicketNum     uint16 `orm:"column:ticket_num"` //票券数量
	IsHide        uint8  `orm:"column:is_hide"` //1显示 2隐藏活动
	IsCancel      uint8  `orm:"column:is_cancel"` //1不取消 2取消活动
	CreateAt      uint   `orm:"column:create_at"` //创建时间
	UpdateAt      uint   `orm:"column:update_at"` //修改时间
}

func (t *ZxgjAct) TableName() string {
	return "goouc_xmf_zxgj_act"
}

func init() {
	orm.RegisterModel(new(ZxgjAct))
}

// AddZxgjAct insert a new ZxgjAct into database and returns
// last inserted Id on success.
func AddZxgjAct(m *ZxgjAct) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjActById retrieves ZxgjAct by Id. Returns error if
// Id doesn't exist
func GetZxgjActById(id int) (v *ZxgjAct, err error) {
	o := orm.NewOrm()
	v = &ZxgjAct{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjAct retrieves all ZxgjAct matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjAct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjAct))
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

	var l []ZxgjAct
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

// UpdateZxgjAct updates ZxgjAct by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjActById(m *ZxgjAct) (err error) {
	o := orm.NewOrm()
	v := ZxgjAct{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjAct deletes ZxgjAct by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjAct(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjAct{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjAct{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
