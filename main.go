package main

import (
	_ "GoSpark/routers"
	"github.com/astaxie/beego"
	"GoSpark/controllers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

//入口初始化函数
func init()  {

}
