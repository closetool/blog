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
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/collectionsutils"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	logrus.Debugf("user = %#v", user)
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
		reply.CreateJSONError(c, reply.ParamError)
		return err
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

	logrus.Debugf("user = %#v", userVO)

	//将AuthUser中的数字属性默认值设置为-1
	//避免默认值和真实值相冲突
	if userVO.Id != -1 && userVO.Status != -1 {
		count, err := db.DB.Table(new(po.AuthUser)).ID(userVO.Id).
			Where("role_id = ?", constants.RoleUser).
			Update(map[string]interface{}{"status": userVO.Status})
		logrus.Debugf("count = %v", count)
		logrus.Debugf("err = %v", err)
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
	logrus.Debugf("userVO = %#v", userVO)
	logrus.Debugf("page = %#v", page)

	session := db.DB.Table(new(po.AuthUser))
	if userVO.BaseVO != nil && userVO.Keywords != "" {
		logrus.Debugf("keywords = %s", userVO.Keywords)
		session = session.Where("name like ?", "%"+userVO.Keywords+"%")
	}
	if userVO.Name != "" {
		session = session.Where("name = ?", userVO.Name)
	}
	if userVO.Status != -1 {
		session = session.Where("status = ?", userVO.Status)
	}
	session = session.Limit(pageutils.StartAndEnd(page))

	users := make([]*po.AuthUser, 0)
	count, err := session.FindAndCount(&users)
	if err != nil {
		logrus.Errorf("when selecting in database, an error occurred: %v", err)
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
	logrus.Debugf("github auth url = %v", url)
	c.PureJSON(http.StatusOK, reply.CreateWithModel(map[string]string{"authorizeUrl": url}))
}

func saveUserByGithub(c *gin.Context) error {
	userVO := &vo.AuthUser{}
	err := c.ShouldBindJSON(&userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		logrus.Errorf("save user by github parameters error: %v", err)
		return err
	}

	user := new(po.AuthUser)
	ok, err := db.DB.Where("social_id = ?", userVO.SocialId).Get(user)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		logrus.Errorf("get user by social_id failed: %v", err)
		return err
	}
	hs := hmac.New(md5.New, collectionsutils.RandomString(32))
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
			logrus.Errorf("insert into %s failed: %v", userPO.TableName(), err)
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
		logrus.Errorf("generate token failed: %v", err)
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
		logrus.Errorf("Insert token failed: %v", err)
		return err
	}
	reply.CreateJSONModel(c, userVO)
	return nil
}

func registerAdminByGithub(c *gin.Context) error {
	logrus.Debugln("/auth/admin/v1/register was called")
	userVO := vo.CreateDefaultAuthUser()
	err := c.ShouldBindJSON(userVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		logrus.Errorf("could not bind parameters: %v", err)
		return err
	}

	admin := &po.AuthUser{}
	ok, err := db.DB.Where("role_id = ?", constants.RoleAdmin).Get(admin)
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		logrus.Errorf("select admin from db failed: %v", err)
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
			logrus.Errorf("insert admin from db failed: %v", err)
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
		logrus.Errorf("could not bind parameters: %v", err)
		return
	}

	admin := &po.AuthUser{}
	ok, err := db.DB.Where("role_id=? and email = ?", constants.RoleAdmin, userVO.Email).
		Get(admin)

	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		logrus.Errorf("select from database failed: %v", err)
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
				logrus.Errorf("generate token failed: %v", err)
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
				logrus.Errorf("Insert token failed: %v", err)
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
		logrus.Errorf("could not bind parameters: %v", err)
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
		logrus.Errorf("could not bind parameters: %v", err)
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
		logrus.Errorf("could not bind parameters: %v", err)
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
	socialVO := &vo.AuthUserSocial{}
	err := c.BindJSON(socialVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	logrus.Debugf("social = %#v", socialVO)

	if socialVO.Code == "" {
		reply.CreateJSONError(c, reply.ParamError)
		return nil
	}

	socialPO := &po.AuthUserSocial{
		Code:      socialVO.Code,
		ShowType:  socialVO.ShowType,
		Content:   socialVO.Content,
		Remark:    socialVO.Remark,
		Icon:      socialVO.Icon,
		IsEnabled: socialVO.IsEnabled,
		IsHome:    socialVO.IsHome,
	}

	if count, err := db.DB.InsertOne(socialPO); count == 0 || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		if err != nil {
			return err
		}
		return nil
	}

	reply.CreateJSONsuccess(c)
	return nil
}

func editSocial(c *gin.Context) error {
	socialVO := &vo.AuthUserSocial{}
	err := c.BindJSON(socialVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	if socialVO.Id == 0 {
		reply.CreateJSONError(c, reply.ParamError)
		return nil
	}

	socialPO := &po.AuthUserSocial{
		Code:      socialVO.Code,
		ShowType:  socialVO.ShowType,
		Content:   socialVO.Content,
		Remark:    socialVO.Remark,
		Icon:      socialVO.Icon,
		IsEnabled: socialVO.IsEnabled,
		IsHome:    socialVO.IsHome,
	}
	if count, err := db.DB.ID(socialVO.Id).AllCols().
		Update(socialPO); count == 0 || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		if err != nil {
			return err
		}
		return nil
	}

	reply.CreateJSONsuccess(c)
	return nil
}

func getSocial(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}

	socialPO := &po.AuthUserSocial{}
	ok, err := db.DB.ID(id).Get(socialPO)
	if !ok || err != nil {
		logrus.Debugf("selecting from db failed: %v", err)
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	social := &vo.AuthUserSocial{
		Id:         socialPO.Id,
		Code:       socialPO.Code,
		Content:    socialPO.Content,
		ShowType:   socialPO.ShowType,
		Remark:     socialPO.Remark,
		Icon:       socialPO.Icon,
		IsEnabled:  socialPO.IsEnabled,
		IsHome:     socialPO.IsHome,
		CreateTime: &models.JSONTime{socialPO.CreateTime},
		UpdateTime: &models.JSONTime{socialPO.UpdateTime},
	}

	reply.CreateJSONModel(c, social)
	return
}

func delSocial(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	if count, err := db.DB.ID(id).Delete(&po.AuthUserSocial{}); count == 0 || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}

	reply.CreateJSONsuccess(c)
	return nil
}

func getSocialList(c *gin.Context) {
	socialList(c, -1)
}

func getSocialEnableList(c *gin.Context) {
	socialList(c, 1)
}

func getSocialInfo(c *gin.Context) {
	socialVOs := make([]interface{}, 0)
	socialPOs := make([]*po.AuthUserSocial, 0)
	err := db.DB.Where("is_enabled = ? and is_home = ?", constants.SocialEnabled, constants.SocialIsHome).
		Find(&socialPOs)
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	for _, socialPO := range socialPOs {
		social := vo.AuthUserSocial{
			Id:         socialPO.Id,
			Code:       socialPO.Code,
			Content:    socialPO.Content,
			ShowType:   socialPO.ShowType,
			Remark:     socialPO.Remark,
			Icon:       socialPO.Icon,
			IsEnabled:  socialPO.IsEnabled,
			IsHome:     socialPO.IsHome,
			CreateTime: &models.JSONTime{socialPO.CreateTime},
			UpdateTime: &models.JSONTime{socialPO.UpdateTime},
		}
		socialVOs = append(socialVOs, social)
	}
	reply.CreateJSONModels(c, socialVOs)
}

func socialList(c *gin.Context, enabled int) {
	socialVO := &vo.AuthUserSocial{IsEnabled: -1, IsHome: -1}
	err := c.BindQuery(socialVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
	}

	if enabled == 1 {
		socialVO.IsEnabled = enabled
	}

	page := pageutils.CheckAndInitPage(socialVO.BaseVO)
	session := db.DB.NewSession()

	if socialVO.BaseVO != nil && socialVO.Keywords != "" {
		session = session.Where("code like ?", "%"+socialVO.Keywords+"%")
	}
	if socialVO.Code != "" {
		session = session.Where("code = ?", socialVO.Code)
	}
	if socialVO.Content != "" {
		session = session.Where("content = ?", socialVO.Content)
	}
	if socialVO.ShowType != 0 {
		session = session.Where("show_type = ?", socialVO.ShowType)
	}
	if socialVO.Remark != "" {
		session = session.Where("remark = ?", socialVO.Remark)
	}
	if socialVO.IsEnabled != -1 {
		session = session.Where("is_enabled = ?", socialVO.IsEnabled)
	}
	if socialVO.IsHome != -1 {
		session = session.Where("is_home = ?", socialVO.IsHome)
	}
	session = session.Limit(pageutils.StartAndEnd(page))
	session = session.OrderBy("id")

	lists := make([]po.AuthUserSocial, 0)
	if err = session.Find(&lists); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	results := make([]interface{}, 0)
	for _, socialPO := range lists {
		social := vo.AuthUserSocial{
			Id:         socialPO.Id,
			Code:       socialPO.Code,
			Content:    socialPO.Content,
			ShowType:   socialPO.ShowType,
			Remark:     socialPO.Remark,
			Icon:       socialPO.Icon,
			IsEnabled:  socialPO.IsEnabled,
			IsHome:     socialPO.IsHome,
			CreateTime: &models.JSONTime{socialPO.CreateTime},
			UpdateTime: &models.JSONTime{socialPO.UpdateTime},
		}
		results = append(results, social)
	}

	reply.CreateJSONPaging(c, results, page)
}
