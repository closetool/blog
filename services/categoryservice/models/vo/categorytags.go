package vo

import (
	"time"
)

type CategoryTags struct {
	Id         int64     `form:"id" json:"id,omitempty"`
	TagsId     int64     `form:"tagsId" json:"tagsId,omitempty"`
	CategoryId int64     `form:"categoryId" json:"categoryId,omitempty"`
	Sort       int       `form:"sort" json:"sort,omitempty"`
	CreateTime time.Time `form:"createTime" json:"createTime,omitempty"`
	UpdateTime time.Time `form:"updateTime" json:"updateTime,omitempty"`
}
