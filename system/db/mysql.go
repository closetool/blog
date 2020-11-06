package db

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

var DB *xorm.Engine
var DbNotInitialized = errors.New("mysql isn't initialized")

func DbInit(beans ...interface{}) {
	logrus.Info("xorm.Engine initialized")
	dataSource := fmt.Sprintf(viper.GetString("db_location"), viper.GetString("db_password"))
	var err error
	DB, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		logrus.Panicf("connect to mysql failed: %v", err)
	}

	for _, bean := range beans {
		cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
		DB.MapCacher(bean, cacher)
	}
}

func SyncTables(beans ...interface{}) error {
	if DB == nil {
		return DbNotInitialized
	}
	return DB.Sync2(beans)
}

//var DB *gorm.DB
//
//func DbInit() {
//	dataSource := fmt.Sprintf(viper.GetString("db_location"), viper.GetString("db_password"))
//	var err error
//	DB, err = gorm.Open(viper.GetString("db_type"), dataSource)
//	if err != nil {
//		logrus.Panicf("connect to mysql failed: %v", err)
//	}
//}
//
//func SyncTables(beans ...interface{}) error {
//	if DB == nil {
//		return DbNotInitialized
//	}
//	DB.AutoMigrate(beans...)
//	return nil
//}
