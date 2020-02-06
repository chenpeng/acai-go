package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("datasource"))
	orm.RegisterModel(new(User))
}

type User struct {
	Id         int64
	Username   string
	Password   string
	State      int
	DeleteFlag bool
}

func FindByUsername(username string) (User, error) {
	o := orm.NewOrm()
	user := User{Username: username}
	err := o.Read(&user, "Username")
	return user, err
}

func AddUser(user User) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(&user)
}
