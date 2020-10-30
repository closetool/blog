package amqp

import (
	"github.com/closetool/blog/services/categoryservice/models/po"
	"github.com/closetool/blog/services/categoryservice/models/vo"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetTagsByName() {
	messaging.Client.SubscribeToQueueAndReply("tags.getTagsByName", "tags.getTagsByName", func(d amqp.Delivery) []byte {
		tagName := string(d.Body)
		logrus.Debugln(tagName)
		tagsPO := &po.Tags{}
		if ok, err := db.DB.Where("name = ?", tagName).Get(tagsPO); err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		} else if !ok {
			return reply.ErrorBytes(reply.DataNoExist)
		}
		res := vo.Tags{
			Id:   tagsPO.Id,
			Name: tagsPO.Name,
		}
		return reply.ModelBytes(res)
	})
}

func GetTagsByIds() {
	messaging.Client.SubscribeToQueueAndReply("tags.getTagsByIds", "tags.getTagsByIds", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		logrus.Debugln(string(d.Body))
		jsoniter.Get(d.Body).ToVal(&ids)
		idsInterface := make([]interface{}, len(ids))
		for i, v := range ids {
			idsInterface[i] = v
		}

		tags := make([]po.Tags, 0)
		if err := db.DB.In("id", idsInterface...).Find(&tags); err != nil {
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}
		tagsInterface := make([]interface{}, 0)
		for _, v := range tags {
			tagsInterface = append(tagsInterface, vo.Tags{
				Id:   v.Id,
				Name: v.Name,
			})
		}
		return reply.ModelsBytes(tagsInterface)
	})
}
