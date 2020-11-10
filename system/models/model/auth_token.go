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


CREATE TABLE `auth_token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `token` varchar(256) NOT NULL COMMENT 'token',
  `expire_time` datetime NOT NULL COMMENT '过期时间',
  `user_id` bigint(20) NOT NULL COMMENT '创建人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='标签表'

JSON Sample
-------------------------------------
{    "id": 12,    "token": "ZaiNcxXWCLmbZwUXfqZahntsl",    "expire_time": "2112-11-07T09:26:53.997830476+08:00",    "user_id": 65}



*/

// AuthToken struct is a row record of the auth_token table in the test database
type AuthToken struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"`
	//[ 1] token                                          varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	Token string `gorm:"column:token;type:varchar;size:256;" json:"token,omitempty" form:"token"` // token
	//[ 2] expire_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ExpireTime time.Time `gorm:"column:expire_time;type:datetime;" json:"expireTime,omitempty" form:"expireTime"` // 过期时间
	//[ 3] user_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	UserID int64 `gorm:"column:user_id;type:bigint;" json:"userId,omitempty" form:"userId"` // 创建人

}

var auth_tokenTableInfo = &TableInfo{
	Name: "auth_token",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
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
			Name:               "token",
			Comment:            `token`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "Token",
			GoFieldType:        "string",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "expire_time",
			Comment:            `过期时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ExpireTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "expire_time",
			ProtobufFieldName:  "expire_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "user_id",
			Comment:            `创建人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int64",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthToken) TableName() string {
	return "auth_token"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthToken) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthToken) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthToken) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthToken) TableInfo() *TableInfo {
	return auth_tokenTableInfo
}
