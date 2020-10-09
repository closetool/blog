package po

import "time"

type AuthUserSocial struct {
	Id         int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Code       string    `xorm:"not null comment('qq、csdn、wechat、weibo、email等') VARCHAR(32)"`
	Content    string    `xorm:"comment('社交内容') VARCHAR(100)"`
	ShowType   int       `xorm:"not null comment('展示类型( 1、显示图片，2、显示账号，3、跳转链接)') SMALLINT(6)"`
	Remark     string    `xorm:"comment('备注') VARCHAR(150)"`
	Icon       string    `xorm:"comment('图标') VARCHAR(100)"`
	IsEnabled  int       `xorm:"not null default 0 comment('是否启用') SMALLINT(6)"`
	IsHome     int       `xorm:"default 0 comment('是否主页社交信息') SMALLINT(6)"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME created"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME updated"`
}

func (t AuthUserSocial) TableName() string {
	return "closetool_auth_user_social"
}
