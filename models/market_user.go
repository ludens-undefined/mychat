package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MarketUser struct {
	Id          int    `orm:"column:id;autoIncrement"`
	ShopId      uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name        string `orm:"column:name;size:50"` //信息采集名称
	IsCync      uint8  `orm:"column:is_cync"` //同步用户资料 1开启 2关闭
	UserContent string `orm:"column:user_content;size:2000"` //联系人组件信息
	IsStart     uint8  `orm:"column:is_start"` //开启采集 1开启采集 2不开启采集
	Content     string `orm:"column:content;size:2000"` //商品内容，包含内容如type:1,id:1
	CreateAt    uint   `orm:"column:create_at"` //创建时间
	UpdateAt    uint   `orm:"column:update_at"` //修改时间
}

func (t *MarketUser) TableName() string {
	return "goouc_xmf_market_user"
}

func init() {
	orm.RegisterModel(new(MarketUser))
}

// AddMarketUser insert a new MarketUser into database and returns
// last inserted Id on success.
func AddMarketUser(m *MarketUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMarketUserById retrieves MarketUser by Id. Returns error if
// Id doesn't exist
func GetMarketUserById(id int) (v *MarketUser, err error) {
	o := orm.NewOrm()
	v = &MarketUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMarketUser retrieves all MarketUser matches certain condition. Returns empty list if
// no records exist
func GetAllMarketUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MarketUser))
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

	var l []MarketUser
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

// UpdateMarketUser updates MarketUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateMarketUserById(m *MarketUser) (err error) {
	o := orm.NewOrm()
	v := MarketUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMarketUser deletes MarketUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMarketUser(id int) (err error) {
	o := orm.NewOrm()
	v := MarketUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MarketUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
