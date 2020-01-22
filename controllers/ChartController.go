package controllers

import "github.com/astaxie/beego"

type ChartController struct {
	beego.Controller
}

// @Title GetChart
// @Description get chart
// @Success 200 {object} models.User
// @router / [get]
func (mrc *ChartController) Get() {

}
