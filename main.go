package main

import (
	"github.com/kataras/iris"
	"ims/models"
	"ims/router"
	"ims/socket"
)


func Cors(ctx iris.Context) {
	origin := ctx.GetHeader("Origin")
	allowHeaders := "X-Requested-With, Access-Control-Allow-Origin, X-HTTP-Method-Override, Content-Type, Authorization, Accept"
	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Vary", "origin")
	ctx.Header("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", allowHeaders)
	ctx.Next()
}

func main()  {
	app := iris.New()

	//全局api 跨域
	app.Use(Cors)
	//控制台日志
	//app.Use(ConseLog)

	// 设置日志级别，开发阶段为 debug
	app.Logger().SetLevel("debug")

	//注册路由
	router.SetRouter(app)

	//websocket注册
	socket.InitWsSocket(app)

	//socket.io
	socket.InitSocketIo(app)

	//init事件注册
	initEvent()

	app.Run(iris.Addr(":9090"),
		iris.WithoutPathCorrection,
		iris.WithoutServerError(iris.ErrServerClosed))
}


func initEvent()  {
	//orm 数据库注册
	models.InitDbConn()
}

