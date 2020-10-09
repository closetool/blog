package log

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var Logger logrus.Logger

func InitLog() {
	Logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%",
	})
	Logger.SetLevel(logrus.Level(viper.GetUint32("log_level")))
	logPath := viper.GetString("log_file_path")
	logName := viper.GetString("log_file_name")

	err := os.MkdirAll(logPath, 0755)
	if err != nil {
		logrus.Panicf("create log file path failed: %v", err)
	}
	filePath := path.Join(logPath, fmt.Sprintf("app_%s.log", logName))
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Panicf("create log file failed: %v", err)
	}

	Logger.SetOutput(io.MultiWriter(file, os.Stdout))
	Logger.Info("logger initialized")
}
