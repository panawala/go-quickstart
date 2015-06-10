package controllers

import (
	"fmt"
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
	file, file_header, err := this.GetFile("upload")
	fmt.Println(file_header.Filename)
	fmt.Println(file_header.Header["Content-Type"][0])
	if err != nil {
		msg := "file upload error" + err.Error()
		fmt.Println(msg)
		return
	}

	cdn_url := helper.UploadFile(file, file.(Sizer).Size(), file_header.Header["Content-Type"][0])
	defer file.Close()
	this.Data["json"] = map[string]string{"file_url": cdn_url}
	this.ServeJson()
}
