package service

import "ims/models"

func GetVideoIds()  []uint {
	videoModel := models.Video{}
	return videoModel.GetIds()
}