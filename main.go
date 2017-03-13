package main

import (
	_ "go_admin/service"
	_ "go_admin/models"
	_ "go_admin/routers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
)

func init(){
	//var FilterUser = func(ctx *context.Context) {
	//	_, ok := ctx.Input.Session("sunny_user_uid").(int)
	//	if !ok && ctx.Request.RequestURI != "/admin/login" {
	//		ctx.Redirect(302, "/admin/login")
	//	}
	//}
	//
	//beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
}

func main() {
	beego.Run()
}

