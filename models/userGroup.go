/**
 * Created by sunny.
 * User: sunny
 * always for darling!
 * Date: 2017/3/10
 * Time: 16:04
 */
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type UserGroup struct{
Group_id int `orm:"pk"`
Group_name string
Group_role string
Owner_id int
Group_desc string
}

func init(){
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(UserGroup))
}

func (usergroup *UserGroup)GetUserGroupById() (UserGroup,error){
	o := orm.NewOrm()
	o.Using("default")
	// read
	err := o.Read(usergroup)
	return *usergroup,err
}

func (userGroup *UserGroup)GetALLGroups() ([]UserGroup,error){
	o := orm.NewOrm()
	o.Using("default")
	var groups []UserGroup
	_,err:= o.Raw("SELECT * FROM "+beego.AppConfig.String("mysqldbprefix")+"user_group  ORDER BY group_id ASC").QueryRows(&groups)
	//未查询到数据或者查询出错
	return groups,err
}