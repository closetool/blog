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

func GetTagsIDAndCount() {
	messaging.Client.SubscribeToQueueAndReply("posts.tags.getTagsIDAndCount", "posts.tags.getTagsIDAndCount", func(d amqp.Delivery) []byte {
		rows, err := db.Gorm.Raw("select tags_id,count(*) as count from posts_tags group by tags_id").Rows()
		if err != nil {
			logrus.Errorf("find posts tags failed: %v", err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}
		res := make(map[int64]int64)

		for rows.Next() {
			var (
				count  int64
				tagsID int64
			)

			if err := rows.Scan(&tagsID, &count); err != nil {
				logrus.Errorf("scan params failed: %v", err)
				return reply.ErrorBytes(reply.DatabaseSqlParseError)
			}

			res[tagsID] = count
		}
		return reply.ModelBytes(res)
	})
}

func GetCategoryIDAndCount() {
	messaging.Client.SubscribeToQueueAndReply("posts.getCategoryIDAndCount", "posts.getCategoryIDAndCount", func(d amqp.Delivery) []byte {
		rows, err := db.Gorm.Raw("select category_id,count(*) count from posts group by category_id").Rows()
		if err != nil {
			logrus.Errorf("find categories failed: %v", err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		res := map[int64]int64{}

		for rows.Next() {
			var (
				categoryID int64
				count      int64
			)
			if err := rows.Scan(&categoryID, &count); err != nil {
				logrus.Errorf("scan params failed: %v", err)
				return reply.ErrorBytes(reply.DatabaseSqlParseError)
			}

			res[categoryID] = count
		}

		return reply.ModelBytes(res)
	})
}

func GetCategoryIDAndCountInUse() {
	messaging.Client.SubscribeToQueueAndReply("posts.getCategoryIDAndCountInUse", "posts.getCategoryIDAndCountInUse", func(d amqp.Delivery) []byte {
		rows, err := db.Gorm.Raw("select category_id,count(*) count from posts where status = 2 group by category_id").Rows()
		if err != nil {
			logrus.Errorf("find categories failed: %v", err)
			return reply.ErrorBytes(reply.DatabaseSqlParseError)
		}

		res := map[int64]int64{}

		for rows.Next() {
			var (
				categoryID int64
				count      int64
			)
			if err := rows.Scan(&categoryID, &count); err != nil {
				logrus.Errorf("scan params failed: %v", err)
				return reply.ErrorBytes(reply.DatabaseSqlParseError)
			}

			res[categoryID] = count
		}

		return reply.ModelBytes(res)
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
