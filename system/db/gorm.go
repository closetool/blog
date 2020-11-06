package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func GormInit() {
	dsn := fmt.Sprintf(viper.GetString("gorm_locaion"), viper.GetString("db_password"))
	var err error
	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("can not connect to database")
	}
}
