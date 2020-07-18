package api

import (
	"encoding/json"
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

//中间件
func (this *Base)AnyMiddlewareHere(ctx iris.Context)  {
	if this.validate(ctx){
		ctx.Next()
	}else {
		errJson ,_ :=json.Marshal(lib.ErrMsg("token错误"))
		ctx.WriteString(string(errJson))
	}
}