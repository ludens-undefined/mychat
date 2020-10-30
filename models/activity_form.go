package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ActivityForm struct {
	Id        int       `orm:"column:id;autoIncrement;"`
	ShopId    int       `orm:"column:shop_id"`
	Type      int       `orm:"column:type"` //1自定义表单
	CateId    int       `orm:"column:cate_id"` //分类
	Title     string    `orm:"column:title;size:255"` //标题
	Content   string    `orm:"column:content"` //介绍
	Price     float64   `orm:"column:price;scale:10;precision:2"` //价格
	Banner    string    `orm:"column:banner"` //轮播图
	BannerImg string    `orm:"column:banner_img;size:2000"`
	Agreement string    `orm:"column:agreement"` //协议
	Starttime int       `orm:"column:starttime"` //开始参与时间
	Startdate time.Time `orm:"column:startdate;type:time"` //结束时间
	Enddate   time.Time `orm:"column:enddate;type:time"` //结束时间
	KfPhone   string    `orm:"column:kf_phone;size:32"`
	Master    string    `orm:"column:master;size:255"`
	Address   string    `orm:"column:address;size:255"`
	Smsnotice int8      `orm:"column:smsnotice"` //是否短信通知
	IsHot     int       `orm:"column:is_hot"` //推荐 1是 2否
	Sort      int       `orm:"column:sort"` //排序 由小到大
	Status    int8      `orm:"column:status"` //状态 1成功 2失败
	IsDelete  int8      `orm:"column:is_delete"`
}

func (t *ActivityForm) TableName() string {
	return "goouc_xmf_activity_form"
}

func init() {
	orm.RegisterModel(new(ActivityForm))
}

// AddActivityForm insert a new ActivityForm into database and returns
// last inserted Id on success.
func AddActivityForm(m *ActivityForm) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetActivityFormById retrieves ActivityForm by Id. Returns error if
// Id doesn't exist
func GetActivityFormById(id int) (v *ActivityForm, err error) {
	o := orm.NewOrm()
	v = &ActivityForm{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllActivityForm retrieves all ActivityForm matches certain condition. Returns empty list if
// no records exist
func GetAllActivityForm(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ActivityForm))
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

	var l []ActivityForm
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

// UpdateActivityForm updates ActivityForm by Id and returns error if
// the record to be updated doesn't exist
func UpdateActivityFormById(m *ActivityForm) (err error) {
	o := orm.NewOrm()
	v := ActivityForm{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteActivityForm deletes ActivityForm by Id and returns error if
// the record to be deleted doesn't exist
func DeleteActivityForm(id int) (err error) {
	o := orm.NewOrm()
	v := ActivityForm{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ActivityForm{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
