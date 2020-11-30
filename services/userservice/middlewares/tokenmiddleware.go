package middlewares

import (
	"net/http"
	"time"

	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UserToken(c *gin.Context) {
	ok := getSession(c)
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
	ok := getSession(c)
	if !ok {
		noPrivilege(c)
		return
	}
	if value, exist := c.Get("session"); exist {
		if user, ok := value.(*model.AuthUser); !ok {
			noPrivilege(c)
		} else {
			if user.RoleID == constants.RoleAdmin {
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
	logrus.Debugf("no privilege %v\n", value)
	c.AbortWithStatusJSON(http.StatusOK, reply.CreateWithErrorX(reply.AccessNoPrivilege))
}

func checkToken(token string) bool {
	if token == "" {
		return false
	}
	/*TODO:引入redis依赖，通过判断redis的token set是否含有该token决定是否成功， 如果含有，则重设该token的过期时间 */
	return true
}

func getSession(c *gin.Context) bool {
	tokenString := c.Request.Header.Get(constants.AuthHeader)

	if !checkToken(tokenString) {
		return false
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &model.AuthUser{})
	if err != nil {
		return false
	}

	var (
		claim *model.AuthUser
		ok    bool
	)

	if claim, ok = token.Claims.(*model.AuthUser); !ok {
		return false
	} else if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()).Seconds() <= 0 {
		return false
	}

	logrus.Debugf("claim = %#v\n", claim)
	logrus.Debugf("standardclaim = %#v\n", claim.StandardClaims)

	claim, err = dao.GetAuthUser(db.Gorm, claim.ID)
	if err != nil {
		return false
	}

	ok = utils.VerifyToken(tokenString, claim.Password)
	if !ok {
		return false
	}

	c.Set("session", claim)
	return true
}
