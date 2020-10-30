package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberRefundLog struct {
	Id          int     `orm:"column:id;autoIncrement"`
	ShopId      int     `orm:"column:shop_id"`
	Userid      int     `orm:"column:userid"`
	Mcid        int     `orm:"column:mcid"` //用户课程id
	OitemId     int     `orm:"column:oitem_id"` //订单详情id
	RefundSn    string  `orm:"column:refund_sn;size:32"` //退款单号
	Createtime  int     `orm:"column:createtime"`
	Updatetime  int     `orm:"column:updatetime"`
	Checktime   int     `orm:"column:checktime"`
	Paytime     int     `orm:"column:paytime"`
	Price       float64 `orm:"column:price;scale:10;precision:2"` //应退金额
	TruePrice   float64 `orm:"column:true_price;scale:10;precision:2"` //核实金额
	IsMaterial  int8    `orm:"column:is_material"` //材料费 1无 2有
	MaterialFee float64 `orm:"column:material_fee;scale:10;precision:2"` //资料费
	Refundnum   int     `orm:"column:refundnum"` //退费课次
	Updatehave  int16   `orm:"column:updatehave"` //修正课次
	Status      int8    `orm:"column:status"` //状态 1待审核 2已拒绝 3待复核 4待退款 5已退款
	IsErr       int8    `orm:"column:is_err"` //状态 1正常退款 2特殊退费
	IsAfter     int8    `orm:"column:is_after"` //课后退
	IsRefund    int8    `orm:"column:is_refund"` //是否计入统计 1计入 2不计入
	Way         string  `orm:"column:way;size:32"` //退款方式
	Applicant   string  `orm:"column:applicant;size:32"` //申请人
	Operator    string  `orm:"column:operator;size:32"` //操作人
	Platform    string  `orm:"column:platform;size:32"` //操作平台
	Msg         string  `orm:"column:msg;size:2000"` //退款原因
}

func (t *MemberRefundLog) TableName() string {
	return "goouc_xmf_member_refund_log"
}

func init() {
	orm.RegisterModel(new(MemberRefundLog))
}

// AddMemberRefundLog insert a new MemberRefundLog into database and returns
// last inserted Id on success.
func AddMemberRefundLog(m *MemberRefundLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberRefundLogById retrieves MemberRefundLog by Id. Returns error if
// Id doesn't exist
func GetMemberRefundLogById(id int) (v *MemberRefundLog, err error) {
	o := orm.NewOrm()
	v = &MemberRefundLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberRefundLog retrieves all MemberRefundLog matches certain condition. Returns empty list if
// no records exist
func GetAllMemberRefundLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberRefundLog))
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

	var l []MemberRefundLog
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

// UpdateMemberRefundLog updates MemberRefundLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberRefundLogById(m *MemberRefundLog) (err error) {
	o := orm.NewOrm()
	v := MemberRefundLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberRefundLog deletes MemberRefundLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberRefundLog(id int) (err error) {
	o := orm.NewOrm()
	v := MemberRefundLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberRefundLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
