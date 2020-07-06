package datamodels

import (
	"github.com/dgrijalva/jwt-go"
)

type UserLoginData struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type UserJwt struct {
	//User models.User
	jwt.StandardClaims
}
