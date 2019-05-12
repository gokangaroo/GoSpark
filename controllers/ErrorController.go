package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}
