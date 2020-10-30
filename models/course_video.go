package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CourseVideo struct {
	Id              int     `orm:"column:id;autoIncrement"`
	ShopId          uint    `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name            string  `orm:"column:name;size:50"` //视频标题
	VideoName       string  `orm:"column:video_name;size:255"` //视频名称
	Video           string  `orm:"column:video;size:255"` //视频
	Img             string  `orm:"column:img;size:255"` //视频封面
	VideoImg        string  `orm:"column:video_img;size:255"` //视频贴片
	Detail          string  `orm:"column:detail;size:100"` //视频详情路径
	IsCopy          uint8   `orm:"column:is_copy"` //1可以复制 2不可以复制
	IsAloneSale     uint8   `orm:"column:is_alone_sale"` //1支持单独售卖 2不支持单独售卖
	AloneSaleIsFree uint8   `orm:"column:alone_sale_is_free"` //对应is_alone_sale为1(单独售卖)，1付费 2免费
	Price           float64 `orm:"column:price;scale:10;precision:2"` //商品价格，对应alone_sale_is_free为1时，商品售卖价格
	LinePrice       float64 `orm:"column:line_price;scale:10;precision:2"` //划线价格
	IsRelated       uint8   `orm:"column:is_related"` //1关联售卖 2不关联售卖
	GoodTypeIds     string  `orm:"column:good_type_ids;size:255"` //商品分组id,对应goouc_xet_cource_good_type表中的id
	MessId          uint    `orm:"column:mess_id"` //0不选择信息采集，对应goouc_ext_cource_mess表中的id
	IsSale          uint8   `orm:"column:is_sale"` //0暂不上架 1立即上架 2定时上架
	TimeSale        uint    `orm:"column:time_sale"` //定时上架时间，对应is_sale为2时的时间
	Hide            uint8   `orm:"column:hide"` //1隐藏 2不设置，对应is_sale为1和2时判断是否设置
	Stop            uint8   `orm:"column:stop"` //1停售 2不设置，对应is_sale为1和2时判断是否设置
	IsJoin          uint8   `orm:"column:is_join"` //引导加群，1开启 2关闭
	JoinWay1        uint8   `orm:"column:join_way1"` //引导方式：详情页引导加群1开启 2关闭
	Label           string  `orm:"column:label;size:16"` //引导标签内容设置，对应join_way1为1
	JoinWay2        uint8   `orm:"column:join_way2"` //引导方式：购买成功页引导加群1开启 2关闭
	Desp            string  `orm:"column:desp;size:40"` //引导描述
	CodeTitle       string  `orm:"column:code_title;size:30"` //二维码标题
	Code            string  `orm:"column:code;size:255"` //二维码
	IsRecommend     uint8   `orm:"column:is_recommend"` //1不推荐 2推荐
	IsDelete        uint8   `orm:"column:is_delete"` //状态 1启用 2删除
	ZybId           uint    `orm:"column:zyb_id"` //关联作业本id，对应goouc_xet_zxgj_zyb表中的id
	CreateAt        uint    `orm:"column:create_at"` //创建时间
	UpdateAt        uint    `orm:"column:update_at"` //修改时间
	DeleteAt        uint    `orm:"column:delete_at"` //删除时间
}

func (t *CourseVideo) TableName() string {
	return "goouc_xmf_course_video"
}

func init() {
	orm.RegisterModel(new(CourseVideo))
}

// AddCourseVideo insert a new CourseVideo into database and returns
// last inserted Id on success.
func AddCourseVideo(m *CourseVideo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourseVideoById retrieves CourseVideo by Id. Returns error if
// Id doesn't exist
func GetCourseVideoById(id int) (v *CourseVideo, err error) {
	o := orm.NewOrm()
	v = &CourseVideo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourseVideo retrieves all CourseVideo matches certain condition. Returns empty list if
// no records exist
func GetAllCourseVideo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CourseVideo))
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

	var l []CourseVideo
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

// UpdateCourseVideo updates CourseVideo by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourseVideoById(m *CourseVideo) (err error) {
	o := orm.NewOrm()
	v := CourseVideo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseVideo deletes CourseVideo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseVideo(id int) (err error) {
	o := orm.NewOrm()
	v := CourseVideo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseVideo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
