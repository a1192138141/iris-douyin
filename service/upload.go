package service

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"ims/models"
	"ims/utils"
	"os"
)

//用于文件上传的service
type UploadInterface interface {
	WsUploadFile()
}

type UploadService struct {
	VideoFilePath string
	VideoFileName string
	FfmpegPath string
	FfmpegName string
}

func NewUpload(fileName string) *UploadService {
	fileConfig, err := config.NewConfig("ini","./conf/file.conf")
	if err != nil {
		fmt.Print(err)
	}

	videoFilePath := fileConfig.String("file::path")
	ffmpegPath :=fileConfig.String("ffmpeg::save_path")
	ffmpegName := utils.UniqueId()+".jpg"
	return &UploadService{VideoFilePath:videoFilePath,VideoFileName:fileName,FfmpegPath:ffmpegPath,FfmpegName:ffmpegName}
}

//websocket文件上传
func(this *UploadService) WsUploadFile(userId int,title string ,status string , data []byte) (tip string,err error){
	writeErr :=this.WriteFileByAppend(data)



	if  writeErr != nil{
		return  "append",writeErr
	}
	if status == "start"  {
		return  "append" , nil
	}else {
		tip :="success"
		info := VideoToImg(this)
		fmt.Print(info)
		videoModel := models.Video{
			UserId:userId,
			Title:title,
			Thumbnail:this.FfmpegPath+this.FfmpegName,
			Video:this.VideoFilePath+this.VideoFileName,
			PlayNum:0,
			LikeNum:0,
			ForwardNum:0,
			CommentNum:0,
		}
		res :=videoModel.InsertOne()

		fmt.Print(res)

		if !res {
			tip ="视屏转换到jpg失败"
			err =errors.New("插入数据库失败")
		}

		return  tip,err
	}
}

//查询文件是否存在
func (this *UploadService) GetFileExits() (bool,int64) {
	fmt.Print(this.VideoFilePath+this.VideoFileName)
	fileInfo , err :=os.Stat(this.VideoFilePath+this.VideoFileName)
	if err != nil ||os.IsExist(err){
		return false , 0
	}
	return true,fileInfo.Size()
}



func(this *UploadService) WriteFileByAppend( data []byte)  error {
	fp , err :=os.OpenFile(this.VideoFilePath+this.VideoFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)

	defer fp.Close()
	if err !=nil{
		return nil
	}
	_, err = fp.Write(data)

	if err != nil {
		return err
	}
	return nil
}

