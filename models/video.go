package models

import (
	"github.com/jinzhu/gorm"
)

type VideoInterface interface {
	GetTableName() string
}

type Video struct {
	gorm.Model
	UserId     int    //用户id
	Title      string `gorm:"type:text"` //标题
	Thumbnail  string //缩略图
	Video      string
	PlayNum    int //播放次数
	LikeNum    int //喜欢次数
	ForwardNum int //转发次数
	CommentNum int //评论条数
}

func (this *Video) GetTableName() string {
	return "videos"
}

//插入一条
func (this *Video) InsertOne() bool {
	DbConn.Create(this)
	if this.ID == 0 {
		return false
	}
	return true
}

func (this *Video) GetIds() []uint {
	var videos []Video
	DbConn.Select("id").Find(&videos)
	var ids []uint

	for _, value := range videos {
		ids = append(ids, value.ID)
	}
	return ids
}
