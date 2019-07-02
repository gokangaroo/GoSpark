package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type User struct {
	Id       int       `orm:"column(id);auto"`
	Name     string    `orm:"column(name);size(32);unique"`                   //用户名
	Username string    `orm:"column(username);size(64);unique"`               //用户名
	Email    string    `orm:"column(email);size(255);unique;default();"`      //邮箱
	Phone    string    `orm:"column(phone);size(64);default()"`               //邮箱
	Avatar   string    `orm:"column(avatar);size(255);unique;default();"`     //头像
	Password string    `orm:"column(password);size(64)"`                      //密码
	Intro    string    `orm:"column(intro);size(255);default()"`              //个性签名
	Post     []*Post   `orm:"reverse(many);on_delete(set_null)"`              //一对多的反向关系 (用户-文章)
	Profile  *Profile  `orm:"rel(one);"`                                      // 一对一(用户-概况)
	IsActive int       `orm:"column(is_active);default(0)"`                   //是否激活
	Created  time.Time `orm:"column(created_at);auto_now_add;type(datetime)"` //创建时间
	Updated  time.Time `orm:"column(updated_at);auto_now;type(datetime)"`     //更新时间
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
