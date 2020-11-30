package service

import (
	"bytes"
	"strconv"

	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func getTagsList(c *gin.Context) {
	tag := model.Tags{}
	c.ShouldBindQuery(&tag)
	logrus.Debugf("tag = %#v", tag)
	tags := make([]model.Tags, 0)
	page := pageutils.CheckAndInitPage(tag.BaseVO)

	if err := db.Gorm.Model(&tag).Scopes(dao.TagsCond(&tag)).Count(&page.Total).Scopes(dao.Paginate(page)).Find((&tags)).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONPaging(c, model.Tags2Interfaces(tags), page)
}

//TODO: test
func getTagsAndArticleQuantityList(c *gin.Context) {
	tags := []model.Tags{}
	if err := db.Gorm.Find(&tags); err != nil {
		logrus.Errorf("could not select tags from db: %v", tags)
		panic(reply.DatabaseSqlParseError)
	}

	IDAndCount := make(map[int64]int64)
	rpl, err := messaging.Client.PublishOnQueueWaitReply([]byte{}, "posts.tags.getTagsIDAndCount")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		logrus.Errorf("can not get id mapped count: %v", err)
		panic(reply.Error)
	}

	for i, tag := range tags {
		tags[i].PostsTotal = IDAndCount[tag.ID]
	}
	reply.CreateJSONModels(c, model.Tags2Interfaces(tags))
}

func getTags(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}
	logrus.Debugf("id = %v", id)

	var tag *model.Tags
	if tag, err = dao.GetTags(db.Gorm, id); err != nil {
		logrus.Debugln(err)
		switch err {
		case dao.ErrNotFound:
			panic(reply.DataNoExist)
		default:
			panic(reply.DatabaseSqlParseError)
		}
	}
	reply.CreateJSONModel(c, tag)
}

func saveTags(c *gin.Context, tx *gorm.DB) {
	tag := model.Tags{}
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		panic(reply.ParamError)
	}
	count := int64(0)
	tx.Model(&model.Tags{}).Where("name = ?", tag.Name).Count(&count)
	if count == 0 {
		if _, _, err := dao.AddTags(tx, &tag); err != nil {
			logrus.Debugln(err)
			panic(reply.DatabaseSqlParseError)
		}
	} else {
		panic(reply.DataNoExist)
	}
	reply.CreateJSONsuccess(c)
}

func updateTags(c *gin.Context, tx *gorm.DB) {
	tag := model.Tags{}
	err := c.ShouldBindJSON(&tag)
	if err != nil || tag.ID == 0 {
		panic(reply.ParamError)
	}

	if _, _, err := dao.UpdateTags(tx, tag.ID, &tag); err != nil {
		logrus.Debugln(err)
		switch err {
		case dao.ErrNotFound:
			panic(reply.DataNoExist)
		default:
			panic(reply.DatabaseSqlParseError)
		}
	}
	reply.CreateJSONsuccess(c)
}

//TODO:test
func deleteTags(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}
	logrus.Debugf("id = %v", id)

	if _, err := dao.DeleteTags(tx, id); err != nil {
		logrus.Debugln(err)
		switch err {
		case dao.ErrNotFound:
			panic(reply.DataNoExist)
		default:
			panic(reply.DatabaseSqlParseError)
		}
	}
	bts, _ := jsoniter.Marshal(map[string]int64{"id": id})
	if err = messaging.Client.PublishOnQueue(bts, "posts.tags.deletePostsTagsById"); err != nil {
		logrus.Errorf("could not delete poststags: %v", err)
		panic(reply.Error)
	}
	reply.CreateJSONsuccess(c)
}
