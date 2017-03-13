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

func (MenuUrl *MenuUrl)GetTrees() []AllMenuUrl{
	//获取所有module的菜单
	module:=&Module{}
	result,err:=module.GetAllModules(1)
	//获取当前用户的菜单
	//var access=[]int{1,5,108,112,113,122,109,114,115,123}
	var newAllMenuUrl []AllMenuUrl
	//var k=0
	if err==nil {
		for _, v := range result {
			_, newerr := getListByModuleId(v.Module_id, "demo")
			if (newerr == nil) {
				newAllMenuUrl[0].Module_id=1
				return newAllMenuUrl
				//var newAllMenu []MenuUrl
				//var newk=0
				//for _,mv:=range newresult {
				//	for _, v := range access {
				//		if (mv.Menu_id==v) {
				//			newAllMenu[newk].Menu_id=mv.Menu_id
				//			newAllMenu[newk].Menu_name=mv.Menu_name
				//			newAllMenu[newk].Menu_url=mv.Menu_url
				//			newAllMenu[newk].Module_id=mv.Module_id
				//			newAllMenu[newk].Is_show=mv.Is_show
				//			newAllMenu[newk].Online=mv.Online
				//			newAllMenu[newk].Shortcut_allowed=mv.Shortcut_allowed
				//			newAllMenu[newk].Sort=mv.Sort
				//			newAllMenu[newk].Menu_desc=mv.Menu_desc
				//			newAllMenu[newk].Father_menu=mv.Father_menu
				//			newk++
				//			break
				//		}else{
				//			continue
				//		}
				//	}
				//}
				//newAllMenuUrl[k].Module_id=v.Module_id
				//newAllMenuUrl[k].Module_name=v.Module_name
				//newAllMenuUrl[k].Module_url=v.Module_url
				//newAllMenuUrl[k].Module_sort=v.Module_sort
				//newAllMenuUrl[k].Module_desc=v.Module_desc
				//newAllMenuUrl[k].Module_icon=v.Module_icon
				//newAllMenuUrl[k].Online=v.Online
				//newAllMenuUrl[k].MenuList=newresult
				//k++
			} else {
				continue
			}
		}
	        return nil
	}else{
		return nil
	}
}

//func (menuUrl *MenuUrl)GetMenusByRole(roles string) ([]MenuUrl,error){
//	o := orm.NewOrm()
//	o.Using("default")
//	var menus []MenuUrl
//	_,err:= o.Raw("SELECT * FROM "+beego.AppConfig.String("mysqldbprefix")+"menu_url  WHERE online=1 AND menu_id in (?)",roles).QueryRows(&menus)
//	return menus,err
//}

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