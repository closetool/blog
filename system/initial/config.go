package initial

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(appName string) {
	viper.Set("log_level", logrus.InfoLevel)
	//viper.Set("config_server_url", "http://39.108.114.242:8888")
	//viper.Set("profile", "test")
	//viper.Set("branch", "blog")
	viper.Set("log_file_path", "./logs/")
	viper.Set("log_file_name", fmt.Sprintf("%s_%s.log", appName, time.Now().Format("2006-01-02_15:04:05")))

	viper.SetConfigType("yml")
	viper.SetConfigName(appName)
	viper.SetConfigName("config")
	configLoc := []string{"$HOME/.config/", fmt.Sprintf("/etc/%s/", appName), fmt.Sprintf("$HOME/.%s/", appName), "./"}
	for _, loc := range configLoc {
		viper.AddConfigPath(loc)
	}

	logrus.Infof("service will read %s.yml or config.yml configuration file from %s", appName, strings.Join(configLoc, ","))

	_ = viper.ReadInConfig()
	//if err != nil {
	//logrus.Errorf("Read config file failed. Error: %v", err)
	//}
}
