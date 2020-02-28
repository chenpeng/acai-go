package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["acai-go/controllers:ChartController"] = append(beego.GlobalControllerRouter["acai-go/controllers:ChartController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:ClassificationController"] = append(beego.GlobalControllerRouter["acai-go/controllers:ClassificationController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "Download",
			Router:           `/download`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
		beego.ControllerComments{
			Method:           "Upload",
			Router:           `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:OauthController"] = append(beego.GlobalControllerRouter["acai-go/controllers:OauthController"],
		beego.ControllerComments{
			Method:           "Ceshi",
			Router:           `/ceshi`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:OauthController"] = append(beego.GlobalControllerRouter["acai-go/controllers:OauthController"],
		beego.ControllerComments{
			Method:           "Decrypt",
			Router:           `/decrypt`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:OauthController"] = append(beego.GlobalControllerRouter["acai-go/controllers:OauthController"],
		beego.ControllerComments{
			Method:           "Encrypt",
			Router:           `/encrypt`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:OauthController"] = append(beego.GlobalControllerRouter["acai-go/controllers:OauthController"],
		beego.ControllerComments{
			Method:           "PublicKey",
			Router:           `/publicKey`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:OauthController"] = append(beego.GlobalControllerRouter["acai-go/controllers:OauthController"],
		beego.ControllerComments{
			Method:           "SignIn",
			Router:           `/signIn`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:OauthController"] = append(beego.GlobalControllerRouter["acai-go/controllers:OauthController"],
		beego.ControllerComments{
			Method:           "SignUp",
			Router:           `/signUp`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["acai-go/controllers:UserController"] = append(beego.GlobalControllerRouter["acai-go/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
