package models

//用户概况表
type Profile struct {
	Id       int    `orm:"column(id);auto"`
	Position string `orm:"column(position);size(16);default();"` //职位
	Gender   string `orm:"column(gender);size(16);default();"`   //性别
	User     *User  `orm:"reverse(one)"`              //设置一对一反向关系(可选)
}

func (t *Profile) TableName() string {
	return "profiles"
}
