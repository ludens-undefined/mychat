package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ReviewRecord struct {
	Id         int    `orm:"column:id;autoIncrement"`
	ShopId     int    `orm:"column:shop_id"`
	Memberid   int    `orm:"column:memberid"` //学员id
	Xhcourseid int    `orm:"column:xhcourseid"` //课程id
	Chapterid  int    `orm:"column:chapterid"` //章节id
	Date       string `orm:"column:date;size:11"` //当天日期
	Daytime    int    `orm:"column:daytime"` //当天观看总时长
	Studyrate  int    `orm:"column:studyrate"` //当前章节学习到的进度
	Week       string `orm:"column:week;size:255"` //本周第一天日期
	Weektime   int    `orm:"column:weektime"` //本周观看总时长
	Month      string `orm:"column:month;size:255"` //当月第一天日期
	Monthtime  int    `orm:"column:monthtime"` //当月观看总时长
	Daynum     int    `orm:"column:daynum"` //连续学习天数
	Num        int    `orm:"column:num"` //当天领取晓禾币次数
	Createtime int    `orm:"column:createtime"`
}

func (t *ReviewRecord) TableName() string {
	return "goouc_xmf_review_record"
}

func init() {
	orm.RegisterModel(new(ReviewRecord))
}

// AddReviewRecord insert a new ReviewRecord into database and returns
// last inserted Id on success.
func AddReviewRecord(m *ReviewRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReviewRecordById retrieves ReviewRecord by Id. Returns error if
// Id doesn't exist
func GetReviewRecordById(id int) (v *ReviewRecord, err error) {
	o := orm.NewOrm()
	v = &ReviewRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReviewRecord retrieves all ReviewRecord matches certain condition. Returns empty list if
// no records exist
func GetAllReviewRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReviewRecord))
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

	var l []ReviewRecord
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

// UpdateReviewRecord updates ReviewRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateReviewRecordById(m *ReviewRecord) (err error) {
	o := orm.NewOrm()
	v := ReviewRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReviewRecord deletes ReviewRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReviewRecord(id int) (err error) {
	o := orm.NewOrm()
	v := ReviewRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReviewRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
