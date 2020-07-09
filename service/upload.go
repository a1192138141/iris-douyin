package service

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"os"
)

//用于文件上传的service
type UploadInterface interface {
	WsUploadFile()
}

type UploadService struct {
	FilePath string
}

func NewUpload(fileName string) *UploadService {
	fileConfig, err := config.NewConfig("ini","././conf/file.conf")
	if err != nil {
		fmt.Print(err)
	}
	path := fileConfig.String("file::path")
	return &UploadService{FilePath:path+fileName}
}

//websocket文件上传
func WsUploadFile(fileName string,status string , data []byte) (err error){
	//获取配置文件地址
	//fileConf ,err  :=config.NewConfig("ini","./conf/file.conf")
	//if err != nil {
	//	return  err
	//}
	//
	//if status =="start" {
	//	saveSize :=GetFileSize(fileName)
	//}else {
	//
	//
	//
	//}




	//if status == ""
	//
	//
	//
	//path :=fileConf.String("file::path")
	//fileSize :=GetFileSize(path)
	//fmt.Print("======文件大小===")
	//fmt.Print(fileSize)

	return  nil
}

func (this *UploadService) GetFileExits() (bool,int64) {
	fileInfo , err :=os.Stat(this.FilePath)
	fmt.Print("============")
	fmt.Print(err)
	if err != nil ||os.IsExist(err){
		return false , 0
	}
	return true,fileInfo.Size()
}

func GetFileSize(path string ) int {
	fileInfo , err :=os.Stat(path)
	if err != nil {
		return  0
	}
	return int(fileInfo.Size())
}



func WriteFileByAppend(path string, data []byte)  bool {
	fp , err :=os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	defer fp.Close()
	if err !=nil{
		return false
	}

	_, err = fp.Write(data)
	if err != nil {
		return false
	}
	return true
}

