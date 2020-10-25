package main

import (
	"fmt"

	"github.com/closetool/blog/services/categoryservice/models/po"
	"github.com/closetool/blog/services/categoryservice/service"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "categoryservice"

func main() {
	initial.InitConfig(appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	//viper.Set("log_level", )

	initial.InitLog()

	db.DbInit(&po.Category{}, &po.CategoryTags{}, &po.Tags{})
	db.SyncTables(&po.Category{}, &po.CategoryTags{}, &po.Tags{})

	r := initial.InitServer()
	routeutils.RegisterRoute(service.CategoryRoutes, r.Group("/category"))
	routeutils.RegisterRoute(service.TagsRoutes, r.Group("/tags"))

	messaging.Client = new(messaging.MessagingClient)
	messaging.Client.ConnectToBroker(viper.GetString("amqp_location"))

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}
