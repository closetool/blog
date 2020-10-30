package amqp

import (
	"github.com/closetool/blog/services/categoryservice/models/po"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetCategoryNameById() {
	messaging.Client.SubscribeToQueueAndReply("category.getCategoryNameById", "category.getCategoryNameById", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		jsoniter.Get(d.Body).ToVal(&ids)

		idsInterface := make([]interface{}, 0)
		for _, id := range ids {
			idsInterface = append(idsInterface, id)
		}

		categorys := make([]*po.Category, 0)
		if err := db.DB.In("id", idsInterface...).Find(&categorys); err != nil {
			logrus.Debugln(err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		result := make(map[int64]string)
		for _, category := range categorys {
			result[category.Id] = category.Name
		}
		logrus.Debugln(result)
		return reply.ModelBytes(result)
	})
}
