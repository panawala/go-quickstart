package routers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"html/template"
	"io"
	"net/http"
	"quickstart/controllers"
)

type ResponseMsg struct {
	Errcode int    `json:errcode`
	Errmsg  string `json:errmsg`
}

func unauthorized(rw http.ResponseWriter, r *http.Request) {
	// rw.Header().Set("Server", "beego test")
	// rw.WriteHeader(200)
	// rw.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(ResponseMsg{400, "unauthorized"})
	io.WriteString(rw, string(result))
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["Content"] = "page not found"
	t.Execute(rw, data)
}

func init() {
	// beego.Errorhandler("401", unauthorized)
	// beego.Errorhandler("404", page_not_found)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
}
