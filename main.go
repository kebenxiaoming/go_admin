package main

import (
	_ "go_admin/service"
	_ "go_admin/models"
	_ "go_admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

