package controllers

import (
	"acai-go/models"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"path"
)

type NasController struct {
	beego.Controller
}

// @Title Download
// @Description Download files
// @Success 200 {int}
// @Failure 403 body is empty
// @router /download [get]
func (mrc *NasController) Download() {

}

// @Title Download
// @Description Download files
// @Success 200 {int}
// @Failure 403 body is empty
// @router /upload [post]
func (mrc *NasController) Upload() {
	_, info, _ := mrc.GetFile("file")
	fileSuffix := path.Ext(path.Base(info.Filename))
	uuid, _ := uuid.NewV1()
	fileName := uuid.String() + fileSuffix

	err := mrc.SaveToFile("file", path.Join(beego.AppConfig.String("filepath"), fileName))
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "上传失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: fileName, Message: "上传成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}
