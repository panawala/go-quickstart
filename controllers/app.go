package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
)

func init() {

}

type Validator interface {
	validate()
}

type baseApiController struct {
	beego.Controller
	controller_name string
}

func (this *baseApiController) validate() {
	fmt.Println("validate from base api")
}

func (this *baseApiController) GetCurrentUser(auth string) (auth_str map[string]string, err error) {
	return map[string]string{"name": auth}, nil
}

func (this *baseApiController) validate_input() {
	method := this.Ctx.Request.Method
	fmt.Println(reflect.TypeOf(this))
	if method == "POST" {
		fmt.Println("post .....")
	}
	if method == "GET" {
		fmt.Println("get ....")
	}
	this.validate()
}

func (this *baseApiController) Prepare() {
	header_auth := this.Ctx.Request.Header.Get("Authorization")
	if header_auth == "" {
		this.Abort("401")
	} else {
		current_user := map[string]string{"name": "name1"}
		fmt.Println("the header_Auth is ")
		fmt.Println(header_auth)
		current_user, err := this.GetCurrentUser(header_auth)
		if err != nil {
			this.Data["json"] = map[string]string{"error": err.Error()}
			this.ServeJson()
		}
		this.Data["current_user"] = current_user
	}

	this.validate_input()
}

type ResponseMsg struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error401() {
	this.EnableRender = false
	result := ResponseMsg{400, "unauthorized"}
	this.Data["json"] = result
	this.ServeJson()
}
