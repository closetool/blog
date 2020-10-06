package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/closetool/blog/services/userservice/models"
	"github.com/closetool/blog/services/userservice/service"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/utils/dbutils"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/core"
	"xorm.io/xorm"
)

var appName = "userservice"

func init() {

}

func main() {
	//logrus.SetLevel(logrus.DebugLevel)

	initConfig()

	parseFlag()

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	logrus.SetLevel(logrus.Level(viper.GetUint32("log_level")))

	initDatabase()

	r := initServer()

	listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}

func parseFlag() {
	h := flag.Bool("h", false, "help")
	configServer := flag.String("configServer", "", "config server's address and port")
	profile := flag.String("profile", "", "point out which profile you want to use")
	branch := flag.String("branch", "", "which branch in github")
	flag.Parse()

	if *h {
		flag.Usage()
		os.Exit(0)
	}
	if *configServer != "" {
		viper.Set("config_server_url", *configServer)
	}

	if *profile != "" {
		viper.Set("profile", *profile)
	}

	if *branch != "" {
		viper.Set("branch", *branch)
	}
}

func listen(handleExit func()) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-s
		handleExit()
		logrus.Infof("shutdown service %s", appName)
		os.Exit(0)
	}()
}

func initConfig() {
	viper.Set("log_level", logrus.InfoLevel)
	//viper.Set("config_server_url", "http://39.108.114.242:8888")
	//viper.Set("profile", "test")
	//viper.Set("branch", "blog")
	viper.Set("log_file_path", "./")
	viper.Set("log_file_name", fmt.Sprintf("%s_%s.log", appName, time.Now().Format("2006-01-02_15:04:05")))

	viper.SetConfigType("yml")
	viper.SetConfigName("userservice")
	configLoc := []string{"$HOME/config/", "/etc/userservice/", "$HOME/.musicservice/", "./"}
	for _, loc := range configLoc {
		viper.AddConfigPath(loc)
	}

	logrus.Infof("service will read userservice.yml configuration file from %s", strings.Join(configLoc, ","))

	_ = viper.ReadInConfig()
	//if err != nil {
	//logrus.Errorf("Read config file failed. Error: %v", err)
	//}
}

func initServer() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.LogToFile())
	r.Use(middlewares.Recover())
	group := r.Group("/music")

	routeutils.RegisterRoute(service.Routes, group)
	return r
}

func initDatabase() {

	url := fmt.Sprintf("%s:%s@%s",
		viper.Get("mysql_user"),
		viper.Get("mysql_password"),
		viper.Get("mysql_location"))

	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		logrus.Panicf("Connect to db failed: %v", err)
	}
	engine.SetTableMapper(core.NewPrefixMapper(
		core.SnakeMapper{},
		viper.GetString("mysql_prefix"),
	))

	if err != nil {
		logrus.Panicf("Connect to database by %s failed. Error: %v", url, err)
	}

	err = dbutils.InitTables(engine, &models.AuthUser{}, &models.AuthUserSocial{}, &models.AuthToken{}, &models.AuthUserLog{})
	if err != nil {
		logrus.Errorf("An error occurred when initial tables: %v", err)
	}

	service.Engine = engine
}
