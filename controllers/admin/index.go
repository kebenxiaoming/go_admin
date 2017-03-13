package controllers

import (
	//"go_admin/models"
)
type IndexController struct {
	BaseController
}

type User_info struct{
	User_id  int
	User_type int
	User_name string
	Real_name string
	Mobile string
	Email string
	Login_time int64
	Login_ip string
}

func (c *IndexController) Get() {
	//c.Data["ss"]=c.GetSession("uid")
	//menuurl:=&models.MenuUrl{Menu_id:2}
	//result,err:=menuurl.GetMenuUrlById()
	user_info:=&User_info{1,2,"sunny","hehe","15866863307","xiaoyao_xiao@126.com",14567864321,"192.168.0.1"}
	c.Data["user_info"]=user_info
	c.TplName = "admin/index/index.tpl"
}
