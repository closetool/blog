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
	"gorm.io/gorm"
)

func Health(c *gin.Context) {
	if db.Gorm == nil {
		c.JSON(http.StatusOK, map[string]bool{"health": false})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

func getLogs(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}

	log, err := dao.GetAuthUserLog(db.Gorm, id)
	if err != nil {
		panic(reply.DataNoExist)
	}
	reply.CreateJSONModel(c, log)
}

func deleteLogs(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}
	_, err = dao.DeleteAuthUserLog(db.Gorm, id)
	switch err {
	case dao.ErrNotFound:
		panic(reply.DataNoExist)
	case dao.ErrDeleteFailed:
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func getLogsList(c *gin.Context) {
	log := model.AuthUserLog{}
	c.ShouldBindQuery(&log)
	page := pageutils.CheckAndInitPage(log.BaseVO)

	logs := make([]model.AuthUserLog, 0)
	if err := db.Gorm.Scopes(dao.LogsCond(&log)).
		Count(&page.Total).Scopes(dao.Paginate(page)).Find(&logs).Error; err != nil {
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONPaging(c, model.Logs2Intefaces(logs), page)
}
