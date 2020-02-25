package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	State      int    `json:"state"`
	DeleteFlag bool   `json:"delete_flag"`
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
