package vo

import "time"

type AuthUserSocial struct {
	Id         int64     `json:"id" form:"id"`
	Code       string    `json:"code,omitempty" form:"code"`
	Content    string    `json:"content,omitempty" form:"content"`
	ShowType   int       `json:"showType,omitempty" form:"showType"`
	Remark     string    `json:"remark,omitempty" form:"remark"`
	Icon       string    `json:"icon,omitempty" form:"icon"`
	IsEnabled  int       `json:"isEnabled,omitempty" form:"isEnabled"`
	IsHome     int       `json:"isHome,omitempty" form:"isHome"`
	CreateTime time.Time `json:"createTime,omitempty" form:"createTime"`
	UpdateTime time.Time `json:"updateTime,omitempty" form:"updateTime"`
}

func (t AuthUserSocial) TableName() string {
	return "closetool_auth_user_social"
}
