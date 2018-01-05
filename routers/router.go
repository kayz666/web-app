package routers

import (
	"github.com/kayz666/web-app/controllers"
	"github.com/astaxie/beego"
)


func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})

	beego.Router("/console", &controllers.ConsoleCtl{})

	beego.Router("/login", &controllers.LoginCtl{})

	beego.Router("/register", &controllers.RegisterCtl{})

	beego.Router("/email/reg",&controllers.EmailCtl{})


}
