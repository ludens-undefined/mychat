package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjExamUser struct {
	Id           int     `orm:"column:id;autoIncrement"`
	ShopId       int     `orm:"column:shop_id"`
	ExamId       uint    `orm:"column:exam_id"` //考试id，对应goouc_xet_zxgj_exam表中的id
	Memberid     uint    `orm:"column:memberid"` //学员id
	Score        float64 `orm:"column:score;scale:10;precision:1"` //学员考试分值
	Kgscore      float64 `orm:"column:kgscore;scale:10;precision:1"` //客观题得分
	Zgscore      float64 `orm:"column:zgscore;scale:10;precision:1"` //主观题得分
	Truenum      uint8   `orm:"column:truenum"` //正确数量
	Falsenum     uint8   `orm:"column:falsenum"` //错题数量
	Notdonum     uint8   `orm:"column:notdonum"` //没做的数量
	IsPy         uint8   `orm:"column:is_py"` //1批阅 2未批阅
	OperatorName string  `orm:"column:operator_name;size:255"` //批阅人姓名
	PyOperator   int     `orm:"column:py_operator"` //批阅人，user表id
	UseTime      int     `orm:"column:use_time"` //考试用时（秒）
	CreateAt     uint    `orm:"column:create_at"` //提交时间
	PyTime       int     `orm:"column:py_time"` //批阅时间
	ZsId         int     `orm:"column:zs_id"` //证书路径
}

func (t *ZxgjExamUser) TableName() string {
	return "goouc_xmf_zxgj_exam_user"
}

func init() {
	orm.RegisterModel(new(ZxgjExamUser))
}

// AddZxgjExamUser insert a new ZxgjExamUser into database and returns
// last inserted Id on success.
func AddZxgjExamUser(m *ZxgjExamUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjExamUserById retrieves ZxgjExamUser by Id. Returns error if
// Id doesn't exist
func GetZxgjExamUserById(id int) (v *ZxgjExamUser, err error) {
	o := orm.NewOrm()
	v = &ZxgjExamUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjExamUser retrieves all ZxgjExamUser matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjExamUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjExamUser))
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

	var l []ZxgjExamUser
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

// UpdateZxgjExamUser updates ZxgjExamUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjExamUserById(m *ZxgjExamUser) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjExamUser deletes ZxgjExamUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjExamUser(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjExamUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjExamUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
