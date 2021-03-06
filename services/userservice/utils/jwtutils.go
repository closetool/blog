package utils

import (
	"fmt"
	"time"

	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/models/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func GenerateToken(userPO *model.AuthUser) (string, error, int64) {
	user := model.AuthUser{}
	user.StandardClaims = new(jwt.StandardClaims)
	user.ExpiresAt = time.Now().Add(constants.TokenTTL).Unix()
	user.ID = userPO.ID
	user.Name = userPO.Name
	psw := userPO.Password
	user.Password = ""

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	ss, err := token.SignedString([]byte(psw))
	if err != nil {
		logrus.Errorf("generate jwt string failed: %v", err)
	}
	return ss, err, user.ExpiresAt
}

func VerifyToken(tokenString, key string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &model.AuthUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		logrus.Errorf("token is invalid: %v", err)
		return false
	}
	user, ok := token.Claims.(*model.AuthUser)
	if !ok {
		return false
	}

	if token.Valid && user.VerifyExpiresAt(time.Now().Unix(), true) {
		return true
	}
	return false
}
