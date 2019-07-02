package AdminControllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	TplTheme  string //模板主题
	TplStatic string //模板静态文件
	AdminId   int    //管理员是否已登录，如果已登录，则管理员ID大于0
}
