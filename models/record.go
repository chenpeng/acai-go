package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("datasource"))
	orm.RegisterModel(new(MoneyRecord))
}

type MoneyRecord struct {
	Id                 int64
	ClassificationCode int32
	ClassificationName string
	Money              float64
	Type               int32
	Remark             string
	RecordDateTime     time.Time
	PicUrl             string
	DeleteFlag         bool
	CreateUserId       int64
	CreateUserName     string
	CreateDateTime     time.Time
	UpdateUserId       int64
	UpdateUserName     string
	UpdateDateTime     time.Time
}

func AddMoneyRecord(mr MoneyRecord) (id int64, err error) {
	o := orm.NewOrm()
	mr.CreateDateTime = time.Now().UTC().Add(8 * time.Hour)
	mr.UpdateDateTime = time.Now().UTC().Add(8 * time.Hour)
	mr.RecordDateTime = time.Now().UTC().Add(8 * time.Hour)
	return o.Insert(&mr)
}

func GetMoneyRecord(id int64) (mr *MoneyRecord, err error) {
	o := orm.NewOrm()
	record := MoneyRecord{Id: id}
	return &record, o.Read(&record)
}

func GetAllMoneyRecord(pageIndex int, pageSize int) (list []*MoneyRecord, err error) {
	o := orm.NewOrm()
	var mrList []*MoneyRecord
	num, err := o.QueryTable("money_record").OrderBy("-id").Limit(pageSize, (pageIndex-1)*pageSize).All(&mrList)
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
