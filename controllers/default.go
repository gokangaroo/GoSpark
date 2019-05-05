package controllers

import (
	. "GoSpark/controllers/HomeControllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"strings"
)

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

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//中英文转换
	c.Data["welcome"] = c.Tr("welcome")
	c.Data["descripition"] = c.Tr("descripition")
	c.Data["site"] = c.Tr("site")
	c.Data["contact"] = c.Tr("contact")
}
