package service

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/collectionsutils"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//func Health(c *gin.Context) {
//	if db.DB == nil {
//		c.JSON(http.StatusOK, map[string]bool{"health": false})
//	}
//	c.JSON(http.StatusOK, map[string]bool{"health": true})
//}

func Health(c *gin.Context) {
	if db.Gorm == nil {
		c.JSON(http.StatusOK, map[string]bool{"health": false})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

func getUserInfo(c *gin.Context) {
	value, _ := c.Get("session")
	user, _ := value.(*model.AuthUser)
	logrus.Debugf("user = %#v", user)
	user.Password = ""
	user.PasswordOld = ""
	user.Roles = []string{constants.Roles[user.RoleID]}
	reply.CreateJSONModel(c, user)
}

func deleteUser(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("id converting failed: %v", err)
		panic(reply.ParamError)
	}

	user, err := dao.GetAuthUser(tx, id)
	if err != nil {
		panic(reply.AccountNotExist)
	}

	if user.RoleID != constants.RoleAdmin {
		_, err := dao.DeleteAuthUser(tx, id)
		if err != nil {
			panic(reply.DatabaseSqlParseError)
		}
		reply.CreateJSONsuccess(c)
		return
	}
	reply.CreateJSONError(c, reply.Error)
}

func saveAuthUserStatus(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	logrus.Debugf("user = %#v", user)

	//将AuthUser中的数字属性默认值设置为-1
	//避免默认值和真实值相冲突
	if user.ID != 0 && user.Status.Valid {
		if err := tx.Where("id = ? and role_id = ?", user.ID, constants.RoleUser).Update("status", user.Status).Error; err != nil {
			logrus.Errorf("update db failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
		reply.CreateJSONsuccess(c)
		return
	}
	panic(reply.ParamError)
}

func getMasterUserInfo(c *gin.Context) {
	admin := model.AuthUser{}
	if session := db.Gorm.First(&admin, "role_id = ?", constants.RoleAdmin); session.Error != nil {
		logrus.Errorf("find db failed: %v", session.Error)
		panic(reply.DatabaseSqlParseError)
	} else if session.RowsAffected == 0 {
		panic(reply.AccountNotExist)
	}
	reply.CreateJSONModel(c, admin)
}

func getUserList(c *gin.Context) {
	user := model.AuthUser{}
	if err := c.ShouldBindQuery(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}
	page := pageutils.CheckAndInitPage(user.BaseVO)
	logrus.Debugf("userVO = %#v", user)
	logrus.Debugf("page = %#v", page)

	users := make([]model.AuthUser, 0)

	if err := db.Gorm.
		Scopes(dao.UserCond(&user)).
		Count(&page.Total).
		Scopes(dao.Paginate(page)).
		Find(&users).Error; err != nil {
		panic(reply.DatabaseSqlParseError)
	}

	for i := range users {
		users[i].Password = ""
	}
	ints := model.Users2Interfaces(users)
	reply.CreateJSONPaging(c, ints, page)
}

func oathLoginByGithub(c *gin.Context) {
	url := viper.GetString("github_auth_url")
	logrus.Debugf("github auth url = %v", url)
	c.PureJSON(http.StatusOK, reply.CreateWithModel(map[string]string{"authorizeUrl": url}))
}

func saveUserByGithub(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	if err := tx.First(&user, "social_id = ?", user.SocialID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//数据库中不存在，插入相应数据
			hs := hmac.New(md5.New, collectionsutils.RandomString(32))
			user.Password = hex.EncodeToString(hs.Sum([]byte(user.SocialID.String)))
			user.RoleID = constants.RoleUser
			if err := tx.Create(&user).Error; err != nil {
				logrus.Errorf("can not insert a user: %v", user)
				panic(reply.DatabaseSqlParseError)
			}
		} else {
			logrus.Errorf("find user from db failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	} else {
		//数据库中存在，应该检查status判断是否能登录
		if user.Status.Int64 == constants.AccountLocked {
			logrus.Infof("account %s has been locked", user.SocialID.String)
			panic(reply.LoginDisable)
		}
	}

	token, err, expire := utils.GenerateToken(&user)
	if err != nil {
		logrus.Errorf("generate token failed: %v", err)
		panic(reply.Error)
	}
	user.Token = token

	//TODO:将修改存入数据库改为存入redis
	userToken := model.AuthToken{
		UserID:     user.ID,
		Token:      token,
		ExpireTime: time.Unix(expire, 0),
	}
	if _, _, err := dao.AddAuthToken(tx, &userToken); err != nil {
		logrus.Errorf("insert token failed: %v", err)
		panic(reply.Error)
	}
	reply.CreateJSONModel(c, user)
}

func registerAdminByGithub(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	admin := model.AuthUser{}

	if err := tx.First(&admin, "role_id = ?", constants.RoleAdmin).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Errorf("get admin account failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	} else {
		panic(reply.AccountExist)
	}

	passwdHash := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))

	user.Name = user.Email
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(passwdHash)))
	user.RoleID = constants.RoleAdmin

	if _, _, err := dao.AddAuthUser(tx, &user); err != nil {
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func login(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	admin := model.AuthUser{}

	if err := tx.Where("role_id = ? and email = ?", constants.RoleAdmin, user.Email).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(reply.AccountNotExist)
		}
		logrus.Errorf("get admin failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	psw := fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	if strings.EqualFold(admin.Password, psw) {
		token, err, expire := utils.GenerateToken(&admin)
		if err != nil {
			logrus.Errorf("generate token failed: %v", err)
			panic(reply.Error)
		}

		user.Roles = []string{constants.Roles[admin.RoleID]}
		user.Token = token

		//TODO:将修改存入数据库改为存入redis
		userToken := model.AuthToken{
			UserID:     admin.ID,
			Token:      token,
			ExpireTime: time.Unix(expire, 0),
		}
		if _, _, err := dao.AddAuthToken(tx, &userToken); err != nil {
			logrus.Errorf("save token failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
		reply.CreateJSONModel(c, user)
	} else {
		panic(reply.PasswordError)
	}
}

func updatePassword(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	s, _ := c.Get("session")
	admin, ok := s.(*model.AuthUser)
	if !ok {
		panic(reply.Error)
	}

	psw := fmt.Sprintf("%x", md5.Sum([]byte(user.PasswordOld)))

	if !strings.EqualFold(admin.Password, psw) {
		panic(reply.PasswordError)
	}

	if err := tx.Model(&model.AuthUser{}).Update("password", fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))); err != nil {
		logrus.Errorf("can not update password: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func updateAdmin(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	s, exist := c.Get("session")
	admin, ok := s.(*model.AuthUser)
	if !exist || !ok {
		panic(reply.InvalidToken)
	}

	m := map[string]interface{}{}
	if user.Email.Valid {
		m["email"] = user.Email.String
	}
	if user.Avatar.Valid {
		m["avatar"] = user.Avatar.String
	}
	if user.Name.Valid {
		m["name"] = user.Name.String
	}
	if user.Introduction.Valid {
		m["name"] = user.Introduction.String
	}

	if err := tx.Model(&user).Where("id = ?", admin.ID).Updates(m).Error; err != nil {
		logrus.Errorf("update admin profile failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONsuccess(c)
}

func updateUser(c *gin.Context, tx *gorm.DB) {
	user := model.AuthUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	if user.ID == 0 {
		panic(reply.ParamError)
	}

	m := map[string]interface{}{}
	if user.Email.Valid {
		m["email"] = user.Email.String
	}
	if user.Avatar.Valid {
		m["avatar"] = user.Avatar.String
	}
	if user.Name.Valid {
		m["name"] = user.Name.String
	}
	if user.Introduction.Valid {
		m["name"] = user.Introduction.String
	}

	if user.Status.Valid {
		m["status"] = user.Status.Int64
	}

	if err := tx.Model(&model.AuthUser{}).Where("id = ?", user.ID).Updates(m).Error; err != nil {
		logrus.Errorf("update user failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONsuccess(c)
}

func logout(c *gin.Context, tx *gorm.DB) {
	//TODO:修改为删除redis缓存
	reply.CreateJSONsuccess(c)
}

func getAvatar(c *gin.Context) {
	user := model.AuthUser{}

	if err := db.Gorm.Where("role_id = ?", constants.RoleAdmin).First(&user).Error; err != nil {
		logrus.Errorf("can not get admin: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	if !user.Avatar.Valid {
		c.Data(http.StatusOK, "image/jpg", []byte{})
		return
	}

	resp, err := http.Get(user.Avatar.String)
	if err != nil {
		panic(reply.Error)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(reply.Error)
	}
	c.Data(http.StatusOK, resp.Header.Get("Content-Type"), bytes)
}

func saveSocial(c *gin.Context, tx *gorm.DB) {
	social := model.AuthUserSocial{}
	if err := c.ShouldBindJSON(&social); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}
	logrus.Debugf("social = %#v", social)

	if social.Code == "" {
		panic(reply.ParamError)
	}

	if _, _, err := dao.AddAuthUserSocial(tx, &social); err != nil {
		logrus.Errorf("can not add social: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func editSocial(c *gin.Context, tx *gorm.DB) {
	social := model.AuthUserSocial{}
	if err := c.ShouldBindJSON(&social); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	if social.ID == 0 {
		panic(reply.ParamError)
	}

	if _, _, err := dao.UpdateAuthUserSocial(tx, social.ID, &social); err != nil {
		logrus.Errorf("can not update social: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func getSocial(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("id converting failed: %v", err)
		panic(reply.ParamError)
	}

	social := model.AuthUserSocial{}
	if err := db.Gorm.First(&social, id).Error; err != nil {
		logrus.Errorf("can not get social: %v", social)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONModel(c, social)
	return
}

func delSocial(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("id converting failed: %v", err)
		panic(reply.ParamError)
	}

	if _, err := dao.DeleteAuthUserSocial(tx, id); err != nil {
		logrus.Errorf("can not delete social: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func getSocialList(c *gin.Context) {
	socialList(c, -1)
}

func getSocialEnableList(c *gin.Context) {
	socialList(c, 1)
}

func socialList(c *gin.Context, enabled int32) {
	social := model.AuthUserSocial{IsEnabled: -1}
	if err := c.ShouldBindQuery(&social); err != nil {
		logrus.Errorf("can not parse params: %v", err)
		panic(reply.ParamError)
	}

	if enabled == 1 {
		social.IsEnabled = enabled
	}

	page := pageutils.CheckAndInitPage(social.BaseVO)

	socials := make([]model.AuthUserSocial, 0)

	if err := db.Gorm.Scopes(dao.SocialCond(&social)).Count(&page.Total).Scopes(dao.Paginate(page)).Find(&socials).Error; err != nil {
		logrus.Errorf("can not get social list: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	ints := model.Socials2Interfaces(socials)

	reply.CreateJSONPaging(c, ints, page)
}

func getSocialInfo(c *gin.Context) {
	socials := make([]model.AuthUserSocial, 0)

	if err := db.Gorm.Where("is_enabled = ? and is_home = ?", constants.SocialEnabled, constants.SocialIsHome).Find(&socials).Error; err != nil {
		logrus.Errorf("can not get socials: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	ints := model.Socials2Interfaces(socials)
	reply.CreateJSONModels(c, ints)
}

func sendEmail(c *gin.Context) {
	reply.CreateJSONsuccess(c)
}
