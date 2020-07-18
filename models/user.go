package models

import (
	"github.com/jinzhu/gorm"
)

//定义user的方法
type UserInterface interface{
	GetUserInfoById(UserId int) *User
	GetUserInfoByPhone(Phone string) *User
}


//user 结构体
type User struct {
	gorm.Model
	Phone string //手机号码
	Password string //密码
	NickName string  //昵称
	Avatar string  //头像
	Synopsis string  //简介
	Sex int //性别
	Age int //年龄
	Birthday string //生日
	Country string //国家
	Province string //省
	City string  //市
}


//获取表名
func (this *User) GetTableName() string  {
	return "users"
}

func (this *User) GetUserInfoById(UserId int) *User  {
	DbConn.Where("id = ?" ,UserId).First(this)
	return this
}

func (this *User) GetUserInfoByPhone(Phone string) *User  {
	DbConn.Where("phone = ?",Phone).First(this)
	return  this
}







