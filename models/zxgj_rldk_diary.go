package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZxgjRldkDiary struct {
	Id       int       `orm:"column:id;autoIncrement"`
	RldkId   uint      `orm:"column:rldk_id"` //对应goouc_xet_zxgj_rldk表中的id
	UserId   uint      `orm:"column:user_id"` //打卡者，对应goouc_xet_user表中的id
	Date     time.Time `orm:"column:date;type(date)"` //打卡日期
	Content  string    `orm:"column:content;size:255"` //打卡日记内容路径
	IsGood   uint8     `orm:"column:is_good"` //1加入精选 2不加入精选
	IsDelete uint8     `orm:"column:is_delete"` //状态 1启用 2删除
	CreateAt uint      `orm:"column:create_at"` //打卡时间
	DeleteAt uint      `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjRldkDiary) TableName() string {
	return "goouc_xmf_zxgj_rldk_diary"
}

func init() {
	orm.RegisterModel(new(ZxgjRldkDiary))
}

// AddZxgjRldkDiary insert a new ZxgjRldkDiary into database and returns
// last inserted Id on success.
func AddZxgjRldkDiary(m *ZxgjRldkDiary) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjRldkDiaryById retrieves ZxgjRldkDiary by Id. Returns error if
// Id doesn't exist
func GetZxgjRldkDiaryById(id int) (v *ZxgjRldkDiary, err error) {
	o := orm.NewOrm()
	v = &ZxgjRldkDiary{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjRldkDiary retrieves all ZxgjRldkDiary matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjRldkDiary(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjRldkDiary))
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

	var l []ZxgjRldkDiary
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

// UpdateZxgjRldkDiary updates ZxgjRldkDiary by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjRldkDiaryById(m *ZxgjRldkDiary) (err error) {
	o := orm.NewOrm()
	v := ZxgjRldkDiary{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjRldkDiary deletes ZxgjRldkDiary by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjRldkDiary(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjRldkDiary{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjRldkDiary{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
