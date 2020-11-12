package service

import (
	"bytes"
	"net/http"
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
	"gorm.io/gorm/clause"
)

func health(c *gin.Context) {
	if db.Gorm != nil {
		c.JSON(http.StatusOK, map[string]bool{"health": true})
	}
	c.JSON(http.StatusOK, map[string]bool{"health": false})
}

func saveCategory(c *gin.Context, tx *gorm.DB) {
	category := model.Category{}
	err := c.ShouldBindJSON(&category)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	logrus.Debugf("%#v", category)

	if tx.Where("name", category.Name).First(&category).RowsAffected == 0 {
		if err := tx.Create(&category).Error; err != nil {
			logrus.Debugln(err)
			panic(reply.DatabaseSqlParseError)
		}
	} else {
		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&category).Error; err != nil {
			logrus.Debugln(err)
			panic(reply.DatabaseSqlParseError)
		}
	}
	reply.CreateJSONsuccess(c)
}

//TODO:test
func statisticsList(c *gin.Context) {
	rpl, err := messaging.Client.PublishOnQueueWaitReply(nil, "posts.getCategoryIDAndCountInUse")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		panic(reply.ParamError)
	}

	IdAndCount := make(map[int64]int64)
	jsoniter.Get(rpl, "model").ToVal(&IdAndCount)

	categories := make([]model.Category, 0)
	if err := db.Gorm.Order("id").Find(&categories).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}

	for i, category := range categories {
		categories[i].Total = IdAndCount[category.ID]
	}

	reply.CreateJSONModels(c, model.Category2Interfaces(categories))
}

func updateCategory(c *gin.Context, tx *gorm.DB) {
	category := model.Category{}
	err := c.ShouldBindJSON(&category)
	if err != nil || category.ID == 0 {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	logrus.Debugf("%#v", category)

	if err := tx.Model(&category).Updates(&category).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}

	logrus.Debugf("%#v", category)

	if err := tx.Model(&category).Association("Tags").Replace(category.Tags); err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONsuccess(c)
}

func getCategoryTags(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}

	category := model.Category{}
	if err := db.Gorm.Preload(clause.Associations).Where("id = ?", id).First(&category).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONModel(c, category)
}

func getCategoryTagsList(c *gin.Context) {
	category := model.Category{}
	c.ShouldBindQuery(&category)
	page := pageutils.CheckAndInitPage(category.BaseVO)
	logrus.Debugln(page)

	categories := []model.Category{}
	if err := db.Gorm.Model(&category).
		Preload(clause.Associations).
		Scopes(dao.CategoryCond(&category)).
		Count(&page.Total).
		Scopes(dao.Paginate(page)).
		Find(&categories).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONPaging(c, model.Category2Interfaces(categories), page)
}

//TODO:test
func getCategoryList(c *gin.Context) {
	rpl, err := messaging.Client.PublishOnQueueWaitReply(nil, "posts.getCategoryIDAndCount")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		panic(reply.ParamError)
	}

	IdAndCount := make(map[int64]int64)
	jsoniter.Get(rpl, "model").ToVal(&IdAndCount)

	categories := make([]model.Category, 0)
	if err := db.Gorm.Order("id").Find(&categories).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}

	for i, category := range categories {
		categories[i].Total = IdAndCount[category.ID]
	}

	reply.CreateJSONModels(c, model.Category2Interfaces(categories))
}

func deleteCategory(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}

	if err := tx.Where("id = ?", id).Delete(&model.Category{ID: id}).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}
