package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `auth_user_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` varchar(20) NOT NULL COMMENT '记录用户id(游客取系统id：-1)',
  `ip` varchar(32) NOT NULL COMMENT 'ip地址',
  `url` varchar(255) NOT NULL COMMENT '请求的url',
  `parameter` varchar(5000) DEFAULT NULL COMMENT '需要记录的参数',
  `device` varchar(255) DEFAULT NULL COMMENT '来自于哪个设备 eg 手机 型号 电脑浏览器',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `code` varchar(10) DEFAULT NULL COMMENT '日志类型',
  `run_time` bigint(20) NOT NULL COMMENT '执行时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `browser_name` varchar(100) DEFAULT NULL COMMENT '浏览器名称',
  `browser_version` varchar(100) DEFAULT NULL COMMENT '浏览器版本号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户行为日志记录表'

JSON Sample
-------------------------------------
{    "browser_name": "cwNYdTIJfXMWaeJSkIuNLCopi",    "browser_version": "EhcfghpqBUudBFGjFjPwgTYSp",    "id": 19,    "ip": "LKQXOuWymqrqYaMnkaKpICuIP",    "parameter": "mciUeIMlaYWrmFewKlDSgSevN",    "device": "snWKbOEZAEOrFFOVYBmZPbmpE",    "code": "aKBrgrxtJHfgsOSvBMPryGstF",    "run_time": 44,    "user_id": "kGnuOOaehIrtoxLtmoblEhZdO",    "url": "QlgdoQjfAZTSPoKNYpSqvaMOs",    "description": "JdfwHLyWixRMfHYOyHjaSqGjF",    "create_time": "2207-06-12T19:12:45.209291572+08:00"}



*/

// AuthUserLog struct is a row record of the auth_user_log table in the test database
type AuthUserLog struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"` // 主键
	//[ 1] user_id                                        varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	UserID string `gorm:"column:user_id;type:varchar(20);size:20;" json:"userId,omitempty" form:"userId"` // 记录用户id(游客取系统id：-1)
	//[ 2] ip                                             varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	IP string `gorm:"column:ip;type:varchar(32);size:32;" json:"ip,omitempty" form:"ip"` // ip地址
	//[ 3] url                                            varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	URL string `gorm:"column:url;type:varchar(255);size:255;" json:"url,omitempty" form:"url"` // 请求的url
	//[ 4] parameter                                      varchar(5000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 5000    default: []
	Parameter null.String `gorm:"column:parameter;type:varchar(5000);size:5000;" json:"parameter,omitempty" form:"parameter"` // 需要记录的参数
	//[ 5] device                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Device null.String `gorm:"column:device;type:varchar(255);size:255;" json:"device,omitempty" form:"device"` // 来自于哪个设备 eg 手机 型号 电脑浏览器
	//[ 6] description                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Description null.String `gorm:"column:description;type:varchar(255);size:255;" json:"description,omitempty" form:"description"` // 描述
	//[ 7] code                                           varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	Code null.String `gorm:"column:code;type:varchar(10);size:10;" json:"code,omitempty" form:"code"` // 日志类型
	//[ 8] run_time                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	RunTime int64 `gorm:"column:run_time;type:bigint;" json:"runTime,omitempty" form:"runTime"` // 执行时间
	//[ 9] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:datetime;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"` // 创建时间
	//[10] browser_name                                   varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	BrowserName null.String `gorm:"column:browser_name;type:varchar(100);size:100;" json:"browserName,omitempty" form:"browserName"` // 浏览器名称
	//[11] browser_version                                varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	BrowserVersion null.String `gorm:"column:browser_version;type:varchar(100);size:100;" json:"browserVersion,omitempty" form:"browserVersion"` // 浏览器版本号

}

var auth_user_logTableInfo = &TableInfo{
	Name: "auth_user_log",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "user_id",
			Comment:            `记录用户id(游客取系统id：-1)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ip",
			Comment:            `ip地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "IP",
			GoFieldType:        "string",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "url",
			Comment:            `请求的url`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "URL",
			GoFieldType:        "string",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "parameter",
			Comment:            `需要记录的参数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(5000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       5000,
			GoFieldName:        "Parameter",
			GoFieldType:        "null.String",
			JSONFieldName:      "parameter",
			ProtobufFieldName:  "parameter",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "device",
			Comment:            `来自于哪个设备 eg 手机 型号 电脑浏览器`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Device",
			GoFieldType:        "null.String",
			JSONFieldName:      "device",
			ProtobufFieldName:  "device",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "description",
			Comment:            `描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Description",
			GoFieldType:        "null.String",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "code",
			Comment:            `日志类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "Code",
			GoFieldType:        "null.String",
			JSONFieldName:      "code",
			ProtobufFieldName:  "code",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "run_time",
			Comment:            `执行时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RunTime",
			GoFieldType:        "int64",
			JSONFieldName:      "run_time",
			ProtobufFieldName:  "run_time",
			ProtobufType:       "int64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "create_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "browser_name",
			Comment:            `浏览器名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "BrowserName",
			GoFieldType:        "null.String",
			JSONFieldName:      "browser_name",
			ProtobufFieldName:  "browser_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "browser_version",
			Comment:            `浏览器版本号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "BrowserVersion",
			GoFieldType:        "null.String",
			JSONFieldName:      "browser_version",
			ProtobufFieldName:  "browser_version",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthUserLog) TableName() string {
	return "auth_user_log"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthUserLog) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthUserLog) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthUserLog) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthUserLog) TableInfo() *TableInfo {
	return auth_user_logTableInfo
}
