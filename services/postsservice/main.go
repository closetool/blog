package main

import (
	"fmt"

	"github.com/closetool/blog/services/postsservice/models/po"
	"github.com/closetool/blog/services/postsservice/service"
	"github.com/closetool/blog/services/postsservice/service/amqp"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "postsservice"

func main() {
	initial.InitConfig(appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))

	initial.InitLog()

	db.DbInit(&po.Posts{}, &po.PostsTags{}, &po.PostsAttribute{}, &po.PostsComments{})
	db.SyncTables(&po.Posts{}, &po.PostsTags{}, &po.PostsAttribute{}, &po.PostsComments{})

	r := initial.InitServer()
	//TODO
	routeutils.RegisterRoute(service.PostsRoutes, r.Group("posts"))

	messaging.Client = new(messaging.MessagingClient)
	messaging.Client.ConnectToBroker(viper.GetString("amqp_location"))

	amqp.GetTagsIDAndCount()
	amqp.GetCategoryIDAndCount()

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}
