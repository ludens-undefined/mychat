package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Poster struct {
	Id             int    `orm:"column:id;autoIncrement"`
	DateShow       int8   `orm:"column:dateShow"` //日期   1开启   2不开启
	Datenumwidth   int    `orm:"column:datenumwidth"` //日期宽度
	Datenumtop     int    `orm:"column:datenumtop"` //日期上边距
	Datenumleft    int    `orm:"column:datenumleft"` //日期左边距
	Datenumheight  int    `orm:"column:datenumheight"` //日期高度
	HeaderShow     int8   `orm:"column:headerShow"` //头像   1开启   2不开启
	Headerheight   int    `orm:"column:headerheight"` //头像高度
	Headerleft     int    `orm:"column:headerleft"` //头像左边距
	Headertop      int    `orm:"column:headertop"` //头像上边距
	Headerwidth    int    `orm:"column:headerwidth"` //头像宽度
	NicknameShow   int8   `orm:"column:nicknameShow"` //昵称   1开启  2不开启
	Nicknameheight int    `orm:"column:nicknameheight"` //昵称高度
	Nicknameleft   int    `orm:"column:nicknameleft"` //昵称左边距
	Nicknametop    int    `orm:"column:nicknametop"` //昵称上边距
	Nicknamewidth  int    `orm:"column:nicknamewidth"` //昵称宽度
	Fontcolor      string `orm:"column:fontcolor;size:10"` //昵称字体颜色
	Fontsize       int    `orm:"column:fontsize"` //昵称字体大小
	QrShow         int8   `orm:"column:qrShow"` //二维码  1开启   2不开启
	Qrheight       int    `orm:"column:qrheight"` //二维码高度
	Qrleft         int    `orm:"column:qrleft"` //二维码左边距
	Qrtop          int    `orm:"column:qrtop"` //二维码上边距
	Qrwidth        int    `orm:"column:qrwidth"` //二维码宽度
	StudynumShow   string `orm:"column:studynumShow;size:10"` //学习天数  1开启   2不开启
	Studynumheight int    `orm:"column:studynumheight"` //学习天数高度
	Studynumleft   int    `orm:"column:studynumleft"` //学习天数左边距
	Studynumtop    int    `orm:"column:studynumtop"` //学习天数上边距
	Studynumwidth  int    `orm:"column:studynumwidth"` //学习天数宽度
	Bgimg          string `orm:"column:bgimg;size:1000"` //海报背景图
	Title          string `orm:"column:title;size:255"` //海报标题
	Status         int8   `orm:"column:status"` //1：开启   2：关闭
	ShopId         int    `orm:"column:shop_id"` //店铺id
	Createtime     int    `orm:"column:createtime"` //创建时间
	Updatetime     int    `orm:"column:updatetime"` //更新时间
	IsDelete       int8   `orm:"column:is_delete"` //1：启用    2：禁用
}

func (t *Poster) TableName() string {
	return "goouc_xmf_poster"
}

func init() {
	orm.RegisterModel(new(Poster))
}

// AddPoster insert a new Poster into database and returns
// last inserted Id on success.
func AddPoster(m *Poster) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPosterById retrieves Poster by Id. Returns error if
// Id doesn't exist
func GetPosterById(id int) (v *Poster, err error) {
	o := orm.NewOrm()
	v = &Poster{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPoster retrieves all Poster matches certain condition. Returns empty list if
// no records exist
func GetAllPoster(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Poster))
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

	var l []Poster
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

// UpdatePoster updates Poster by Id and returns error if
// the record to be updated doesn't exist
func UpdatePosterById(m *Poster) (err error) {
	o := orm.NewOrm()
	v := Poster{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePoster deletes Poster by Id and returns error if
// the record to be deleted doesn't exist
func DeletePoster(id int) (err error) {
	o := orm.NewOrm()
	v := Poster{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Poster{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
