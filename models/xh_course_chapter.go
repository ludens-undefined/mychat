package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhCourseChapter struct {
	Id                  int     `orm:"column:id;autoIncrement"`
	Xhcourseid          int     `orm:"column:xhcourseid"` //课程id
	Title               string  `orm:"column:title;size:255"` //章节标题
	Num                 int     `orm:"column:num"` //已学习人数 默认为0
	Status              int8    `orm:"column:status"` //默认下架 1:上架   2：下架
	Sort                int     `orm:"column:sort"` //排序默认为0 越大越靠前
	Xnum                int     `orm:"column:xnum"` //虚拟浏览量
	Videotime           float64 `orm:"column:videotime;scale:10;precision:2"` //视频时长
	UpdateStatus        int8    `orm:"column:update_status"` //是否更新过视频1:更新过   2：没更新
	StartTime           int     `orm:"column:start_time"` //开始上课时间
	ChapterType         int8    `orm:"column:chapter_type"` //2为直播课程 1 为录播课程 3伪直播
	CreateTime          int     `orm:"column:create_time"` //创建时间
	UpdateTime          int     `orm:"column:update_time"` //更新时间
	ShopId              int     `orm:"column:shop_id"` //店铺id
	IsDelete            int8    `orm:"column:is_delete"` //状态 1启用 2删除
	ClickNum            uint    `orm:"column:click_num"` //点击人数
	LiveUrl             string  `orm:"column:live_url"` //直播推流地址
	BroadUrl            string  `orm:"column:broad_url"` //直播播放地址
	CreateUrlTime       int     `orm:"column:create_url_time"` //创建url时间
	UrlTime             string  `orm:"column:url_time;size:255"` //url有效期
	Vodtranscodegroupid string  `orm:"column:vodtranscodegroupid;size:255"` //直播转点播的转码模板id
	Videoid             string  `orm:"column:videoid;size:255"` //上传视频id
	Livestatus          int8    `orm:"column:livestatus"` //直播状态     1：未开始 2：直播中   3：直播结束
	Xoldid              int     `orm:"column:xoldid"`
}

func (t *XhCourseChapter) TableName() string {
	return "goouc_xmf_xh_course_chapter"
}

func init() {
	orm.RegisterModel(new(XhCourseChapter))
}

// AddXhCourseChapter insert a new XhCourseChapter into database and returns
// last inserted Id on success.
func AddXhCourseChapter(m *XhCourseChapter) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhCourseChapterById retrieves XhCourseChapter by Id. Returns error if
// Id doesn't exist
func GetXhCourseChapterById(id int) (v *XhCourseChapter, err error) {
	o := orm.NewOrm()
	v = &XhCourseChapter{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhCourseChapter retrieves all XhCourseChapter matches certain condition. Returns empty list if
// no records exist
func GetAllXhCourseChapter(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhCourseChapter))
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

	var l []XhCourseChapter
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

// UpdateXhCourseChapter updates XhCourseChapter by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhCourseChapterById(m *XhCourseChapter) (err error) {
	o := orm.NewOrm()
	v := XhCourseChapter{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhCourseChapter deletes XhCourseChapter by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhCourseChapter(id int) (err error) {
	o := orm.NewOrm()
	v := XhCourseChapter{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhCourseChapter{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
