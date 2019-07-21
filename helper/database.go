package helper

import (
	"GoSpark/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitDatabase() error {
	beego.Info("database config...")
	//连接Mysql
	user := beego.AppConfig.String("db_user")
	password := beego.AppConfig.String("db_password")
	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")
	database := beego.AppConfig.String("db_database")
	dbType := beego.AppConfig.String("db_type")

	maxIdleConn, _ := beego.AppConfig.Int("mysql_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql_max_open_conn")

	switch dbType {
	case "mysql":
		dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database) + "&loc=Asia%2FShanghai"
		//注册数据库驱动
		if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
			return err
		}
		//注册数据库连接 默认default
		if err := orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn); err != nil {
			return err
		}
	case "postgresql":
		dbLink := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, database)
		//注册数据库驱动
		if err := orm.RegisterDriver("postgresql", orm.DRPostgres); err != nil {
			return err
		}
		//注册数据库连接 默认default
		if err := orm.RegisterDataBase("default", "postgresql", dbLink, maxIdleConn, maxOpenConn); err != nil {
			return err
		}
	default:
		return fmt.Errorf("please define the db type")
	}

	//统一注册model
	orm.RegisterModel(
		new(models.User),
		new(models.Profile),
		new(models.Post),
		new(models.Tag))

	//终端打印sql
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

	return nil
}
