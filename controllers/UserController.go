package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /login [post]
func (mrc *UserController) login() {

}
