package controllers

import "github.com/astaxie/beego"

type OauthController struct {
	beego.Controller
}

// @Title 下发公钥
// @Description get publicKey
// @Success 200 {object} models.User
// @router /publicKey [get]
func (mrc *OauthController) PublicKey() {

}
