package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/helper"
)

func init() {

}

type ImageController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
}

func (this *ImageController) Post() {
	beego.MaxMemory = 1 << 22
	file, file_header, err := this.GetFile("avatar")
	if err != nil {
		msg := "file upload error" + err.Error()
		beego.Error(msg)
		return
	}

	cdn_url := helper.UploadFile(file, file.(Sizer).Size(), file_header.Header["Content-Type"][0])
	defer file.Close()

	this.Data["json"] = map[string]string{"file_url": cdn_url}
	this.ServeJson()
}
