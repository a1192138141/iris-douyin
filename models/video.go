package models

import "github.com/jinzhu/gorm"

type video struct {
	gorm.Model
	userId int //用户id
	title string `gorm:"type:text"` //标题
	thumbnail string  //缩略图
	playNum int //播放次数
	likeNum int  //喜欢次数
	forwardNum int //转发次数
	commentNum int //评论条数
}