package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
	_ "quickstart/models"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/blog/:catId", &controllers.MainController{})
	beego.Router("/upload", &controllers.ImageController{})
	beego.Router("/redis", &controllers.RedisController{})
	beego.Router("/goschema", &controllers.GoSchemaController{})
}
