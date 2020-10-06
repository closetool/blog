package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/closetool/blog/services/userservice/handlers"
	"github.com/closetool/blog/services/userservice/models"
	"github.com/closetool/blog/services/userservice/routes"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/utils/dbutils"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/core"
	"xorm.io/xorm"
)

var appName = "userservice"

func init() {

}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	initConfig()

	parseFlag()

	//logrus.SetLevel(logrus.Level(viper.GetUint32("log_level")))

	initDatabase()

	r := initServer()

	go listen()

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}

func parseFlag() {
	configServer := flag.String("configServer", "", "config server's address and port")
	profile := flag.String("profile", "", "point out which profile you want to use")
	branch := flag.String("branch", "", "which branch in github")
	logFilePath := flag.String("log", "", "path to store logs")
	logLevel := flag.String("level", "", "log level")
	flag.Usage()
	flag.Parse()
	if *configServer != "" {
		logrus.Debug(configServer)
		viper.Set("config_server_url", *configServer)
	}

	if *profile != "" {
		viper.Set("profile", *profile)
	}

	if *branch != "" {
		viper.Set("branch", *branch)
	}

	if *logFilePath != "" {
		viper.Set("log_file_path", *logFilePath)
	}

	if *logLevel != "" {
		viper.Set("log_level", *logLevel)
	}
}

func listen() {
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM|syscall.SIGINT)
	<-s
	logrus.Infof("shutdown service %s", viper.GetString("app_name"))
}

func initConfig() {
	viper.Set("log_level", logrus.InfoLevel)
	viper.Set("config_server_url", "http://39.108.114.242:8888")
	viper.Set("profile", "test")
	viper.Set("branch", "blog")
	viper.Set("log_file_path", "./")
	viper.Set("log_file_name", fmt.Sprintf("%s_%s.log", appName, time.Now().Format("2006-01-02_15:04:05")))

	viper.SetConfigType("yml")
	viper.SetConfigName("userservice")
	configLoc := []string{"/etc/userservice", "$HOME/.musicservice", "./"}
	for _, loc := range configLoc {
		viper.AddConfigPath(loc)
	}

	logrus.Infof("service will read userservice.yml configuration file from %s", strings.Join(configLoc, ","))

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Read config file failed. Error: %v", err)
	}

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)
}

func initServer() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.LogToFile())
	r.Use(middlewares.Recover())
	group := r.Group("/auth")

	routeutils.RegisterRoute(routes.Routes, group)
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

	handlers.Engine = engine
}
