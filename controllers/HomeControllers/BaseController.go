package HomeControllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/astaxie/beego/logs"
	"strings"
	"html/template"
)

type BaseController struct {
	beego.Controller
	TplTheme  string //模板主题
	TplStatic string //静态文件
	IsLogin   int    // 用户是否登录
	i18n.Locale      // 国际化
}

func (this *BaseController) Prepare() {
	// Reset language option.
	this.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	// 这个根据用户浏览器的语言设置走.
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		logs.Trace("Browser's setup of language is : " + al)
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	// 2. Default language is America English.
	if len(this.Lang) == 0 {
		this.Lang = "en-US"
	}

	// 3.模板中显示语言，不通过控制器
	this.Data["Lang"] = this.Lang

	//当前模板静态文件
	this.Data["TplStatic"] = "/static/"
}

//是否已经登录，如果已经登录 则返回用户的id
func (this *BaseController) CheckLogin() int {
	uid := this.GetSession("uid")
	if uid != nil{
		id,ok := uid.(int)
		if ok && id>0{
			return id
		}
	}
	return 0
}

//防止跨站攻击 在有表单控制器使用  不需要直接在base控制器使用
func (this *BaseController) Xsrf(){
	//使用时 直接在表单添加{{.xsrfdata}}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
}

//重置cookie
func (this *BaseController) ResetCookie(){
	this.Ctx.SetCookie("uid","")
	this.Ctx.SetCookie("token","")
}

var LangTypes []string // Languages that are supported.
func init() {
	// Initialize language type list.
	LangTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range LangTypes {
		logs.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file:", err)
			return
		}
	}
}
