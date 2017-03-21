package controllers

import (
	"go_admin/models"
	"crypto/md5"
	"io"
	"encoding/hex"
)

type UserController struct {
	BaseController
}

func (c *UserController)Profile(){
	userInfo:=c.GetSession("sunny_user")
	if(userInfo!=nil){
		user_info:=userInfo.(models.User)
		c.Data["user_info"]=user_info
	}
	c.TplName = "admin/index/index.tpl"
}
//用户信息列表
func (c *UserController)Index(){
	user:=&models.User{}
	users,error:=user.GetAllUsers()
	if error==nil{
		c.Data["user_infos"]=users
		c.TplName = "admin/user/index.tpl"
		return
	}else{
		c.Data["code"]=0
		c.Data["msg"] = "未获取到用户信息！"
		c.Data["url"]="/admin/Index"
		c.Data["wait"]=3
		c.TplName = "admin/public/redirect.tpl"
		return
	}
}
//编辑用户信息
func (c *UserController)Edit(){
	//当前登陆用户
	nowUser:=c.GetSession("sunny_user")
	var user_info models.User
	if nowUser!=nil{
		user_info=nowUser.(models.User)
	}
	group:=models.UserGroup{}
	groups,err:=group.GetALLGroups()
	if err==nil{
		c.Data["group_option_list"]=groups
	}
	currentMethod:=c.Controller.Ctx.Request.Method
	if currentMethod=="GET"{
		id,_ := c.GetInt(":uid")
		user:=&models.User{User_id:id}
		myUser,err:=user.GetUserById()
		if err==nil{
			c.Data["group_id"]=myUser.User_group
			c.Data["user"]=myUser
			c.TplName = "admin/user/edit.tpl"
			return
		}else{
			c.Data["code"]=0
			c.Data["msg"] = "未获取到id！"
			c.Data["url"]="/admin/Index"
			c.Data["wait"]=3
			c.TplName = "admin/public/redirect.tpl"
			return
		}
	}else if currentMethod=="POST"{
		user_id,_ := c.GetInt(":uid")
		uid:= c.Input().Get("user_id")
		username := c.Input().Get("user_name")
		user_desc := c.Input().Get("user_desc")
		password := c.Input().Get("password")
		real_name := c.Input().Get("real_name")
		mobile := c.Input().Get("mobile")
		email := c.Input().Get("email")
		user_group,_ := c.GetInt("user_group")

		//查询用户是否存在
		user:=&models.User{User_id:user_id}
		myUser,err:=user.GetUserById()
		if err==nil {
			c.Data["group_id"]=myUser.User_group
			if user_id == 1&&password != "" {
				c.Data["code"] = 0
				c.Data["msg"] = "初始管理员禁止修改密码！"
				c.Data["url"] = "/admin/User/edit/uid/" + uid
				c.Data["wait"] = 3
				c.TplName = "admin/public/redirect.tpl"
				return
			}
			myUser.User_id=user_id
			myUser.User_name=username
			myUser.User_desc=user_desc
			myUser.User_group=user_group
			if password != "" {
				m := md5.New()
				io.WriteString(m, password)
				keyMd5 := hex.EncodeToString(m.Sum(nil))
				myUser.Password = keyMd5
			}
			myUser.Real_name=real_name
			myUser.Email=email
			myUser.Mobile=mobile
			myUser.User_group=user_group
			userInfo := &myUser
			refreshUser, err := userInfo.UpdateUser()
				if err == nil {
					//如果是当前用户更新Session
					if user_info.User_id==myUser.User_id {
						c.SetSession("sunny_user", refreshUser)
					}
					c.Data["code"] = 1
					c.Data["msg"] = "修改成功！"
					c.Data["url"] = "/admin/Index"
					c.Data["wait"] = 1
					c.TplName = "admin/public/redirect.tpl"
				} else {
					c.Data["code"] = 0
					c.Data["msg"] = "修改失败！"
					c.Data["url"] = "/admin/User/edit/uid/" + uid
					c.Data["wait"] = 3
					c.TplName = "admin/public/redirect.tpl"
				}
				return
			}else{
			c.Data["code"] = 0
			c.Data["msg"] = "未获取到用户信息！"
			c.Data["url"] = "/admin/User/edit/uid/" + uid
			c.Data["wait"] = 3
			c.TplName = "admin/public/redirect.tpl"
			return
		}
	}
}
//添加用户
//编辑用户信息
func (c *UserController)Add(){
	group:=models.UserGroup{}
	groups,err:=group.GetALLGroups()
	if err==nil{
		c.Data["group_option_list"]=groups
	}
	currentMethod:=c.Controller.Ctx.Request.Method
	if currentMethod=="GET"{
		c.Data["group_id"]=1
		c.TplName = "admin/user/add.tpl"
		return
	}else if currentMethod=="POST"{
		username := c.Input().Get("user_name")
		user_desc := c.Input().Get("user_desc")
		password := c.Input().Get("password")
		real_name := c.Input().Get("real_name")
		mobile := c.Input().Get("mobile")
		email := c.Input().Get("email")
		user_group,_ := c.GetInt("user_group")

		//查询用户是否存在
		user:=&models.User{User_name:username}
		myUser,err:=user.GetUserById()
		if err==nil {
			c.Data["code"] = 0
			c.Data["msg"] = "已经存在该用户名的用户，禁止重复！！"
			c.Data["url"] = "/admin/User/add"
			c.Data["wait"] = 3
			c.TplName = "admin/public/redirect.tpl"
			return

		}else{
			myUser.User_name=username
			myUser.User_desc=user_desc
			myUser.User_group=user_group
			if password != "" {
				m := md5.New()
				io.WriteString(m, password)
				keyMd5 := hex.EncodeToString(m.Sum(nil))
				myUser.Password = keyMd5
			}
			myUser.Real_name=real_name
			myUser.Email=email
			myUser.Mobile=mobile
			myUser.User_group=user_group
			myUser.Status=1
			userInfo := &myUser
			_, err := userInfo.AddUser()
			if err == nil {
				c.Data["code"] = 1
				c.Data["msg"] = "添加成功！"
				c.Data["url"] = "/admin/User/index"
				c.Data["wait"] = 1
				c.TplName = "admin/public/redirect.tpl"
			} else {
				c.Data["code"] = 0
				c.Data["msg"] = "添加失败！"
				c.Data["url"] = "/admin/User/add"
				c.Data["wait"] = 3
				c.TplName = "admin/public/redirect.tpl"
			}
			return
		}
	}
}

func (c *UserController)Del(){
	user_id,_ := c.GetInt(":uid")
	user:=&models.User{User_id:user_id}
	_,err:=user.DelUser()
	if err==nil{
		c.Data["code"] = 1
		c.Data["msg"] = "删除成功！"
		c.Data["url"] = "/admin/User/index"
		c.Data["wait"] = 1
		c.TplName = "admin/public/redirect.tpl"
	}else{
		c.Data["code"] = 0
		c.Data["msg"] = "删除失败！"
		c.Data["url"] = "/admin/User/index"
		c.Data["wait"] = 3
		c.TplName = "admin/public/redirect.tpl"
	}
}