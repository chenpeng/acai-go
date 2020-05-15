package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("datasource"))
	orm.SetMaxIdleConns("default", 1)
	orm.SetMaxOpenConns("default", 5)
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Classification))
	orm.RegisterModel(new(MoneyRecord))
}
