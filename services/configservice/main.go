package main

import (
	"fmt"
	"sync"

	"github.com/closetool/blog/services/configservice/service"
	"github.com/closetool/blog/services/configservice/service/amqp"
	"github.com/closetool/blog/system/config"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/exit"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var appName = "configservice"

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
	db.Migrate(&model.Config{})
	dao.DB = db.Gorm

	cache := &sync.Map{}
	amqp.ConfigCache = cache
	service.ConfigCache = cache

	getConfig(cache)

	r := initial.InitServer()
	routeutils.RegisterRoute(service.ConfigRoutes, r.Group("/config"))

	messaging.Client = new(messaging.MessagingClient)
	messaging.Client.ConnectToBroker(viper.GetString("amqp_location"))

	amqp.GetConfig()

	exit.Listen(func() {})

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("service_port")))
	if err != nil {
		logrus.Errorf("An err occurred when service running: %v", err)
	}
}

func getConfig(cache *sync.Map) {
	var (
		configs []*model.Config
		err     error
	)
	if configs, _, err = dao.GetAllConfig(db.Gorm, 1, 100, ""); err != nil {
		logrus.Panicf("get configs failed: %v", err)
	}
	for _, config := range configs {
		logrus.Infof("load config %s => %s", config.ConfigKey, config.ConfigValue)
		cache.Store(config.ConfigKey, config.ConfigValue)
	}
}
