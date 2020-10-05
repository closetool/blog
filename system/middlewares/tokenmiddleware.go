package middlewares

import (
	"github.com/closetool/blog/system/blogerrors"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserToken(c *gin.Context){
	if !checkToken(c.GetHeader(constants.AUTH_HEADER)) {
		noPrivilege(c)
	}else{
		c.Next()
	}
}

func AdminToken(c *gin.Context){
	if !checkToken(c.GetHeader(constants.AUTH_HEADER)) {
		noPrivilege(c)
	}else{

	}
}

func noPrivilege(c *gin.Context){
	c.AbortWithStatusJSON(http.StatusOK,reply.CreateWithError(blogerrors.ACCESS_NO_PRIVILEGE))
}

func checkToken(token string)bool{
	if token == "" {
		return false
	}
	/*TODO:引入redis依赖，通过判断redis的token set是否含有该token决定是否成功， 如果含有，则重设该token的过期时间 */
	return true
}
