package main

import (
	"fmt"
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
	"os"
	"os/signal"
	"syscall"
	"xorm.io/core"
	"xorm.io/xorm"
)

func init(){
	viper.Set("app_name","userservice")
	viper.Set("config_server_url","39.108.114.242:8888/")
	viper.Set("profile","test")
	viper.Set("profile","blog")

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/userservice")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Read config file failed. Error: %v",err)
	}
}

func main(){

	logrus.SetLevel(logrus.Level(viper.GetInt("log_level")))

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		viper.GetString("app_name"),
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	initDatabase()

	r := gin.New()
	r.Use(middlewares.LogToFile())
	r.Use(middlewares.Recover())
	group := r.Group("/auth")

	routeutils.RegisterRoute(routes.Routes,group)

	go listen()

	err := r.Run(fmt.Sprintf(":%d",viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v",err)
	}
}

func listen(){
	s := make(chan os.Signal)
	signal.Notify(s,syscall.SIGTERM|syscall.SIGINT)
	<-s
	logrus.Infof("shutdown service %s",viper.GetString("app_name"))
}

func initDatabase(){

	url := fmt.Sprintf("%s:%s@%s",
		viper.Get("mysql_user"),
		viper.Get("mysql_password"),
		viper.Get("mysql_location"))

	engine,err := xorm.NewEngine("mysql",url)
	if err != nil {
		logrus.Panicf("Connect to db failed: %v",err)
	}
	engine.SetTableMapper(core.NewPrefixMapper(
		core.SnakeMapper{},
		viper.GetString("mysql_prefix"),
	))

	if err != nil {
		logrus.Panicf("Connect to database by %d failed. Error: %v",url,err)
	}

	err = dbutils.InitTables(engine,&models.AuthUser{},&models.AuthUserSocial{},&models.AuthToken{},&models.AuthUserLog{})
	if err != nil {
		logrus.Errorf("An error occurred when initial tables: %v",err)
	}

	handlers.Engine = engine
}

