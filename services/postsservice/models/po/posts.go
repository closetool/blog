package models

import (
	"time"
)

type Posts struct {
	Id         int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	AuthorId   int64     `xorm:"comment('文章创建人') BIGINT(255)"`
	Title      string    `xorm:"not null comment('文章标题') VARCHAR(64)"`
	Thumbnail  string    `xorm:"comment('封面图') VARCHAR(255)"`
	Comments   int       `xorm:"not null default 0 comment('评论数') INT(11)"`
	IsComment  int       `xorm:"default 1 comment('是否打开评论 (0 不打开 1 打开 )') SMALLINT(6)"`
	CategoryId int64     `xorm:"comment('分类主键') BIGINT(20)"`
	SyncStatus int       `xorm:"not null default 0 comment('同步到byteblogs状态') SMALLINT(6)"`
	Status     int       `xorm:"not null default 1 comment('状态 1 草稿 2 发布') INT(11)"`
	Summary    string    `xorm:"not null comment('摘要') VARCHAR(255)"`
	Views      int       `xorm:"not null default 0 comment('浏览次数') INT(11)"`
	Weight     int       `xorm:"not null default 0 comment('文章权重') INT(11)"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME"`
}

func (p Posts) TableName() string {
	return "closetool_posts"
}
