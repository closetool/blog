package service

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/closetool/blog/services/fileservice/upload"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

func uploadFile(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}
	keys := []string{constants.StoreType, constants.DefaultPathKey, constants.DefaultImageDomain}
	configs := GetConfig(keys)
	uploadService := upload.Services[configs[constants.StoreType]]
	if configs[constants.StoreType] == "" {
		uploadService = upload.Services[constants.DefaultType]
	}
	filePath := configs[constants.DefaultPathKey]
	if filePath == "" {
		filePath = constants.DefaultPathValue
	}

	fileName := ""
	if uploadService.Check(f) {
		fileName = uploadService.SaveFile(f, filePath)
	}
	fileURL := fmt.Sprintf("%s%s%s", configs[constants.DefaultImageDomain], constants.FileURL, fileName)
	reply.CreateJSONExtra(c, fileURL)
}

func GetConfig(keys []string) map[string]string {
	bts, _ := jsoniter.Marshal(keys)
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, "config.getConfig")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		logrus.Debugln(err)
		panic(reply.Error)
	}

	configs := map[string]string{}
	jsoniter.Get(rpl, "model").ToVal(&configs)
	return configs
}
