package utils

import (
	"time"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) bool {
	tokenString := c.Request.Header.Get(constants.AuthHeader)
	if tokenString == "" {
		return false
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &vo.AuthUser{})
	if err != nil {
		return false
	}

	var (
		claim *vo.AuthUser
		ok    bool
	)

	if claim, ok = token.Claims.(*vo.AuthUser); !ok {
		return false
	} else if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()).Seconds() <= 0 {
		return false
	}

	log.Logger.Debugf("claim = %#v\n", claim)
	log.Logger.Debugf("standardclaim = %#v\n", claim.StandardClaims)

	//登录逻辑，数据库中需要存在相应的token
	//userToken := &po.AuthToken{}
	//ok, err = db.DB.Where("user_id = ?", claim.Id).Get(userToken)
	//if !ok || err != nil || !strings.EqualFold(userToken.Token, tokenString) {
	//	return false
	//}

	user := &po.AuthUser{
		Id: claim.Id,
	}
	ok, err = db.DB.Get(user)
	if !ok || err != nil {
		return false
	}

	ok = VerifyToken(tokenString, user.Password)
	if !ok {
		return false
	}

	c.Set("session", user)
	return true
}
