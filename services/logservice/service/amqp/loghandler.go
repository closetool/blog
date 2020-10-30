package amqp

import (
	"github.com/closetool/blog/services/logservice/model/po"
	"github.com/closetool/blog/services/logservice/model/vo"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func SaveLogs() {
	messaging.Client.SubscribeToQueue("logs.saveLogs", "logs.saveLogs", func(d amqp.Delivery) {
		logsVO := vo.AuthUserLog{}
		if err := jsoniter.Unmarshal(d.Body, &logsVO); err != nil {
			logrus.Errorf("while unmarshal message body, encounterd an error: %v", err)
		}
		logsPO := &po.AuthUserLog{
			Description:    logsVO.Description,
			Device:         logsVO.Device,
			Parameter:      logsVO.Parameter,
			Url:            logsVO.Url,
			Code:           logsVO.Code,
			UserId:         logsVO.UserId,
			RunTime:        logsVO.RunTime,
			BrowserName:    logsVO.BrowserName,
			BrowserVersion: logsVO.BrowserVersion,
		}
		db.DB.InsertOne(logsPO)
	})
}
