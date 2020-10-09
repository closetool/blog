package po

import "time"

type AuthUserLog struct {
	Id             int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	UserId         string    `xorm:"not null comment('记录用户id(游客取系统id：-1)') VARCHAR(20)"`
	Ip             string    `xorm:"not null comment('ip地址') VARCHAR(32)"`
	Url            string    `xorm:"not null comment('请求的url') VARCHAR(255)"`
	Parameter      string    `xorm:"comment('需要记录的参数') VARCHAR(5000)"`
	Device         string    `xorm:"comment('来自于哪个设备 eg 手机 型号 电脑浏览器') VARCHAR(255)"`
	Description    string    `xorm:"comment('描述') VARCHAR(255)"`
	Code           string    `xorm:"comment('日志类型') VARCHAR(10)"`
	RunTime        int64     `xorm:"not null comment('执行时间') BIGINT(20)"`
	CreateTime     time.Time `xorm:"not null comment('创建时间') DATETIME created"`
	BrowserName    string    `xorm:"comment('浏览器名称') VARCHAR(100)"`
	BrowserVersion string    `xorm:"comment('浏览器版本号') VARCHAR(100)"`
}

func (t AuthUserLog) TableName() string {
	return "closetool_auth_user_log"
}
