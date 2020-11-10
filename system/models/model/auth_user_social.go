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


CREATE TABLE `auth_user_social` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(32) NOT NULL COMMENT 'qq、csdn、wechat、weibo、email等',
  `content` varchar(100) DEFAULT NULL COMMENT '社交内容',
  `show_type` smallint(6) NOT NULL COMMENT '展示类型( 1、显示图片，2、显示账号，3、跳转链接)',
  `remark` varchar(150) DEFAULT NULL COMMENT '备注',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `is_enabled` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否启用',
  `is_home` smallint(6) DEFAULT '0' COMMENT '是否主页社交信息',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户表社交信息表'

JSON Sample
-------------------------------------
{    "content": "dydKsmMEllMRxlAoLdNPygyPr",    "remark": "FjkhJdaBPtSffSHVHRVDAeFxC",    "is_enabled": 44,    "is_home": 80,    "update_time": "2147-11-12T21:30:48.255886314+08:00",    "id": 72,    "code": "vEtDBXgONuiHVpinMjfhVebGW",    "show_type": 11,    "icon": "nkMXjwailYlOyBphUhfiMMEoe",    "create_time": "2139-01-20T19:26:07.733470852+08:00"}



*/

// AuthUserSocial struct is a row record of the auth_user_social table in the test database
type AuthUserSocial struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"` // 主键
	//[ 1] code                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Code string `gorm:"column:code;type:varchar(32);size:32;" json:"code,omitempty" form:"code"` // qq、csdn、wechat、weibo、email等
	//[ 2] content                                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Content null.String `gorm:"column:content;type:varchar(100);size:100;" json:"content,omitempty" form:"content"` // 社交内容
	//[ 3] show_type                                      smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: []
	ShowType int32 `gorm:"column:show_type;type:smallint;" json:"showType,omitempty" form:"showType"` // 展示类型( 1、显示图片，2、显示账号，3、跳转链接)
	//[ 4] remark                                         varchar(150)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 150     default: []
	Remark null.String `gorm:"column:remark;type:varchar(150);size:150;" json:"remark,omitempty" form:"remark"` // 备注
	//[ 5] icon                                           varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Icon null.String `gorm:"column:icon;type:varchar(100);size:100;" json:"icon,omitempty" form:"icon"` // 图标
	//[ 6] is_enabled                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsEnabled int32 `gorm:"column:is_enabled;type:smallint;default:0;" json:"isEnabled,omitempty" form:"isEnabled"` // 是否启用
	//[ 7] is_home                                        smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsHome null.Int `gorm:"column:is_home;type:smallint;default:0;" json:"isHome,omitempty" form:"isHome"` // 是否主页社交信息
	//[ 8] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:datetime;autoCreateTime:milli" json:"createTime,omitempty" form:"createTime"` // 创建时间
	//[ 9] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;autoUpdateTime:milli" json:"updateTime,omitempty" form:"updateTime"` // 更新时间

}

var auth_user_socialTableInfo = &TableInfo{
	Name: "auth_user_social",
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
			Name:               "code",
			Comment:            `qq、csdn、wechat、weibo、email等`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Code",
			GoFieldType:        "string",
			JSONFieldName:      "code",
			ProtobufFieldName:  "code",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "content",
			Comment:            `社交内容`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Content",
			GoFieldType:        "null.String",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "show_type",
			Comment:            `展示类型( 1、显示图片，2、显示账号，3、跳转链接)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "ShowType",
			GoFieldType:        "int32",
			JSONFieldName:      "show_type",
			ProtobufFieldName:  "show_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "remark",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(150)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       150,
			GoFieldName:        "Remark",
			GoFieldType:        "null.String",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "icon",
			Comment:            `图标`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Icon",
			GoFieldType:        "null.String",
			JSONFieldName:      "icon",
			ProtobufFieldName:  "icon",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "is_enabled",
			Comment:            `是否启用`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsEnabled",
			GoFieldType:        "int32",
			JSONFieldName:      "is_enabled",
			ProtobufFieldName:  "is_enabled",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "is_home",
			Comment:            `是否主页社交信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsHome",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_home",
			ProtobufFieldName:  "is_home",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "update_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthUserSocial) TableName() string {
	return "auth_user_social"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthUserSocial) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthUserSocial) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthUserSocial) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthUserSocial) TableInfo() *TableInfo {
	return auth_user_socialTableInfo
}
