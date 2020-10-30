package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZs struct {
	Id       int    `orm:"column:id;autoIncrement"`
	ShopId   uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	ExamId   uint   `orm:"column:exam_id"` //考试id，对应goouc_xet_zxgj_exam表中的id
	IsStart  uint8  `orm:"column:is_start"` //1开启考试证书 2不开启
	Score    uint8  `orm:"column:score"` //考试成绩最小值
	Name     string `orm:"column:name;size:50"` //名称
	BakImg   string `orm:"column:bak_img;size:255"` //证书背景图片
	BakType  int8   `orm:"column:bak_type"` //1:固定  2：自选
	Content  string `orm:"column:content;size:2000"` //得分设置，格式：{{lower:,upper:,comment:},{}}
	Code     string `orm:"column:code;size:255"` //二维码
	Share    string `orm:"column:share;size:50"` //分享语
	IsDelete uint8  `orm:"column:is_delete"` //状态 1启用 2删除
	CreateAt uint   `orm:"column:create_at"` //创建时间
	UpdateAt uint   `orm:"column:update_at"` //修改时间
}

func (t *ZxgjZs) TableName() string {
	return "goouc_xmf_zxgj_zs"
}

func init() {
	orm.RegisterModel(new(ZxgjZs))
}

// AddZxgjZs insert a new ZxgjZs into database and returns
// last inserted Id on success.
func AddZxgjZs(m *ZxgjZs) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZsById retrieves ZxgjZs by Id. Returns error if
// Id doesn't exist
func GetZxgjZsById(id int) (v *ZxgjZs, err error) {
	o := orm.NewOrm()
	v = &ZxgjZs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZs retrieves all ZxgjZs matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZs))
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

	var l []ZxgjZs
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

// UpdateZxgjZs updates ZxgjZs by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZsById(m *ZxgjZs) (err error) {
	o := orm.NewOrm()
	v := ZxgjZs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZs deletes ZxgjZs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZs(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
