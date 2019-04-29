package HomeControllers

import "github.com/astaxie/beego"

type BaseController struct{
	beego.Controller
	TplTheme string //模板主题
	TplStatic string //静态文件
	IsLogin int // 用户是否登录
}
