package service

import (
	"errors"
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
func(this *UploadService) WsUploadFile(userId int,title string ,status string , data []byte) (tip string,err error){
	writeRes :=this.WriteFileByAppend(data)

	if !writeRes{
		return  "append",errors.New("文件上传失败")
	}
	if status == "start"  {
		return  "append" , nil
	}else {
		//this.FilePath
		ffmpeg := NewVideoFfmepg(this.FilePath)
		fmt.Print(ffmpeg)
		//todo video 转缩略图 插入数据库
		//fmt.Print("=======")



		return  "success",nil
	}
}

//查询文件是否存在
func (this *UploadService) GetFileExits() (bool,int64) {
	fileInfo , err :=os.Stat(this.FilePath)
	if err != nil ||os.IsExist(err){
		return false , 0
	}
	return true,fileInfo.Size()
}



func(this *UploadService) WriteFileByAppend( data []byte)  bool {
	fp , err :=os.OpenFile(this.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
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

