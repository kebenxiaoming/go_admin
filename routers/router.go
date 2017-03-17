package routers

import (
	"go_admin/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	//首页
    beego.Router("/", &controllers.IndexController{},"*:Index")
    beego.Router("/admin", &controllers.IndexController{},"*:Index")
    beego.Router("/admin/Index", &controllers.IndexController{},"*:Index")
	//登陆退出
    beego.Router("/admin/Login", &controllers.LoginController{})
    beego.Router("/admin/Login/index", &controllers.LoginController{})
    beego.Router("/admin/Logout",&controllers.LogoutController{})


	//用户相关
	beego.Router("/admin/User/profile", &controllers.UserController{},"*:Profile")
	beego.Router("/admin/User/edit/uid/:uid:int", &controllers.UserController{},"*:Edit")
	beego.Router("/admin/User/index", &controllers.UserController{},"*:Index")
}
