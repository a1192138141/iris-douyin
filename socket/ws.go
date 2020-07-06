package socket

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
	"ims/lib"
	"ims/models"
	"reflect"
	"ims/utils"
	"net/http"
)

//这是websocket 控制器
var ws = websocket.New(websocket.Config{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
})


//websocket mvc
func configureWs(m *mvc.Application) {
	m.Register(ws.Upgrade)
	m.Handle(new(Ws))
}

//ws基类
type Ws struct {
	Ctx iris.Context
	User *models.User
	Conn websocket.Connection
}

//ws message
type WsMessage struct {
	Function string `json:"function"`
	Param interface{} `json:"param"`
}


//用于注册所有的websocket事件
func (this *Ws) Get()  {
	this.Conn.OnMessage(this.message)
	this.Conn.Wait()
}


//token 检验
func (this *Ws) validateToken(token string) bool{
	res , err :=lib.ParseUserToken(token)
	if !err {
		return  false
	}
	errors  := utils.MapToSturct(res,&this.User)
	if errors != nil {
		 return false
	}

	fmt.Print(this)
	return  true
}

//获取message
func (this *Ws) message(bytes []byte)  {
	//获取所有传输的字节
	var message WsMessage
	utils.BytesToStruct(bytes,&message)
	if message.Function =="" {
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("参数错误"))
	}
	token := this.validateToken(this.Ctx.FormValue("token"))
	if !token {
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("token不正确"))
	}

	//接下来干想干的事情 (直接字符转方法)
	function := message.Function
	//reflect.ValueOf(NewRealize(&message,this)).MethodByName(function).Call(nil)


	defer func() {
		err := recover()
		fmt.Print("==============")
		fmt.Print(err)
		reflect.ValueOf(NewRealize(&message,this)).MethodByName(function).Call(nil)

	}()

}










