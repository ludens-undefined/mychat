package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhIcon struct {
	Id            int    `orm:"column:id;autoIncrement"`
	Exctimes      int8   `orm:"column:exctimes"` //每年兑课券兑换次数
	Excrules      string `orm:"column:excrules;size:500"` //兑换币使用规则
	Regstatus     int8   `orm:"column:regstatus"` //注册奖励设置 1：开启  2：关闭
	Regnum        int8   `orm:"column:regnum"` //注册奖励数量
	Signstatus    int8   `orm:"column:signstatus"` //签到奖励设置 1：开启  2：关闭
	Signrules     string `orm:"column:signrules;size:500"` //签到奖励规则
	Signday       int8   `orm:"column:signday"` //签到奖励/每日
	Monthtimes    int8   `orm:"column:monthtimes"` //连续签到次数/每月
	Monthnum      int8   `orm:"column:monthnum"` //连续签到奖励/每月
	Sharestatus   int8   `orm:"column:sharestatus"` //分享设置
	SharedayTimes int8   `orm:"column:shareday_times"` //分享次数/每日
	SharedayNum   int8   `orm:"column:shareday_num"` //分享奖励/每日 兑换币数量
	Sharelimit    int8   `orm:"column:sharelimit"` //分享每日获币上限
	Imagenum      int8   `orm:"column:imagenum"` //头像送币数量
	Informnum     int8   `orm:"column:informnum"` //完善信息送币数量
	Buycourse     int8   `orm:"column:buycourse"` //购买课程送币数量
	Attendclass   uint8  `orm:"column:attendclass"` //按时上课送币数量
	Attendweek    int8   `orm:"column:attendweek"` //按时上课每周送币上限
	Playback      int8   `orm:"column:playback"` //回放单词货币数量
	Backweek      int8   `orm:"column:backweek"` //回放每周送币上限
	ShopId        int    `orm:"column:shop_id"`
	Type          int8   `orm:"column:type"` //1:兑换币设置    2：每日任务设置
	ExchangeNum   int    `orm:"column:exchange_num"` //1元需要多少兑换币
}

func (t *XhIcon) TableName() string {
	return "goouc_xmf_xh_icon"
}

func init() {
	orm.RegisterModel(new(XhIcon))
}

// AddXhIcon insert a new XhIcon into database and returns
// last inserted Id on success.
func AddXhIcon(m *XhIcon) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhIconById retrieves XhIcon by Id. Returns error if
// Id doesn't exist
func GetXhIconById(id int) (v *XhIcon, err error) {
	o := orm.NewOrm()
	v = &XhIcon{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhIcon retrieves all XhIcon matches certain condition. Returns empty list if
// no records exist
func GetAllXhIcon(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhIcon))
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

	var l []XhIcon
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

// UpdateXhIcon updates XhIcon by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhIconById(m *XhIcon) (err error) {
	o := orm.NewOrm()
	v := XhIcon{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhIcon deletes XhIcon by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhIcon(id int) (err error) {
	o := orm.NewOrm()
	v := XhIcon{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhIcon{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
