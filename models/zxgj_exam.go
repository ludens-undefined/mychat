package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjExam struct {
	Id         int     `orm:"column:id;autoIncrement"`
	ShopId     uint    `orm:"column:shop_id"` //对应goouc_xet_business_shop表中的id
	Name       string  `orm:"column:name;size:16"` //考试名称
	Img        string  `orm:"column:img;size:255"` //封面图
	Detail     string  `orm:"column:detail"` //详情路径
	MessId     uint    `orm:"column:mess_id"` //信息采集id，0不选择信息采集，对应goouc_ext_cource_mess表中的id
	IsSale     uint8   `orm:"column:is_sale"` //0暂不发布 1立即发布 2定时发布
	TimeSale   uint    `orm:"column:time_sale"` //对应pub为2时，发布考试时间
	IsRemind   uint8   `orm:"column:is_remind"` //服务号通知 1开启 2关闭
	Time       uint8   `orm:"column:time"` //0无限制 单位分
	Count      uint8   `orm:"column:count"` //考试次数 1仅一次 2不限
	CkTime     uint    `orm:"column:ck_time"` //重考间隔 0无限制 单位时
	IsShow     uint8   `orm:"column:is_show"` //答案展示设置 1立即展示 2批改完成后展示 3隐藏
	IsStop     uint8   `orm:"column:is_stop"` //1发布 2停止考试
	CreateAt   uint    `orm:"column:create_at"` //创建时间
	UpdateAt   uint    `orm:"column:update_at"` //修改时间
	IsDelete   uint8   `orm:"column:is_delete"` //1启用 2禁用
	DeleteAt   uint    `orm:"column:delete_at"` //删除时间
	TotalScore float64 `orm:"column:total_score;scale:10;precision:1"` //总分值
	TotalNum   uint16  `orm:"column:total_num"` //总题数
	TotalCanyu int     `orm:"column:total_canyu"` //参与人数
	ZsId       int     `orm:"column:zs_id"` //证书id
}

func (t *ZxgjExam) TableName() string {
	return "goouc_xmf_zxgj_exam"
}

func init() {
	orm.RegisterModel(new(ZxgjExam))
}

// AddZxgjExam insert a new ZxgjExam into database and returns
// last inserted Id on success.
func AddZxgjExam(m *ZxgjExam) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjExamById retrieves ZxgjExam by Id. Returns error if
// Id doesn't exist
func GetZxgjExamById(id int) (v *ZxgjExam, err error) {
	o := orm.NewOrm()
	v = &ZxgjExam{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjExam retrieves all ZxgjExam matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjExam(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjExam))
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

	var l []ZxgjExam
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

// UpdateZxgjExam updates ZxgjExam by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjExamById(m *ZxgjExam) (err error) {
	o := orm.NewOrm()
	v := ZxgjExam{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjExam deletes ZxgjExam by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjExam(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjExam{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjExam{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
