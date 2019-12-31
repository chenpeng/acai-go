package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"] = append(beego.GlobalControllerRouter["acai-go/controllers:MoneyRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["acai-go/controllers:NasController"] = append(beego.GlobalControllerRouter["acai-go/controllers:NasController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/download`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["acai-go/controllers:NasController"] = append(beego.GlobalControllerRouter["acai-go/controllers:NasController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
