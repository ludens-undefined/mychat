package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ConfigSet struct {
	Id          int    `orm:"column:id;autoIncrement"`
	ShopId      uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	IsSub       int8   `orm:"column:is_sub"` //隐藏订阅数 1显示 2隐藏
	IsUpdate    uint8  `orm:"column:is_update"` //隐藏更新期数 1显示 2隐藏
	IsView      uint8  `orm:"column:is_view"` //隐藏播放量/浏览量 1显示 2隐藏
	IsRemark    uint8  `orm:"column:is_remark"` //评论审核设置 1关闭 2隐藏
	IsSend      uint8  `orm:"column:is_send"` //服务号通知 1开启 2关闭
	ForbId      uint   `orm:"column:forb_id"` //对应goouc_xet_config_set_forb表中的id,0没设置
	IsShuiyin   uint8  `orm:"column:is_shuiyin"` //水印显示 1显示 2隐藏
	Img         string `orm:"column:img;size:255"` //水印图片
	Pos         uint8  `orm:"column:pos"` //水印位置1左上角 2右上角 3左下角 4右下角
	Transparent uint8  `orm:"column:transparent"` //透明度
	Size        uint8  `orm:"column:size"` //水印大小
}

func (t *ConfigSet) TableName() string {
	return "goouc_xmf_config_set"
}

func init() {
	orm.RegisterModel(new(ConfigSet))
}

// AddConfigSet insert a new ConfigSet into database and returns
// last inserted Id on success.
func AddConfigSet(m *ConfigSet) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetConfigSetById retrieves ConfigSet by Id. Returns error if
// Id doesn't exist
func GetConfigSetById(id int) (v *ConfigSet, err error) {
	o := orm.NewOrm()
	v = &ConfigSet{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllConfigSet retrieves all ConfigSet matches certain condition. Returns empty list if
// no records exist
func GetAllConfigSet(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ConfigSet))
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

	var l []ConfigSet
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

// UpdateConfigSet updates ConfigSet by Id and returns error if
// the record to be updated doesn't exist
func UpdateConfigSetById(m *ConfigSet) (err error) {
	o := orm.NewOrm()
	v := ConfigSet{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteConfigSet deletes ConfigSet by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConfigSet(id int) (err error) {
	o := orm.NewOrm()
	v := ConfigSet{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ConfigSet{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
