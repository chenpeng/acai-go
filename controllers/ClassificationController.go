package controllers

import (
	"acai-go/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about MoneyRecordList
type ClassificationController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Classification
// @Success 200 {object} models.Classification
// @router / [get]
func (mrc *ClassificationController) GetAll() {
	typeStr := mrc.Ctx.Input.Query("type")
	classificationType, _ := strconv.Atoi(typeStr)
	classificationList, err := models.GetAllClassification(classificationType)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "查询失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: classificationList, Message: "查询成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}
