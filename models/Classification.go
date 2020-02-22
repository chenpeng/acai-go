package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Classification struct {
	Id             int64     `json:"id"`
	Code           int32     `json:"code"`
	Name           string    `json:"name"`
	Type           int32     `json:"type"`
	Remark         string    `json:"remark"`
	DeleteFlag     bool      `json:"delete_flag"`
	CreateUserId   int64     `json:"create_user_id"`
	CreateUserName string    `json:"create_user_name"`
	CreateDateTime time.Time `json:"create_date_time"`
	UpdateUserId   int64     `json:"update_user_id"`
	UpdateUserName string    `json:"update_user_name"`
	UpdateDateTime time.Time `json:"update_date_time"`
}

func GetAllClassification(classificationType int) (list []*Classification, err error) {
	o := orm.NewOrm()
	var cfList []*Classification
	num, err := o.QueryTable("classification").Filter("type", classificationType).All(&cfList)
	println(num)
	return cfList, err
}
