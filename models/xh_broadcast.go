package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type XhBroadcast struct {
	Id          int    `orm:"column:id;autoIncrement"`
	PushDomain  string `orm:"column:push_domain;size:255"` //推流域名
	PalyDomain  string `orm:"column:paly_domain;size:255"` //播流域名
	AppName     string `orm:"column:app_name;size:255"` //Appname
	Ip          string `orm:"column:ip;size:255"` //聊天服务IP设置
	ShopId      int    `orm:"column:shop_id"`
	Type        int8   `orm:"column:type"` //1.视频剪辑开关  2.直播参数设置
	Classify    int8   `orm:"column:classify"` //1：阿里云
	Bucketname  string `orm:"column:bucketname;size:255"`
	Endpoint    string `orm:"column:endpoint;size:500"`
	LiveTime    int    `orm:"column:live_time"`
	BroadTime   int    `orm:"column:broad_time"`
	MpsRegionId string `orm:"column:mps_region_id;size:255"` //剪辑location
	PipelineId  string `orm:"column:pipeline_id;size:500"` //管道id
	TemplateId  string `orm:"column:template_id;size:500"` //转码模板id
	OssLocation string `orm:"column:oss_location;size:500"` //oss区域oss-cn-shenzhen
	OssBucket   string `orm:"column:oss_bucket;size:500"` //oss_bucket
	OssStatus   int8   `orm:"column:oss_status"` //视频剪辑1:开启  2：关闭
}

func (t *XhBroadcast) TableName() string {
	return "goouc_xmf_xh_broadcast"
}

func init() {
	orm.RegisterModel(new(XhBroadcast))
}

// AddXhBroadcast insert a new XhBroadcast into database and returns
// last inserted Id on success.
func AddXhBroadcast(m *XhBroadcast) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetXhBroadcastById retrieves XhBroadcast by Id. Returns error if
// Id doesn't exist
func GetXhBroadcastById(id int) (v *XhBroadcast, err error) {
	o := orm.NewOrm()
	v = &XhBroadcast{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllXhBroadcast retrieves all XhBroadcast matches certain condition. Returns empty list if
// no records exist
func GetAllXhBroadcast(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(XhBroadcast))
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

	var l []XhBroadcast
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

// UpdateXhBroadcast updates XhBroadcast by Id and returns error if
// the record to be updated doesn't exist
func UpdateXhBroadcastById(m *XhBroadcast) (err error) {
	o := orm.NewOrm()
	v := XhBroadcast{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteXhBroadcast deletes XhBroadcast by Id and returns error if
// the record to be deleted doesn't exist
func DeleteXhBroadcast(id int) (err error) {
	o := orm.NewOrm()
	v := XhBroadcast{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&XhBroadcast{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
