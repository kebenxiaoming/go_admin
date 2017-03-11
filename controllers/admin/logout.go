package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.SetSession("uid","")
	c.Data["result"] = "已经退出"
	c.TplName = "admin/public/error.tpl"
}


