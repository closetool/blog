package po

import (
	"time"
)

type Tags struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)" form:"id"`
	Name       string    `xorm:"not null comment('名称') VARCHAR(32) unique" form:"name"`
	Sort       int       `xorm:"not null default 0 comment('排序') SMALLINT(6)" form:"sort"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME created" form:"createTime"`
	CreateBy   int64     `xorm:"comment('创建人') BIGINT(20)" form:"createBy"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME updated" form:"updateTime"`
	UpdateBy   int64     `xorm:"comment('更新人') BIGINT(20)" form:"updateBy"`
}

func (c Tags) TableName() string {
	return "closetool_tags"
}
