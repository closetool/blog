package constants

import (
	"time"
)

const (
	AuthHeader = "Authorization"
	TokenTTL   = time.Duration(time.Second * 7 * 24 * 60 * 60)

	RoleUser  = 1
	RoleAdmin = 2

	AccountUnlocked = 0
	AccountLocked   = 1

	SocialIsHome  = 1
	SocialEnabled = 1

	TreePath = "."

	ConfigTypeBase      = 0
	ConfigTypeQiniu     = 1
	ConfigTypeMusicId   = 2
	ConfigTypeStoreType = 3
	ConfigTypeAliyun    = 4
	ConfigTypeCross     = 5
	ConfigTypeDefault   = 6

	DefaultPathKey     = "default_path"
	DefaultPathValue   = "~/file"
	DefaultImageDomain = "default_image_domain"
	FileURL            = "/files/"

	StoreType   = "store_type"
	DefaultType = "default"
	AliyunOss   = "aliyun_oss"
	QiNiu       = "qiniu"
	COS         = "cos"
)

var (
	Roles = map[int64]string{RoleUser: "user", RoleAdmin: "admin"}
)
