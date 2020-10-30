package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Slide struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     int    `orm:"column:shop_id"`
	Model      int    `orm:"column:model"` //1教务系统 2视频学院
	Title      string `orm:"column:title;size:32"`
	Image      string `orm:"column:image;size:2000"`
	Position   int8   `orm:"column:position"` //显示位置 1PC 2H5 3小程序
	PcPosition int8   `orm:"column:pc_position"` //官网显示位置 1首页,2教学体系,3师资力量,4加入我们
	Status     int8   `orm:"column:status"` //是否显示 1显示 2隐藏
	Type       int8   `orm:"column:type"` //跳转方式
	Link       string `orm:"column:link;size:255"` //跳转地址   1：外部链接 2：课程   3:排行榜
	Grade      string `orm:"column:grade;size:255"` //关联年级
	Sort       int    `orm:"column:sort"`
	Createtime int    `orm:"column:createtime"`
	Updatetime int    `orm:"column:updatetime"`
	IsDelete   int8   `orm:"column:is_delete"` //状态 1显示 2隐藏
}

func (t *Slide) TableName() string {
	return "goouc_xmf_slide"
}

func init() {
	orm.RegisterModel(new(Slide))
}

// AddSlide insert a new Slide into database and returns
// last inserted Id on success.
func AddSlide(m *Slide) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSlideById retrieves Slide by Id. Returns error if
// Id doesn't exist
func GetSlideById(id int) (v *Slide, err error) {
	o := orm.NewOrm()
	v = &Slide{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSlide retrieves all Slide matches certain condition. Returns empty list if
// no records exist
func GetAllSlide(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Slide))
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

	var l []Slide
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

// UpdateSlide updates Slide by Id and returns error if
// the record to be updated doesn't exist
func UpdateSlideById(m *Slide) (err error) {
	o := orm.NewOrm()
	v := Slide{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSlide deletes Slide by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSlide(id int) (err error) {
	o := orm.NewOrm()
	v := Slide{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Slide{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
