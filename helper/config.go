package helper

import (
	"fmt"
	"strconv"
)

//获取配置
func GetConfig(cate string, key string, def ...string) string {
	if val, ok := ConfigMap.Load(fmt.Sprintf("%v.%v", cate, key)); ok {
		return val.(string)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

//获取配置 int64
func GetConfigInt64(cate string, key string) (val int64) {
	val, _ = strconv.ParseInt(GetConfig(cate, key), 10, 64)
	return
}

//获取配置 float64
func GetConfigFloat64(cate string, key string) (val float64) {
	val, _ = strconv.ParseFloat(GetConfig(cate, key), 64)
	return
}
