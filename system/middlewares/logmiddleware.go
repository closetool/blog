package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

func LogToFile() gin.HandlerFunc{
	logFilePath :=  viper.GetString("log_file_path")
	logFileName := viper.GetString("log_file_name")

	fileName := path.Join(logFilePath,logFileName)
	dst,err := os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY|os.O_CREATE,os.ModeAppend)
	if err != nil{
		logrus.Panicf("Open log file failed. Terminating. Error: %v",err)
	}

	logger := logrus.New()
	logger.Out = dst
	logger.SetLevel(logrus.Level(viper.GetUint32("log_level")))
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context){
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