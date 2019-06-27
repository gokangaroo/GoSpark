package HomeControllers

type AuthController struct {
	BaseController
}

func (c *AuthController) Login(){
	if c.IsLogin > 0{
		c.Redirect("/user",302)
		return
	}

	//get请求
	if c.Ctx.Request.Method == "GET"{
		c.Data["isUser"]  = true
		c.Layout = "layouts/app.html"
		c.TplName = "auth/login.html"
	}
}

func (c *AuthController) Register(){
	//get请求
	if c.Ctx.Request.Method == "GET"{
		c.Data["isUser"]  = true
		c.Layout = "layouts/app.html"
		c.TplName = "auth/register.html"
		return
	}
}
func (c *AuthController) PasswordReset(){

}