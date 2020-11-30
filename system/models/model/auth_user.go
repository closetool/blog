package model

import (
	"database/sql"
	"time"

	"github.com/closetool/blog/system/models"
	"github.com/dgrijalva/jwt-go"
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


CREATE TABLE `auth_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `social_id` varchar(255) DEFAULT NULL COMMENT '社交账户ID',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `name` varchar(255) DEFAULT NULL COMMENT '别名',
  `role_id` bigint(20) NOT NULL COMMENT '角色主键 1 普通用户 2 admin',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `introduction` varchar(255) DEFAULT NULL COMMENT '个人简介',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `create_time` datetime NOT NULL COMMENT '注册时间',
  `access_key` varchar(255) DEFAULT NULL COMMENT 'ak',
  `secret_key` varchar(255) DEFAULT NULL COMMENT 'sk',
  `status` int(1) DEFAULT '0' COMMENT '0 正常 1 锁定 ',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_email` (`email`) COMMENT '邮箱唯一'
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT

JSON Sample
-------------------------------------
{    "name": "UNMuptQXOLVsbaaMivnlheURv",    "email": "hMBxUpkSlVGLGBNsMAEKpjNWX",    "avatar": "MidUSwYWoRbUTceaOwSEikgAX",    "secret_key": "NeOoFNemlmTFwsRKPFNrCkyNs",    "status": 23,    "social_id": "PcWquBiyKZVxhqHBLXpwnrXdt",    "password": "muMPfhwqwbBHVDrGtWwAAdFcY",    "introduction": "XrADJWNIqJeOLggmMnPXRmCfo",    "create_time": "2092-10-31T19:57:47.780142044+08:00",    "access_key": "PEKCfQhoqkEZIdfJhrXDYKMBF",    "id": 41,    "role_id": 55}



*/

// AuthUser struct is a row record of the auth_user table in the test database
type AuthUser struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"` // 主键
	//[ 1] social_id                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SocialID null.String `gorm:"column:social_id;type:varchar(255);size:255;" json:"socialId,omitempty" form:"socialId"` // 社交账户ID
	//[ 2] password                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Password string `gorm:"column:password;type:varchar(255);size:255;" json:"password,omitempty" form:"password"` // 密码

	PasswordOld string `gorm:"-" json:"passwordOld,omitempty" form:"passwordOld"`
	//[ 3] name                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name null.String `gorm:"column:name;type:varchar(255);size:255;" json:"name,omitempty" form:"name"` // 别名
	//[ 4] role_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	RoleID int64 `gorm:"column:role_id;type:bigint;" json:"roleId,omitempty" form:"roleId"` // 角色主键 1 普通用户 2 admin
	//[ 5] email                                          varchar(128)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	Email null.String `gorm:"column:email;type:varchar(128);size:128;" json:"email,omitempty" form:"email"` // 邮箱
	//[ 6] introduction                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Introduction null.String `gorm:"column:introduction;type:varchar(255);size:255;" json:"introduction,omitempty" form:"introduction"` // 个人简介
	//[ 7] avatar                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Avatar null.String `gorm:"column:avatar;type:varchar(255);size:255;" json:"avatar,omitempty" form:"avatar"` // 头像
	//[ 8] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime int64 `gorm:"column:create_time;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"` // 注册时间
	//[ 9] access_key                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AccessKey null.String `gorm:"column:access_key;type:varchar(255);size:255;" json:"accessKey,omitempty" form:"accessKey"` // ak
	//[10] secret_key                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SecretKey null.String `gorm:"column:secret_key;type:varchar(255);size:255;" json:"secretKey,omitempty" form:"secretKey"` // sk
	//[11] status                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status null.Int `gorm:"column:status;type:int;default:0;not null;" json:"status,omitempty" form:"status"` // 0 正常 1 锁定

	Token string `gorm:"-" json:"token,omitempty" form:"token"`

	Roles []string `gorm:"-" json:"roles,omitempty" form:"roles"`

	VerifyCode string `gorm:"-" json:"verifyCode,omitempty" form:"verifyCode"`

	*jwt.StandardClaims `gorm:"-"`

	*models.BaseVO

	AuthToken AuthToken `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func Users2Interfaces(users []AuthUser) []interface{} {
	ints := make([]interface{}, len(users))
	for i := range users {
		ints[i] = users[i]
	}
	return ints
}

var auth_userTableInfo = &TableInfo{
	Name: "auth_user",
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
			Name:               "social_id",
			Comment:            `社交账户ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SocialID",
			GoFieldType:        "null.String",
			JSONFieldName:      "social_id",
			ProtobufFieldName:  "social_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "password",
			Comment:            `密码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
			Comment:            `别名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "role_id",
			Comment:            `角色主键 1 普通用户 2 admin`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RoleID",
			GoFieldType:        "int64",
			JSONFieldName:      "role_id",
			ProtobufFieldName:  "role_id",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "email",
			Comment:            `邮箱`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "introduction",
			Comment:            `个人简介`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Introduction",
			GoFieldType:        "null.String",
			JSONFieldName:      "introduction",
			ProtobufFieldName:  "introduction",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "avatar",
			Comment:            `头像`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Avatar",
			GoFieldType:        "null.String",
			JSONFieldName:      "avatar",
			ProtobufFieldName:  "avatar",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "create_time",
			Comment:            `注册时间`,
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
			Name:               "access_key",
			Comment:            `ak`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "AccessKey",
			GoFieldType:        "null.String",
			JSONFieldName:      "access_key",
			ProtobufFieldName:  "access_key",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "secret_key",
			Comment:            `sk`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SecretKey",
			GoFieldType:        "null.String",
			JSONFieldName:      "secret_key",
			ProtobufFieldName:  "secret_key",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "status",
			Comment:            `0 正常 1 锁定 `,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "null.Int",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthUser) TableName() string {
	return "auth_user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthUser) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthUser) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthUser) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthUser) TableInfo() *TableInfo {
	return auth_userTableInfo
}
