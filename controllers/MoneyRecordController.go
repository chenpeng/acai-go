package controllers

import (
	"acai-go/dto"
	"acai-go/models"
	"acai-go/util"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
	"path"
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
	user := util.GetUser(mrc.Ctx)
	moneyRecord.CreateUserId = user.Id
	moneyRecord.UpdateUserId = user.Id
	moneyRecord.CreateUserName = user.Nickname
	moneyRecord.UpdateUserName = user.Nickname
	row, err := models.AddMoneyRecord(moneyRecord)
	if err != nil {
		result := dto.Result{Code: 1, Data: nil, Message: "新增失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: row, Message: "新增成功"}
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
	user := util.GetUser(mrc.Ctx)
	userId := user.Id
	moneyRecordList, err := models.GetAllMoneyRecord(pageIndex, pageSize, userId)
	if err != nil {
		result := dto.Result{Code: 1, Data: nil, Message: "查询失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: moneyRecordList, Message: "查询成功"}
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
		result := dto.Result{Code: 1, Data: nil, Message: "查询失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: moneyRecord, Message: "查询成功"}
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
	user := util.GetUser(mrc.Ctx)
	moneyRecord.Id = id
	moneyRecord.UpdateUserId = user.Id
	moneyRecord.UpdateUserName = user.Nickname
	hehe, err := models.UpdateMoneyRecord(id, &moneyRecord)
	if err != nil {
		result := dto.Result{Code: 1, Data: nil, Message: "更新失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: hehe, Message: "更新成功"}
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
		result := dto.Result{Code: 1, Data: nil, Message: "删除失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: row, Message: "删除成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title Download
// @Description Download files
// @Success 200 {int}
// @Failure 403 body is empty
// @router /upload [post]
func (mrc *MoneyRecordController) Upload() {
	_, info, _ := mrc.GetFile("file")
	fileSuffix := path.Ext(path.Base(info.Filename))
	uuid := uuid.NewV1()
	fileName := uuid.String() + fileSuffix

	err := mrc.SaveToFile("file", path.Join(beego.AppConfig.String("filepath"), fileName))
	if err != nil {
		result := dto.Result{Code: 1, Data: nil, Message: "上传失败"}
		mrc.Data["json"] = result
	} else {
		result := dto.Result{Code: 0, Data: fileName, Message: "上传成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title Download
// @Description Download files
// @Param	picUrl		query 	string	true		"The key for picUrl"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /download [get]
func (mrc *MoneyRecordController) Download() {
	picUrl := mrc.Ctx.Input.Query("picUrl")
	w := mrc.Ctx.ResponseWriter
	//打开文件
	file, err := os.Open(path.Join(beego.AppConfig.String("filepath"), picUrl))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	//结束后关闭文件
	defer file.Close()

	//设置响应的header头
	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+picUrl+"\"")
	//将文件写至responseBody
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
}
