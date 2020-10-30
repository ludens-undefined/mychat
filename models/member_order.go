package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberOrder struct {
	Id            int     `orm:"column:id;autoIncrement"`
	ShopId        int     `orm:"column:shop_id"`
	OutTradeNo    string  `orm:"column:out_trade_no;size:255"`
	Userid        int     `orm:"column:userid"`
	Xhopenid      string  `orm:"column:xhopenid;size:255"`
	Status        int8    `orm:"column:status"` //订单状态 1待支付 2已付款
	AllMoney      float64 `orm:"column:all_money;scale:10;precision:2"` //总金额
	TrueMoney     float64 `orm:"column:true_money;scale:10;precision:2"` //实付款
	LessMoney     float64 `orm:"column:less_money;scale:10;precision:2"` //优惠金额
	CouponType    int8    `orm:"column:coupon_type"` //优惠类型 0无优惠 1优惠券 2联报优惠 3续报优惠
	Couponid      int     `orm:"column:couponid"` //优惠id
	Createtime    int     `orm:"column:createtime"`
	Paytime       int     `orm:"column:paytime"`
	Msg           string  `orm:"column:msg;size:255"`
	TransactionSn string  `orm:"column:transaction_sn;size:255"` //第三方流水号
	PayWay        string  `orm:"column:pay_way;size:255"` //支付方式
	PayCode       string  `orm:"column:pay_code;size:32"`
	Paytype       int8    `orm:"column:paytype"` //1银联扫码 2银联H5
	Platform      string  `orm:"column:platform;size:255"` //操作平台
	Operator      string  `orm:"column:operator;size:32"` //操作人
	Type          int8    `orm:"column:type"` //1：教务系统订单  2：视频学院订单
	OrderType     int8    `orm:"column:order_type"` //订单类型    1:原价  2：开团团购   3：拼团团购 4：兑课券+人民币 5：兑换币 6：后台导入 
	IconMoney     int     `orm:"column:icon_money"` //所需兑换币数量
	Ticketid      int     `orm:"column:ticketid"` //兑课券id
	Groupid       int     `orm:"column:groupid"` //团id
	Oldid         int     `orm:"column:oldid"`
	Xoldid        int     `orm:"column:xoldid"`
	Remarks       string  `orm:"column:remarks;size:2000"` //备注
	Oriwtorderid  string  `orm:"column:oriwtorderid;size:50"`
}

func (t *MemberOrder) TableName() string {
	return "goouc_xmf_member_order"
}

func init() {
	orm.RegisterModel(new(MemberOrder))
}

// AddMemberOrder insert a new MemberOrder into database and returns
// last inserted Id on success.
func AddMemberOrder(m *MemberOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberOrderById retrieves MemberOrder by Id. Returns error if
// Id doesn't exist
func GetMemberOrderById(id int) (v *MemberOrder, err error) {
	o := orm.NewOrm()
	v = &MemberOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberOrder retrieves all MemberOrder matches certain condition. Returns empty list if
// no records exist
func GetAllMemberOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberOrder))
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

	var l []MemberOrder
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

// UpdateMemberOrder updates MemberOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberOrderById(m *MemberOrder) (err error) {
	o := orm.NewOrm()
	v := MemberOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberOrder deletes MemberOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberOrder(id int) (err error) {
	o := orm.NewOrm()
	v := MemberOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
