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
	"strings"
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
	//赋值用户信息
	userInfo:=c.GetSession("sunny_user")
	if(userInfo!=nil){
		user_info:=userInfo.(models.User)
		c.Data["user_info"]=user_info
	}
	oldAccess:=c.GetSession("sunny_user_role")
	if(oldAccess!=nil){
		menu:=&models.MenuUrl{}
		access:=oldAccess.(string)
		result := menu.GetTrees(access)
		c.Data["sidebar"] = result
	}
	//获取当前的路由
	oldController,action:=c.GetControllerAndAction()
	if action==""{
		action="index"
	}
	//替换Controller
	controller:=strings.Replace(oldController,"Controller","",-1)
	nowUrl:=controller+"/"+strings.ToLower(action)
	//根据url获取对应的当前菜单
	if nowUrl!="" {
		urlMenu := &models.MenuUrl{Menu_url:nowUrl}
		nowMenu,err:=urlMenu.GetMenuUrlByUrl()
		if err==nil{
			c.Data["content_header"]=nowMenu
			c.Data["current_module_id"]=nowMenu.Module_id
		}
	}
}
