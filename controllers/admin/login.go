package controllers

import (
	"github.com/astaxie/beego"
)
type LoginController struct {
	beego.Controller
}

type loginUser struct {
	Username    string         `form:"user_name"`
	Password  string `form:"password"`
}

func (c *LoginController) Get() {
	c.TplName = "admin/login/index.tpl"
}

func (c *LoginController) Post() {
	u := loginUser{}
	if err := c.ParseForm(&u); err != nil {
		c.Data["result"] = err
		c.TplName = "admin/public/error.tpl"
	}else{
		if(u.Username=="sunny"&&u.Password=="111111"){
			c.SetSession("uid",1)
			c.Ctx.Redirect(302, "/admin")
		}else{
			c.Data["err"]="用户名或者密码错误"
			c.TplName = "admin/login/index.tpl"
		}
	}
}
