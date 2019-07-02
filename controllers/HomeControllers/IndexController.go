package HomeControllers

import (
	"GoSpark/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.TplName = "home/index.html"
	c.Layout = "layouts/app.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Styles"] = "home/_styles.html"
	c.LayoutSections["Scripts"] = "home/_scripts.html"
}

func (c *IndexController) Test() {
	o := orm.NewOrm()
	o.Using("default")

	profile := new(models.Profile)
	profile.Position = "Java开发"

	user := new(models.User)
	user.Username = "geekghc"
	user.Email = "2438462863@qq.com"
	user.Profile = profile
	user.Phone = "13151568306"

	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	c.Ctx.WriteString("success")

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//中英文转换
	//c.Data["welcome"] = c.Tr("welcome")
	//c.Data["description"] = c.Tr("description")
	//c.Data["site"] = c.Tr("site")
	//c.Data["contact"] = c.Tr("contact")
}
