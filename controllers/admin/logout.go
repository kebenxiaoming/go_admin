package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.DestroySession()
	c.Data["code"]=1
	c.Data["msg"] = "已经退出"
	c.Data["url"]="/admin/Login/index"
	c.Data["wait"]=1
	c.TplName = "admin/public/redirect.tpl"
}


