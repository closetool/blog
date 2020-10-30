package amqp

import (
	"strconv"

	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
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
