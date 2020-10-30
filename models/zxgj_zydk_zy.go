package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZydkZy struct {
	Id          int    `orm:"column:id;autoIncrement"`
	ZydkId      uint   `orm:"column:zydk_id"` //对应goouc_xet_zxgj_rldk表中的id
	ZyTitle     string `orm:"column:zy_title;size:50"` //主题标题
	ZyStartTime int    `orm:"column:zy_start_time"` //作业有效期开始时间
	ZyEndTime   uint   `orm:"column:zy_end_time"` //作业有效期结束时间
	ZyContent   string `orm:"column:zy_content;size:100"` //主题正文，格式：[{type:1,desc:'图文'},{type:2,audio_name:'音频名称'，audio_url:'音频路径'},{type:3,video_name:'视频名称',video_url:'视频路径'},{type:4,course_type:'课程类型','course_id':'课程id','course_title':'课程标题','course_img':'课程图片'}，{'type':5,'audio_name':'音频名称','audio_url':'音频路径','follow_read_text':'跟读文本','remark':'备注'}]
	CreateAt    uint   `orm:"column:create_at"` //创建时间
	UpdateAt    uint   `orm:"column:update_at"` //修改时间
	IsDelete    uint8  `orm:"column:is_delete"` //1启用 2禁用
	DeleteAt    uint   `orm:"column:delete_at"`
}

func (t *ZxgjZydkZy) TableName() string {
	return "goouc_xmf_zxgj_zydk_zy"
}

func init() {
	orm.RegisterModel(new(ZxgjZydkZy))
}

// AddZxgjZydkZy insert a new ZxgjZydkZy into database and returns
// last inserted Id on success.
func AddZxgjZydkZy(m *ZxgjZydkZy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZydkZyById retrieves ZxgjZydkZy by Id. Returns error if
// Id doesn't exist
func GetZxgjZydkZyById(id int) (v *ZxgjZydkZy, err error) {
	o := orm.NewOrm()
	v = &ZxgjZydkZy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZydkZy retrieves all ZxgjZydkZy matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZydkZy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZydkZy))
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

	var l []ZxgjZydkZy
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

// UpdateZxgjZydkZy updates ZxgjZydkZy by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZydkZyById(m *ZxgjZydkZy) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydkZy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZydkZy deletes ZxgjZydkZy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZydkZy(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydkZy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZydkZy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
