package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZydk struct {
	Id     int    `orm:"column:id;autoIncrement"`
	ShopId uint   `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name   string `orm:"column:name;size:50"` //打卡名称
	Img    string `orm:"column:img;size:255"` //封面图路径
	Intro string `orm:"column:intro;size:100" description:"打卡介绍路径 type1图文 2音频 3视频 4课程
[{type:1,desc:'图文'},{type:2,audio_name:'音频名称'，audio_url:'音频路径'},{type:3,video_name:'视频名称',video_url:'视频路径'},{type:4,course_type:'课程类型','course_id':'课程id','course_title':'课程标题','course_img':'课程图片'}]"`
	Condition  uint8   `orm:"column:condition"` //1免费参与 2需要购买课程参与 3付费参与
	Price      float64 `orm:"column:price;scale:10;precision:2"` //condition为3，对应的价格
	IsZb       uint8   `orm:"column:is_zb"` //防作弊模式 1开启 2关闭
	IsRepeatDk uint8   `orm:"column:is_repeat_dk"` //不允许重新打卡 1开启 2关闭
	IsRemind   uint8   `orm:"column:is_remind"` //打卡提醒 1开启 2关闭
	IsShow     uint8   `orm:"column:is_show"` //展示设置 1显示 2隐藏
	IsBk       uint8   `orm:"column:is_bk"` //补打卡 1开启 2关闭
	BkCount    uint8   `orm:"column:bk_count"` //对应is_bk为1是，允许补打卡的次数
	WordMax    uint16  `orm:"column:word_max"` //文字最多次数
	ImgMax     uint16  `orm:"column:img_max"` //图片最多张数
	AudioMax   uint16  `orm:"column:audio_max"` //音频最长秒数
	ZsNickname string  `orm:"column:zs_nickname;size:50"` //小助手昵称
	ZsImg      string  `orm:"column:zs_img;size:255"` //小助手头像
	ZsWeixin   string  `orm:"column:zs_weixin;size:255"` //小助手微信号
	IsDelete   uint8   `orm:"column:is_delete"` //状态 1启用 2删除
	CreateAt   uint    `orm:"column:create_at"` //创建时间
	UpdateAt   uint    `orm:"column:update_at"` //修改时间
	DeleteAt   uint    `orm:"column:delete_at"`
}

func (t *ZxgjZydk) TableName() string {
	return "goouc_xmf_zxgj_zydk"
}

func init() {
	orm.RegisterModel(new(ZxgjZydk))
}

// AddZxgjZydk insert a new ZxgjZydk into database and returns
// last inserted Id on success.
func AddZxgjZydk(m *ZxgjZydk) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZydkById retrieves ZxgjZydk by Id. Returns error if
// Id doesn't exist
func GetZxgjZydkById(id int) (v *ZxgjZydk, err error) {
	o := orm.NewOrm()
	v = &ZxgjZydk{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZydk retrieves all ZxgjZydk matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZydk(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZydk))
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

	var l []ZxgjZydk
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

// UpdateZxgjZydk updates ZxgjZydk by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZydkById(m *ZxgjZydk) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydk{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZydk deletes ZxgjZydk by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZydk(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZydk{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZydk{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
