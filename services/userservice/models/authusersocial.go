package models

import "time"

type AuthUserSocial struct{
	Id int64
	Code string `xorm:"varchar(32) not null comment('qq,csdn,wechat,weibo,email,etc')"`
	Content string `xorm:"varchar(100) default null comment('社交内容')"`
	ShowType int64 `xorm:"smallint(6) not null comment('展示类型(1、显示图片，2、显示账号，3、跳转链接)')"`
	Remark string `xorm:"varchar(150) default null comment('备注')"`
	Icon string `xorm:"varchar(100) default null comment('图标')"`
	IsEnabled int64 `xorm:"smallint(6) not null default(0) comment('是否启用')"`
	IsHome int64 `xorm:"smallint(6) default(0) comment('是否主页社交信息')"`
	CreateTime time.Time `xorm:"not null created comment('创建时间')"`
	UpdateTime time.Time `xorm:"not null updated comment('更新时间')"`
}
