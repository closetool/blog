package vo

import (
	"github.com/closetool/blog/services/categoryservice/models/vo"
	"github.com/closetool/blog/system/models"
)

type Posts struct {
	Id                int64            `json:"id,omitempty" form:"id"`
	AuthorId          int64            `json:"authorId,omitempty" form:"authorId"`
	Author            string           `json:"author,omitempty" form:"author"`
	Title             string           `json:"title,omitempty" form:"title" binding:"required"`
	Content           string           `json:"content,omitempty" form:"content" binding:"required"`
	Thumbnail         string           `json:"thumbnail,omitempty" form:"thumbnail"` //封面图
	Comments          int              `json:"comments,omitempty" form:"comments"`
	CommentsTotal     int64            `json:"commentsTotal,omitempty" form:"commentsTotal"`
	IsComment         int              `json:"isComment,omitempty" form:"isComment"`
	CategoryId        int64            `json:"categoryId,omitempty" form:"categoryId"`
	SyncStatus        int              `json:"syncStatus,omitempty" form:"syncStatus"`
	Status            int              `json:"status,omitempty" form:"status"`
	Summary           string           `json:"summary,omitempty" form:"summary"`
	Views             int              `json:"views,omitempty" form:"views"`
	ViewsTotal        int64            `json:"viewsTotal,omitempty" form:"viewsTotal"`
	ArticleTotal      int64            `xorm:"articleTotal" json:"articleTotal,omitempty" form:"articleTotal"`
	DraftTotal        int64            `json:"draftTotal,omitempty" form:"draftTotal"`
	SyncTotal         int64            `json:"syncTotal,omitempty" form:"syncTotal"`
	TodayPublishTotal int64            `json:"todayPublishTotal,omitempty" form:"todayPublishTotal"`
	Weight            int              `json:"weight,omitempty" form:"weight"`
	TagList           []string         `json:"tagList,omitempty" form:"tagList"`
	SocialId          string           `json:"socialId,omitempty" form:"socialId"`
	Year              int64            `json:"year,omitempty" form:"year"`
	TagsName          string           `json:"tagsName,omitempty" form:"tagsName"`
	IsWeight          int64            `json:"isWeight,omitempty" form:"isWeight"`
	CategoryName      string           `json:"categoryName,omitempty" form:"categoryName"`
	ArchivePosts      []*Posts         `json:"archivePosts,omitempty" form:"archivePosts"`
	PostsTagsId       int64            `json:"postsTagsId,omitempty" form:"postsTagsId"`
	TagsList          []*vo.Tags       `json:"tagsList,omitempty"`
	ArchiveDate       *models.JSONTime `xorm:"archiveDate" json:"archiveDate,omitempty" form:"archiveDate"`
	CreateTime        *models.JSONTime `json:"createTime,omitempty" form:"createTime"`
	UpdateTime        *models.JSONTime `json:"updateTime,omitempty" form:"updateTime"`
	*models.BaseVO
}

func (p Posts) TableName() string {
	return "closetool_posts"
}
