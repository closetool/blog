package initial

import (
	"github.com/closetool/blog/system/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitServer() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.LogToFile())
	r.Use(middlewares.Recover())
	r.Use(gin.Recovery())

	logrus.Infof("server initialized")
	return r
}
