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
			return true,*user
		}
	}
	return false,*user
}

func (user *User) GetUserById() (User){
	o := orm.NewOrm()
	o.Using("default")
	// read
	err := o.Read(user)
	if err!=nil{
		return *user
	}else{
		return *user
	}
}


