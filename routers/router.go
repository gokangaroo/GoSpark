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
	beego.Router("/login", &HomeControllers.AuthController{}, "get:ShowLogin")
	beego.Router("/login", &HomeControllers.AuthController{}, "post:Login")
	beego.Router("/register", &HomeControllers.AuthController{}, "get:ShowRegister")
	beego.Router("/register", &HomeControllers.AuthController{}, "post:Register")
	beego.Router("/password/reset", &HomeControllers.AuthController{}, "get:ShowPasswordReset")
	beego.Router("/password/reset", &HomeControllers.AuthController{}, "post:PasswordReset")
}

//后台路由
func back() {
	beego.Router("/admin", &AdminControllers.IndexController{})
}
