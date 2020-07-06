package api

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/datamodels"
	"ims/lib"
	"ims/service"
)

type User struct {
	Ctx iris.Context
	Base *Base
}

//使用 new来重写构造函数
func NewUser() *User  {
	return &User{Base:NewBase()}
}

//前置操作 中间件
func (this *User) BeforeActivation(b mvc.BeforeActivation)  {
	anyMiddlewareHere := func(ctx iris.Context) {
		if this.Base.validate(ctx){
			ctx.Next()
		}else {
			errJson ,_ :=json.Marshal(lib.ErrMsg("token错误"))
			ctx.WriteString(string(errJson))
		}
	}
	b.Handle("POST","/getUserInfo","GetUserInfo",anyMiddlewareHere)
	b.Handle("POST","/login","Login") //登录操作
}

//登录
func (this *User) Login() interface{}  {
	loginData := &datamodels.UserLoginData{}
	_ =this.Ctx.ReadJSON(&loginData)
	//进行登录操作
	token , err :=service.GetUserInfoByPhone(loginData.Phone,loginData.Password)
	if err != nil {
		return 	lib.ErrMsg(err.Error())
	}
	return  lib.SuccessData(token)
}


func (this *User) GetUserInfo() interface{} {
	return lib.SuccessData(this.Base.User)
}





