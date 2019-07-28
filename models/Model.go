package models

import (
	"github.com/astaxie/beego"
	"strings"
)

//获取带表前缀的数据表
//@param            table               数据表
func getTable(table string) string {
	prefix := beego.AppConfig.DefaultString("db::prefix", "gs_")
	if !strings.HasPrefix(table, prefix) {
		table = prefix + table
	}
	return table
}