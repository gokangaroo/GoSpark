package HomeControllers

import "GoSpark/models"

type UserController struct {
	BaseController
}

func (c *UserController) UserInfo(){
	id := c.Ctx.Input.Param(":id")
	user := models.NewUser().UserInfo(id)
	c.ResponseJson(true, "success",user)
}
