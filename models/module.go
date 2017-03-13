/**
 * Created by sunny.
 * User: sunny
 * always for darling!
 * Date: 2017/3/10
 * Time: 15:56
 */
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Module struct{
Module_id int `orm:"pk"`
Module_name string
Module_url string
Module_sort int
Module_desc string
Module_icon string
Online int
}

func init(){
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(Module))
}

func(module *Module)GetAllModules(online int) ([]Module,error){
	o := orm.NewOrm()
	o.Using("default")

	var modules []Module
	_,err:= o.Raw("SELECT * FROM "+beego.AppConfig.String("mysqldbprefix")+"module  WHERE online = ? ORDER BY  module_sort ASC,module_id ASC", online).QueryRows(&modules)
	//未查询到数据或者查询出错
	return modules,err
}

