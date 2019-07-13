package helper

import (
	"github.com/astaxie/beego"
	"sync"
)

var (
	//develop mode
	Debug = beego.AppConfig.String("runmode") == "dev"

	//允许直接访问的文件扩展名
	StaticExt = make(map[string]bool)

	//配置文件的全局map
	ConfigMap sync.Map

	//程序是否已经安装
	IsInstalled = false

	//允许上传的文档扩展名
	AllowedUploadExt = ",doc,docx,rtf,wps,odt,ppt,pptx,pps,ppsx,dps,odp,pot,xls,xlsx,et,ods,txt,pdf,chm,epub,umd,mobi,"
)