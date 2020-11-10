package po

import (
	"time"
)

type PostsTags struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	TagsId     int64     `xorm:"not null comment('名称') BIGINT(32)"`
	PostsId    int64     `xorm:"not null comment('文章主键') BIGINT(20)"`
	Sort       int       `xorm:"not null default 0 comment('排序') SMALLINT(6)"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME created"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME updated"`
}

func (p PostsTags) TableName() string {
	return "closetool_posts_tags"
}
