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

func saveMenu(c *gin.Context, tx *gorm.DB) {
	menu := model.Menu{}
	c.ShouldBindJSON(&menu)
	logrus.Debugf("%#v", menu)
	if err := tx.Create(&menu).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func getMenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	menu, err := dao.GetMenu(db.Gorm, id)
	switch err {
	case dao.ErrNotFound:
		panic(reply.DataNoExist)
	case nil:
	default:
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONModel(c, menu)
}

func getMenuList(c *gin.Context) {
	menu := model.Menu{}
	c.ShouldBindQuery(&menu)
	logrus.Debugf("%#v", menu)
	page := pageutils.CheckAndInitPage(menu.BaseVO)

	menus := []model.Menu{}
	if err := db.Gorm.Model(&model.Menu{}).
		Scopes(dao.MenuCond(&menu)).
		Count(&page.Total).
		Scopes(dao.MenuCond(&menu), dao.Paginate(page)).
		Find(&menus).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	menusInterface := model.MenuToInterface(menus)
	reply.CreateJSONPaging(c, menusInterface, page)
}

func getFrontMenuList(c *gin.Context) {
	menus := []model.Menu{}
	if err := db.Gorm.Model(&model.Menu{}).Where("parent_id is null").Find(&menus).Error; err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}

	for i := range menus {
		menus[i].Children = []model.Menu{}
		if err := db.Gorm.Model(&menus[i]).Association("Children").Find(&menus[i].Children); err != nil {
			logrus.Debugf("%s", err.Error())
			panic(reply.DatabaseSqlParseError)
		}
		logrus.Debugln(menus[i].Children)
	}
	menusInterface := model.MenuToInterface(menus)
	reply.CreateJSONModels(c, menusInterface)
}

func updateMenu(c *gin.Context, tx *gorm.DB) {
	menu := model.Menu{}
	c.ShouldBindJSON(&menu)
	logrus.Debugf("%#v", menu)
	if _, _, err := dao.UpdateMenu(tx, menu.ID, &menu); err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func deleteMenu(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Debugln(err)
		panic(reply.ParamError)
	}

	if _, err := dao.DeleteMenu(tx, id); err != nil {
		logrus.Debugln(err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}
