package vo

import (
	"github.com/closetool/blog/system/models"
)

type PostsTags struct {
	Id         int64            `form:"id" json:"id,omitempty"`
	TagsId     int64            `form:"tagsId" json:"tagsId,omitempty"`
	PostsId    int64            `form:"postsId" json:"postsId,omitempty"`
	Sort       int              `form:"sort" json:"sort,omitempty"`
	CreateTime *models.JSONTime `form:"createTime" json:"createTime,omitempty"`
	UpdateTime *models.JSONTime `form:"updateTime" json:"updateTime,omitempty"`
}
