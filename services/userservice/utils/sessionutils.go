package utils

import (
	"time"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/reply"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) {
	tokenString := c.Request.Header.Get(constants.AuthHeader)
	if tokenString == "" {
		panic(reply.InvalidToken)
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &vo.AuthUser{})
	if err != nil {
		panic(reply.InvalidToken)
	}

	var (
		claim *vo.AuthUser
		ok    bool
	)

	if claim, ok = token.Claims.(*vo.AuthUser); !ok {
		panic(reply.InvalidToken)
	} else if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()).Seconds() <= 0 {
		panic(reply.LoginOverTime)
	}

	log.Logger.Debugf("claim = %#v\n", claim)
	log.Logger.Debugf("standardclaim = %#v\n", claim.StandardClaims)

	user := &po.AuthUser{
		Id: claim.Id,
	}
	ok, err = db.DB.Get(user)
	if !ok || err != nil {
		panic(reply.InvalidToken)
	}

	ok = VerifyToken(tokenString, user.Password)
	if !ok {
		panic(reply.InvalidToken)
	}

	c.Set("session", user)
}
