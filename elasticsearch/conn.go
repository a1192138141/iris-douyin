package elasticsearch

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/olivere/elastic/v7"
	"ims/logs"
	"strconv"
)

type EsStruct struct {
	EsConn *elastic.Client
	host   string
	port   int
}

var EsConn *EsStruct

func GetEsConn() {

	var err error
	logs := logs.NewLogs()
	exConfig, err := config.NewConfig("ini", "./conf/datasource.conf")

	if err != nil {
		logs.Print(err)
		return
	}
	host := exConfig.String("elasticsearch::host")
	port := exConfig.String("elasticsearch::port")
	linkUrl := fmt.Sprintf("http://%s:%s", host, port)
	//创建连接
	es := EsStruct{}
	es.port, err = strconv.Atoi(port)
	es.host = host

	es.EsConn, err = elastic.NewClient(elastic.SetURL(linkUrl), elastic.SetSniff(false))
	EsConn = &es
	if err != nil {
		logs.Print("es连接错误：" + err.Error())
		return
	}

}
