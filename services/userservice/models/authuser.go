package models

import "time"

type AuthUser struct{
	Id int64
	SocialId string `xorm:"default(null) comment('社交账户ID')"`
	Password string `xorm:"not null comment('密码')"`
	Name string `xorm:"default(null) comment('别名')"`
	RoleId int64 `xorm:"not null comment('角色主键 1 普通用户 2 admin')"`
	Email string `xorm:"varchar(128) default(null) unique comment('邮箱')"`
	Introduction string `xorm:"default(null) comment('个人简介')"`
	Avatar string `xorm:"default(null) comment('头像')"`
	CreateTime time.Time `xorm:"created not null comment('创建时间')"`
	AccessKey string `xorm:"default(null) comment('ak')"`
	SecretKey string `xorm:"default(null) comment('sk')"`
	Status int64 `xorm:"int(1) default(0) comment('0 正常 1 锁定')"`
}
