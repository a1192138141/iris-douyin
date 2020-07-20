package api

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type Friend struct {
	Ctx iris.Context
	Base *Base
}
//使用 new来重写构造函数
func NewFriend() *Friend  {
	return &Friend{Base:NewBase()}
}


//前置操作 中间件
func (this *Friend) BeforeActivation(b mvc.BeforeActivation)  {
	//b.Handle("POST","/","GetUserInfo",this.Base.AnyMiddlewareHere)
	b.Handle("POST","/friend/getVideoInfo","GetFriendVideo",this.Base.AnyMiddlewareHere)
}

func (this *Friend) GetFriendVideo() interface{}  {



	fmt.Print("=======")

	return  nil
}
