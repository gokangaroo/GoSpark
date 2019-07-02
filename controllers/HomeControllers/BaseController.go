package HomeControllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"html/template"
	"strings"
)

type BaseController struct {
	beego.Controller
	TplTheme  string //模板主题
	TplStatic string //静态文件
	IsLogin   int    // 用户是否登录
	i18n.Locale      // 国际化
}

func (c *BaseController) Prepare() {
	// Reset language option.
	c.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	// 这个根据用户浏览器的语言设置走.
	al := c.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		logs.Trace("Browser's setup of language is : " + al)
		if i18n.IsExist(al) {
			c.Lang = al
		}
	}

	// 2. Default language is America English.
	if len(c.Lang) == 0 {
		c.Lang = "en-US"
	}

	// 3.模板中显示语言，不通过控制器
	c.Data["Lang"] = c.Lang

	//当前模板静态文件
	c.Data["TplStatic"] = "/static/"
}

//是否已经登录，如果已经登录 则返回用户的id
func (c *BaseController) CheckLogin() int {
	uid := c.GetSession("uid")
	if uid != nil {
		id, ok := uid.(int)
		if ok && id > 0 {
			return id
		}
	}
	return 0
}

//防止跨站攻击 在有表单控制器使用  不需要直接在base控制器使用
func (c *BaseController) Xsrf() {
	//使用时 直接在表单添加{{.xsrfdata}}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
}

//重置cookie
func (c *BaseController) ResetCookie() {
	c.Ctx.SetCookie("uid", "")
	c.Ctx.SetCookie("token", "")
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

//响应json
func (c *BaseController) ResponseJson(isSuccess bool, msg string, data ...interface{}) {
	status := 0
	if isSuccess {
		status = 1
	}
	ret := map[string]interface{}{"status": status, "message": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	c.Data["json"] = ret
	c.ServeJSON()
	c.StopRun()
}
