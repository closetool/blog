package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UniversialAuthUser struct {
	Id           int64     `json:"id,omitempty"`
	SocialId     string    `json:"socialId,omitempty"`
	Password     string    `json:"-"`
	Name         string    `json:"name,omitempty"`
	RoleId       int64     `json:"roleId,omitempty"`
	Email        string    `json:"email,omitempty"`
	Introduction string    `json:"introduction,omitempty"`
	Avatar       string    `json:"avatar,omitempty"`
	CreateTime   time.Time `json:"-"`
	AccessKey    string    `json:"accessKey,omitempty"`
	SecretKey    string    `json:"secretKey,omitempty"`
	Status       int       `json:"status,omitempty"`
	*jwt.StandardClaims
}
