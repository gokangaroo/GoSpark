package models

import "github.com/astaxie/beego/orm"

//用户概况表
type Profile struct {
	Id       int    `orm:"column(id);auto"`
	Position string `orm:"size(16);column(position)"` //职位
	User     *User  `orm:"reverse(one)"`              //设置一对一反向关系(可选)
}

func (t *Profile) TableName() string {
	return "profile"
}

func init() {
	orm.RegisterModel(new(Profile))
}
