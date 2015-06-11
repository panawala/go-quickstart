package main

import (
	"flag"
	"github.com/astaxie/beego"
	_ "quickstart/routers"
)

var env string

func init() {
	flag.StringVar(&env, "env", "dev", "application environment")
}

func main() {
	flag.Parse()
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	if env == "dev" {
		beego.AppConfigPath = "conf/app_dev.conf"
	} else {
		beego.AppConfigPath = "conf/app_prod.conf"
	}

	beego.Run()
}
