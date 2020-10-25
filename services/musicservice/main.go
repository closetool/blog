package main

import (
	"fmt"

	"github.com/closetool/blog/services/musicservice/service"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "musicservice"

func main() {
	//logrus.SetLevel(logrus.DebugLevel)

	initial.InitConfig(appName)

	//parseFlag()

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	initial.InitLog()

	r := initial.InitServer()
	routeutils.RegisterRoute(service.Routes, r.Group("/music"))

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}
