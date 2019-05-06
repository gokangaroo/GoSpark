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
}
//后台路由
func back()  {
	beego.Router("/admin",&AdminControllers.IndexController{})
}