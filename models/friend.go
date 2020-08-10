package models

import "github.com/jinzhu/gorm"

//好友表
type Friend struct {
	gorm.Model
	UserId   int //用户id
	FriendId int //好友id
}

func GetVideoInfo(userId int) {
	//var friends []Friend
	//DbConn.Where("")
	//friend := []Friend{}
	//DbConn.

}
