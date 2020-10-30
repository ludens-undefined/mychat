package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjTkTest struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     uint   `orm:"column:shop_id"`
	Type       int    `orm:"column:type"` //题库分类,对应goouc_xmf_zxgj_tk_type表中的id,0未分类
	Title      string `orm:"column:title;size:1000"` //试题题目
	OptionType uint8  `orm:"column:option_type"` //1单选题 2多选题 3判断题 4填空题 5问答题
	Option     string `orm:"column:option"` //选项，格式：{[id:A,content:],[]}
	Answer     string `orm:"column:answer"` //正确答案
	Analysis   string `orm:"column:analysis"` //解析
	CreateAt   uint   `orm:"column:create_at"`
	UpdateAt   uint   `orm:"column:update_at"`
	IsDelete   uint8  `orm:"column:is_delete"` //1启用 2禁用
	DeleteAt   uint   `orm:"column:delete_at"`
	IsOrder    uint8  `orm:"column:is_order"` //1有序 2乱序
	AudioInfo  string `orm:"column:audio_info;size:5000"` //音频{[audio_name:,audio:]}
}

func (t *ZxgjTkTest) TableName() string {
	return "goouc_xmf_zxgj_tk_test"
}

func init() {
	orm.RegisterModel(new(ZxgjTkTest))
}

// AddZxgjTkTest insert a new ZxgjTkTest into database and returns
// last inserted Id on success.
func AddZxgjTkTest(m *ZxgjTkTest) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjTkTestById retrieves ZxgjTkTest by Id. Returns error if
// Id doesn't exist
func GetZxgjTkTestById(id int) (v *ZxgjTkTest, err error) {
	o := orm.NewOrm()
	v = &ZxgjTkTest{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjTkTest retrieves all ZxgjTkTest matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjTkTest(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjTkTest))
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

	var l []ZxgjTkTest
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

// UpdateZxgjTkTest updates ZxgjTkTest by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjTkTestById(m *ZxgjTkTest) (err error) {
	o := orm.NewOrm()
	v := ZxgjTkTest{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjTkTest deletes ZxgjTkTest by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjTkTest(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjTkTest{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjTkTest{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
