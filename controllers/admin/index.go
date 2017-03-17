package controllers


type IndexController struct {
	BaseController
}

type User_info struct{
	User_id  int
	User_type int
	User_name string
	Real_name string
	Mobile string
	Email string
	Login_time int64
	Login_ip string
}

func (c *IndexController) Get() {
	//c.Data["ss"]=c.GetSession("uid")
	//menuurl:=&models.MenuUrl{Menu_id:2}
	//result,err:=menuurl.GetMenuUrlById()
}

func (c *IndexController)Index(){
	c.TplName = "admin/index/index.tpl"
}