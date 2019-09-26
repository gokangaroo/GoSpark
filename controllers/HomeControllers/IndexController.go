package HomeControllers

import (
	"GoSpark/helper"
	"GoSpark/library/message"
	"fmt"
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
