package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("datasource"))
	orm.RegisterModel(new(Classification))
}

type Classification struct {
	Id             int64
	Code           int32
	Name           string
	Type           int32
	Remark         string
	DeleteFlag     bool
	CreateUserId   int64
	CreateUserName string
	CreateDateTime time.Time
	UpdateUserId   int64
	UpdateUserName string
	UpdateDateTime time.Time
}

func GetAllClassification(classificationType int) (list []*Classification, err error) {
	o := orm.NewOrm()
	var cfList []*Classification
	num, err := o.QueryTable("classification").Filter("type", classificationType).All(&cfList)
	println(num)
	return cfList, err
}
