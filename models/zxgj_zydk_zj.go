package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZydkZj struct {
	Id       int   `orm:"column:id;autoIncrement"`
	UserId   uint  `orm:"column:user_id"` //用户id，对应goouc_xet_user表中的id
	ZydkId   uint  `orm:"column:zydk_id"` //对应goouc_xet_zxgj_zydk表中的id
	Zan      uint  `orm:"column:zan"` //总点赞数
	Dp       uint  `orm:"column:dp"` //总点评数
	Jx       uint  `orm:"column:jx"` //精选数
	IsDelete uint8 `orm:"column:is_delete"` //状态 1启用 2删除
	DeleteAt uint  `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjZydkZj) TableName() string {
	return "goouc_xmf_zxgj_zydk_zj"
}

func init() {
	orm.RegisterModel(new(ZxgjZydkZj))
}

// AddZxgjZydkZj insert a new ZxgjZydkZj into database and returns
// last inserted Id on success.
func AddZxgjZydkZj(m *ZxgjZydkZj) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZydkZjById retrieves ZxgjZydkZj by Id. Returns error if
// Id doesn't exist
func GetZxgjZydkZjById(id int) (v *ZxgjZydkZj, err error) {
	o := orm.NewOrm()
	v = &ZxgjZydkZj{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZydkZj retrieves all ZxgjZydkZj matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZydkZj(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZydkZj))
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

	var l []ZxgjZydkZj
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

// UpdateZxgjZydkZj updates ZxgjZydkZj by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZydkZjById(m *ZxgjZydkZj) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydkZj{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZydkZj deletes ZxgjZydkZj by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZydkZj(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydkZj{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZydkZj{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
