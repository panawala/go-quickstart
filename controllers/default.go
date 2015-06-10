package controllers

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
