package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZxgjZsDk struct {
	Id         int       `orm:"column:id;autoIncrement"`
	ShopId     uint      `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	DkType     uint8     `orm:"column:dk_type"` //1日历打卡 2作业打卡 3闯关打卡 4训练营
	DkId       uint      `orm:"column:dk_id"` //dk_type为1时，日历打卡id；2作业打卡id 3闯关打卡id 4训练营id
	IsStart    uint8     `orm:"column:is_start"` //1发放结课证书 2不开启
	Count      uint16    `orm:"column:count"` //dk_type为1日历打卡打卡次数  2作业打卡作业个数 3闯关打卡课时个数
	ModelId    uint8     `orm:"column:model_id"` //选择模式 1使用固定模式 2上传证书
	ZsImg      string    `orm:"column:zs_img;size:255"` //model_id为2时，结课证书图片
	Name       string    `orm:"column:name;size:50"` //证书名称
	ShowData   string    `orm:"column:show_data;size:50"` //展示数据，格式：{1,2,3,4}1打卡天数 2连续打卡|完成作业|解锁课时 3评为精选 4收获点赞；训练营：1完成任务数量 2学习课程数 3打卡完成数 4考试完成数 5考试总分
	Content    string    `orm:"column:content;size:100"` //鼓励文案，格式:{text1:,text2:}
	OrganName  string    `orm:"column:organ_name;size:50"` //机构名称
	QrcodeType uint8     `orm:"column:qrcode_type"` //1跳转当前打卡 2选择其他课程 3使用自定义二维码
	Class      string    `orm:"column:class;size:50"` //当qrcode_type为2课程，格式为:{type:,id:},type对应goouc_xet_cource_type中的id
	Qrcode     string    `orm:"column:qrcode;size:255"` //当qrcode_type为3，自定义二维码
	Text       string    `orm:"column:text;size:50"` //二维码文案
	BgImg      string    `orm:"column:bg_img;size:255"` //背景图
	TimeType   uint8     `orm:"column:time_type"` //发放证书时间 1训练营营期结束 2自定义时间
	Time       time.Time `orm:"column:time;type(date)"` //对应time_type为2时，自定义发放证书时间
	TaskType   uint8     `orm:"column:task_type"` //训练营任务完成数 1开启 2关闭
	TaskNum    uint16    `orm:"column:task_num"` //任务完成数
	ScoreType  uint8     `orm:"column:score_type"` //考试成绩规则 1开启 2关闭
	Score      uint8     `orm:"column:score"` //对应score_type为1，分数
	DakaType   uint8     `orm:"column:daka_type"` //打卡次数规则 1开启 2关闭
	DkNum      uint16    `orm:"column:dk_num"` //对应daka_type为1，打卡次数
	ZsText     string    `orm:"column:zs_text;size:50"` //证书入口文案
	IsDelete   uint8     `orm:"column:is_delete"` //1启用 2删除
	CreateAt   uint      `orm:"column:create_at"` //创建时间
	UpdateAt   uint      `orm:"column:update_at"` //编辑时间
	DeleteAt   uint      `orm:"column:delete_at"` //删除时间
}

func (t *ZxgjZsDk) TableName() string {
	return "goouc_xmf_zxgj_zs_dk"
}

func init() {
	orm.RegisterModel(new(ZxgjZsDk))
}

// AddZxgjZsDk insert a new ZxgjZsDk into database and returns
// last inserted Id on success.
func AddZxgjZsDk(m *ZxgjZsDk) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZsDkById retrieves ZxgjZsDk by Id. Returns error if
// Id doesn't exist
func GetZxgjZsDkById(id int) (v *ZxgjZsDk, err error) {
	o := orm.NewOrm()
	v = &ZxgjZsDk{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZsDk retrieves all ZxgjZsDk matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZsDk(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZsDk))
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

	var l []ZxgjZsDk
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

// UpdateZxgjZsDk updates ZxgjZsDk by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZsDkById(m *ZxgjZsDk) (err error) {
	o := orm.NewOrm()
	v := ZxgjZsDk{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZsDk deletes ZxgjZsDk by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZsDk(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZsDk{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZsDk{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
