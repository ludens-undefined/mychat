package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MarketMessGood struct {
	Id     int  `orm:"column:id;autoIncrement"`
	MessId uint `orm:"column:mess_id"`
	Type   uint `orm:"column:type"` //对应goouc_xmf_cource_type表中的id
	GoodId uint `orm:"column:good_id"` //与type想对应的商品id
}

func (t *MarketMessGood) TableName() string {
	return "goouc_xmf_market_mess_good"
}

func init() {
	orm.RegisterModel(new(MarketMessGood))
}

// AddMarketMessGood insert a new MarketMessGood into database and returns
// last inserted Id on success.
func AddMarketMessGood(m *MarketMessGood) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMarketMessGoodById retrieves MarketMessGood by Id. Returns error if
// Id doesn't exist
func GetMarketMessGoodById(id int) (v *MarketMessGood, err error) {
	o := orm.NewOrm()
	v = &MarketMessGood{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMarketMessGood retrieves all MarketMessGood matches certain condition. Returns empty list if
// no records exist
func GetAllMarketMessGood(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MarketMessGood))
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

	var l []MarketMessGood
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

// UpdateMarketMessGood updates MarketMessGood by Id and returns error if
// the record to be updated doesn't exist
func UpdateMarketMessGoodById(m *MarketMessGood) (err error) {
	o := orm.NewOrm()
	v := MarketMessGood{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMarketMessGood deletes MarketMessGood by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMarketMessGood(id int) (err error) {
	o := orm.NewOrm()
	v := MarketMessGood{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MarketMessGood{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
