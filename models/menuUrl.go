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
	"strings"
	"strconv"
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

type AllMenuUrl struct{
	Module_id int
	Module_name string
	Module_url string
	Module_sort int
	Module_desc string
	Module_icon string
	Online int
	MenuList []MenuUrl
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

func (menuurl *MenuUrl) GetMenuUrlByUrl() (MenuUrl,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	err := o.Read(menuurl,"Menu_url")

	return *menuurl,err
}

func (menuUrl *MenuUrl)GetTrees(oldAccess string) []AllMenuUrl{
	//获取所有module的菜单
	module:=&Module{}
	result,err:=module.GetAllModules(1)
	newAllMenuUrl:=make([]AllMenuUrl,0)
	access:=strings.Split(oldAccess,",")
	if err==nil {
		for _, v := range result {
			newResult, newErr := getListByModuleId(v.Module_id, "demo")
			if (newErr == nil) {
				newMenus:=make([]MenuUrl,0)
				for _,mv:=range newResult {
					for _, newV := range access {
						iv,_:=strconv.Atoi(newV)
						if (mv.Menu_id==iv) {
							newMenus=append(newMenus,mv)
							break
						}else{
							continue
						}
					}
				}
				newModule:=AllMenuUrl{v.Module_id,v.Module_name,v.Module_url,v.Module_sort,v.Module_desc,v.Module_icon,v.Online,newMenus}
				newAllMenuUrl=append(newAllMenuUrl,newModule)
			} else {
				continue
			}
		}
	        return newAllMenuUrl
	}else{
		return newAllMenuUrl
	}
}

func getListByModuleId(id int,flag string)([]MenuUrl,error){
	o := orm.NewOrm()
	o.Using("default")
	var menus []MenuUrl
	if flag=="role"{
		_,err:= o.Raw("SELECT * FROM "+beego.AppConfig.String("mysqldbprefix")+"menu_url  WHERE module_id=? AAND is_show=1 ORDER BY sort DESC",id).QueryRows(&menus)
		return menus,err
	}else{
		_,err:= o.Raw("SELECT * FROM "+beego.AppConfig.String("mysqldbprefix")+"menu_url  WHERE module_id=? AND is_show=1 AND online=1 ORDER BY sort DESC",id).QueryRows(&menus)
		return menus,err
	}
}