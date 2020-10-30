package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZydkContent struct {
	Id      int    `orm:"column:id;autoIncrement"`
	ZydkId  uint   `orm:"column:zydk_id"` //对应goouc_xet_zxgj_zydk表中的id
	Type    uint8  `orm:"column:type"` //1图文 2音频 3视频 4课程
	Content string `orm:"column:content;size:255"` //type为1,2,3对应路径，若为4课程格式为：type:1,id:1，其中type对应goouc_xet_cource_type表中的id
}

func (t *ZxgjZydkContent) TableName() string {
	return "goouc_xmf_zxgj_zydk_content"
}

func init() {
	orm.RegisterModel(new(ZxgjZydkContent))
}

// AddZxgjZydkContent insert a new ZxgjZydkContent into database and returns
// last inserted Id on success.
func AddZxgjZydkContent(m *ZxgjZydkContent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZydkContentById retrieves ZxgjZydkContent by Id. Returns error if
// Id doesn't exist
func GetZxgjZydkContentById(id int) (v *ZxgjZydkContent, err error) {
	o := orm.NewOrm()
	v = &ZxgjZydkContent{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZydkContent retrieves all ZxgjZydkContent matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZydkContent(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZydkContent))
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

	var l []ZxgjZydkContent
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

// UpdateZxgjZydkContent updates ZxgjZydkContent by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZydkContentById(m *ZxgjZydkContent) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydkContent{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZydkContent deletes ZxgjZydkContent by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZydkContent(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydkContent{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZydkContent{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
