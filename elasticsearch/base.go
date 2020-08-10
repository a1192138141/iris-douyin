package elasticsearch

import (
	"context"
	"fmt"
	"ims/logs"
)

type BaseInterface interface {
	GetIndexName() string
	CreateIndex(index string, indexJson string) bool
	SearchIndex(index string) bool
	DeleteIndex(index string) bool
	InsertData(index string, data string) bool
	GetTyp() string
	Proxy(func(),...interface{})
}

type EsBase struct {
	indexName string
	typ string
}


func NewEsBase(indexName string) BaseInterface {
	return &EsBase{indexName:indexName,typ: "_doc"}
}

type TestDemoStruct struct {
	Name         string `json:"name"`
	Age          string `json:"age"`
	Married      bool   `json:"married"`
	Created      int64  `json:"created"`
	Tags         string `json:"tags"`
	Location     string `json:"location"`
	SuggestField string `json:"suggest_field"`
}

const TestDemo = `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
			"properties":{
				"name":{
					"type":"text"
				},
				"age":{
					"type":"long"
				},
				"married":{
				"type":"boolean"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
		}
	}
}`

func (self *EsBase) Proxy(function func() , args ...interface{}) {
	logs.NewLogs().Print(fmt.Sprintf("插入数据结构为%s\n",args))
	function()
	logs.NewLogs().Print(fmt.Sprintf("结束插入数据:%s",args))
}

func (self *EsBase) GetIndexName() string  {
	return  self.indexName
}

func (self *EsBase) GetTyp() string  {
	return  self.typ
}

func (self *EsBase) CreateIndex(index string, indexJson string) bool {
	_, err := EsConn.EsConn.CreateIndex(index).BodyJson(indexJson).Do(context.Background())
	return err != nil
}

func (self *EsBase) SearchIndex(index string) bool {
	_, err := EsConn.EsConn.Search(index).Do(context.Background())
	return err != nil
}

func (self *EsBase) DeleteIndex(index string) bool {
	_, err := EsConn.EsConn.DeleteIndex(index).Do(context.Background())
	return err != nil
}

func (self *EsBase) InsertData(index string, data string) bool {
	_, err := EsConn.EsConn.Index().Index(index).Type(self.typ).BodyJson(data).Do(context.Background())
	return err != nil
}
