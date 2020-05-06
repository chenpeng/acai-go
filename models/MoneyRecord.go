package models

import (
	"acai-go/dto"
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type MoneyRecord struct {
	Id                 int64     `json:"id"`
	ClassificationCode int32     `json:"classification_code"`
	ClassificationName string    `json:"classification_name"`
	Money              float64   `json:"money"`
	Type               int32     `json:"type"`
	Remark             string    `json:"remark"`
	RecordDateTime     string    `json:"record_date_time"`
	PicUrl             string    `json:"pic_url"`
	DeleteFlag         bool      `json:"delete_flag"`
	CreateUserId       int64     `json:"create_user_id"`
	CreateUserName     string    `json:"create_user_name"`
	CreateDateTime     time.Time `json:"create_date_time"`
	UpdateUserId       int64     `json:"update_user_id"`
	UpdateUserName     string    `json:"update_user_name"`
	UpdateDateTime     time.Time `json:"update_date_time"`
}

func AddMoneyRecord(mr MoneyRecord) (id int64, err error) {
	o := orm.NewOrm()
	mr.CreateDateTime = time.Now().UTC().Add(8 * time.Hour)
	mr.UpdateDateTime = time.Now().UTC().Add(8 * time.Hour)
	return o.Insert(&mr)
}

func GetMoneyRecord(id int64) (mr *MoneyRecord, err error) {
	o := orm.NewOrm()
	record := MoneyRecord{Id: id}
	return &record, o.Read(&record)
}

func GetAllMoneyRecord(year int, month int, userId int64) (list []*MoneyRecord, err error) {
	o := orm.NewOrm()
	var mrList []*MoneyRecord
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)
	num, err := o.Raw(`select id,
       classification_code,
       classification_name,
       money,
       type,
       remark,
       DATE_FORMAT( record_date_time, '%Y-%m-%d' ) as record_date_time,
       pic_url,
       delete_flag,
       create_user_id,
       create_user_name,
       create_date_time,
       update_user_id,
       update_user_name,
       update_date_time
from money_record
where create_user_id = ?
  and delete_flag = 0
  and record_date_time >= ?
  and record_date_time < ?
order by record_date_time desc, id desc`, userId, startDate, endDate).QueryRows(&mrList)
	println(num)
	return mrList, err
}

func UpdateMoneyRecord(id int64, mr *MoneyRecord) (num int64, err error) {
	o := orm.NewOrm()
	mr.UpdateDateTime = time.Now().UTC().Add(8 * time.Hour)
	return o.Update(mr)
}

func DeleteMoneyRecord(id int64) (num int64, err error) {
	o := orm.NewOrm()
	record := MoneyRecord{Id: id}
	if o.Read(&record) == nil {
		record.DeleteFlag = true
		record.UpdateDateTime = time.Now().UTC().Add(8 * time.Hour)
		if num, err := o.Update(&record); err == nil {
			return num, nil
		}
	}
	return 0, errors.New("没有查到要删除的数据")
}

func GetAllMoneyRecordChart(year int, month int, classificationType int, userId int64) (list []*dto.MoneyRecordChartDto, err error) {
	o := orm.NewOrm()
	var mrList []*dto.MoneyRecordChartDto
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)
	num, err := o.Raw(`select classification_code, classification_name, sum(money) as money
from money_record
where record_date_time >= ?
  and record_date_time < ?
  and type = ?
  and create_user_id = ?
group by classification_code, classification_name
order by money desc`, startDate, endDate, classificationType, userId).QueryRows(&mrList)
	println(num)
	return mrList, err
}
