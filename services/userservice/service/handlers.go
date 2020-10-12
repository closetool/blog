package service

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/collectionsutils"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func deleteUser(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(reply.ParamError)
	}
	user := &po.AuthUser{Id: id}
	ok, err := db.DB.Get(user)
	if !ok || err != nil {
		reply.CreateJSONError(c, reply.AccountNotExist)
		return err
	}

	if user.RoleId != constants.RoleAdmin {
		_, err := db.DB.Delete(user)
		c.JSON(http.StatusOK, reply.CreateWithSuccess())
		return err
	}
	reply.CreateJSONError(c, reply.Error)
	return err
}

func saveAuthUserStatus(c *gin.Context) error {
	userVO := vo.CreateDefaultAuthUser()
	err := c.ShouldBindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	log.Logger.Debugf("user = %#v\n", userVO)

	//将AuthUser中的数字属性默认值设置为-1
	//避免默认值和真实值相冲突
	if userVO.Id != -1 && userVO.Status != -1 {
		count, err := db.DB.Table(new(po.AuthUser)).ID(userVO.Id).
			Where("role_id = ?", constants.RoleUser).
			Update(map[string]interface{}{"status": userVO.Status})
		log.Logger.Debugf("count = %v\n", count)
		log.Logger.Debugf("err = %v\n", err)
		if err == nil {
			c.JSON(http.StatusOK, reply.CreateWithSuccess())
			return nil
		}
	}
	reply.CreateJSONError(c, reply.Error)
	return err
}

func getMasterUserInfo(c *gin.Context) {
	user := &po.AuthUser{
		RoleId: constants.RoleAdmin,
	}
	ok, err := db.DB.Get(user)
	if !ok || err != nil {
		reply.CreateJSONError(c, reply.AccountNotExist)
		return
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
	userVO := vo.CreateDefaultAuthUser()
	err := c.ShouldBindQuery(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}
	page := pageutils.CheckAndInitPage(userVO.BaseVO)
	log.Logger.Debugf("userVO = %#v\n", userVO)
	log.Logger.Debugf("page = %#v\n", page)

	session := db.DB.Table(new(po.AuthUser))
	if userVO.BaseVO != nil && userVO.Keywords != "" {
		log.Logger.Debugf("keywords = %s\n", userVO.Keywords)
		session = session.Where("name like ?", "%"+userVO.Keywords+"%")
	}
	if userVO.Name != "" {
		session = session.Where("name = ?", userVO.Name)
	}
	if userVO.Status != -1 {
		session = session.Where("status = ?", userVO.Status)
	}
	session.Limit(pageutils.StartAndEnd(page))

	users := make([]*po.AuthUser, 0)
	count, err := session.FindAndCount(&users)
	if err != nil {
		log.Logger.Errorf("when selecting in database, an error occurred: %v\n", err)
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	page.Total = count

	userVOs := make([]interface{}, 0)
	for _, user := range users {
		tmp := &vo.AuthUser{
			Id:           user.Id,
			Status:       user.Status,
			Name:         user.Name,
			RoleId:       user.RoleId,
			Introduction: user.Introduction,
		}
		userVOs = append(userVOs, tmp)
	}
	c.JSON(http.StatusOK, reply.CreateWithPaging(userVOs, page))
}

func oathLoginByGithub(c *gin.Context) {
	url := viper.GetString("github_auth_url")
	log.Logger.Debugf("github auth url = %v\n", url)
	c.PureJSON(http.StatusOK, reply.CreateWithModel(map[string]string{"authorizeUrl": url}))
}

func saveUserByGithub(c *gin.Context) error {
	userVO := &vo.AuthUser{}
	err := c.ShouldBindJSON(&userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("save user by github parameters error: %v", err)
		return err
	}

	user := new(po.AuthUser)
	ok, err := db.DB.Where("social_id = ?", userVO.SocialId).Get(user)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("get user by social_id failed: %v", err)
		return err
	}
	hs := hmac.New(md5.New, collectionsutils.RandomBytes(32))
	if !ok {
		userPO := po.AuthUser{
			SocialId: userVO.SocialId,
			Avatar:   userVO.Avatar,
			Name:     userVO.Name,
			RoleId:   constants.RoleUser,
			Password: hex.EncodeToString(hs.Sum([]byte(userVO.SocialId))),
		}
		_, err := db.DB.InsertOne(&userPO)
		if err != nil {
			reply.CreateJSONError(c, reply.Error)
			log.Logger.Errorf("insert into %s failed: %v", userPO.TableName(), err)
			return err
		}

		user.Name = userPO.Name
		user.Password = userPO.Password
		user.Id = userPO.Id
		user.CreateTime = userPO.CreateTime

	} else {
		if user.Status == constants.AccountLocked {
			reply.CreateJSONError(c, reply.LoginDisable)
			return nil
		}
	}

	token, err, expire := utils.GenerateToken(user)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		log.Logger.Errorf("generate token failed: %v", err)
	}

	userVO.CreateTime = &models.JSONTime{user.CreateTime}
	userVO.Token = token

	//TODO:将修改存入数据库改为存入redis
	userToken := &po.AuthToken{
		UserId:     user.Id,
		Token:      token,
		ExpireTime: time.Unix(expire, 0),
	}
	_, err = db.DB.InsertOne(userToken)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		log.Logger.Errorf("Insert token failed: %v", err)
		return err
	}
	reply.CreateJSONModel(c, userVO)
	return nil
}

func registerAdminByGithub(c *gin.Context) error {
	userVO := vo.CreateDefaultAuthUser()
	err := c.BindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("could not bind parameters: %v\n", err)
		return err
	}

	admin := &po.AuthUser{}
	ok, err := db.DB.Where("role_id = ?", constants.RoleAdmin).Get(admin)
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		log.Logger.Errorf("select admin from db failed: %v\n", err)
		return err
	}

	if !ok {
		userPO := &po.AuthUser{
			Name:     userVO.Email,
			Email:    userVO.Email,
			RoleId:   constants.RoleAdmin,
			Password: fmt.Sprintf("%x", md5.Sum([]byte(userVO.Password))),
		}
		_, err := db.DB.InsertOne(userPO)
		if err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			log.Logger.Errorf("insert admin from db failed: %v\n", err)
			return err
		}
	} else {
		reply.CreateJSONError(c, reply.AccountExist)
		return nil
	}

	reply.CreateJSONsuccess(c)
	return nil
}

func login(c *gin.Context) {
	userVO := &vo.AuthUser{}
	err := c.BindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("could not bind parameters: %v\n", err)
		return
	}

	admin := &po.AuthUser{}
	ok, err := db.DB.Where("role_id=? and email = ?", constants.RoleAdmin, userVO.Email).
		Get(admin)

	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		log.Logger.Errorf("select from database failed: %v\n", err)
		return
	}

	if !ok {
		reply.CreateJSONError(c, reply.AccountNotExist)
		return
	} else {
		psw := fmt.Sprintf("%x", md5.Sum([]byte(userVO.Password)))
		if strings.EqualFold(admin.Password, psw) {
			token, err, expire := utils.GenerateToken(admin)
			if err != nil {
				reply.CreateJSONError(c, reply.Error)
				log.Logger.Errorf("generate token failed: %v", err)
			}

			userVO.Roles = []string{constants.Roles[admin.RoleId]}
			userVO.Token = token

			//TODO:将修改存入数据库改为存入redis
			userToken := &po.AuthToken{
				UserId:     admin.Id,
				Token:      token,
				ExpireTime: time.Unix(expire, 0),
			}
			_, err = db.DB.InsertOne(userToken)
			if err != nil {
				reply.CreateJSONError(c, reply.Error)
				log.Logger.Errorf("Insert token failed: %v", err)
				return
			}
		} else {
			reply.CreateJSONError(c, reply.PasswordError)
			return
		}
	}

	reply.CreateJSONModel(c, userVO)

}

func updatePassword(c *gin.Context) error {
	userVO := &vo.AuthUser{}
	err := c.BindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("could not bind parameters: %v\n", err)
		return err
	}

	session, _ := c.Get("session")
	admin, ok := session.(*po.AuthUser)
	if !ok {
		reply.CreateJSONError(c, reply.Error)
	}

	psw := fmt.Sprintf("%x", md5.Sum([]byte(userVO.PasswordOld)))

	if !strings.EqualFold(admin.Password, psw) {
		reply.CreateJSONError(c, reply.PasswordError)
		return nil
	}

	_, err = db.DB.Table(&po.AuthUser{}).ID(admin.Id).
		Update(map[string]string{
			"password": fmt.Sprintf("%x", md5.Sum([]byte(userVO.Password))),
		})

	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
	} else {
		reply.CreateJSONsuccess(c)
	}
	return err
}

func updateAdmin(c *gin.Context) error {
	userVO := &vo.AuthUser{}
	err := c.BindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("could not bind parameters: %v\n", err)
		return err
	}

	sessionInter, exist := c.Get("session")
	admin, ok := sessionInter.(*po.AuthUser)
	if !exist || !ok {
		reply.CreateJSONError(c, reply.AccountNotExist)
		return nil
	}

	userPO := &po.AuthUser{
		Email:        userVO.Email,
		Avatar:       userVO.Avatar,
		Name:         userVO.Name,
		Introduction: userVO.Introduction,
	}

	if count, err := db.DB.ID(admin.Id).
		Cols("email", "avatar", "name", "introduction").
		Update(userPO); err != nil || count == 0 {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func updateUser(c *gin.Context) error {
	userVO := &vo.AuthUser{}
	if err := contextBindAuthUser(c, userVO); err != nil {
		return err
	}

	if userVO.Id == 0 {
		reply.CreateJSONError(c, reply.ParamError)
		return nil
	}

	userPO := &po.AuthUser{
		Email:        userVO.Email,
		Avatar:       userVO.Avatar,
		Name:         userVO.Name,
		Introduction: userVO.Introduction,
		Status:       userVO.Status,
	}
	if count, err := db.DB.ID(userVO.Id).
		Cols("email", "avatar", "name", "introduction", "status").
		Update(userPO); count == 0 || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}

	reply.CreateJSONsuccess(c)
	return nil
}

func contextBindAuthUser(c *gin.Context, userVO *vo.AuthUser) error {
	err := c.BindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		log.Logger.Errorf("could not bind parameters: %v\n", err)
		return err
	}
	return nil
}

func logout(c *gin.Context) error {
	//TODO:修改为删除redis缓存
	return nil
}

func getAvatar(c *gin.Context) {
	userPO := *&po.AuthUser{}
	if count, err := db.DB.Where("role_id = ?", constants.RoleAdmin).
		Get(userPO); count == false || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	resp, err := http.Get(userPO.Avatar)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}

	c.Data(http.StatusOK, resp.Header.Get("Content-Type"), bytes)
}

func saveSocial(c *gin.Context) error {
	return nil
}
