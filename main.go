package main

import (
	_ "github.com/kayz666/web-app/routers"
	"github.com/astaxie/beego"
	//"github.com/kayz666/web-app/controllers"
)

func main() {
	//beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

