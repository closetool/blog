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


CREATE TABLE `config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `config_key` varchar(128) NOT NULL COMMENT '配置key',
  `config_value` varchar(255) NOT NULL COMMENT '配置值',
  `type` smallint(6) NOT NULL DEFAULT '0' COMMENT '配置类型',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UK_99vo6d7ci4wlxruo3gd0q2jq8` (`config_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT

JSON Sample
-------------------------------------
{    "type": 61,    "id": 92,    "config_key": "LPrTKGWWfhrUDTQFBtLJcFnSf",    "config_value": "HbrwEFchLlUOakWkdebcAlGrq"}



*/

// Config struct is a row record of the config table in the test database
type Config struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"`
	//[ 1] config_key                                     varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	ConfigKey string `gorm:"column:config_key;type:varchar(128);size:128;" json:"configKey,omitempty" form:"configKey" binding:"required"` // 配置key
	//[ 2] config_value                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ConfigValue string `gorm:"column:config_value;type:varchar(255);size:255;" json:"configValue,omitempty" form:"configValue" binding:"required"` // 配置值
	//[ 3] type                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Type int32 `gorm:"column:type;type:smallint;default:0;" json:"type,omitempty" form:"type"` // 配置类型

}

func Configs2Interface(configs []Config) []interface{} {
	inSlice := make([]interface{}, len(configs))
	for i, config := range configs {
		inSlice[i] = config
	}
	return inSlice
}

var configTableInfo = &TableInfo{
	Name: "config",
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
			Name:               "config_key",
			Comment:            `配置key`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "ConfigKey",
			GoFieldType:        "string",
			JSONFieldName:      "config_key",
			ProtobufFieldName:  "config_key",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "config_value",
			Comment:            `配置值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ConfigValue",
			GoFieldType:        "string",
			JSONFieldName:      "config_value",
			ProtobufFieldName:  "config_value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
			Comment:            `配置类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Config) TableName() string {
	return "config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Config) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Config) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Config) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Config) TableInfo() *TableInfo {
	return configTableInfo
}
