package models

import (
	"time"
)

type ClosetoolPostsComments struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	AuthorId   int64     `xorm:"not null BIGINT(20)"`
	Content    string    `xorm:"not null VARCHAR(255)"`
	ParentId   int64     `xorm:"not null default 0 BIGINT(20)"`
	Status     int       `xorm:"not null default 0 INT(11)"`
	PostsId    int64     `xorm:"not null BIGINT(20)"`
	TreePath   string    `xorm:"comment('层级结构') VARCHAR(128)"`
	CreateTime time.Time `xorm:"DATETIME"`
}
