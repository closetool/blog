package main

import (
	"fmt"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/service"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/messaging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "userservice"

func main() {
	//logrus.SetLevel(logrus.DebugLevel)

	initial.InitConfig(appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	initial.InitLog()

	db.DbInit(&po.AuthUser{}, &po.AuthToken{}, &po.AuthUserLog{}, &po.AuthUserSocial{})
	db.SyncTables(&po.AuthUser{}, &po.AuthToken{}, &po.AuthUserLog{}, &po.AuthUserSocial{})

	r := initial.InitServer(service.Routes, "/auth")

	messaging.Client = new(messaging.MessagingClient)
	messaging.Client.ConnectToBroker(viper.GetString("amqp_location"))
	//FIXME
	service.TokenHandler()

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}
