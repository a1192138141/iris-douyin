package service

import (
	"errors"
	"ims/lib"
	"ims/models"
)

func GetUserInfoByPhone(phone string, password string) (map[string]string, error) {
	var maps = make(map[string]string)
	user := models.User{}
	userInfo := user.GetUserInfoByPhone(phone)
	if userInfo.ID == 0 {
		return maps, errors.New("账号或者密码错误")
	}

	//进行密码对别
	pass := lib.EncryptionPassword(password)
	if userInfo.Password != pass {
		return maps, errors.New("账号或者密码错误")
	}

	//对比成功生成token
	//jwt :=datamodels.UserJwt{User:userInfo}
	token , err :=lib.GetJwtToken(userInfo)

	maps["token"] = token
	return maps, err

}
