package controllers

type MainController struct {
	baseApiController
}

func (this *MainController) Get() {
	result := map[string]string{}
	result["email"] = "astaxie@gmail.com"
	result["website"] = "beego.me"
	this.Data["json"] = result
	this.ServeJson()
}
