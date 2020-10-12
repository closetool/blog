package utils

import (
	"fmt"
	"time"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/log"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userPO *po.AuthUser) (string, error, int64) {
	user := vo.AuthUser{}
	user.StandardClaims = new(jwt.StandardClaims)
	user.ExpiresAt = time.Now().Add(constants.TokenTTL).Unix()
	user.Id = userPO.Id
	user.Name = userPO.Name
	psw := userPO.Password
	user.Password = ""

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	ss, err := token.SignedString([]byte(psw))
	if err != nil {
		log.Logger.Errorf("generate jwt string failed: %v", err)
	}
	return ss, err, user.ExpiresAt
}

func VerifyToken(tokenString, key string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &vo.AuthUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		log.Logger.Errorf("token is invalid: %v", err)
		return false
	}
	user, ok := token.Claims.(*vo.AuthUser)
	if !ok {
		return false
	}

	if token.Valid && user.VerifyExpiresAt(time.Now().Unix(), true) {
		return true
	}
	return false
}
