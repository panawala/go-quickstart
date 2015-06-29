package controllers

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"quickstart/models"
)

type MainController struct {
	baseApiController
}

func (this *MainController) Get() {
	result := map[string]interface{}{}
	result["email"] = "astaxie@gmail.com"
	result["website"] = "beego.me"
	result["user"] = this.Data["current_user"].(map[string]string)["name"]
	this.Data["json"] = result

	this.ServeJson()
}

func (this *MainController) Prepare() {
	this.baseApiController.Prepare()

	schemaString := map[string]string{
		"POST": `
		{
			"properties": {
				"website": {
					"type": "string"
				}
			},
			"required": ["website"],
			"type": "object"
		}`,
		"GET": `
		{
			"properties": {
				"get_key": {
					"type": "string"
				}
			},
			"required": ["get_key"],
			"type": "object"
		}`,
	}

	schemaLoader := gojsonschema.NewStringLoader(schemaString[this.Ctx.Request.Method])
	documentLoader := gojsonschema.NewGoLoader(this.Data["input_map"])
	validate_result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println(err.Error())
		this.Data["error"] = err.Error()
	}

	if validate_result.Valid() {
		fmt.Println("The document is valid\n")
	} else {
		fmt.Println("The document is not valid \n")
		for _, desc := range validate_result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		this.Abort("400")
	}
}

func (this *MainController) Post() {
	result := map[string]interface{}{}
	result["email"] = "astaxie@gmail.com"
	result["website"] = "beego.me"
	result["user"] = this.Data["current_user"].(map[string]string)["name"]
	result["input_map"] = this.Data["input_map"]
	this.Data["json"] = result

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
