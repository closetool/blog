package amqp

import (
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetCategoryNameById() {
	messaging.Client.SubscribeToQueueAndReply("category.getCategoryNameById", "category.getCategoryNameById", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		jsoniter.Get(d.Body).ToVal(&ids)

		categories := make([]model.Category, 0)

		if err := db.Gorm.Where("id in ?", ids).Find(&categories); err != nil {
			logrus.Debugln(err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}
		result := make(map[int64]string)
		for _, category := range categories {
			result[category.ID] = category.Name
		}
		logrus.Debugln(result)
		t := reply.ModelBytes(result)
		logrus.Debugln(string(t))
		return t
	})
}
