package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjTk struct {
	Id       int    `orm:"column:id;autoIncrement"`
	ShopId   uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	TypeId   uint   `orm:"column:type_id"` //题目类型，对应goouc_xet_zxgj_tk_type表中的id
	TtypeId  uint   `orm:"column:ttype_id"` //题目分类id，对应goouc_xet_zxgj_tk_ttype表中的id
	Title    string `orm:"column:title;size:255"` //题目
	Content  string `orm:"column:content;size:2000"` //选项内容，格式为：{{option:1,content:,answer:},{}}
	Jiexi    string `orm:"column:jiexi;size:2000"` //解析
	Audio    string `orm:"column:audio;size:255"` //类型为问答题，音频
	IsOrder  uint8  `orm:"column:is_order"` //1不乱序匹配答案 2乱序匹配答案，对应类型为填空题时
	CreateAt uint   `orm:"column:create_at"` //创建时间
	UpdateAt uint   `orm:"column:update_at"` //修改时间
}

func (t *ZxgjTk) TableName() string {
	return "goouc_xmf_zxgj_tk"
}

func init() {
	orm.RegisterModel(new(ZxgjTk))
}

// AddZxgjTk insert a new ZxgjTk into database and returns
// last inserted Id on success.
func AddZxgjTk(m *ZxgjTk) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjTkById retrieves ZxgjTk by Id. Returns error if
// Id doesn't exist
func GetZxgjTkById(id int) (v *ZxgjTk, err error) {
	o := orm.NewOrm()
	v = &ZxgjTk{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjTk retrieves all ZxgjTk matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjTk(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjTk))
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

	var l []ZxgjTk
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

// UpdateZxgjTk updates ZxgjTk by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjTkById(m *ZxgjTk) (err error) {
	o := orm.NewOrm()
	v := ZxgjTk{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjTk deletes ZxgjTk by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjTk(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjTk{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjTk{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
