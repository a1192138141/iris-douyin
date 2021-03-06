package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/controller/api"
)

func SetRouter(app *iris.Application) {
	mvc.Configure(app.Party("/api"), UserMvc) //注册UserMvc
	mvc.Configure(app.Party("/api"), VideoMvc)
	mvc.Configure(app.Party("/api"), FriendMvc)
	mvc.Configure(app.Party("/api"), SearchMvc)
}

//use mvc
func UserMvc(app *mvc.Application) {
	app.Handle(api.NewUser())
}

func FriendMvc(app *mvc.Application) {
	app.Handle(api.NewFriend())
}

//video mvc
func VideoMvc(app *mvc.Application) {
	app.Handle(api.NewVideo())
}

func SearchMvc(app *mvc.Application) {
	app.Handle(api.NewSearch())
}
