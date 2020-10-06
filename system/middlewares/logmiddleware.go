package middlewares

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LogToFile() gin.HandlerFunc {
	logFilePath := viper.GetString("log_file_path")
	logFileName := viper.GetString("log_file_name")

	fileName := path.Join(logFilePath, logFileName)
	dst, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		logrus.Panicf("Open log file failed. Terminating. Error: %v", err)
	}

	logger := logrus.New()
	logger.Out = io.MultiWriter(dst, os.Stdout)
	logger.SetLevel(logrus.Level(viper.GetUint32("log_level")))
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		QuoteEmptyFields: true,
	})

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.RequestURI)
	}
}
