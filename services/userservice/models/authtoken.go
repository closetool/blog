package models

import "time"

type AuthToken struct{
	Id int64
	Token string `xorm:"not null comment('token')"`
	ExpireTime time.Time `xorm:"not null comment('过期时间')"`
	UserId int64 `xorm:"not null comment('创建人')"`
}
