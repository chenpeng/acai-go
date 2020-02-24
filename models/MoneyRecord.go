package models

import (
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
	//mr.RecordDateTime = time.Now().UTC().Add(8 * time.Hour)
	return o.Insert(&mr)
}

func GetMoneyRecord(id int64) (mr *MoneyRecord, err error) {
	o := orm.NewOrm()
	record := MoneyRecord{Id: id}
	return &record, o.Read(&record)
}

func GetAllMoneyRecord(pageIndex int, pageSize int, userId int64) (list []*MoneyRecord, err error) {
	o := orm.NewOrm()
	var mrList []*MoneyRecord
	num, err := o.QueryTable("money_record").Filter("create_user_id", userId).Filter("delete_flag", false).OrderBy("-record_date_time").Limit(pageSize, (pageIndex-1)*pageSize).All(&mrList)
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
