package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberOrderItem struct {
	Id         int     `orm:"column:id;autoIncrement"`
	ShopId     int     `orm:"column:shop_id"`
	OrderId    int     `orm:"column:order_id"`
	Type       int     `orm:"column:type"` //购买类型 1课程报名  2 视频课程
	Mcid       int     `orm:"column:mcid"` //用户班级id
	Dataid     int     `orm:"column:dataid"` //班级id或课程id（即：商品id）
	Payhave    int16   `orm:"column:payhave"` //购买课次
	Userid     int     `orm:"column:userid"`
	Status     int8    `orm:"column:status"` //订单状态 1待支付 2已付款 3申请退款 4已退款
	RefundType int8    `orm:"column:refund_type"` //退款类型：1：手动退款  2：自动退款
	Updown     int8    `orm:"column:updown"` //状态 1课程上 2课程下
	AllMoney   float64 `orm:"column:all_money;scale:10;precision:2"` //单项总金额
	TrueMoney  float64 `orm:"column:true_money;scale:10;precision:2"` //单项实付款
	AllPrice   float64 `orm:"column:all_price;scale:10;precision:2"` //单项原单价
	TruePrice  float64 `orm:"column:true_price;scale:10;precision:2"` //单项实际单价
	Createtime int     `orm:"column:createtime"`
	Paytime    int     `orm:"column:paytime"` //付款时间
	Applytime  int     `orm:"column:applytime"` //申请退款时间
	Endtime    int     `orm:"column:endtime"` //退款时间
	Msg        string  `orm:"column:msg;size:2000"`
	PayWay     string  `orm:"column:pay_way;size:255"` //支付方式
	Platform   string  `orm:"column:platform;size:255"` //操作平台
	CouponType int8    `orm:"column:coupon_type"` //优惠类型 0无优惠 1优惠券 2联报优惠 3续报优惠
	Couponid   int     `orm:"column:couponid"` //优惠券id
	IconMoney  int     `orm:"column:icon_money"` //所需兑换币数量
	OrderType  int8    `orm:"column:order_type"` //订单类型    1:原价  2：开团团购   3：拼团团购 4：兑课券+人民币 5：兑换币 6：后台导入
	Ticketid   int     `orm:"column:ticketid"` //兑课券id
	Groupid    int     `orm:"column:groupid"` //团id
	Oldid      int     `orm:"column:oldid"`
	Xoldid     int     `orm:"column:xoldid"`
}

func (t *MemberOrderItem) TableName() string {
	return "goouc_xmf_member_order_item"
}

func init() {
	orm.RegisterModel(new(MemberOrderItem))
}

// AddMemberOrderItem insert a new MemberOrderItem into database and returns
// last inserted Id on success.
func AddMemberOrderItem(m *MemberOrderItem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberOrderItemById retrieves MemberOrderItem by Id. Returns error if
// Id doesn't exist
func GetMemberOrderItemById(id int) (v *MemberOrderItem, err error) {
	o := orm.NewOrm()
	v = &MemberOrderItem{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberOrderItem retrieves all MemberOrderItem matches certain condition. Returns empty list if
// no records exist
func GetAllMemberOrderItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberOrderItem))
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

	var l []MemberOrderItem
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

// UpdateMemberOrderItem updates MemberOrderItem by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberOrderItemById(m *MemberOrderItem) (err error) {
	o := orm.NewOrm()
	v := MemberOrderItem{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberOrderItem deletes MemberOrderItem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberOrderItem(id int) (err error) {
	o := orm.NewOrm()
	v := MemberOrderItem{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberOrderItem{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
