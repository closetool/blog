package models

import "time"

type AuthUserLog struct{
	Id int64
	UserId int64 `xorm:"not null comment('记录用户id')"`
	Ip string `xorm:"varchar(32) not null comment('IP地址')"`
	Url string `xorm:"not null comment('请求的url')"`
	Parameter string `xorm:"varchar(5000) default(null) comment('需要记录的参数')"`
	Device string `xorm:"default(null) comment('来自于哪个设备')"`
	Description string `xorm:"default null comment('描述')"`
	Code string`xorm:"varchar(10) default null comment('日志类型')"`
	RunTime int64 `xorm:"bigint(20) not null comment('执行时间')"`
	CreateTime time.Time `xorm:"created not null comment('创建时间')"`
	BrowserName string `xorm:"varchar(100) default null comment('浏览器名称')"`
	BrowserVersion string `xorm:"varchar(100) default null comment('浏览器版本号')"`
}
