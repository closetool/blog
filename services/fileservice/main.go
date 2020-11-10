package main

import (
	"fmt"
	"net/http"

	"github.com/closetool/blog/services/fileservice/service"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "fileservice"

func main() {
	initial.InitConfig(appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("config_server_url"),
		appName,
		viper.GetString("profile"),
		viper.GetString("branch"),
	)

	initial.InitLog()

	r := initial.InitServer()
	routeutils.RegisterRoute(service.FileRoutes, r.Group("/file"))

	messaging.Client = new(messaging.MessagingClient)
	messaging.Client.ConnectToBroker(viper.GetString("amqp_location"))

	initFileServer(r)

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}

func initFileServer(r *gin.Engine) {
	keys := []string{constants.StoreType, constants.DefaultPathKey, constants.DefaultImageDomain}
	configs := service.GetConfig(keys)
	filePath := configs[constants.DefaultPathKey]
	if filePath == "" {
		filePath = constants.DefaultPathValue
	}
	r.StaticFS("/files", http.Dir(filePath))
}
