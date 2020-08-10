package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/datamodels"
	"ims/lib"
	"ims/service"
)

type Video struct {
	Base *Base
	Ctx  iris.Context
}

func NewVideo() *Video {
	return &Video{Base: NewBase()}
}

//前置操作 中间件
func (this *Video) BeforeActivation(b mvc.BeforeActivation) {
	//b.Handle("POST","/getUserInfo","GetUserInfo",anyMiddlewareHere)
	//b.Handle("POST","/login","Login") //登录操作
	b.Handle("POST", "/video/getIds", "GetVideoIds")
	b.Handle("POST", "/video/getInfo", "GetVideoInfo")
}

func (this *Video) GetVideoIds() interface{} {
	//获取所有的ids
	ids := service.GetVideoIds()
	return lib.SuccessData(ids)
}

func (this *Video) GetVideoInfo() interface{} {
	id := datamodels.GetVideoInfoData{}
	_ = this.Ctx.ReadJSON(&id)
	//获取id的详细信息
	result, err := service.GetVideoInfoById(id.Id)
	if err != nil {
		return lib.ErrMsg(err.Error())
	}
	return lib.SuccessData(result)
}
