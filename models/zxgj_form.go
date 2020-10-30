package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjForm struct {
	Id          int    `orm:"column:id;autoIncrement"`
	ShopId      uint   `orm:"column:shop_id"`
	Name        string `orm:"column:name;size:50"` //表单名称
	Dep         string `orm:"column:dep;size:50"` //描述内容
	BgImg       string `orm:"column:bg_img;size:255"` //背景图片
	Content     string `orm:"column:content;size:2000"` //表单内容
	ShareTitle  string `orm:"column:share_title;size:50"` //分享标题
	ShareDep    string `orm:"column:share_dep;size:255"` //分享描述
	SharePic    string `orm:"column:share_pic;size:255"` //分享图片
	CollectType uint8  `orm:"column:collect_type"` //1定量回收 2定时回收
	Num         uint   `orm:"column:num"` //collect_type为1时，收集表单份数
	Tips        string `orm:"column:tips;size:255"` //collect_type为1时，提示信息
	TimeType    uint8  `orm:"column:time_type"` //对应collect_type为2时，1永久有效 2定时回收
	Time        uint   `orm:"column:time"` //回收时间
	IsSub       uint8  `orm:"column:is_sub"` //1发布 2暂不发布 3已结束
	IsDelete    uint8  `orm:"column:is_delete"` //1启用 2删除
	CreateAt    uint   `orm:"column:create_at"` //创建时间
	UpdateAt    uint   `orm:"column:update_at"` //修改时间
	DeleteAt    uint   `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjForm) TableName() string {
	return "goouc_xmf_zxgj_form"
}

func init() {
	orm.RegisterModel(new(ZxgjForm))
}

// AddZxgjForm insert a new ZxgjForm into database and returns
// last inserted Id on success.
func AddZxgjForm(m *ZxgjForm) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjFormById retrieves ZxgjForm by Id. Returns error if
// Id doesn't exist
func GetZxgjFormById(id int) (v *ZxgjForm, err error) {
	o := orm.NewOrm()
	v = &ZxgjForm{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjForm retrieves all ZxgjForm matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjForm(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjForm))
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

	var l []ZxgjForm
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

// UpdateZxgjForm updates ZxgjForm by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjFormById(m *ZxgjForm) (err error) {
	o := orm.NewOrm()
	v := ZxgjForm{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjForm deletes ZxgjForm by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjForm(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjForm{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjForm{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
