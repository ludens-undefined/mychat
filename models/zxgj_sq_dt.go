package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjSqDt struct {
	Id       int    `orm:"column:id;autoIncrement"`
	SqId     uint   `orm:"column:sq_id"` //社群id,对应goouc_xet_zxgj_sq表中的id
	Name     string `orm:"column:name;size:50"` //动态名称
	Content  string `orm:"column:content;size:255"` //动态内容路径，多个用，分隔
	IsSend   uint8  `orm:"column:is_send"` //服务号通知 1开启通知 2关闭
	IsGood   uint8  `orm:"column:is_good"` //1加入精选 2不加入精选
	IsGg     uint8  `orm:"column:is_gg"` //1设为公告 2不设为公告
	Remark   uint16 `orm:"column:remark"` //评论数
	Zan      uint   `orm:"column:zan"` //点赞数
	IsDelete uint8  `orm:"column:is_delete"` //状态 1启用 2删除
	CreateAt uint   `orm:"column:create_at"` //创建时间
	UpdateAt uint   `orm:"column:update_at"` //修改时间
}

func (t *ZxgjSqDt) TableName() string {
	return "goouc_xmf_zxgj_sq_dt"
}

func init() {
	orm.RegisterModel(new(ZxgjSqDt))
}

// AddZxgjSqDt insert a new ZxgjSqDt into database and returns
// last inserted Id on success.
func AddZxgjSqDt(m *ZxgjSqDt) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjSqDtById retrieves ZxgjSqDt by Id. Returns error if
// Id doesn't exist
func GetZxgjSqDtById(id int) (v *ZxgjSqDt, err error) {
	o := orm.NewOrm()
	v = &ZxgjSqDt{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjSqDt retrieves all ZxgjSqDt matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjSqDt(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjSqDt))
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

	var l []ZxgjSqDt
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

// UpdateZxgjSqDt updates ZxgjSqDt by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjSqDtById(m *ZxgjSqDt) (err error) {
	o := orm.NewOrm()
	v := ZxgjSqDt{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjSqDt deletes ZxgjSqDt by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjSqDt(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjSqDt{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjSqDt{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
