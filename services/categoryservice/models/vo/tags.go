package vo

import "github.com/closetool/blog/system/models"

type Tags struct {
	Id         int64  `form:"id" json:"id,omitempty"`
	Name       string `form:"name" json:"name,omitempty" binding:"required"`
	PostsTotal int64  `form:"postsTotal" json:"postsTotal,omitempty"`
	*models.BaseVO
}
