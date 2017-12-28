package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

type ConsoleCtl struct{
	beego.Controller
}
func (c *ConsoleCtl) Get(){
	c.TplName="console.html"
}

type LoginCtl struct {
	beego.Controller
}

func (c *LoginCtl) Get(){
	c.TplName="login.html"
}

type RegisterCtl struct {
	beego.Controller
}

func (c *RegisterCtl) Get(){
	c.TplName="register.html"
}