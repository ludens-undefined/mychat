package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Member struct {
	Id             int       `orm:"column:id;autoIncrement"`
	ShopId         int       `orm:"column:shop_id"`
	Openid         string    `orm:"column:openid;size:255"` //公众号openid
	Unionid        string    `orm:"column:unionid;size:255"`
	SessionKey     string    `orm:"column:session_key;size:255"`
	Username       string    `orm:"column:username;size:255"`
	Avatarurl      string    `orm:"column:avatarurl;size:255"`
	WxAvatarurl    string    `orm:"column:wx_avatarurl;size:500"` //微信头像
	WxNickname     string    `orm:"column:wx_nickname;size:255"` //微信昵称
	Mobile         string    `orm:"column:mobile;size:32"`
	StudentSn      string    `orm:"column:student_sn;size:32"` //学号
	StudentCode    string    `orm:"column:student_code;size:32"` //学号code
	Truename       string    `orm:"column:truename;size:32"`
	Password       string    `orm:"column:password;size:32"`
	School         string    `orm:"column:school;size:255"` //在读学校
	Province       string    `orm:"column:province;size:32"` //省份
	ProvinceCode   string    `orm:"column:province_code;size:32"` //省份
	City           string    `orm:"column:city;size:32"` //城市
	CityCode       string    `orm:"column:city_code;size:32"` //城市
	District       string    `orm:"column:district;size:32"` //区
	DistrictCode   string    `orm:"column:district_code;size:32"` //区
	Birthday       time.Time `orm:"column:birthday;type(date)"` //出生年月
	Addr           string    `orm:"column:addr;size:255"` //地址
	Grade          int       `orm:"column:grade"` //年级
	Pid            int       `orm:"column:pid"` //推荐人
	Status         int8      `orm:"column:status"` //状态 1学生 2黑名单
	Sex            int8      `orm:"column:sex"` //状态 1男生 2女生
	Createtime     int       `orm:"column:createtime"`
	Logintime      int       `orm:"column:logintime"`
	IsDelete       int8      `orm:"column:is_delete"`
	TicketNum      int       `orm:"column:ticket_num"` //听课券数量
	IconNum        int       `orm:"column:icon_num"` //币的数量
	RegisterSource int8      `orm:"column:register_source"` //来源 1：教务系统   2：视频学院
	Xhopenid       string    `orm:"column:xhopenid;size:255"` //视频学院openid
	Email          string    `orm:"column:email;size:500"` //邮箱
	Oldid          int       `orm:"column:oldid"`
	Xoldid         int       `orm:"column:xoldid"` //晓禾学院会员id
}

func (t *Member) TableName() string {
	return "goouc_xmf_member"
}

func init() {
	orm.RegisterModel(new(Member))
}

// AddMember insert a new Member into database and returns
// last inserted Id on success.
func AddMember(m *Member) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberById retrieves Member by Id. Returns error if
// Id doesn't exist
func GetMemberById(id int) (v *Member, err error) {
	o := orm.NewOrm()
	v = &Member{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMember retrieves all Member matches certain condition. Returns empty list if
// no records exist
func GetAllMember(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Member))
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

	var l []Member
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

// UpdateMember updates Member by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberById(m *Member) (err error) {
	o := orm.NewOrm()
	v := Member{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMember deletes Member by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMember(id int) (err error) {
	o := orm.NewOrm()
	v := Member{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Member{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
