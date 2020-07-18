package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/controller/api"
)

func SetRouter(app *iris.Application)  {
	mvc.Configure(app.Party("/api"),UserMvc) //注册UserMvc
	mvc.Configure(app.Party("/api"),VideoMvc)

}


//use mvc
func UserMvc(app *mvc.Application)  {
	app.Handle(api.NewUser())
}

//video mvc
func VideoMvc(app *mvc.Application)  {
	app.Handle(api.NewVideo())
}

