package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title 登录
// @Description 用户登录
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /signIn [post]
func (mrc *UserController) signIn() {

}

// @Title 登出
// @Description 用户登出，销毁token
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /signOut [post]
func (mrc *UserController) signOut() {

}

// @Title 注册
// @Description 用户注册账号
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /signUp [post]
func (mrc *UserController) signUp() {

}
