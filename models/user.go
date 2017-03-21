/**
 * Created by sunny.
 * User: sunny
 * always for darling!
 * Date: 2017/3/10
 * Time: 16:01
 */
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"crypto/md5"
	"io"
	"encoding/hex"
	"time"
	"strconv"
)

type User struct{
User_id int `orm:"pk"`
User_name string
Password string
Real_name string
Mobile string
Email string
User_desc string
Login_time int
Status int
Login_ip string
User_group int
Template string
Shortcuts string
Show_quicknote int
}

func init(){
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(User))
}

func (user *User) Login(username string,password string) (bool,User){
	o := orm.NewOrm()
	o.Using("default")
	// read
	err := o.Read(user,"User_name")
	if err!=nil{
		return false,*user
	}else{
		m := md5.New()
		io.WriteString(m, password)
		keyMd5 := hex.EncodeToString(m.Sum(nil))
		if user.Password!=keyMd5{
			return false,*user
		}else{
			//登陆成功更新登录信息
			user.Login_time,_= strconv.Atoi(strconv.FormatInt(time.Now().Unix(),10))
			_,err := o.Update(user)
			if err!=nil{
				return false,*user
			}
			return true,*user
		}
	}
	return false,*user
}

func (user *User)GetAllUsers()([]User,error){
	o := orm.NewOrm()
	o.Using("default")
	var users []User
	_,err:= o.Raw("SELECT * FROM "+beego.AppConfig.String("mysqldbprefix")+"user  WHERE status = 1").QueryRows(&users)
	//未查询到数据或者查询出错
	return users,err
}

func (user *User) GetUserById() (User,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	err := o.Read(user)
	return *user,err
}

func (user *User) UpdateUser() (User,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	_,err := o.Update(user)
	return *user,err
}

func (user *User) AddUser() (User,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	_,err := o.Insert(user)
	return *user,err
}

func (user *User) DelUser() (User,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	_,err := o.Delete(user)
	return *user,err
}



