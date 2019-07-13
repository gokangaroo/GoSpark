package main

import (
	"GoSpark/controllers"
	"GoSpark/helper"
	_ "GoSpark/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

//入口初始化函数
func init() {
	beego.Info("开始启动")

	//初始化日志
	helper.InitLogs()

	//连接Mysql
	user := beego.AppConfig.String("db_user")
	password := beego.AppConfig.String("db_password")
	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")
	database := beego.AppConfig.String("db_database")

	maxIdleConn, _ := beego.AppConfig.Int("mysql_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql_max_open_conn")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database) + "&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	//orm.RunSyncdb("default",true,true)

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
