package service

import (
	"time"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func TokenHandler() {
	messaging.Client.SubscribeToQueueAndReply(viper.GetString("amqp_token"), viper.GetString("amqp_token"), func(a amqp.Delivery) []byte {
		header := jsoniter.Get(a.Body, constants.AuthHeader).ToString()
		if header == "" {
			return errorReply()
		}

		token, _, err := new(jwt.Parser).ParseUnverified(header, &vo.AuthUser{})
		if err != nil {
			return errorReply()
		}

		var (
			claim *vo.AuthUser
			ok    bool
		)

		if claim, ok = token.Claims.(*vo.AuthUser); !ok {
			return errorReply()
		} else if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()).Seconds() <= 0 {
			return errorReply()
		}

		logrus.Debugf("claim = %#v", claim)
		logrus.Debugf("standardclaim = %#v", claim.StandardClaims)

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
			return errorReply()
		}

		ok = utils.VerifyToken(header, user.Password)
		if !ok {
			return errorReply()
		}

		userVO := &vo.AuthUser{
			Name:   user.Name,
			Email:  user.Email,
			Id:     user.Id,
			RoleId: user.RoleId,
		}
		bytes, err := jsoniter.Marshal(reply.CreateWithModel(userVO))
		if err != nil {
			return nil
		}
		return bytes

	})
}

func errorReply() []byte {
	bytes, err := jsoniter.Marshal(reply.CreateWithErrorX(reply.InvalidToken))
	if err != nil {
		return nil
	}
	return bytes
}
