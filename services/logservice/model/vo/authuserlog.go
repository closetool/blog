package vo

import (
	"github.com/closetool/blog/system/models"
)

type AuthUserLog struct {
	Id             int64            `json:"id,omitempty" form:"id"`
	UserId         string           `json:"userId,omitempty" form:"userId"`
	Ip             string           `json:"ip,omitempty" form:"ip"`
	Url            string           `json:"url,omitempty" form:"url"`
	Parameter      string           `json:"parameter,omitempty" form:"parameter"`
	Device         string           `json:"device,omitempty" form:"device"`
	Description    string           `json:"description,omitempty" form:"description"`
	Code           string           `json:"code,omitempty" form:"code"`
	CodeName       string           `json:"codeName,omitempty" form:"codeName"`
	RunTime        int64            `json:"runTime,omitempty" form:"runTime"`
	CreateTime     *models.JSONTime `json:"createTime,omitempty" form:"createTime"`
	BrowserName    string           `json:"browserName,omitempty" form:"browserName"`
	BrowserVersion string           `json:"browserVersion,omitempty" form:"browserVersion"`
	Count          int64            `json:"count,omitempty" form:"count"`
	UserTotal      int64            `json:"userTotal,omitempty" form:"userTotal"`
	ViewTotal      int64            `json:"viewTotal,omitempty" form:"viewTotal"`
	StartTime      *models.JSONTime `json:"startTime,omitempty" form:"startTime"`
	EndTime        *models.JSONTime `json:"endTime,omitempty" form:"endTime"`
	Index          int64            `json:"index,omitempty" form:"index"`
	Type           string           `json:"type,omitempty" form:"type"`
	*models.BaseVO
}

func (t AuthUserLog) TableName() string {
	return "closetool_auth_user_log"
}
