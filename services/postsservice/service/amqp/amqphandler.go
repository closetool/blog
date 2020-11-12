package amqp

import (
	"strconv"

	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetTagsIDAndCount() {
	messaging.Client.SubscribeToQueueAndReply("posts.tags.getTagsIDAndCount", "posts.tags.getTagsIDAndCount", func(d amqp.Delivery) []byte {
		var bytes []byte = reply.ErrorBytes(reply.DatabaseSqlParseError)
		if res, err := db.DB.SQL("select tags_id,count(*) as count from closetool_posts_tags group by tags_id").QueryString(); err != nil {
			logrus.Debugln(err)
			return bytes
		} else {
			resMap := make(map[int64]int64)
			for _, r := range res {
				logrus.Debugln(r)
				tagsIdStr := r["tags_id"]
				countStr := r["count"]
				tagsId, err := strconv.ParseInt(tagsIdStr, 10, 64)
				if err != nil {
					return reply.ErrorBytes(reply.Error)
				}

				count, err := strconv.ParseInt(countStr, 10, 64)
				if err != nil {
					return reply.ErrorBytes(reply.Error)
				}
				resMap[tagsId] = count
			}
			bytes = reply.ModelBytes(resMap)
		}
		return bytes
	})
}

func GetCategoryIDAndCount() {
	messaging.Client.SubscribeToQueueAndReply("posts.getCategoryIDAndCount", "posts.getCategoryIDAndCount", func(d amqp.Delivery) []byte {
		var bytes []byte = reply.ErrorBytes(reply.DatabaseSqlParseError)
		if res, err := db.DB.SQL("select category_id,count(*) count from closetool_posts group by category_id").QueryString(); err != nil {
			logrus.Debugln(err)
			return bytes
		} else {
			resMap := make(map[int64]int64)
			for _, r := range res {
				categoryIdStr := r["category_id"]
				countStr := r["count"]
				categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
				if err != nil {
					return reply.ErrorBytes(reply.Error)
				}

				count, err := strconv.ParseInt(countStr, 10, 64)
				if err != nil {
					return reply.ErrorBytes(reply.Error)
				}
				resMap[categoryId] = count
			}
			bytes = reply.ModelBytes(resMap)
			logrus.Debugln(string(bytes))
		}
		return bytes
	})
}

func GetCategoryIDAndCountInUse() {
	messaging.Client.SubscribeToQueueAndReply("posts.getCategoryIDAndCountInUse", "posts.getCategoryIDAndCountInUse", func(d amqp.Delivery) []byte {
		var bytes []byte = reply.ErrorBytes(reply.DatabaseSqlParseError)
		if res, err := db.DB.SQL("select category_id,count(*) count from closetool_posts where status = 2 group by category_id").QueryString(); err != nil {
			logrus.Debugln(err)
			return bytes
		} else {
			resMap := make(map[int64]int64)
			for _, r := range res {
				categoryIdStr := r["category_id"]
				countStr := r["count"]
				categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
				if err != nil {
					return reply.ErrorBytes(reply.Error)
				}

				count, err := strconv.ParseInt(countStr, 10, 64)
				if err != nil {
					return reply.ErrorBytes(reply.Error)
				}
				resMap[categoryId] = count
			}
			bytes = reply.ModelBytes(resMap)
			logrus.Debugln(string(bytes))
		}
		return bytes
	})
}

func DeletePostsTagsById() {
	messaging.Client.SubscribeToQueue("posts.tags.deletePostsTagsById", "posts.tags.deletePostsTagsById", func(a amqp.Delivery) {
		id := map[string]int64{}
		if err := jsoniter.Unmarshal(a.Body, &id); err != nil {
			logrus.Errorf("param incorrect: %v", err)
			return
		}

		if err := db.Gorm.Where("tags_id = ?", id["id"]).Delete(&model.PostsTags{}).Error; err != nil {
			logrus.Errorf("could not delete from db: %v", err)
			return
		}
	})
}
