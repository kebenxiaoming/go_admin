package controllers


type LogoutController struct {
	BaseController
}
//退出
func (c *LogoutController) Get() {
	c.DestroySession()
	c.redirect("/admin/Login/index","已经退出！",1,1)
	return
}


