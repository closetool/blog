package amqp

import (
	"time"

	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func VerifyToken() {
	messaging.Client.SubscribeToQueueAndReply("auth.verifyToken", "auth.verifyToken", func(a amqp.Delivery) []byte {
		header := string(a.Body)
		logrus.Debugln(header)
		if header == "" {
			return reply.ErrorBytes(reply.InvalidToken)
		}

		token, _, err := new(jwt.Parser).ParseUnverified(header, &model.AuthUser{})
		if err != nil {
			logrus.Debug(err)
			return reply.ErrorBytes(reply.InvalidToken)
		}

		var (
			claim *model.AuthUser
			ok    bool
		)

		if claim, ok = token.Claims.(*model.AuthUser); !ok {
			logrus.Debug("can not convert interface{} to *vo.AuthUser")
			return reply.ErrorBytes(reply.InvalidToken)
		} else if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()).Seconds() <= 0 {
			logrus.Debug("Time Excced")
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

		user, err := dao.GetAuthUser(db.Gorm, claim.ID)
		if err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		ok = utils.VerifyToken(header, user.Password)
		if !ok {
			return reply.ErrorBytes(reply.InvalidToken)
		}

		user.Password = ""

		bytes, err := jsoniter.Marshal(reply.CreateWithModel(user))
		if err != nil {
			return reply.ErrorBytes(reply.Error)
		}
		return bytes

	})
}

func GetUserNameById() {
	messaging.Client.SubscribeToQueueAndReply("auth.getUserNameById", "auth.getUserNameById", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		jsoniter.Get(d.Body).ToVal(&ids)

		users := make([]model.AuthUser, 0)

		if err := db.Gorm.Where("id in ?", ids).Find(&users).Error; err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		result := make(map[int64]string)
		for _, user := range users {
			if user.Name.Valid {
				result[user.ID] = user.Name.String
			}
		}
		return reply.ModelBytes(result)
	})
}

func SelectAdmin() {
	messaging.Client.SubscribeToQueueAndReply("auth.selectAdmin", "auth.selectAdmin", func(d amqp.Delivery) []byte {
		admin := model.AuthUser{}
		if err := db.Gorm.Where("role_id = ?", constants.RoleAdmin).First(&admin).Error; err != nil {
			logrus.Errorf("can not get admin: %v", err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}
		admin.Password = ""
		return reply.ModelBytes(admin)
	})
}

func GetUserById() {
	messaging.Client.SubscribeToQueueAndReply("auth.getUserById", "auth.getUserById", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		jsoniter.Get(d.Body).ToVal(&ids)

		users := make([]model.AuthUser, 0)
		if err := db.Gorm.Where("id in ?", ids).Find(&users).Error; err != nil {
			logrus.Errorf("can not get users: %v", err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		result := make(map[int64]model.AuthUser)
		for _, user := range users {
			user.Password = ""
			result[user.ID] = user
		}
		return reply.ModelBytes(result)
	})
}
