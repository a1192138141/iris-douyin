package lib

import "encoding/json"

//websocket 请求返回
type ResponseWs struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Cmd string `json:"cmd"`
	Data interface{} `json:"data"`
}

func ErrWsResponseMsg(message string) []byte {
	errRes := &ResponseWs{Code:500,Cmd:"",Data:make(map[string]string),Message:message}
	res ,_ :=json.Marshal(errRes)
	return res
}

func SuccessWsResponseMsg(message string) []byte  {
	errRes := &ResponseWs{Code:200,Cmd:"",Data:make(map[string]string),Message:message}
	res ,_ :=json.Marshal(errRes)
	return res
}

func SuccessSuccessWsResponseData(data interface{}) []byte  {
	errRes := &ResponseWs{Code:200,Cmd:"",Data:data,Message:"success"}
	res ,_ :=json.Marshal(errRes)
	return res
}





