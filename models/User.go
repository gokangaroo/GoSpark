package models

import "github.com/astaxie/beego/orm"

//用户表
type User struct {
	Id       int      `orm:"column(id);auto"`
	Email    string   `orm:"size(50);unique;column(email);default();"` //邮箱
	Password string   `orm:"size(32);column(password)"`                //密码
	Username string   `orm:"size(16);unique;column(username)"`         //用户名
	Intro    string   `orm:"size(255);default();column(intro)"`        //个性签名
	Post     []*Post  `orm:"reverse(many)"` //一对多的反向关系 (用户-文章)
	Profile  *Profile `orm:"rel(one)"`      // 一对一(用户-概况)
}

func getUser() *User {
	return &User{}
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}
