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


CREATE TABLE `black_list` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ip_range` varchar(150) DEFAULT NULL COMMENT 'ip范围',
  `is_enable` int(1) DEFAULT '0' COMMENT '是否启用 0 启用，1不启用',
  `create_user` varchar(255) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_user` datetime DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='黑名单拦截表'

JSON Sample
-------------------------------------
{    "id": 48,    "ip_range": "fxwMUgvpAJZdiVQlPiTfgtHZw",    "is_enable": 18,    "create_user": "VYnkXPdsjasYtbVMMmDJQaqFO",    "create_time": "2146-06-25T20:46:19.65833932+08:00",    "update_user": "2106-04-12T01:46:01.420592466+08:00",    "update_time": "2234-12-24T20:16:19.402607302+08:00"}



*/

// BlackList struct is a row record of the black_list table in the test database
type BlackList struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"` // 主键
	//[ 1] ip_range                                       varchar(150)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 150     default: []
	IPRange null.String `gorm:"column:ip_range;type:varchar(150);size:150;" json:"ipRange,omitempty" form:"ipRange"` // ip范围
	//[ 2] is_enable                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsEnable null.Int `gorm:"column:is_enable;type:int;default:0;" json:"isEnable,omitempty" form:"isEnable"` // 是否启用 0 启用，1不启用
	//[ 3] create_user                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CreateUser null.String `gorm:"column:create_user;type:varchar(255);size:255;" json:"createUser,omitempty" form:"createUser"` // 创建者
	//[ 4] create_time                                    datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime null.Time `gorm:"column:create_time;type:datetime;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"` // 创建时间
	//[ 5] update_user                                    datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateUser null.Time `gorm:"column:update_user;type:datetime;" json:"updateUser,omitempty" form:"updateUser"` // 更新者
	//[ 6] update_time                                    datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime null.Time `gorm:"column:update_time;type:datetime;autoUpdateTime:milli;" json:"updateTime,omitempty" form:"updateTime"` // 更新时间

}

var black_listTableInfo = &TableInfo{
	Name: "black_list",
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
			Name:               "ip_range",
			Comment:            `ip范围`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(150)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       150,
			GoFieldName:        "IPRange",
			GoFieldType:        "null.String",
			JSONFieldName:      "ip_range",
			ProtobufFieldName:  "ip_range",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "is_enable",
			Comment:            `是否启用 0 启用，1不启用`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsEnable",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_enable",
			ProtobufFieldName:  "is_enable",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "create_user",
			Comment:            `创建者`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "CreateUser",
			GoFieldType:        "null.String",
			JSONFieldName:      "create_user",
			ProtobufFieldName:  "create_user",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "create_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "null.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "update_user",
			Comment:            `更新者`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdateUser",
			GoFieldType:        "null.Time",
			JSONFieldName:      "update_user",
			ProtobufFieldName:  "update_user",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "update_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "null.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BlackList) TableName() string {
	return "black_list"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BlackList) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BlackList) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BlackList) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BlackList) TableInfo() *TableInfo {
	return black_listTableInfo
}
