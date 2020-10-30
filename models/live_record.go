package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type LiveRecord struct {
	Id         int    `orm:"column:id;autoIncrement"`
	Chapterid  int    `orm:"column:chapterid"` //章节id
	Stream     string `orm:"column:stream;size:128"` //直播码
	App        string `orm:"column:app;size:128"` //频道名
	ShopId     int    `orm:"column:shop_id"` //店铺id
	Ctime      int    `orm:"column:ctime"` //创建时间
	Utime      int    `orm:"column:utime"` //修改时间
	Flv        string `orm:"column:flv"`
	Mp4        string `orm:"column:mp4;size:500"`
	M3u8       string `orm:"column:m3u8"`
	Duration   int    `orm:"column:duration"` //录制时间 秒
	StartTime  int    `orm:"column:start_time"` //开始时间
	StopTime   int    `orm:"column:stop_time"` //结束时间
	Domain     string `orm:"column:domain;size:255"` //播流域名
	BroadUrl   string `orm:"column:broad_url"` //最终播放地址
	ClipStatus int8   `orm:"column:clip_status"` //是否剪辑处理过 1;剪辑过  2：没剪辑
}

func (t *LiveRecord) TableName() string {
	return "goouc_xmf_live_record"
}

func init() {
	orm.RegisterModel(new(LiveRecord))
}

// AddLiveRecord insert a new LiveRecord into database and returns
// last inserted Id on success.
func AddLiveRecord(m *LiveRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLiveRecordById retrieves LiveRecord by Id. Returns error if
// Id doesn't exist
func GetLiveRecordById(id int) (v *LiveRecord, err error) {
	o := orm.NewOrm()
	v = &LiveRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLiveRecord retrieves all LiveRecord matches certain condition. Returns empty list if
// no records exist
func GetAllLiveRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LiveRecord))
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

	var l []LiveRecord
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

// UpdateLiveRecord updates LiveRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateLiveRecordById(m *LiveRecord) (err error) {
	o := orm.NewOrm()
	v := LiveRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLiveRecord deletes LiveRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLiveRecord(id int) (err error) {
	o := orm.NewOrm()
	v := LiveRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LiveRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
