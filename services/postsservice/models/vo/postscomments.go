package vo

import (
	"github.com/closetool/blog/system/models"
)

type PostsComments struct {
	Id         int64            `form:"id" json:"id,omitempty"`
	AuthorId   int64            `form:"authorId" json:"authorId,omitempty"`
	Content    string           `form:"content" json:"content,omitempty"`
	ParentId   int64            `form:"parentId" json:"parentId,omitempty"`
	Status     int              `form:"status" json:"status,omitempty"`
	PostsId    int64            `form:"postsId" json:"postsId,omitempty"`
	TreePath   string           `form:"treePath" json:"treePath,omitempty"`
	CreateTime *models.JSONTime `form:"createTime" json:"createTime,omitempty"`
}
