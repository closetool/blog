package service

import (
	"net/http"
	"sync"

	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ConfigCache *sync.Map

func health(c *gin.Context) {
	if db.Gorm != nil {
		c.JSON(http.StatusOK, map[string]bool{"health": true})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"health": false})
}

func updateConfig(c *gin.Context, tx *gorm.DB) {
	configs := []model.Config{}
	c.ShouldBindJSON(&configs)
	for _, config := range configs {
		if config.ConfigKey == "" || config.ConfigValue == "" {
			panic(reply.ParamIncorrect)
		}
		if err := tx.Model(&model.Config{}).
			Where("config_key = ?", config.ConfigKey).
			Update("config_value", config.ConfigValue).Error; err != nil {
			logrus.Debugln(err)
			panic(reply.DatabaseSqlParseError)
		}
		ConfigCache.Store(config.ConfigKey, config.ConfigValue)
	}
	reply.CreateJSONsuccess(c)
}

func getConfigList(c *gin.Context) {
	config := model.Config{}
	c.ShouldBindQuery(&config)
	configList(c, &config)
}

func getConfigBaseList(c *gin.Context) {
	config := model.Config{}
	c.ShouldBindQuery(&config)
	config.Type = constants.ConfigTypeBase
	configList(c, &config)
}

func configList(c *gin.Context, config *model.Config) {
	configs := []model.Config{}
	if config.Type == constants.ConfigTypeAliyun || config.Type == constants.ConfigTypeQiniu || config.Type == constants.ConfigTypeCross {
		if err := db.Gorm.Model(config).Find(&configs, "type in ?", []int32{config.Type, constants.ConfigTypeStoreType}).Error; err != nil {
			logrus.Debugln(err)
			panic(reply.DatabaseSqlParseError)
		}
	} else {
		if err := db.Gorm.Model(config).Find(&configs, "type = ?", config.Type).Error; err != nil {
			logrus.Debugln(err)
			panic(reply.DatabaseSqlParseError)
		}
	}
	ints := model.Configs2Interface(configs)
	reply.CreateJSONModels(c, ints)
}
