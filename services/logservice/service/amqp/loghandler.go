package amqp

import (
	"strconv"

	"github.com/closetool/blog/services/logservice/models/po"
	"github.com/closetool/blog/services/logservice/models/vo"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
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

func GetParamGroupByCode() {
	messaging.Client.SubscribeToQueueAndReply("logs.getParamGroupByCode", "logs.getParamGroupByCode", func(d amqp.Delivery) []byte {
		params, err := db.DB.SQL("select parameter,count(1) as count from closetool_auth_user_log where code = ? group by parameter", string(d.Body)).QueryString()
		if err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}
		logrus.Debugf("%#v", params)
		logVOs := make([]interface{}, 0)
		for _, obj := range params {
			count, err := strconv.ParseInt(obj["count"], 10, 64)
			if err != nil {
				return reply.ErrorBytes(reply.Error)
			}
			logVOs = append(logVOs, vo.AuthUserLog{
				Parameter: obj["parameter"],
				Count:     count,
			})
		}
		logrus.Debugf("%#v", logVOs)
		return reply.ModelsBytes(logVOs)
	})
}
