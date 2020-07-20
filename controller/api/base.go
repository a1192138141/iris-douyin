package api

import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"github.com/kataras/iris"
	"ims/lib"
	"ims/models"
)

type Base struct {
	User *models.User
}

func NewBase() *Base  {
	return  &Base{User:&models.User{}}
}


func (this *Base) validate(ctx iris.Context) bool {
	//获取token
	token :=ctx.GetHeader("token")
	if token == "" {
		return  false
	}
	tokenRes , err :=lib.ParseUserToken(token)

	if err == false {
		return  false
	}

	//this.User
	errors := mapstructure.Decode(tokenRes,this.User)
	res := tokenRes.(map[string]interface{})
	this.User.ID = uint(res["ID"].(float64))

	if errors != nil {
		return  false
	}
	return  true
}

//中间件
func (this *Base)AnyMiddlewareHere(ctx iris.Context)  {
	if this.validate(ctx){
		ctx.Next()
	}else {
		errJson ,_ :=json.Marshal(lib.ErrMsg("token错误"))
		ctx.WriteString(string(errJson))
	}
}