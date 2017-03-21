package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/cache"
	"go_admin/models"
)
type LoginController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store) //一定要写在构造函数里面，要不然第一次打开页面有可能是X
	cpt.ChallengeNums = 2
	cpt.StdWidth = 80
	cpt.StdHeight = 25
}

//登陆页
func (c *LoginController) Get() {
	c.TplName = "admin/login/index.tpl"
}
//登陆
func (c *LoginController) Post() {
	verify:= cpt.VerifyReq(c.Ctx.Request)
	if verify {
		username := c.Input().Get("user_name")
		password:=c.Input().Get("password")
		if username=="" || password=="" {
			c.Data["err"] = "请填写完整再提交!"
			c.TplName = "admin/login/index.tpl"
		} else {
			u := &models.User{User_name:username}
			ok,user:=u.Login(username,password)
			if ok {
				c.SetSession("sunny_user", user)
				//在这里查出对应的权限的内容
				userGroup:=&models.UserGroup{Group_id:user.User_group}
				ug,err:=userGroup.GetUserGroupById()
				if(err==nil){
					c.SetSession("sunny_user_role",ug.Group_role)
				}
				c.Ctx.Redirect(302, "/admin")
			} else {
				c.Data["err"] = "用户名或者密码错误！"
				c.TplName = "admin/login/index.tpl"
			}
		}
	}else{
		c.Data["err"] = "验证码错误！"
		c.TplName = "admin/login/index.tpl"
	}
}

