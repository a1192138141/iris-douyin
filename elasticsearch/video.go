package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"reflect"
	"time"
)

/**
用于保存视屏的测试
*/


const videoMappings = `{
	"mappings":{
			"properties":{
				"user_id":{
					"type":"integer"
				},
				
				"title":{
					"type":"string"
				},
				"desc":{
					"type":"text"
				},
				"name":{
					"type":"string"
				},
				"url":{
					"type":"string"
				},
				"created":{
					"type":"date"
				},
				"updated":{
					"type":"date"
				},
		}
	}
}`

/**
保存的数据结构
*/
type SaveVideoStruct struct {
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Created int  `json:"created"`
	Updated int    `json:"updated"`
}

type EsVideo struct {
	BaseInterface
	saveData  SaveVideoStruct
}

func NewEsVideo() *EsVideo {
	return &EsVideo{NewEsBase("es_video"), SaveVideoStruct{}}
}
/**
	插入测试
 */
func (self *EsVideo) InsertEsVideo()  bool{
	save := SaveVideoStruct{
		1,
		"golang ok",
		"世界上最好的语言",
		"php",
		"",
		int(time.Now().Unix()),
		int(time.Now().Unix()),
	}
	saveJson,_ := json.Marshal(save)
	data := false
	self.BaseInterface.Proxy(func(){
		data =  self.InsertData(self.BaseInterface.GetIndexName(),string(saveJson))
	},self.BaseInterface.GetIndexName(),string(saveJson),"video.InsertData")
	return  true
}

func (self *EsVideo) SearchKeyWord(indexName string , value string) []SaveVideoStruct {
	searchQuery := indexName + ":*"+value+"*"
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewQueryStringQuery(searchQuery))
	searchResult, _ :=EsConn.EsConn.Search().Index(self.GetIndexName()).Query(boolQ).Do(context.Background())
	total := searchResult.TotalHits()
	Subject := []SaveVideoStruct{}
	if total > 0 {
		for _, item := range searchResult.Each(reflect.TypeOf(SaveVideoStruct{})) {
			Subject = append(Subject ,item.(SaveVideoStruct))
		}

	} else {
	}
	return Subject
}



