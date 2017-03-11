package routers

import (
	"go_admin/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/admin", &controllers.IndexController{})
    beego.Router("/admin/login", &controllers.LoginController{})
    beego.Router("/admin/logout",&controllers.LogoutController{})
}
