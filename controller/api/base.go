package api

import (
	"github.com/kataras/iris"
	"ims/lib"
	"ims/models"
	"ims/utils"
)

type Base struct {
	User *models.User
}

func NewBase() *Base  {
	return  &Base{User:&models.User{}}
}


func (this *Base) validate(ctx iris.Context) bool {
	//获取token
	token :=ctx.GetHeader("token")
	if token == "" {
		return  false
	}

	res , err :=lib.ParseUserToken(token)

	if err == false {
		return  false
	}

	userModel := models.User{}
	errors :=utils.MapToSturct(res,&userModel)
	if errors != nil {
		return  false
	}
	this.User = &userModel
	return  true
}