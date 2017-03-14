/**
 * Created by sunny.
 * User: sunny
 * always for darling!
 * Date: 2017/3/13
 * Time: 11:28
 */
package controllers

import (
	"github.com/astaxie/beego"
	"go_admin/models"
)

type BaseController struct {
	beego.Controller
}

type ContentHeader struct{
	Menu_id int
	Menu_name string
}
//初始化整个的权限和目录
func (c *BaseController)Prepare(){
	menu:=&models.MenuUrl{}
	result:=menu.GetTrees()
	c.Data["sidebar"] = result
	c.Data["content_header"]=models.MenuUrl{21,"hehe","Index/index",1,1,1,1,1,"asdf",0}
	c.Data["current_module_id"]=1
}
