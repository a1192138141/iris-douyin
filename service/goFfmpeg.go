package service

import (
	"github.com/xfrr/goffmpeg/transcoder"
)

func VideoToImg(service *UploadService)  error {
	inputPath := service.FfmpegPath+service.FfmpegName
	outputPath := service.VideoFilePath+service.VideoFileName

	trans := new(transcoder.Transcoder)

	err := trans.Initialize( outputPath,inputPath )

	done := trans.Run(false)

	err = <-done

	return  err
}