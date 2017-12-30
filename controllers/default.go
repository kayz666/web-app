package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type ConsoleCtl struct{
	beego.Controller
}
type LoginCtl struct {
	beego.Controller
}
type RegisterCtl struct {
	beego.Controller
}



func (c *MainController) Get() {
	c.Layout="iotconsole/layout/html_layout.html"
	c.TplName="iotconsole/index.html"
	c.LayoutSections= make(map[string]string)
	c.LayoutSections["Modal"]="iotconsole/module/ic_loginmodal.html"
	c.LayoutSections["Nav"]="iotconsole/module/ic_nav.html"
	//c.LayoutSections["SiderBar"] = "iotconsole/module/ic_sidebar.html"
}


func (c *ConsoleCtl) Get(){
	c.Layout="iotconsole/layout/html_layout.html"
	c.TplName="iotconsole/console.html"
	c.LayoutSections= make(map[string]string)
	c.LayoutSections["Modal"]="iotconsole/module/ic_loginmodal.html"
	c.LayoutSections["Nav"]="iotconsole/module/ic_nav.html"
	c.LayoutSections["SiderBar"] = "iotconsole/module/ic_sidebar.html"
}



func (c *LoginCtl) Post(){

}



func (c *RegisterCtl) Get(){
	//c.TplName="register.html"
	c.Layout="iotconsole/layout/html_layout.html"
	c.TplName="iotconsole/register.html"
	c.LayoutSections= make(map[string]string)
	c.LayoutSections["Modal"]="iotconsole/module/ic_loginmodal.html"
	c.LayoutSections["Nav"]="iotconsole/module/ic_nav.html"

}
func (c *RegisterCtl) Post(){

}