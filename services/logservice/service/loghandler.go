package service

import (
	"net/http"
	"strconv"

	"github.com/closetool/blog/services/logservice/model/po"
	"github.com/closetool/blog/services/logservice/model/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	if db.DB == nil {
		c.JSON(http.StatusOK, map[string]bool{"health": false})
	}
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

func getLogs(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}

	logPO := &po.AuthUserLog{}
	_, err = db.DB.ID(id).Get(logPO)
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	logVO := &vo.AuthUserLog{
		Ip:             logPO.Ip,
		CodeName:       constants.OperationNames[logPO.Code],
		CreateTime:     &models.JSONTime{logPO.CreateTime},
		Parameter:      logPO.Parameter,
		UserId:         logPO.UserId,
		RunTime:        logPO.RunTime,
		BrowserName:    logPO.BrowserName,
		BrowserVersion: logPO.BrowserVersion,
		Device:         logPO.Device,
		Description:    logPO.Description,
		Url:            logPO.Url,
	}
	reply.CreateJSONModel(c, logVO)
}

func deleteLogs(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	db.DB.ID(id).Delete(&po.AuthUserLog{})
	return
}

func getLogsList(c *gin.Context) {
	logVO := &vo.AuthUserLog{}
	c.ShouldBindQuery(logVO)
	page := pageutils.CheckAndInitPage(logVO.BaseVO)

	session := db.DB.NewSession()
	if logVO.UserId != "" {
		session = session.Where("user_id = ?", logVO.UserId)
	}
	if logVO.Ip != "" {
		session = session.Where("ip = ?", logVO.Ip)
	}
	if logVO.Url != "" {
		session = session.Where("url like ?", "%"+logVO.Url+"%")
	}
	if logVO.Parameter != "" {
		session = session.Where("parameter like ?", "%"+logVO.Parameter+"%")
	}
	if logVO.Device != "" {
		session = session.Where("device like ?", "%"+logVO.Device+"%")
	}
	if logVO.Description != "" {
		session = session.Where("description like ?", "%"+logVO.Description+"%")
	}
	if logVO.Code != "" {
		session = session.Where("code = ?", logVO.Code)
	}
	if logVO.BrowserName != "" {
		session = session.Where("browset_name like ?", "%"+logVO.BrowserName+"%")
	}
	if logVO.BrowserVersion != "" {
		session = session.Where("browser_version = ?", logVO.BrowserVersion)
	}
	if logVO.CreateTime != nil {
		session = session.Where("DATE_FORMAT( ? ,'%Y-%m-%d')=DATE_FORMAT(create_time, '%Y-%m-%d')", logVO.CreateTime)
	}
	session = session.Limit(pageutils.StartAndEnd(page))
	var (
		err     error
		logsPOs = make([]po.AuthUserLog, 0)
	)
	if page.Total, err = session.FindAndCount(&logsPOs); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	logVOs := make([]interface{}, 0)
	for _, logPO := range logsPOs {
		logVO := &vo.AuthUserLog{
			Ip:             logPO.Ip,
			CodeName:       constants.OperationNames[logPO.Code],
			CreateTime:     &models.JSONTime{logPO.CreateTime},
			Parameter:      logPO.Parameter,
			UserId:         logPO.UserId,
			RunTime:        logPO.RunTime,
			BrowserName:    logPO.BrowserName,
			BrowserVersion: logPO.BrowserVersion,
			Device:         logPO.Device,
			Description:    logPO.Description,
			Url:            logPO.Url,
		}
		logVOs = append(logVOs, logVO)
	}
	reply.CreateJSONPaging(c, logVOs, page)
}
