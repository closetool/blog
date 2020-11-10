package service

import (
	"net/http"
	"strconv"

	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func health(c *gin.Context) {
	if db.Gorm != nil {
		c.JSON(http.StatusOK, map[string]bool{"health": true})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"health": false})
}

func getFriendshipLinkList(c *gin.Context) {
	link := model.FriendshipLink{}
	c.ShouldBindQuery(&link)
	page := pageutils.CheckAndInitPage(link.BaseVO)

	links := []model.FriendshipLink{}
	if err := db.Gorm.Model(&model.FriendshipLink{}).
		Scopes(dao.LinkCond(&link)).
		Count(&page.Total).
		Scopes(dao.Paginate(page)).
		Find(&links).Error; err != nil {
		panic(reply.DatabaseSqlParseError)
	}
	ints := model.Links2Interfaces(links)
	reply.CreateJSONPaging(c, ints, page)
}

func getFriendshipLinkMap(c *gin.Context) {
	link := model.FriendshipLink{}
	c.ShouldBindQuery(&link)

	links := []model.FriendshipLink{}
	if err := db.Gorm.Model(&model.FriendshipLink{}).
		Scopes(dao.LinkCond(&link)).
		Find(&links).Error; err != nil {
		panic(reply.DatabaseSqlParseError)
	}

	linkMap := map[string]model.FriendshipLink{}
	for _, link := range links {
		if link.Title.Valid {
			linkMap[link.Title.String] = link
		}
	}
	reply.CreateJSONExtra(c, linkMap)
}

func saveFriendshipLink(c *gin.Context, tx *gorm.DB) {
	link := model.FriendshipLink{}
	err := c.ShouldBindJSON(&link)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	if _, _, err := dao.AddFriendshipLink(tx, &link); err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONsuccess(c)
}

func updateFriendshipLink(c *gin.Context, tx *gorm.DB) {
	link := model.FriendshipLink{}
	err := c.ShouldBindJSON(&link)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	if _, _, err := dao.UpdateFriendshipLink(tx, link.ID, &link); err != nil {
		switch err {
		case dao.ErrNotFound:
			panic(reply.DataNoExist)
		default:
			panic(reply.DatabaseSqlParseError)
		}
	}
	reply.CreateJSONsuccess(c)
}

func getFriendshipLink(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	link, err := dao.GetFriendshipLink(db.Gorm, id)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.DataNoExist)
	}
	reply.CreateJSONModel(c, link)
}

func deleteFriendshipLink(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	if _, err = dao.DeleteFriendshipLink(tx, id); err != nil {
		logrus.Debug(err)
		switch err {
		case dao.ErrNotFound:
			panic(reply.DataNoExist)
		default:
			panic(reply.DatabaseSqlParseError)
		}
	}
	reply.CreateJSONsuccess(c)
}
