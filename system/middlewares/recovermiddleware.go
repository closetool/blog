package middlewares

import (
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recover() func (ctx *gin.Context){
	return func(c *gin.Context) {
		defer func(){
			if r := recover();r != nil {
				errCode,ok := r.(int)
				if ok {
					c.JSON(http.StatusOK,reply.CreateWithError(errCode))
				}
			}
		}()
	}
}
