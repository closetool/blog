package vo

import (
	"github.com/closetool/blog/system/models"
	"github.com/dgrijalva/jwt-go"
)

type AuthUser struct {
	Id           int64            `json:"id,omitempty" form:"id"`
	SocialId     string           `json:"socialId,omitempty" form:"socialId"`
	Password     string           `json:"password,omitempty" form:"password"`
	PasswordOld  string           `json:"passwordOld,omitempty" form:"passwordOld"`
	Name         string           `json:"name,omitempty" form:"name"`
	RoleId       int64            `json:"roleId,omitempty" form:"roleId"`
	Email        string           `json:"email,omitempty" form:"email"`
	Introduction string           `json:"introduction,omitempty" form:"introduction"`
	Avatar       string           `json:"avatar,omitempty" form:"avatar"`
	CreateTime   *models.JSONTime `json:"createTime,omitempty" form:"createTime"`
	AccessKey    string           `json:"accessKey,omitempty" form:"accessKey"`
	SecretKey    string           `json:"secretKey,omitempty" form:"secretKey"`
	Status       int              `json:"status,omitempty" form:"status"`
	Roles        []string         `json:"roles,omitempty" form:"roles"`
	Token        string           `json:"token,omitempty" form:"token"`
	VerifyCode   string           `json:"verifyCode,omitempty" form:"verifyCode"`
	*jwt.StandardClaims
	*models.BaseVO
}

func CreateDefaultAuthUser() *AuthUser {
	return &AuthUser{
		Id:     -1,
		RoleId: -1,
		Status: -1,
	}
}
