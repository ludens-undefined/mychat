package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZybZj struct {
	Id       int   `orm:"column:id;autoIncrement"`
	UserId   uint  `orm:"column:user_id"` //助教，对应goouc_xet_user表中的id
	ZybId    uint  `orm:"column:zyb_id"` //对应goouc_xet_zxgj_zyb表中的id
	Zan      uint  `orm:"column:zan"` //总点赞数
	Dp       uint  `orm:"column:dp"` //总点评数
	Jx       uint  `orm:"column:jx"` //精选数
	IsDelete uint8 `orm:"column:is_delete"` //状态 1启用 2删除
	DeleteAt uint  `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjZybZj) TableName() string {
	return "goouc_xmf_zxgj_zyb_zj"
}

func init() {
	orm.RegisterModel(new(ZxgjZybZj))
}

// AddZxgjZybZj insert a new ZxgjZybZj into database and returns
// last inserted Id on success.
func AddZxgjZybZj(m *ZxgjZybZj) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZybZjById retrieves ZxgjZybZj by Id. Returns error if
// Id doesn't exist
func GetZxgjZybZjById(id int) (v *ZxgjZybZj, err error) {
	o := orm.NewOrm()
	v = &ZxgjZybZj{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZybZj retrieves all ZxgjZybZj matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZybZj(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZybZj))
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

	var l []ZxgjZybZj
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

// UpdateZxgjZybZj updates ZxgjZybZj by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZybZjById(m *ZxgjZybZj) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZj{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZybZj deletes ZxgjZybZj by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZybZj(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZj{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZybZj{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
