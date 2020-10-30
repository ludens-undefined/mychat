package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type BusinessShopRz struct {
	Id       int    `orm:"column:id;autoIncrement"`
	Type     uint8  `orm:"column:type"` //1个人认证 2企业认证 3其他机构认证
	Name     string `orm:"column:name;size:30"` //姓名|企业名称|机构名称
	Leading  string `orm:"column:leading;size:30"` //type为3，负责人
	ZjType   uint8  `orm:"column:zj_type"` //证件类型 1中国大陆居民身份证
	Faren    string `orm:"column:faren;size:30"` //对应type为2,法人
	Card     string `orm:"column:card;size:30"` //身份证号码
	CardPic  string `orm:"column:card_pic;size:500"` //身份证照片
	HandPic  string `orm:"column:hand_pic;size:255"` //手持证件照
	Code     string `orm:"column:code;size:50"` //type为2，统一社会信用代码
	License  string `orm:"column:license;size:255"` //type为2营业执照
	Pictures string `orm:"column:pictures;size:1000"` //type为3，机构证件或其他证明材料
}

func (t *BusinessShopRz) TableName() string {
	return "goouc_xmf_business_shop_rz"
}

func init() {
	orm.RegisterModel(new(BusinessShopRz))
}

// AddBusinessShopRz insert a new BusinessShopRz into database and returns
// last inserted Id on success.
func AddBusinessShopRz(m *BusinessShopRz) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBusinessShopRzById retrieves BusinessShopRz by Id. Returns error if
// Id doesn't exist
func GetBusinessShopRzById(id int) (v *BusinessShopRz, err error) {
	o := orm.NewOrm()
	v = &BusinessShopRz{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBusinessShopRz retrieves all BusinessShopRz matches certain condition. Returns empty list if
// no records exist
func GetAllBusinessShopRz(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BusinessShopRz))
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

	var l []BusinessShopRz
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

// UpdateBusinessShopRz updates BusinessShopRz by Id and returns error if
// the record to be updated doesn't exist
func UpdateBusinessShopRzById(m *BusinessShopRz) (err error) {
	o := orm.NewOrm()
	v := BusinessShopRz{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBusinessShopRz deletes BusinessShopRz by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBusinessShopRz(id int) (err error) {
	o := orm.NewOrm()
	v := BusinessShopRz{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BusinessShopRz{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
