package controllers

import (
	"github.com/astaxie/beego"
	"go_admin/models"
)
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	menuurl:=&models.MenuUrl{Menu_id:2}
	result,err:=menuurl.GetMenuUrlById()
	c.Data["Website"] = err
	c.Data["Email"] = result
	c.TplName = "admin/index.tpl"
}
