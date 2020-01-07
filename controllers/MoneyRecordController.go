package controllers

import (
	"acai-go/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about MoneyRecordList
type MoneyRecordController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (mrc *MoneyRecordController) Post() {
	var moneyRecord models.MoneyRecord
	json.Unmarshal(mrc.Ctx.Input.RequestBody, &moneyRecord)
	row, err := models.AddMoneyRecord(moneyRecord)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "新增失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: row, Message: "新增成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (mrc *MoneyRecordController) GetAll() {
	pageIndexStr := mrc.Ctx.Input.Query("pageIndex")
	pageSizeStr := mrc.Ctx.Input.Query("pageSize")
	pageIndex, _ := strconv.Atoi(pageIndexStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	moneyRecordList, err := models.GetAllMoneyRecord(pageIndex, pageSize)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "查询失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: moneyRecordList, Message: "查询成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:id [get]
func (mrc *MoneyRecordController) Get() {
	id, _ := mrc.GetInt64(":id")
	moneyRecord, err := models.GetMoneyRecord(id)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "查询失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: moneyRecord, Message: "查询成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:id [put]
func (mrc *MoneyRecordController) Put() {
	id, _ := mrc.GetInt64(":id")
	var moneyRecord models.MoneyRecord
	json.Unmarshal(mrc.Ctx.Input.RequestBody, &moneyRecord)
	hehe, err := models.UpdateMoneyRecord(id, &moneyRecord)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "更新失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: hehe, Message: "更新成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:id [delete]
func (mrc *MoneyRecordController) Delete() {
	id, _ := mrc.GetInt64(":id")
	row, err := models.DeleteMoneyRecord(id)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "删除失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: row, Message: "删除成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}
