package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjBook struct {
	Id              int     `orm:"column:id;autoIncrement"`
	Path            string  `orm:"column:path;size:255"` //电子书路径
	BookName        string  `orm:"column:book_name;size:50"` //书名
	Author          string  `orm:"column:author;size:50"` //作者
	Img             string  `orm:"column:img;size:255"` //封面
	Brief           string  `orm:"column:brief;size:100"` //简介
	Type            uint8   `orm:"column:type"` //是否支持试读 1支持试读 2不支持试读
	Chapter         uint16  `orm:"column:chapter"` //连载章数，对应type=1
	IsAloneSale     uint8   `orm:"column:is_alone_sale"` //1支持单独售卖 2不支持单独售卖
	AloneSaleIsFree uint8   `orm:"column:alone_sale_is_free"` //对应is_alone_sale为1(单独售卖)，1付费 2免费
	Price           float64 `orm:"column:price;scale:10;precision:2"` //商品价格，对应alone_sale_is_free为1时，商品售卖价格
	LinePrice       float64 `orm:"column:line_price;scale:10;precision:2"` //划线价格
	IsRelated       uint8   `orm:"column:is_related"` //1关联售卖 2不关联售卖
	ZlId            string  `orm:"column:zl_id;size:255"` //专栏id，对应is_related为1时，对应goouc_xet_cource_zl表中的id
	HyId            string  `orm:"column:hy_id;size:255"` //会员id，对应is_related为1时，对应goouc_ext_cource_hy表中的id
	XlyId           string  `orm:"column:xly_id;size:255"` //训练营id，对应is_related为1时，对应goouc_ext_cource_xly表中的id
	MessId          uint    `orm:"column:mess_id"` //0不选择信息采集，对应goouc_ext_cource_mess表中的id
	IsSale          uint8   `orm:"column:is_sale"` //0暂不上架 1立即上架 2定时上架
	TimeSale        uint    `orm:"column:time_sale"` //定时上架时间，对应is_sale为2时的时间
	Hide            uint8   `orm:"column:hide"` //1隐藏 2不设置，对应is_sale为1和2时判断是否设置
	Stop            uint8   `orm:"column:stop"` //1停售 2不设置，对应is_sale为1和2时判断是否设置
	IsJoin          uint8   `orm:"column:is_join"` //引导加群，1开启 2关闭
	IsRecommend     uint8   `orm:"column:is_recommend"` //1不推荐 2推荐
	IsDelete        uint8   `orm:"column:is_delete"` //状态 1启用 2删除
	CreateAt        uint    `orm:"column:create_at"` //创建时间
	UpdateAt        uint    `orm:"column:update_at"` //修改时间
	DeleteAt        uint    `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjBook) TableName() string {
	return "goouc_xmf_zxgj_book"
}

func init() {
	orm.RegisterModel(new(ZxgjBook))
}

// AddZxgjBook insert a new ZxgjBook into database and returns
// last inserted Id on success.
func AddZxgjBook(m *ZxgjBook) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjBookById retrieves ZxgjBook by Id. Returns error if
// Id doesn't exist
func GetZxgjBookById(id int) (v *ZxgjBook, err error) {
	o := orm.NewOrm()
	v = &ZxgjBook{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjBook retrieves all ZxgjBook matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjBook(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjBook))
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

	var l []ZxgjBook
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

// UpdateZxgjBook updates ZxgjBook by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjBookById(m *ZxgjBook) (err error) {
	o := orm.NewOrm()
	v := ZxgjBook{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjBook deletes ZxgjBook by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjBook(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjBook{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjBook{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
