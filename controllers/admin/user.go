package controllers

import "go_admin/models"

type UserController struct {
	BaseController
}


func (c *UserController)Profile(){
	userInfo:=c.GetSession("sunny_user")
	if(userInfo!=nil){
		user_info:=userInfo.(models.User)
		c.Data["user_info"]=user_info
	}
	c.TplName = "admin/index/index.tpl"
}