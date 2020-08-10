package lib

type ResponseApi struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessData(data interface{}) *ResponseApi {
	return &ResponseApi{Code: 200, Message: "success", Data: data}
}

func SuccessMsg() *ResponseApi {
	return &ResponseApi{Code: 200, Message: "success", Data: make(map[string]string)}
}

func ErrMsg(message string) *ResponseApi {
	return &ResponseApi{Code: 500, Message: message, Data: make(map[string]string)}
}
