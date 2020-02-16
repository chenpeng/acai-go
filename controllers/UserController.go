package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title 修改
// @Description 修改用户信息
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [put]
func (mrc *UserController) Put() {

}
