package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ZxgjZybZyRecord struct {
	Id           int    `orm:"column:id;autoIncrement"` //ID
	ShopId       int    `orm:"column:shop_id"`
	ZybId        int    `orm:"column:zyb_id"` //试卷id
	ZyId         int    `orm:"column:zy_id"` //学员考试记录表id    对应goouc_xmf_exam_user的id
	StuZyId      int    `orm:"column:stu_zy_id"`
	Memberid     int    `orm:"column:memberid"` //学员id
	TkTestId     int    `orm:"column:tk_test_id"` //试题id,对应goouc_xmf_zxgj_tk_test表中的id
	OptionType   int8   `orm:"column:option_type"` //1单选题 2多选题 3判断题 4填空题 5问答题
	MAnswer      string `orm:"column:m_answer"` //学员答案
	AnswerImg    string `orm:"column:answer_img"` //图片答案
	TrueAnswer   string `orm:"column:true_answer;size:1000"`
	AnswerStatus int8   `orm:"column:answer_status"` //1：正确  2：错误
	Comment      string `orm:"column:comment"` //老师评语（简答）
	CommentImg   string `orm:"column:comment_img"` //评论图片
	IsPy         int8   `orm:"column:is_py"` //批阅状态（针对简答）1：已批阅 2未批阅
	Createtime   int    `orm:"column:createtime"` //提交时间
}

func (t *ZxgjZybZyRecord) TableName() string {
	return "goouc_xmf_zxgj_zyb_zy_record"
}

func init() {
	orm.RegisterModel(new(ZxgjZybZyRecord))
}

// AddZxgjZybZyRecord insert a new ZxgjZybZyRecord into database and returns
// last inserted Id on success.
func AddZxgjZybZyRecord(m *ZxgjZybZyRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZxgjZybZyRecordById retrieves ZxgjZybZyRecord by Id. Returns error if
// Id doesn't exist
func GetZxgjZybZyRecordById(id int) (v *ZxgjZybZyRecord, err error) {
	o := orm.NewOrm()
	v = &ZxgjZybZyRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZxgjZybZyRecord retrieves all ZxgjZybZyRecord matches certain condition. Returns empty list if
// no records exist
func GetAllZxgjZybZyRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZxgjZybZyRecord))
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

	var l []ZxgjZybZyRecord
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

// UpdateZxgjZybZyRecord updates ZxgjZybZyRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateZxgjZybZyRecordById(m *ZxgjZybZyRecord) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZyRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZxgjZybZyRecord deletes ZxgjZybZyRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZxgjZybZyRecord(id int) (err error) {
	o := orm.NewOrm()
	v := ZxgjZybZyRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZxgjZybZyRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
