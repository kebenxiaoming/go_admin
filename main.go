package main

import (
	_ "go_admin/service"
	_ "go_admin/models"
	_ "go_admin/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"go_admin/models"
)

func init(){
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("sunny_user").(models.User)
		if !ok && ctx.Request.RequestURI != "/admin/Login" {
			ctx.Redirect(302, "/admin/Login")
		}
	}

	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
}

func url(in string)(out string){
	out = "/admin/"+in
	return out
}

func main() {
	beego.AddFuncMap("url",url)
	beego.Run()
}

