package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberTicket struct {
	Id        int  `orm:"column:id;autoIncrement"`
	ShopId    int  `orm:"column:shop_id"`
	Memberid  int  `orm:"column:memberid"` //学员id
	Ticketid  int  `orm:"column:ticketid"` //兑课券种类id
	Status    int8 `orm:"column:status"` //使用状态状态 2已使用 1 未使用  3：已过期
	Year      int  `orm:"column:year"` //兑换年份(一年兑换次数用)
	Addtime   int  `orm:"column:addtime"` //兑换时间
	Icon      int  `orm:"column:icon"` //消耗    币的数量
	Courseid  int  `orm:"column:courseid"` //课程id
	Chapterid int  `orm:"column:chapterid"` //章节id
	Usetime   int  `orm:"column:usetime"` //使用时间
	IsDelete  int8 `orm:"column:is_delete"` //1:启用   2：删除
}

func (t *MemberTicket) TableName() string {
	return "goouc_xmf_member_ticket"
}

func init() {
	orm.RegisterModel(new(MemberTicket))
}

// AddMemberTicket insert a new MemberTicket into database and returns
// last inserted Id on success.
func AddMemberTicket(m *MemberTicket) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberTicketById retrieves MemberTicket by Id. Returns error if
// Id doesn't exist
func GetMemberTicketById(id int) (v *MemberTicket, err error) {
	o := orm.NewOrm()
	v = &MemberTicket{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberTicket retrieves all MemberTicket matches certain condition. Returns empty list if
// no records exist
func GetAllMemberTicket(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberTicket))
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

	var l []MemberTicket
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

// UpdateMemberTicket updates MemberTicket by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberTicketById(m *MemberTicket) (err error) {
	o := orm.NewOrm()
	v := MemberTicket{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberTicket deletes MemberTicket by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberTicket(id int) (err error) {
	o := orm.NewOrm()
	v := MemberTicket{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberTicket{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
