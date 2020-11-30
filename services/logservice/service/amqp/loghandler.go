package amqp

import (
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/guregu/null"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func SaveLogs() {
	messaging.Client.SubscribeToQueue("logs.saveLogs", "logs.saveLogs", func(d amqp.Delivery) {
		log := model.AuthUserLog{}
		if err := jsoniter.Unmarshal(d.Body, &log); err != nil {
			logrus.Errorf("while unmarshal message body, encounterd an error: %v", err)
		}
		dao.AddAuthUserLog(db.Gorm, &log)
	})
}

func GetParamGroupByCode() {
	messaging.Client.SubscribeToQueueAndReply("logs.getParamGroupByCode", "logs.getParamGroupByCode", func(d amqp.Delivery) []byte {
		rows, err := db.Gorm.
			Raw("select parameter,count(1) as count from closetool_auth_user_log where code = ? group by parameter", string(d.Body)).Rows()
		if err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}
		logrus.Debugf("%#v", rows)
		logs := make([]interface{}, 0)
		for rows.Next() {
			param := ""
			count := int64(0)
			rows.Scan(&param, &count)
			logs = append(logs, model.AuthUserLog{
				Parameter: null.StringFrom(param),
				Count:     count,
			})
		}
		logrus.Debugf("%#v", logs)
		return reply.ModelsBytes(logs)
	})
}
