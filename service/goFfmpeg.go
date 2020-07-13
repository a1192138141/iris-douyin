package service

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/xfrr/goffmpeg/transcoder"
)


//用于存放视屏转缩略图
type  VideoFfmepg struct {
	TempPath string
	SavePath string
}

//用于注入地址和文件名
func NewVideoFfmepg(tempPath string) *VideoFfmepg {
	info , err :=	config.NewConfig("ini","./config/file.conf")
	if err != nil {
		fmt.Print("获取配置文件错误：",err)
	}

	savePath :=info.String("ffmpeg::save_path")
	return &VideoFfmepg{TempPath:tempPath,SavePath:savePath}
}


func videoToImg()  {

	inputPath := ""

	outputPath := ""

	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize( inputPath, outputPath )
	// Handle error...

	// Start transcoder process without checking progress
	done := trans.Run(false)

	// This channel is used to wait for the process to end
	err = <-done
	// Handle error...

	fmt.Print(err)
}