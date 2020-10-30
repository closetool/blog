package amqp

import (
	"time"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func VerifyToken() {
	messaging.Client.SubscribeToQueueAndReply("auth.verifyToken", "auth.verifyToken", func(a amqp.Delivery) []byte {
		amqpUser := vo.AuthUser{}
		jsoniter.Get(a.Body).ToVal(amqpUser)
		header := amqpUser.Token
		if header == "" {
			return reply.ErrorBytes(reply.InvalidToken)
		}

		token, _, err := new(jwt.Parser).ParseUnverified(header, &vo.AuthUser{})
		if err != nil {
			return reply.ErrorBytes(reply.InvalidToken)
		}

		var (
			claim *vo.AuthUser
			ok    bool
		)

		if claim, ok = token.Claims.(*vo.AuthUser); !ok {
			return reply.ErrorBytes(reply.InvalidToken)
		} else if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()).Seconds() <= 0 {
			return reply.ErrorBytes(reply.InvalidToken)
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
			return reply.ErrorBytes(reply.InvalidToken)
		}

		ok = utils.VerifyToken(header, user.Password)
		if !ok {
			return reply.ErrorBytes(reply.InvalidToken)
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

func GetUserNameById() {
	messaging.Client.SubscribeToQueueAndReply("auth.getUserNameById", "auth.getUserNameById", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		jsoniter.Get(d.Body).ToVal(&ids)

		idsInterface := make([]interface{}, 0)
		for _, id := range ids {
			idsInterface = append(idsInterface, id)
		}

		users := make([]*po.AuthUser, 0)
		if err := db.DB.In("id", idsInterface...).Find(&users); err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		result := make(map[int64]string)
		for _, user := range users {
			result[user.Id] = user.Name
		}
		return reply.ModelBytes(result)
	})
}
