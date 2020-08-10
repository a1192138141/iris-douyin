package lib

import (
	"fmt"
	//"crypto"
	//"fmt"
	"github.com/dgrijalva/jwt-go"
	//"golang.org/x/crypto/cryptobyte"
	//"ims/datamodels"
	"ims/models"
	//"reflect"
	//"time"
)

//const JwtKey  = []byte("douyin")

var (
	key []byte = []byte("douyin")
)

func GetJwtToken(UserInfo *models.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["user"] = UserInfo

	token.Claims = claims
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

func ParseUserToken(tokenString string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user"], true
	} else {
		fmt.Println("=====2=====")
		return "", false
	}

}
