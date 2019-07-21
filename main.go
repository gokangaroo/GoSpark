package main

import (
	"GoSpark/controllers"
	"GoSpark/helper"
	_ "GoSpark/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

//入口初始化函数
func init() {
	//获取全局panic
	defer func() {
		if err := recover(); err != nil {
			beego.Error("Panic error:", err)
		}
	}()

	func(functions ...func() error) {
		for _, f := range functions {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(
		//初始化日志
		helper.InitLogs,
		//初始化db
		helper.InitDatabase,
		//初始化redis
		helper.InitRedis,
	)
}
