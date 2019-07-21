package models

//用户概况表
type Profile struct {
	Id       int    `orm:"column(id);auto"`
	Position string `column(position);orm:"size(16)"` //职位
	Gender   string `orm:"size(16);column(gender)"`   //性别
	User     *User  `orm:"reverse(one)"`              //设置一对一反向关系(可选)
}

func (t *Profile) TableName() string {
	return "profile"
}
