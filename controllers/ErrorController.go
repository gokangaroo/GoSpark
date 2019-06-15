package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

func (c *ErrorController) Error501()  {
	referer := c.Ctx.Request.Referer()
	c.Layout  = ""
	c.Data["content"] = "Server Error"
	if len(referer)>0{
		c.Data["isReferer"] = true
	}
	c.TplName = "501.html"
}

func (c *ErrorController) ErrorDb() {
	c.Layout = ""
	c.Data["content"] = "Database is something error"
	c.TplName = "error.html"
}
