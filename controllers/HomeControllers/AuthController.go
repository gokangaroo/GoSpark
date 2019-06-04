package HomeControllers

import (
	"GoSpark/models"
	"github.com/astaxie/beego/orm"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Login(){
	
}

func (c *AuthController) ShowLogin(){
	user := models.User{Id:1}
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil{
		panic(err)
	}else{
		c.Ctx.WriteString(user.Username)
	}


}

func (c *AuthController) Register(){

}

func (c *AuthController) ShowRegister(){

}

func (c *AuthController) PasswordReset(){

}

func (c *AuthController) ShowPasswordReset(){

}