package model

import (
	"database/sql"
	"time"

	"github.com/closetool/blog/system/models"
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


CREATE TABLE `tags` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL COMMENT '名称',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='标签表'

JSON Sample
-------------------------------------
{    "update_time": "2113-01-06T09:38:31.747069658+08:00",    "update_by": 6,    "id": 18,    "name": "qvJFjbqsHiiIOXWyZctBPRsKc",    "sort": 10,    "create_time": "2170-01-01T05:15:34.047300813+08:00",    "create_by": 29}



*/

// Tags struct is a row record of the tags table in the test database
type Tags struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"`
	//[ 1] name                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Name string `gorm:"column:name;type:varchar(32);size:32;unique;" json:"name,omitempty" form:"name"` // 名称
	//[ 2] sort                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Sort int32 `gorm:"column:sort;type:smallint;default:0;" json:"sort,omitempty" form:"sort"` // 排序
	//[ 3] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime int64 `gorm:"column:create_time;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"` // 创建时间
	//[ 4] create_by                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	CreateBy null.Int `gorm:"column:create_by;type:bigint;" json:"createBy,omitempty" form:"createBy"` // 创建人
	//[ 5] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime int64 `gorm:"column:update_time;autoUpdateTime:milli;" json:"updateTime,omitempty" form:"updateTime"` // 更新时间
	//[ 6] update_by                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	UpdateBy null.Int `gorm:"column:update_by;type:bigint;" json:"updateBy,omitempty" form:"updateBy"` // 更新人

	//Categorys []Category `gorm:"many2many:category_tags;foreignKey:Name;joinForeignKey:TagsName;References:Name;joinReferences:CategoryName;"`

	PostsTotal int64 `gorm:"-" form:"postsTotal" json:"postsTotal,omitempty"`

	*models.BaseVO
}

func Tags2Interfaces(tags []Tags) []interface{} {
	ints := make([]interface{}, len(tags))
	for i, tag := range tags {
		ints[i] = tag
	}
	return ints
}

var tagsTableInfo = &TableInfo{
	Name: "tags",
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
			Name:               "name",
			Comment:            `名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "sort",
			Comment:            `排序`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Sort",
			GoFieldType:        "int32",
			JSONFieldName:      "sort",
			ProtobufFieldName:  "sort",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "create_by",
			Comment:            `创建人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CreateBy",
			GoFieldType:        "null.Int",
			JSONFieldName:      "create_by",
			ProtobufFieldName:  "create_by",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "update_by",
			Comment:            `更新人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "UpdateBy",
			GoFieldType:        "null.Int",
			JSONFieldName:      "update_by",
			ProtobufFieldName:  "update_by",
			ProtobufType:       "int64",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Tags) TableName() string {
	return "tags"
}

// BeforeSave invoked before saving, return an error if field is not populated.
//func (t *Tags) BeforeSave() error {
//	return nil
//}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Tags) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Tags) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Tags) TableInfo() *TableInfo {
	return tagsTableInfo
}
