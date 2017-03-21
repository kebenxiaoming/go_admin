/**
 * Created by sunny.
 * User: sunny
 * always for darling!
 * Date: 2017/3/21
 * Time: 10:40
 */
package controllers

import "go_admin/models"

type GroupController struct {
	BaseController
}

func (c *GroupController)Index(){
	user:=&models.UserGroup{}
	groups,err:=user.GetALLGroups()
	if err==nil{
		c.Data["groups"]=groups
		c.TplName = "admin/user/index.tpl"
		return
	}else{
		c.redirect("/admin/Index","未获取到群组信息！",0,3)
		return
	}
}
