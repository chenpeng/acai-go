// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"acai-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/moneyRecord",
			beego.NSInclude(
				&controllers.MoneyRecordController{},
			),
		),
		beego.NSNamespace("/classification",
			beego.NSInclude(
				&controllers.ClassificationController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/oauth",
			beego.NSInclude(
				&controllers.OauthController{},
			),
		),
		beego.NSNamespace("/chart",
			beego.NSInclude(
				&controllers.ChartController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
