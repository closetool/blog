package middlewares

import (
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				errCode, ok := r.(int)
				if ok {
					reply.CreateJSONError(c, errCode)
				} else {
					panic(r)
					//c.AbortWithError(http.StatusInternalServerError, errors.New(fmt.Sprintf("%v", r)))
				}
			}
		}()
		c.Next()
	}
}
