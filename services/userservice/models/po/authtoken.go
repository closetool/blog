package po

import (
	"time"
)

type AuthToken struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	Token      string    `xorm:"not null comment('token') VARCHAR(256)"`
	ExpireTime time.Time `xorm:"not null comment('过期时间') DATETIME"`
	UserId     int64     `xorm:"not null comment('创建人') BIGINT(20)"`
}

func (t AuthToken) TableName() string {
	return "closetool_auth_token"
}
