package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kayz666/web-app/utils"
)

type EmailCtl struct{
	beego.Controller
}


func (c *EmailCtl)Get(){
	code:=c.GetString("code")
	email:=c.GetString("email")
	if code == ""||email==""{c.Abort("401")}
	err := utils.HTTP_ConfirmWithEmail(email,code)
	if err!= nil{
		c.Abort("401")
		return
	}
	c.Redirect("/login",301)
}