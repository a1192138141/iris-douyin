package service

import (
	"ims/datamodels"
	"ims/models"
)

func GetVideoIds()  []uint {
	videoModel := models.Video{}
	return videoModel.GetIds()
}

func GetVideoInfoById(id int)  (videoInfo datamodels.VideoInfo , err error) {
	videoModel := models.Video{}
	videoName :=videoModel.GetTableName()

	userModel := models.User{}
	userName := userModel.GetTableName()

	joinSql := "left join "+userName + " on " + videoName+".user_id = "+userName+".id"
	rows ,err :=models.DbConn.Table(videoName).Select("*").Joins(joinSql).Where(videoName+".id = ?",int(id)).Rows()

	result := datamodels.VideoInfo{}
	for rows.Next() {
		err = models.DbConn.ScanRows(rows,&result)
	}
	if err != nil {
		return result,err
	}
	return result,err


}