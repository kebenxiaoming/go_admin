package routers

import (
	"go_admin/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{},"*:Index")
    beego.Router("/admin", &controllers.IndexController{},"*:Index")
    beego.Router("/admin/Index", &controllers.IndexController{},"*:Index")
    beego.Router("/admin/User/profile", &controllers.UserController{},"*:Profile")
    beego.Router("/admin/Login", &controllers.LoginController{})
    beego.Router("/admin/Login/index", &controllers.LoginController{})
    beego.Router("/admin/Logout",&controllers.LogoutController{})
}
