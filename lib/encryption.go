package lib

import (
	"crypto/md5"
	"fmt"
)

const PassSlot = "douyin"

//加密密码
func EncryptionPassword(password string) string {
	strByes := []byte(password + PassSlot)
	md5s := md5.Sum(strByes)
	md51 := fmt.Sprint("%x", md5s)
	res := md5.Sum([]byte(md51))
	result := fmt.Sprintf("%x", res)
	return result
}
