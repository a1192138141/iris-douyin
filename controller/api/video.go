package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/lib"
	"ims/service"
)

type Video struct {
	Base *Base
	Ctx iris.Context
}

func NewVideo() *Video {
	return &Video{Base:NewBase()}
}

//前置操作 中间件
func (this *Video) BeforeActivation(b mvc.BeforeActivation)  {
	//b.Handle("POST","/getUserInfo","GetUserInfo",anyMiddlewareHere)
	//b.Handle("POST","/login","Login") //登录操作
	b.Handle("POST","/video/getIds","GetVideoIds")
}

func (this *Video) GetVideoIds() interface{} {
	//获取所有的ids
	ids :=service.GetVideoIds()
	return  lib.SuccessData(ids)
}


