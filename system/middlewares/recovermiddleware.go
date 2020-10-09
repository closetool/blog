package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				errCode, ok := r.(int)
				if ok {
					c.JSON(http.StatusOK, reply.CreateWithErrorX(errCode))
					c.Abort()
				} else {
					logrus.Errorf("Recovered: %v", r)
					c.AbortWithError(http.StatusInternalServerError, errors.New(fmt.Sprintf("%v", r)))
				}
			}
		}()
		c.Next()
	}
}
