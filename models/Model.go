package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"strings"
)

//获取带表前缀的数据表
//@param            table               数据表
func getTable(table string) string {
	prefix := beego.AppConfig.DefaultString("db_prefix", "")
	if !strings.HasPrefix(table, prefix) {
		table = prefix + table
	}
	return table
}

//获取指定struct字段
func GetFields(tableObj interface{}) (fields []string) {
	elem := reflect.ValueOf(tableObj).Elem()
	for i := 0; i < elem.NumField(); i++ {
		fields = append(fields, elem.Type().Field(i).Name)
	}
	return fields
}

//左联合查询创建SQL语句
func LeftJoinSqlBuild(tables []string, on []map[string]string, fields map[string][]string, p, listRows int, orderBy []string, groupBy []string, condition string) (sql string, err error) {
	if len(tables) < 2 || len(tables)-1 != len(on) {
		err = errors.New("参数不规范：联合查询的数据表数量必须在2个或2个以上，同时表数量比on条件多一个")
		return
	}
	var (
		FieldSlice   []string
		StrOrderBy   string
		StrGroupBy   string
		StrCondition string
		joinKV       string
		join         = []string{tables[0]}
		usedTables   = []string{}
	)
	for table, field := range fields {
		for _, f := range field {
			FieldSlice = append(FieldSlice, strings.Trim(fmt.Sprintf("%v.%v", table, f), "."))
		}
	}
	for index, table := range tables {
		slice := strings.Split(strings.TrimSpace(table), " ")
		if len(slice) == 1 {
			slice = append(slice, slice[0])
		}
		usedTables = append(usedTables, slice[1])
		if index > 0 {
			on, joinKV = joinOn(slice[1], usedTables, on)
			join = append(join, "left join "+table+" on "+joinKV)
		}
	}
	if len(orderBy) > 0 {
		StrOrderBy = " order by " + strings.Join(orderBy, ",")
	}
	if len(condition) > 0 {
		StrCondition = " where " + condition
	}
	if len(groupBy) > 0 {
		StrGroupBy = " group by " + strings.Join(groupBy, ",")
	}

	sql = fmt.Sprintf("select %v from %v %v %v %v limit %v offset %v", strings.Join(FieldSlice, ","), strings.Join(join, " "), StrCondition, StrGroupBy, StrOrderBy, listRows, (p-1)*listRows)
	return
}

//只供LeftJoinSqlBuild创建SQL语句使用
func joinOn(table string, usedTables []string, on []map[string]string) (newon []map[string]string, ret string) {
	table = table + "."
	lenon := len(on)
	for index, v := range on {
		for key, val := range v {
			if strings.HasPrefix(key, table) || strings.HasPrefix(val, table) {
				for _, used := range usedTables {
					if strings.HasPrefix(key, used) || strings.HasPrefix(val, used) {
						ret = key + "=" + val
						if index > 0 {
							newon = append(newon, on[0:index]...)
						}
						if index+1 <= lenon {
							newon = append(newon, on[(index+1):]...)
						}
						return
					}
				}
			}
		}
	}
	return
}
