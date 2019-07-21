package models

//标签表
type Tag struct {
	Id    int     `orm:"column(id)"`
	Name  string  `orm:"size(16);unique;column(name)"` //标签名
	Posts []*Post `orm:"reverse(many)"`
}

func (t *Tag) TableName() string {
	return "tag"
}