package models

import "github.com/astaxie/beego/orm"

//文章动态表
type Post struct {
	Id    int    `orm:"column(id);auto"`
	Title string `orm:"size(16);unique;column(title)"` //标题
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m);rel_table(post_tag)"`   //设置多对多
}

func (t *Post) TableName() string {
	return "post"
}

func init() {
	orm.RegisterModel(new(Post))
}