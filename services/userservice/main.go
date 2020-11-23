package main

import (
	"fmt"

	"github.com/closetool/blog/services/userservice/service"
	"github.com/closetool/blog/services/userservice/service/amqp"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "userservice"

func main() {
	initial.InitConfig(appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	initial.InitLog()

	db.GormInit()
	db.Migrate(&model.AuthUser{}, &model.AuthToken{}, &model.AuthUserSocial{})

	r := initial.InitServer()
	routeutils.RegisterRoute(service.Routes, r.Group("/auth"))

	messaging.Client = new(messaging.MessagingClient)
	messaging.Client.ConnectToBroker(viper.GetString("amqp_location"))
	//FIXME
	amqp.VerifyToken()
	amqp.GetUserNameById()
	amqp.SelectAdmin()
	amqp.GetUserById()

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}
