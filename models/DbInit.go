package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("datasource"))
	orm.SetMaxIdleConns("default", 1)
	orm.SetMaxOpenConns("default", 5)
	db, _ := orm.GetDB("default")
	db.SetConnMaxLifetime(3595 * time.Second)
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Classification))
	orm.RegisterModel(new(MoneyRecord))
}
