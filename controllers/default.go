package controllers

import (
	"encoding/json"
	"fmt"
	"quickstart/models"
)

type MainController struct {
	baseApiController
}

func (this *MainController) validate() {
	fmt.Println("validate from main api")
}

func (this *MainController) Get() {
	result := map[string]interface{}{}
	result["email"] = "astaxie@gmail.com"
	result["website"] = "beego.me"
	result["user"] = this.Data["current_user"].(map[string]string)["name"]
	this.Data["json"] = result

	fmt.Println(this.Ctx.Input.Params)
	fmt.Println(this.Ctx.Input.RequestBody)

	this.ServeJson()
}

func (this *MainController) Post() {
	result := map[string]interface{}{}
	result["email"] = "astaxie@gmail.com"
	result["website"] = "beego.me"
	result["user"] = this.Data["current_user"].(map[string]string)["name"]
	this.Data["json"] = result

	fmt.Println(this.Ctx.Input.Params)
	fmt.Println(this.Ctx.Input.Params[":catId"])
	fmt.Println(this.Ctx.Input.Request.Form)
	fmt.Println(this.Ctx.Input.Request.PostForm)
	this.Ctx.Input.Request.ParseMultipartForm(1 << 22)
	fmt.Println(this.Ctx.Input.Request.MultipartForm)
	fmt.Println(this.GetString("aaa"))

	input_json := map[string]interface{}{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &input_json)
	fmt.Println(input_json)

	this.ServeJson()
}

type RedisController struct {
	baseApiController
}

func (this *RedisController) Get() {
	models.RedisInstance()
	models.RedisInstance().Set("name", "williampan")
	models.RedisInstance().Get("name")

	result := map[string]interface{}{}
	result["redis"] = "redis is connected"
	this.Data["json"] = result

	this.ServeJson()
}

type GoSchemaController struct {
	baseApiController
}

func (this *GoSchemaController) Get() {
	// models.TestGoSchema()

	result := map[string]interface{}{}
	result["go"] = "go schema"
	this.Data["json"] = result

	this.ServeJson()
}

func (this *GoSchemaController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	input_json := map[string]interface{}{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &input_json)
	fmt.Println(input_json)

	// models.ValidInput(input_json)

	models.TestTag()

	result := map[string]interface{}{}
	result["go"] = "beego post"
	this.Data["json"] = result

	this.ServeJson()
}
