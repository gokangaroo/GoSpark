package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

//用户表
type User struct {
	Id       int       `orm:"column(id);auto"`
	Username string    `orm:"column(username);size(64);unique"`               //用户名
	Email    string    `orm:"column(email);size(255);unique;default();"`      //邮箱
	Phone    string    `orm:"column(phone);size(64);default()"`               //邮箱
	Avatar   string    `orm:"column(avatar);size(255);default();"`            //头像
	Password string    `orm:"column(password);size(64)"`                      //密码
	Intro    string    `orm:"column(intro);size(255);default()"`              //个性签名
	Post     []*Post   `orm:"reverse(many);on_delete(set_null)"`              //一对多的反向关系 (用户-文章)
	Profile  *Profile  `orm:"rel(one)"`                                       // 一对一(用户-概况)
	IsActive int       `orm:"column(is_active);default(0)"`                   //是否激活
	Created  time.Time `orm:"column(created_at);auto_now_add;type(datetime)"` //创建时间
	Updated  time.Time `orm:"column(updated_at);auto_now;type(datetime)"`     //更新时间
}

func NewUser() *User {
	return &User{}
}

func GetTableUser() string {
	return getTable("users")
}

func (t *User) TableName() string {
	return "users"
}

// 创建用户
func (m *User) CreateUser(username, emil, password string, passwordConfirm string) (error, int) {
	var (
		user User
		o    = orm.NewOrm()
	)
	l := strings.Count(username, "") - 1
	if l < 2 || l > 16 {
		return errors.New("用户名称限制在2-16个字符"), 0
	}
	if o.QueryTable(GetTableUser()).Filter("Username", username).One(&user); user.Id > 0 {
		return errors.New("用户名已被注册"), 0
	}
	if password != passwordConfirm {
		return errors.New("两次密码不一样"), 0
	}

	user = User{Username: username, Email: emil, Password: password}
	// 创建 profile
	profile := new(Profile)
	profile.Gender = "f"

	err := o.Begin()
	user.Profile = profile
	_, err1 := o.Insert(profile)
	_, err2 := o.Insert(&user)
	if err1 != nil || err2 != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err, user.Id
}

// 获取表的字段
func (m *User) Fields() map[string]string {
	var fields map[string]string
	fields = make(map[string]string)
	v := reflect.ValueOf(m).Elem()
	k := v.Type()
	num := v.NumField()
	for i := 0; i < num; i++ {
		key := k.Field(i)
		fields[key.Name] = key.Type.String()
	}
	return fields
}

// 获取用户除密码之外的信息
func (m *User) GetById(id interface{}) (params orm.Params, rows int64, err error){
	var data []orm.Params
	tables := []string{GetTableUser()+"u",GetTableUser()+"ui"}
	on := []map[string]string{
		{"u.Id":"ui.Id"},
	}
	fields := map[string][]string{
		"u":GetFields(NewUser()),
		"ui":GetFields(NewUser()),
	}
	if sql, err := LeftJoinSqlBuild(tables, on, fields, 1, 1, nil, nil, "u.Id=?"); err == nil {
		if rows, err = orm.NewOrm().Raw(sql, id).Values(&data); err == nil && len(data) > 0 {
			params = data[0]
		}
	}
	return
}