package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type User struct {
	Id       int       `orm:"column(id);auto"`
	Email    string    `orm:"size(50);unique;column(email);default();"`       //邮箱
	Avatar   string    `orm:"size(255);unique;column(avatar);default();"`     //头像
	Password string    `orm:"size(32);column(password)"`                      //密码
	Username string    `orm:"size(16);unique;column(username)"`               //用户名
	Intro    string    `orm:"size(255);default();column(intro)"`              //个性签名
	Post     []*Post   `orm:"reverse(many);on_delete(set_null)"`              //一对多的反向关系 (用户-文章)
	Profile  *Profile  `orm:"rel(one);on_delete(set_null)"`                   // 一对一(用户-概况)
	IsActive int       `orm:"column(is_active);default(0)"`                   //是否激活
	Created  time.Time `orm:"auto_now_add;column(created_at);type(datetime)"` //创建时间
	Updated  time.Time `orm:"auto_now;column(updated_at);type(datetime)"`     //更新时间
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
