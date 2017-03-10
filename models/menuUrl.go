/**
 * Created by sunny.
 * User: sunny
 * always for darling!
 * Date: 2017/3/10
 * Time: 15:52
 */
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type MenuUrl struct{
Menu_id   int `orm:"pk"`
Menu_name string
Menu_url  string
Module_id int
Is_show int
Online int
Shortcut_allowed int
Sort int
Menu_desc string
Father_menu int
}

func init(){
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(MenuUrl))
}

func (menuurl *MenuUrl) GetMenuUrlById() (MenuUrl,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	err := o.Read(menuurl)

	return *menuurl,err
}