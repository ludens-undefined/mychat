package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type BusinessShop struct {
	Id          int    `orm:"column:id;autoIncrement"`
	Shopname    string `orm:"column:shopname;size:32"` //名称
	Logo        string `orm:"column:logo;size:255"` //logo
	FootLogo    string `orm:"column:foot_logo;size:255"` //页脚logo
	MLogo       string `orm:"column:m_logo;size:255"`
	StudentCode string `orm:"column:student_code;size:32"` //学号前缀
	Email       string `orm:"column:email;size:255"` //email
	Tel         string `orm:"column:tel;size:32"` //联系电话
	KefuTel     string `orm:"column:kefu_tel;size:32"` //客服电话
	Address     string `orm:"column:address;size:255"` //公司地址
	About       string `orm:"column:about"` //关于我们
	Icp         string `orm:"column:icp;size:255"` //icp备案号
	Appid       string `orm:"column:appid;size:32"` //网站appid
	Secret      string `orm:"column:secret;size:255"` //网站secret
	AppAppid    string `orm:"column:app_appid;size:32"` //小程序
	AppSecret   string `orm:"column:app_secret;size:255"` //小程序
	IsChange    int8   `orm:"column:is_change"` //是否开启转班
	IsAdjust    int8   `orm:"column:is_adjust"` //是否开启调课
	IsRefund    int8   `orm:"column:is_refund"` //是否开启退费
	ShareTitle  string `orm:"column:share_title;size:255"` //分享标题
	ShareImg    string `orm:"column:share_img;size:255"` //分享图片
	ShareTips   string `orm:"column:share_tips;size:255"` //分享提示
	Createtime  uint   `orm:"column:createtime"` //注册时间
	Expiretime  uint   `orm:"column:expiretime"` //过期时间
	Status      int8   `orm:"column:status"` //状态 1启用 2禁用
	IsDelete    int8   `orm:"column:is_delete"` //状态 1启用 2删除
	Xhdomain    string `orm:"column:xhdomain;size:255"` //视频学院域名
	IsRename    int8   `orm:"column:is_rename"` //改名 1开启 2关闭
}

func (t *BusinessShop) TableName() string {
	return "goouc_xmf_business_shop"
}

func init() {
	orm.RegisterModel(new(BusinessShop))
}

// AddBusinessShop insert a new BusinessShop into database and returns
// last inserted Id on success.
func AddBusinessShop(m *BusinessShop) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBusinessShopById retrieves BusinessShop by Id. Returns error if
// Id doesn't exist
func GetBusinessShopById(id int) (v *BusinessShop, err error) {
	o := orm.NewOrm()
	v = &BusinessShop{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBusinessShop retrieves all BusinessShop matches certain condition. Returns empty list if
// no records exist
func GetAllBusinessShop(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BusinessShop))
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

	var l []BusinessShop
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

// UpdateBusinessShop updates BusinessShop by Id and returns error if
// the record to be updated doesn't exist
func UpdateBusinessShopById(m *BusinessShop) (err error) {
	o := orm.NewOrm()
	v := BusinessShop{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBusinessShop deletes BusinessShop by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBusinessShop(id int) (err error) {
	o := orm.NewOrm()
	v := BusinessShop{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BusinessShop{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
