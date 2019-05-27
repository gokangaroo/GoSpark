package routers

import (
		"github.com/astaxie/beego"
	"GoSpark/controllers/HomeControllers"
	"GoSpark/controllers/AdminControllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})

    front()
    back()
}

//前台路由
func front(){
	beego.Router("/",&HomeControllers.IndexController{})
	beego.Router("/login",&HomeControllers.AuthController{},"get:Login")
	beego.Router("/login",&HomeControllers.AuthController{},"post:PostLogin")
	beego.Router("/register",&HomeControllers.AuthController{},"post:register")
	beego.Router("/register",&HomeControllers.AuthController{},"post:PostRegister")
}
//后台路由
func back()  {
	beego.Router("/admin",&AdminControllers.IndexController{})
}