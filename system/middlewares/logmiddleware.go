package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		//log.Logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//	c.Writer.Status(),
		//	latency,
		//	c.ClientIP(),
		//	c.Request.Method,
		//	c.Request.RequestURI)
		logrus.Infof("| %-7s | %-3d | %-20s | %-13v | %-s |", c.Request.Method, c.Writer.Status(), c.Request.URL, latency, c.ClientIP())
	}
}
