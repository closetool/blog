package vo

import "github.com/closetool/blog/system/models"

type Category struct {
	Id         int64            `form:"id" json:"id,omitempty"`
	Name       string           `form:"name" json:"name,omitempty" binding:"required"`
	Sort       int              `form:"sort" json:"sort,omitempty"`
	CreateTime *models.JSONTime `form:"createTime" json:"createTime,omitempty"`
	UpdateTime *models.JSONTime `form:"updateTime" json:"updateTime,omitempty"`
	TagsList   []*Tags          `form:"tagsList" json:"tagsList,omitempty"`
	Total      int64            `form:"total" json:"total,omitempty"`
	*models.BaseVO
}
