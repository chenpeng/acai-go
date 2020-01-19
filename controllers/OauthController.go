package controllers

import "github.com/astaxie/beego"

type OauthController struct {
	beego.Controller
}

// @Title GetPublicKey
// @Description get publicKey
// @Success 200 {object} models.User
// @router /publicKey [get]
func (mrc *OauthController) publicKey() {

}
