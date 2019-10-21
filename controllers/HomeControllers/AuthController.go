package HomeControllers

import (
	"GoSpark/helper"
	"GoSpark/models"
	"html/template"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare() {
	c.EnableXSRF = false
}

func (c *AuthController) Login() {
	if c.IsLogin > 0 {
		c.Redirect("/user", 302)
		return
	}

	//get请求
	if c.Ctx.Request.Method == "GET" {
		c.Data["isUser"] = true
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.Layout = "layouts/app.html"
		c.TplName = "auth/login.html"
	}
}

//用户注册
func (c *AuthController) Register() {
	if c.IsLogin > 0 {
		c.Redirect("/user", 302)
		return
	}

	//get请求
	if c.Ctx.Request.Method == "GET" {
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.Data["isUser"] = true
		c.Layout = "layouts/app.html"
		c.TplName = "auth/register.html"
		return
	}

	// 验证注册信息
	username := c.GetString("username")
	email := c.GetString("email")

	err, uid := models.NewUser().CreateUser(
		username,
		email,
		c.GetString("password"),
		c.GetString("password_confirmation"))
	if err != nil || uid == 0 {
		if err != nil {
			helper.Logger.Error(err.Error())
			c.ResponseJson(false, err.Error())
		}
		c.ResponseJson(false, "注册失败")
	}
	c.IsLogin = uid
	c.SetCookieLogin(uid)
	c.ResponseJson(true, "用户注册成功")
}

//重置密码
func (c *AuthController) PasswordReset() {
	//get请求
	if c.Ctx.Request.Method == "GET" {
		c.Data["isUser"] = true
		c.Layout = "layouts/app.html"
		c.TplName = "auth/password_reset.html"
		return
	}
}
