package logs

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"
	log2 "log"
	"os"
	"time"
)

var Logs *LogsStruct

//日志系统
type LogsStruct struct {
	Path     string
	FilePath string
	Level    string
}

func NewLogs() *LogsStruct {
	timeFileName := time.Now().Format("2006-01-02") + ".log"
	logConf, _ := config.NewConfig("ini", "./conf/logs.conf")

	level := logConf.String("logs::level")
	filePath := logConf.String("logs::path")
	Logs = &LogsStruct{
		Path:     filePath,
		FilePath: filePath + "/" + timeFileName,
		Level:    level,
	}

	return Logs
}

//重写sql log
func (this LogsStruct) Print(v ...interface{}) {
	fileName := "/sql_" + time.Now().Format("2006-01-02") + ".log"
	f, _ := os.OpenFile(this.Path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log := log2.New(f, "", log2.Llongfile)
	message := gorm.LogFormatter(time.Now().Format("2006-01-02 15:04:05")+":\r\n", v)
	fmt.Print(message)
	log.Println(message)
}
