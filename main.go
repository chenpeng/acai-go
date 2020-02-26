package main

import (
	"acai-go/filter"
	_ "acai-go/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	beego.InsertFilter("*", beego.BeforeRouter, filter.OauthFilter)
	beego.Run()
}
