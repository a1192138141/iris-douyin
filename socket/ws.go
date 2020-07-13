package socket

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
	"ims/lib"
	"ims/models"
	"ims/service"
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
	m.Register(ws.Upgrade).Handle(NewWs())
}


//ws基类
type Ws struct {
	Ctx iris.Context
	User *models.User
	Conn websocket.Connection
	ExitChan chan string
	MsgErrChan chan string
}

func NewWs() *Ws  {
	return &Ws{ExitChan:make(chan string),MsgErrChan:make(chan string),User:&models.User{}}
}

//ws message
type WsMessage struct {
	Function string `json:"function"`
	Param interface{} `json:"param"`
}



//用于注册所有的websocket事件
func (this *Ws) Get()  {
	this.Conn.OnMessage(this.message)
	//可以注册一个协程 用于退出
	go this.ExitWs()
	//map注册 思路反射所有的类和方法
	//go this.initFunc()
	this.Conn.Wait()
}






func (this *Ws) ExitWs()  {
	for{
		fmt.Print(this.ExitChan)
		//fmt.Print("============")
		select {
			case  exit := <-this.ExitChan:{
				fmt.Print("=====exie====",exit)
			   err :=	this.Conn.Disconnect()
			   fmt.Print(err)
			}

		}
	}
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
	return  true
}

//消息前置
func (this *Ws)messageHandel(message *WsMessage,bytes []byte)  bool {
	//获取所有传输的字节
	//var message WsMessage
	utils.BytesToStruct(bytes,&message)
	if message.Function =="" {
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("参数错误"))
		return false
	}
	token := this.validateToken(this.Ctx.FormValue("token"))
	if !token {
		fmt.Print("=========")
		this.ExitChan <- "err"
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("token不正确"))
		return  false
	}
	return  true
}

//获取message
func (this *Ws) message(bytes []byte)  {
	var WsMessage WsMessage
	if handel := this.messageHandel(&WsMessage,bytes); !handel {
		return
	}

	//文件上传方法解析
	switch WsMessage.Function {
		case "upload" :
			this.upload(&WsMessage)
		case "selectUpload":
			this.selectUpload(&WsMessage)

	}



}

//查询文件上传
func (this *Ws) selectUpload(message *WsMessage)  {
	fileName, ok :=message.Param.(map[string]interface{})["file_name"]
	if !ok {
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("参数错误"))
		return
	}
	uploadService  :=service.NewUpload(fileName.(string))
	fileExits , size  :=uploadService.GetFileExits()
	result := make(map[string]interface{})
	result["status"] = fileExits
	result["size"] = size
	this.Conn.To(this.Conn.ID()).EmitMessage(lib.SuccessSuccessWsResponseData(result,"selectUpload"))
	return
}

//文件上传操作
func (this *Ws) upload(message *WsMessage)  {
	//获取文件filename
	//类型断言
	messageParamType := message.Param.(map[string]interface{})

	fileName,fileOk := messageParamType["file_name"]
	status , statusOk := messageParamType["status"]
	data , dataOk := messageParamType["data"].(string)
	title , titleOk := messageParamType["title"].(string)

	//字段类型判断
	if !fileOk || !statusOk  || !dataOk  || !titleOk {
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("参数错误"))
		return
	}
	uploadService := service.NewUpload(fileName.(string))

	tip , err :=uploadService.WsUploadFile(int(this.User.ID),title,status.(string),[]byte(data))

	if err != nil {
		this.Conn.To(this.Conn.ID()).EmitMessage(lib.ErrWsResponseMsg("文件上传失败"))
		return
	}

	result := make(map[string]interface{})
	result["tip"] = tip
	this.Conn.To(this.Conn.ID()).EmitMessage(lib.SuccessSuccessWsResponseData(result,"upload"))
	return
}










