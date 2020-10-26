package middlewares

import (
	"bytes"
	"net/http"

	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func UserToken(c *gin.Context) {
	header := c.Request.Header.Get(constants.AuthHeader)
	if header == "" {
		logrus.Debugln("header is empty")
		noPrivilege(c)
		return
	}
	obj := vo.AuthUser{Token: header}
	bt, _ := jsoniter.Marshal(obj)
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bt, viper.GetString("amqp_verifyToken"))
	if err != nil {
		noPrivilege(c)
		return
	}
	if bytes.Contains(rpl, []byte(reply.HandleErrCode(reply.Success))) {
		c.Next()
		user := vo.AuthUser{}
		jsoniter.Get(rpl, "model").ToVal(&user)
		c.Set("session", user)
	} else {
		noPrivilege(c)
	}
}

func AdminToken(c *gin.Context) {
	header := c.Request.Header.Get(constants.AuthHeader)
	if header == "" {
		logrus.Debugln("header is empty")
		noPrivilege(c)
		return
	}
	obj := vo.AuthUser{Token: header}
	bt, _ := jsoniter.Marshal(obj)
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bt, viper.GetString("amqp_verifyToken"))
	if err != nil {
		logrus.Errorf("send message to mq failed: %v", err)
		noPrivilege(c)
		return
	}
	if jsoniter.Get(rpl, "model", "roleId").ToInt() != constants.RoleAdmin {
		logrus.Debugf("reply = %v", string(rpl))
		logrus.Debugln("user has no admin privilege")
		noPrivilege(c)
		return
	} else {
		c.Next()
		user := vo.AuthUser{}
		jsoniter.Get(rpl, "model").ToVal(&user)
		c.Set("session", user)
	}
}

func noPrivilege(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, reply.CreateWithErrorX(reply.AccessNoPrivilege))
}

func checkToken(token string) bool {
	if token == "" {
		return false
	}
	/*TODO:引入redis依赖，通过判断redis的token set是否含有该token决定是否成功， 如果含有，则重设该token的过期时间 */
	return true
}
