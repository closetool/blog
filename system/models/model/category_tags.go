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


CREATE TABLE `category_tags` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tags_id` bigint(32) NOT NULL COMMENT '名称',
  `category_id` bigint(20) NOT NULL COMMENT '分类的主键',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT

JSON Sample
-------------------------------------
{    "id": 42,    "tags_id": 49,    "category_id": 14,    "sort": 33,    "create_time": "2132-04-24T22:17:01.750310986+08:00",    "update_time": "2309-03-01T02:00:02.974479618+08:00"}



*/

// CategoryTags struct is a row record of the category_tags table in the test database
type CategoryTags struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"`
	//[ 1] tags_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	TagsID int64 `gorm:"column:tags_id;type:bigint;" json:"tagsId,omitempty" form:"tagsId"` // 名称
	//[ 2] category_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	CategoryID int64 `gorm:"column:category_id;type:bigint;" json:"categoryId,omitempty" form:"categoryId"` // 分类的主键
	//[ 3] sort                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Sort int32 `gorm:"column:sort;type:smallint;default:0;" json:"sort,omitempty" form:"sort"` // 排序
	//[ 4] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:datetime;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"` // 创建时间
	//[ 5] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;autoUpdateTime:milli;" json:"updateTime,omitempty" form:"updateTime"` // 更新时间

}

var category_tagsTableInfo = &TableInfo{
	Name: "category_tags",
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
			Name:               "tags_id",
			Comment:            `名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "TagsID",
			GoFieldType:        "int64",
			JSONFieldName:      "tags_id",
			ProtobufFieldName:  "tags_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "category_id",
			Comment:            `分类的主键`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "int64",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
	},
}

// TableName sets the insert table name for this struct type
func (c *CategoryTags) TableName() string {
	return "category_tags"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CategoryTags) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CategoryTags) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CategoryTags) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CategoryTags) TableInfo() *TableInfo {
	return category_tagsTableInfo
}
