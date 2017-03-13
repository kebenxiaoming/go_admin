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

//初始化整个的权限和目录
func (c *BaseController)Prepare(){
	menu:=&models.MenuUrl{}
	result:=menu.GetTrees()
	c.Data["hehe"] = result
}
