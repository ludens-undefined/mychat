package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhMemberCourse struct {
	Id         int     `orm:"column:id;autoIncrement"`
	Memberid   int     `orm:"column:memberid"` //学员id
	Xhcourseid int     `orm:"column:xhcourseid"` //课程id
	ShopId     int     `orm:"column:shop_id"`
	Paytype    int8    `orm:"column:paytype"` //付款类型   1：免费 2：付费（晓禾币）3：付费（全款人民币）4：团购  5：后台导入
	Status     int8    `orm:"column:status"` //学习状态 2：未开始学习  1 ：已开始学习
	Iconnum    int     `orm:"column:iconnum"` //晓禾币数量
	Groupprice float64 `orm:"column:groupprice;scale:11;precision:2"` //团购价格
	TrueMoney  float64 `orm:"column:true_money;scale:11;precision:2"` //用户全款购买的价格
	IconMoney  int     `orm:"column:icon_money"` //需要兑换币数量
	Ticketid   int     `orm:"column:ticketid"` //兑换券id
	Ctime      int     `orm:"column:ctime"` //创建时间
	IsDelete   int8    `orm:"column:is_delete"` //1:启用   2：删除
	Xoldid     int     `orm:"column:xoldid"`
}

func (t *XhMemberCourse) TableName() string {
	return "goouc_xmf_xh_member_course"
}

func init() {
	orm.RegisterModel(new(XhMemberCourse))
}

// AddXhMemberCourse insert a new XhMemberCourse into database and returns
// last inserted Id on success.
func AddXhMemberCourse(m *XhMemberCourse) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhMemberCourseById retrieves XhMemberCourse by Id. Returns error if
// Id doesn't exist
func GetXhMemberCourseById(id int) (v *XhMemberCourse, err error) {
	o := orm.NewOrm()
	v = &XhMemberCourse{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhMemberCourse retrieves all XhMemberCourse matches certain condition. Returns empty list if
// no records exist
func GetAllXhMemberCourse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhMemberCourse))
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

	var l []XhMemberCourse
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

// UpdateXhMemberCourse updates XhMemberCourse by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhMemberCourseById(m *XhMemberCourse) (err error) {
	o := orm.NewOrm()
	v := XhMemberCourse{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhMemberCourse deletes XhMemberCourse by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhMemberCourse(id int) (err error) {
	o := orm.NewOrm()
	v := XhMemberCourse{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhMemberCourse{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
