package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/controller/api"
)

func SetRouter(app *iris.Application)  {
	mvc.Configure(app.Party("/api"),UserMvc) //注册UserMvc
}


//use mvc
func UserMvc(app *mvc.Application)  {
	app.Handle(api.NewUser())
}

