package HomeControllers

import "html/template"

type AuthController struct {
	BaseController
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
