package po

import (
	"time"
)

type CategoryTags struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)" form:"id"`
	TagsId     int64     `xorm:"not null comment('名称') BIGINT(32)" form:"tagsId"`
	CategoryId int64     `xorm:"not null comment('分类的主键') BIGINT(20)" form:"categoryId"`
	Sort       int       `xorm:"not null default 0 comment('排序') SMALLINT(6)" form:"sort"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME created" form:"createTime"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME updated" form:"updateTime"`
}

func (c CategoryTags) TableName() string {
	return "closetool_category_tags"
}
