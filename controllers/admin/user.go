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
	users,err:=user.GetAllUsers()
	if err==nil{
		c.Data["user_infos"]=users
		c.TplName = "admin/user/index.tpl"
		return
	}else{
		c.redirect("/admin/Index","未获取到用户信息！",0,3)
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
			c.redirect("/admin/User/index","未获取到id！",0,3)
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
				c.redirect("/admin/User/edit/uid/" + uid,"初始管理员禁止修改密码！",0,3)
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
					c.redirect("/admin/User/index","修改成功！",1,1)
				} else {
					c.redirect("/admin/User/edit/uid/" + uid,"修改失败！",0,3)
				}
				return
			}else{
			c.redirect("/admin/User/edit/uid/" + uid,"未获取到用户信息！",0,3)
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
			c.redirect("/admin/User/add","已经存在该用户名的用户，禁止重复！！",0,3)
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
				c.redirect("/admin/User/index","添加成功！",1,1)
			} else {
				c.redirect("/admin/User/add","添加失败！",0,3)
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
		c.redirect("/admin/User/index","删除成功！",1,1)
	}else{
		c.redirect("/admin/User/index","删除失败！",0,3)
	}
	return
}