package po

import (
	"time"
)

type AuthUser struct {
	Id           int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	SocialId     string    `xorm:"comment('社交账户ID') VARCHAR(255)"`
	Password     string    `xorm:"not null comment('密码') VARCHAR(255)"`
	Name         string    `xorm:"comment('别名') VARCHAR(255)"`
	RoleId       int64     `xorm:"not null comment('角色主键 1 普通用户 2 admin') BIGINT(20)"`
	Email        string    `xorm:"comment('邮箱') unique VARCHAR(128)"`
	Introduction string    `xorm:"comment('个人简介') VARCHAR(255)"`
	Avatar       string    `xorm:"comment('头像') VARCHAR(255)"`
	CreateTime   time.Time `xorm:"not null comment('注册时间') DATETIME created"`
	AccessKey    string    `xorm:"comment('ak') VARCHAR(255)"`
	SecretKey    string    `xorm:"comment('sk') VARCHAR(255)"`
	Status       int       `xorm:"default 0 comment('1 正常 2 锁定 ') INT(1)"`
}

func (t AuthUser) TableName() string {
	return "closetool_auth_user"
}
