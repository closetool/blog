package amqp

import (
	"github.com/closetool/blog/services/categoryservice/models/vo"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetTagsByName() {
	messaging.Client.SubscribeToQueueAndReply("tags.getTagsByName", "tags.getTagsByName", func(d amqp.Delivery) []byte {
		tagName := string(d.Body)
		logrus.Debugln(tagName)
		tag := model.Tags{}

		count := int64(0)

		if err := db.Gorm.Where("name = ?", tagName).Count(&count).First(&tag).Error; err != nil {
			logrus.Errorf("select from db failed: %v", err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		if count != 1 {
			return reply.ErrorBytes(reply.DataNoExist)
		}

		res := vo.Tags{
			Id:   tag.ID,
			Name: tag.Name,
		}
		return reply.ModelBytes(res)
	})
}

func GetTagsByIds() {
	messaging.Client.SubscribeToQueueAndReply("tags.getTagsByIds", "tags.getTagsByIds", func(d amqp.Delivery) []byte {
		ids := make([]int64, 0)
		logrus.Debugln(string(d.Body))
		jsoniter.Get(d.Body).ToVal(&ids)

		tags := make([]model.Tags, 0)

		if err := db.Gorm.Where("id = ?", ids).Find(&tags).Error; err != nil {
			logrus.Errorf("select from db failed: %v", err)
		}
		return reply.ModelsBytes(model.Tags2Interfaces(tags))
	})
}

func AddTags() {
	messaging.Client.SubscribeToQueueAndReply("tags.addTags", "tags.addTags", func(d amqp.Delivery) []byte {
		tags := []*model.Tags{}
		logrus.Debugln(string(d.Body))
		if err := jsoniter.Unmarshal(d.Body, &tags); err != nil {
			logrus.Debugln(err)
			return reply.ErrorBytes(reply.Error)
		}
		ids := []interface{}{}
		tx := db.Gorm.Begin()
		for _, tag := range tags {
			if err := tx.Where("name = ?", tag.Name).FirstOrCreate(&tag).Error; err != nil {
				logrus.Errorf("could not create tag: %v", err)
				tx.Rollback()
				return reply.ErrorBytes(reply.Error)
			}
			ids = append(ids, tag.ID)
		}
		tx.Commit()
		return reply.ModelsBytes(ids)
	})
}
