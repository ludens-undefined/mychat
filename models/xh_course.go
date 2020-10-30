package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhCourse struct {
	Id           int     `orm:"column:id;autoIncrement"`
	Title        string  `orm:"column:title;size:100"` //课程标题
	Cover        string  `orm:"column:cover;size:255"` //课程封面图
	CourseType   int8    `orm:"column:course_type"` //课次类型    1：系列课   2：单节课
	Status       int8    `orm:"column:status"` //上下架状态  2：下架     1：上架
	TeacherId    int     `orm:"column:teacher_id"` //主讲老师id
	ShopId       int     `orm:"column:shop_id"` //店铺id
	AssistantId  int     `orm:"column:assistant_id"` //助教老师id
	SubjectId    int     `orm:"column:subject_id"` //学科id
	CityId       int     `orm:"column:city_id"` //城市id
	TermId       int     `orm:"column:term_id"` //学期id
	ClassId      int     `orm:"column:class_id"` //班型id
	GradeId      string  `orm:"column:grade_id"` //年级id序列
	CoursetypeId int     `orm:"column:coursetype_id"` //课程类型id
	Yearid       int     `orm:"column:yearid"` //年份id
	Sort         int     `orm:"column:sort"` //排序
	PayType      int8    `orm:"column:pay_type"` //购买方式     1：免费    2：付费（晓禾币）3：付费（人民币） 4：兑换币+人民币
	PayIcon      int     `orm:"column:pay_icon"` //兑换币付款时所需兑换币数量
	Desc         string  `orm:"column:desc"` //课程简介 区别于课程介绍
	PcType       int8    `orm:"column:pc_type"` //PC首页分类 1:精品公开课 2:干货短期班 3:热门常规课
	XnNum        int     `orm:"column:xn_num"` //虚拟学习人数
	PayMoney     float64 `orm:"column:pay_money;scale:10;precision:2"` //价格
	PayNum       int     `orm:"column:pay_num"` //报名人数
	StudyNum     int     `orm:"column:study_num"` //学习人数
	FreeNum      int     `orm:"column:free_num"` //免费数量
	Recommend    int8    `orm:"column:recommend"` //课程是否被推荐 1：推荐  2：未推荐
	Qrcode       string  `orm:"column:qrcode;size:255"` //每个课程的静态码
	ShareImage   string  `orm:"column:share_image;size:255"` //课程分享封面图 不能为空 必须有
	CreateTime   int     `orm:"column:create_time"` //开始时间
	IsDelete     int8    `orm:"column:is_delete"` //状态 1启用 2删除
	UpdateTime   int     `orm:"column:update_time"` //修改时间
	IsGroupbuy   int8    `orm:"column:is_groupbuy"` //是否支持团购  2：关闭    1：开启
	GroupPrice   float64 `orm:"column:group_price;scale:10;precision:2"` //团购价格
	GroupNum     int     `orm:"column:group_num"` //该团人数
	GroupValid   int     `orm:"column:group_valid"` //团有效期（小时）
	GroupStart   int     `orm:"column:group_start"` //拼团开始时间
	GroupEnd     int     `orm:"column:group_end"` //拼团结束时间
	StartTime    int     `orm:"column:start_time"` //直播和伪直播开始时间
	Xoldid       int     `orm:"column:xoldid"`
	CatType      int8    `orm:"column:cat_type"` //1:线上课 2：线下课 3：礼品
	IconMax      int     `orm:"column:icon_max"` //兑换币使用最大值
	IconMin      int     `orm:"column:icon_min"` //兑换币使用最小值
	IconStatus   int8    `orm:"column:icon_status"`
}

func (t *XhCourse) TableName() string {
	return "goouc_xmf_xh_course"
}

func init() {
	orm.RegisterModel(new(XhCourse))
}

// AddXhCourse insert a new XhCourse into database and returns
// last inserted Id on success.
func AddXhCourse(m *XhCourse) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhCourseById retrieves XhCourse by Id. Returns error if
// Id doesn't exist
func GetXhCourseById(id int) (v *XhCourse, err error) {
	o := orm.NewOrm()
	v = &XhCourse{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhCourse retrieves all XhCourse matches certain condition. Returns empty list if
// no records exist
func GetAllXhCourse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhCourse))
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

	var l []XhCourse
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

// UpdateXhCourse updates XhCourse by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhCourseById(m *XhCourse) (err error) {
	o := orm.NewOrm()
	v := XhCourse{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhCourse deletes XhCourse by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhCourse(id int) (err error) {
	o := orm.NewOrm()
	v := XhCourse{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhCourse{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
