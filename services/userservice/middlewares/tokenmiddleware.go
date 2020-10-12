package middlewares

import (
	"net/http"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
)

func UserToken(c *gin.Context) {
	ok := utils.GetSession(c)
	if !ok {
		noPrivilege(c)
		return
	}
	if _, exist := c.Get("session"); exist {
		c.Next()
	} else {
		noPrivilege(c)
	}
}

func AdminToken(c *gin.Context) {
	ok := utils.GetSession(c)
	if !ok {
		noPrivilege(c)
		return
	}
	if value, exist := c.Get("session"); exist {
		if user, ok := value.(*po.AuthUser); !ok {
			noPrivilege(c)
		} else {
			if user.RoleId == constants.RoleAdmin {
				c.Next()
			} else {
				noPrivilege(c)
			}
		}
	} else {
	}
}

func noPrivilege(c *gin.Context) {
	value, _ := c.Get("session")
	log.Logger.Debugf("no privilege %v\n", value)
	c.AbortWithStatusJSON(http.StatusOK, reply.CreateWithErrorX(reply.AccessNoPrivilege))
}

func checkToken(token string) bool {
	if token == "" {
		return false
	}
	/*TODO:引入redis依赖，通过判断redis的token set是否含有该token决定是否成功， 如果含有，则重设该token的过期时间 */
	return true
}
