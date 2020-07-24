package main

import (
	"github.com/kataras/iris"
	"ims/logs"
	"ims/models"
	"ims/router"
	"ims/socket"
	"os"
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

	//日志
	Logs := logs.NewLogs()

	f ,err :=os.OpenFile(Logs.FilePath,os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		app.Logger().Info(err)
	}
	app.Logger().SetLevel(Logs.Level)

	app.Logger().SetOutput(f)

	defer  f.Close()

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

