package routers

import (
	"GoSpark/controllers/AdminControllers"
	"GoSpark/controllers/HomeControllers"
	"github.com/astaxie/beego"
)

func init() {
	front()
	back()
}

//前台路由
func front() {
	beego.Router("/", &HomeControllers.IndexController{})
	beego.Router("/test", &HomeControllers.IndexController{}, "get:Test")
	beego.Router("/producer", &HomeControllers.IndexController{}, "get:Producer")
	beego.Router("/consumer", &HomeControllers.IndexController{}, "get:Consumer")
	beego.Router("/login", &HomeControllers.AuthController{}, "*:Login")
	beego.Router("/register", &HomeControllers.AuthController{}, "*:Register")
	beego.Router("/password/reset", &HomeControllers.AuthController{}, "*:PasswordReset")

	beego.Router("/user/:id",&HomeControllers.UserController{},"get:UserInfo")
}

//后台路由
func back() {
	beego.Router("/admin", &AdminControllers.IndexController{})
}
