package helper

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
)

//日志变量
var Logger = logs.NewLogger()

//初始化日志
func InitLogs() {
	//创建日志目录
	if _, err := os.Stat("logs"); err != nil {
		os.Mkdir("logs", os.ModePerm)
	}
	var level = 7
	if Debug {
		level = 4
	}
	maxLines := GetConfigInt64("logs", "max_lines")
	if maxLines <= 0 {
		maxLines = 10000
	}
	maxDays := GetConfigInt64("logs", "max_days")
	if maxDays <= 0 {
		maxDays = 7
	}
	//初始化日志的各种配置
	LogsConf := fmt.Sprintf(`{"filename":"logs/gospark.log","level":%v,"maxlines":%v,"maxsize":0,"daily":true,"maxdays":%v}`, level, maxLines, maxDays)
	Logger.SetLogger(logs.AdapterFile, LogsConf)
	if Debug {
		logs.SetLogger(logs.AdapterConsole)
		beego.Info("日志配置信息: " + LogsConf)
	} else {
		//设置异步写日志
		Logger.Async(1e3)
	}
	//显示文件和行号
	Logger.EnableFuncCallDepth(true)
}
