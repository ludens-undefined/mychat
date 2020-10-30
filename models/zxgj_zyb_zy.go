package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZybZy struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ZybId      int    `orm:"column:zyb_id"` //对应goouc_xet_zxgj_zyb表中的id
	ShopId     int    `orm:"column:shop_id"`
	Type       int8   `orm:"column:type"` //1手动添加 2从题库添加
	Title      string `orm:"column:title;size:45"` //作业名称
	Content    string `orm:"column:content;size:1024"` //作业内容-内容
	Img        string `orm:"column:img;size:5000"` //图片，{1.png,2.png}
	AudioInfo  string `orm:"column:audio_info;size:5000"` //音频{['audio_name':'音频名称','audio':'音频路径']}
	TkId       string `orm:"column:tk_id;size:5000"` //tk_id对应goouc_xet_zxgj_tk表中的id，格式：{1,2}
	IsShow     int8   `orm:"column:is_show"` //展示设置 1显示 2隐藏
	IsDelete   int8   `orm:"column:is_delete"` //状态 1启用 2删除
	Createtime int    `orm:"column:createtime"` //创建时间
	UpdateAt   int    `orm:"column:update_at"` //修改时间
	DeleteAt   int    `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjZybZy) TableName() string {
	return "goouc_xmf_zxgj_zyb_zy"
}

func init() {
	orm.RegisterModel(new(ZxgjZybZy))
}

// AddZxgjZybZy insert a new ZxgjZybZy into database and returns
// last inserted Id on success.
func AddZxgjZybZy(m *ZxgjZybZy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZybZyById retrieves ZxgjZybZy by Id. Returns error if
// Id doesn't exist
func GetZxgjZybZyById(id int) (v *ZxgjZybZy, err error) {
	o := orm.NewOrm()
	v = &ZxgjZybZy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZybZy retrieves all ZxgjZybZy matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZybZy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZybZy))
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

	var l []ZxgjZybZy
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

// UpdateZxgjZybZy updates ZxgjZybZy by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZybZyById(m *ZxgjZybZy) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZybZy deletes ZxgjZybZy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZybZy(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZybZy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
