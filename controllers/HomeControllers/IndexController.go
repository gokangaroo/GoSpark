package HomeControllers

type IndexController struct{
	BaseController
}

func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//中英文转换
	c.Data["welcome"] = c.Tr("welcome")
	c.Data["description"] = c.Tr("description")
	c.Data["site"] = c.Tr("site")
	c.Data["contact"] = c.Tr("contact")
}