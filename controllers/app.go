package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

func init() {

}

type baseApiController struct {
	beego.Controller
}

func (this *baseApiController) GetCurrentUser(auth string) (auth_str map[string]string, err error) {
	return map[string]string{"name": auth}, nil
}

func (this *baseApiController) Prepare() {
	header_auth := this.Ctx.Request.Header.Get("Authorization")
	if header_auth == "" {
		this.Abort("401")
		// fmt.Println("header auth is nil")
		// this.Data["json"] = map[string]string{"error": "authorization is empty"}
		// this.ServeJson()
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
	// this.Ctx.Output.Header("Content-Type", "application/json")
	// result, _ := json.Marshal(ResponseMsg{400, "unauthorized"})
	result := ResponseMsg{400, "unauthorized"}
	// result := map[string]string{"errcode": "401", "errmsg": "unauthorized"}
	// fmt.Println(&result)
	// this.Data["json"] = map[string]string{"errcode": "401", "errmsg": "unauthorized"}
	this.Data["json"] = result
	this.ServeJson()
	// this.Ctx.Output.Body([]byte(result))
}
