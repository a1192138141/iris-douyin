package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ims/elasticsearch"
	"ims/lib"
)

/**
用于搜索的搜索类
*/
type Search struct {
	Ctx  iris.Context
	Base *Base
}

func NewSearch() *Search {
	return &Search{Base: NewBase()}
}

//前置操作 中间件
func (this *Search) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/search", "Search")
}

func (this *Search) Search() interface{} {
	video := elasticsearch.NewEsVideo()
	res := video.SearchKeyWord("desc","世界上")
	return  lib.SuccessData(res)
}
