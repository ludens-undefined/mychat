package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CourseDzl struct {
	Id              int     `orm:"column:id;autoIncrement"`
	ShopId          uint    `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name            string  `orm:"column:name;size:50"` //专栏名称
	Brief           string  `orm:"column:brief;size:100"` //专栏简介路径
	Img             string  `orm:"column:img;size:255"` //专栏封面
	Detail          string  `orm:"column:detail;size:100"` //专栏详情路径
	IsAloneSale     uint8   `orm:"column:is_alone_sale"` //1支持单独售卖 2不支持单独售卖
	AloneSaleIsFree uint8   `orm:"column:alone_sale_is_free"` //对应is_alone_sale为1(单独售卖)，1付费 2免费
	Price           float64 `orm:"column:price;scale:10;precision:2"` //商品价格，对应alone_sale_is_free为1时，商品售卖价格
	LinePrice       float64 `orm:"column:line_price;scale:10;precision:2"` //划线价格
	GoodTypeIds     string  `orm:"column:good_type_ids;size:255"` //商品分组id,对应goouc_xet_cource_good_type表中的id
	MessId          uint    `orm:"column:mess_id"` //0不选择信息采集，对应goouc_ext_cource_mess表中的id
	IsSale          uint8   `orm:"column:is_sale"` //0暂不上架 1立即上架
	TimeSale        uint    `orm:"column:time_sale"` //上架时间
	Hide            uint8   `orm:"column:hide"` //1隐藏 2不设置，对应is_sale为1和2时判断是否设置
	Stop            uint8   `orm:"column:stop"` //1停售 2不设置，对应is_sale为1和2时判断是否设置
	IsJoin          uint8   `orm:"column:is_join"` //引导加群，1开启 2关闭
	JoinWay1        uint8   `orm:"column:join_way1"` //引导方式：详情页引导加群1开启 2关闭
	Label           string  `orm:"column:label;size:16"` //引导标签内容设置，对应join_way1为1
	JoinWay2        uint8   `orm:"column:join_way2"` //引导方式：购买成功页引导加群1开启 2关闭
	Desp            string  `orm:"column:desp;size:40"` //引导描述
	CodeTitle       string  `orm:"column:code_title;size:30"` //二维码标题
	Code            string  `orm:"column:code;size:255"` //二维码
	IsRecommend     uint8   `orm:"column:is_recommend"` //1推荐 2不推荐
	IsDelete        uint8   `orm:"column:is_delete"` //状态 1启用 2删除
	CreateAt        uint    `orm:"column:create_at"` //创建时间
	UpdateAt        uint    `orm:"column:update_at"` //修改时间
	DeleteAt        uint    `orm:"column:delete_at"` //删除时间
	ShareTitle      string  `orm:"column:share_title;size:50"` //页面分享标题
	ShareDesp       string  `orm:"column:share_desp;size:100"` //页面分享描述路径
	ShareImg        string  `orm:"column:share_img;size:255"` //页面分享图片
}

func (t *CourseDzl) TableName() string {
	return "goouc_xmf_course_dzl"
}

func init() {
	orm.RegisterModel(new(CourseDzl))
}

// AddCourseDzl insert a new CourseDzl into database and returns
// last inserted Id on success.
func AddCourseDzl(m *CourseDzl) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseDzlById retrieves CourseDzl by Id. Returns error if
// Id doesn't exist
func GetCourseDzlById(id int) (v *CourseDzl, err error) {
	o := orm.NewOrm()
	v = &CourseDzl{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseDzl retrieves all CourseDzl matches certain condition. Returns empty list if
// no records exist
func GetAllCourseDzl(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseDzl))
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

	var l []CourseDzl
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

// UpdateCourseDzl updates CourseDzl by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseDzlById(m *CourseDzl) (err error) {
	o := orm.NewOrm()
	v := CourseDzl{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseDzl deletes CourseDzl by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseDzl(id int) (err error) {
	o := orm.NewOrm()
	v := CourseDzl{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseDzl{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
