package HomeControllers

import (
	"GoSpark/helper"
	"GoSpark/library/message"
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
}

func (c *IndexController) Producer() {
	val,err :=  helper.Redis.Get("ghc").Result()
	if err != nil{
		helper.Logger.Error("something error.....")
	}
	fmt.Println(val)
	//kafka := message.KafkaServer{}
	//if error := kafka.ProducerSync(); error == nil {
	//	helper.Logger.Error("success")
	//}
	c.Ctx.WriteString("producer")
}

func (c *IndexController) Consumer()  {
	kafka := message.KafkaServer{}
	if error := kafka.Consumer();error == nil{
		helper.Logger.Error("success")
	}
	c.Ctx.WriteString("consumer")

}
