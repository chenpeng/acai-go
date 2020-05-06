package controllers

import (
	"acai-go/dto"
	"acai-go/models"
	"acai-go/util"
	"github.com/astaxie/beego"
	"strconv"
)

type ChartController struct {
	beego.Controller
}

// @Title GetChart
// @Description get chart
// @Success 200 {object} models.User
// @router / [get]
func (mrc *ChartController) Get() {
	yearStr := mrc.Ctx.Input.Query("year")
	monthStr := mrc.Ctx.Input.Query("month")
	classificationTypeStr := mrc.Ctx.Input.Query("classificationType")
	year, _ := strconv.Atoi(yearStr)
	month, _ := strconv.Atoi(monthStr)
	classificationType, _ := strconv.Atoi(classificationTypeStr)
	user := util.GetUser(mrc.Ctx)
	userId := user.Id
	userId = 1
	moneyRecordChartList, err := models.GetAllMoneyRecordChart(year, month, classificationType, userId)
	if err != nil {
		result := dto.Result{Code: 1, Data: nil, Message: "查询失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: moneyRecordChartList, Message: "查询成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}
