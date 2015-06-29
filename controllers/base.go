package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {

}

type baseApiController struct {
	beego.Controller
}

func (this *baseApiController) GetCurrentUser(auth string) (auth_str map[string]string, err error) {
	return map[string]string{"id": "1"}, nil
}

func (this *baseApiController) format_input() {
	// 获取输入参数，依次为route中参数，post参数，url参数，body的json参数

	// InputRequestBody 得到body体中的json字符串
	body_json := map[string]interface{}{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &body_json)

	input_map := map[string]interface{}{}
	// params能得到路径中的命名参数, :key -> val
	for k, v := range this.Ctx.Input.Params {
		input_map[k[1:len(k)]] = v
	}
	// form可以得到get和post中的内容, key -> [val]
	for k, v := range this.Ctx.Input.Request.Form {
		input_map[k] = v[0]
	}
	// Input()得到get url后面的参数
	for k, v := range this.Input() {
		input_map[k] = v[0]
	}
	// body中的json字符串
	for k, v := range body_json {
		input_map[k] = v
	}
	this.Data["input_map"] = input_map
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

	this.format_input()
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
	result := ResponseMsg{401, "unauthorized"}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *ErrorController) Error400() {
	this.EnableRender = false
	result := ResponseMsg{400, "input invalid"}
	this.Data["json"] = result
	this.ServeJson()
}
