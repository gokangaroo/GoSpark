package routers

import (
	"github.com/astaxie/beego"
	"GoSpark/controllers/HomeControllers"
	"GoSpark/controllers/AdminControllers"
)

func init() {

	front()
	back()
}

//前台路由
func front() {
	beego.Router("/", &HomeControllers.IndexController{})
	beego.Router("/test", &HomeControllers.IndexController{}, "get:Test")
	beego.Router("/login", &HomeControllers.AuthController{}, "*:Login")
	beego.Router("/register", &HomeControllers.AuthController{}, "*:Register")
	beego.Router("/password/reset", &HomeControllers.AuthController{}, "*:PasswordReset")
}

//后台路由
func back() {
	beego.Router("/admin", &AdminControllers.IndexController{})
}
