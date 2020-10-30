package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhPc struct {
	Id        int    `orm:"column:id;autoIncrement"`
	Name      string `orm:"column:name;size:255"` //网站名称
	PcAppid   string `orm:"column:pc_appid;size:50"` //pc_appid
	PcSecret  string `orm:"column:pc_secret;size:255"` //pc_secret
	Icon      string `orm:"column:icon;size:255"` //icon图标
	Keyword   string `orm:"column:keyword;size:255"` //搜索关键字
	LogoTitle string `orm:"column:logo_title;size:255"` //logo标题
	Logo      string `orm:"column:logo;size:255"` //logo
	Publicqr  string `orm:"column:publicqr;size:255"` //公众号二维码
	Smallqr   string `orm:"column:smallqr;size:255"` //小程序二维码
	Desc      string `orm:"column:desc;size:255"` //网站描述
	Qq        string `orm:"column:qq;size:255"` //网站qq
	Address   string `orm:"column:address;size:255"` //地址
	Contact   string `orm:"column:contact;size:255"` //网站联系方式
	Beian     string `orm:"column:beian;size:255"` //网站备案信息
	ShopId    int    `orm:"column:shop_id"` //店铺id
}

func (t *XhPc) TableName() string {
	return "goouc_xmf_xh_pc"
}

func init() {
	orm.RegisterModel(new(XhPc))
}

// AddXhPc insert a new XhPc into database and returns
// last inserted Id on success.
func AddXhPc(m *XhPc) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhPcById retrieves XhPc by Id. Returns error if
// Id doesn't exist
func GetXhPcById(id int) (v *XhPc, err error) {
	o := orm.NewOrm()
	v = &XhPc{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhPc retrieves all XhPc matches certain condition. Returns empty list if
// no records exist
func GetAllXhPc(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhPc))
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

	var l []XhPc
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

// UpdateXhPc updates XhPc by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhPcById(m *XhPc) (err error) {
	o := orm.NewOrm()
	v := XhPc{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhPc deletes XhPc by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhPc(id int) (err error) {
	o := orm.NewOrm()
	v := XhPc{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhPc{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
