package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
}
