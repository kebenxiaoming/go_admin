package main

import (
	_ "hello/service"
	_ "hello/models"
	_ "hello/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

