package service

import (
	"net/http"
	"strconv"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	if db.DB == nil {
		c.JSON(http.StatusOK, map[string]bool{"health": false})
	}
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

func getUserInfo(c *gin.Context) {
	value, _ := c.Get("session")
	user, _ := value.(*po.AuthUser)
	log.Logger.Debugf("user = %#v\n", user)
	userVO := &vo.AuthUser{}
	userVO.Status = user.Status
	userVO.Roles = []string{constants.Roles[user.RoleId]}
	userVO.Name = user.Name
	userVO.CreateTime = &models.JSONTime{user.CreateTime}
	userVO.Introduction = user.Introduction
	userVO.Avatar = user.Avatar
	userVO.Email = user.Email
	c.JSON(http.StatusOK, reply.CreateWithModel(userVO))
}

func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}
	user := &po.AuthUser{Id: id}
	ok, err := db.DB.Get(user)
	if !ok || err != nil {
		panic(reply.AccountNotExist)
	}
	if user.RoleId != constants.RoleAdmin {
		db.DB.Delete(user)
		c.JSON(http.StatusOK, reply.CreateWithSuccess())
		c.Abort()
	}
	panic(reply.Error)
}

func saveAuthUserStatus(c *gin.Context) {
	userVO := &vo.AuthUser{}
	err := c.ShouldBindJSON(userVO)
	if err != nil {
		panic(reply.ParamError)
	}

	log.Logger.Debugf("user = %#v\n", userVO)

	//FIXME:无法区别status是否设置
	if userVO.Id != 0 {
		count, err := db.DB.Table(new(po.AuthUser)).ID(userVO.Id).
			Where("role_id = ?", constants.RoleUser).
			Update(map[string]interface{}{"status": userVO.Status})
		log.Logger.Debugf("count = %v\n", count)
		log.Logger.Debugf("err = %v\n", err)
		if err == nil {
			c.JSON(http.StatusOK, reply.CreateWithSuccess())
		}
	}
	panic(reply.Error)
}

func getMasterUserInfo(c *gin.Context) {
	user := &po.AuthUser{
		RoleId: constants.RoleAdmin,
	}
	ok, err := db.DB.Get(user)
	if !ok || err != nil {
		panic(reply.Error)
	}
	userVO := &vo.AuthUser{
		Name:         user.Name,
		Introduction: user.Introduction,
		Email:        user.Email,
		Avatar:       user.Avatar,
	}
	c.JSON(http.StatusOK, reply.CreateWithModel(userVO))
}

func getUserList(c *gin.Context) {
	userVO := new(vo.AuthUser)
	err := c.ShouldBindQuery(userVO)
	if err != nil {
		panic(reply.ParamError)
	}
}
