package vo

import (
	"github.com/closetool/blog/system/models"
	"github.com/dgrijalva/jwt-go"
)

type AuthUser struct {
	Id           int64            `json:"id,omitempty"`
	SocialId     string           `json:"socialId,omitempty"`
	Password     string           `json:"password,omitempty"`
	Name         string           `json:"name,omitempty"`
	RoleId       int64            `json:"roleId,omitempty"`
	Email        string           `json:"email,omitempty"`
	Introduction string           `json:"introduction,omitempty"`
	Avatar       string           `json:"avatar,omitempty"`
	CreateTime   *models.JSONTime `json:"createTime,omitempty"`
	AccessKey    string           `json:"accessKey,omitempty"`
	SecretKey    string           `json:"secretKey,omitempty"`
	Status       int              `json:"status,omitempty"`
	Roles        []string         `json:"roles,omitempty"`
	Token        string           `json:"token,omitempty"`
	*jwt.StandardClaims
	*models.BaseVO
}
