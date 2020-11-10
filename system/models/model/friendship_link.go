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


CREATE TABLE `friendship_link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(30) DEFAULT NULL COMMENT '标题',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `logo` varchar(255) NOT NULL COMMENT '文件',
  `href` varchar(255) NOT NULL COMMENT '跳转的路径',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='友情链接表'

JSON Sample
-------------------------------------
{    "id": 11,    "title": "CtQcMsUfJjRsqpixTeICGfoGO",    "name": "rLciGqDHGuDeLimXeoaiQdCyg",    "logo": "AnCovUGoMDWsfbInSvTsuFPIk",    "href": "bqLjJjoQbBCmrwDAEeoUmebIt",    "sort": 35,    "description": "eHpbrOShBaPWXOIxFFkoGNbgP"}



*/

// FriendshipLink struct is a row record of the friendship_link table in the test database
type FriendshipLink struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"`
	//[ 1] title                                          varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Title null.String `gorm:"column:title;type:varchar(30);size:30;" json:"title,omitempty" form:"title"` // 标题
	//[ 2] name                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Name string `gorm:"column:name;type:varchar(32);size:32;" json:"name,omitempty" form:"name" binding:"required"` // 名称
	//[ 3] logo                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Logo string `gorm:"column:logo;type:varchar(255);size:255;" json:"logo,omitempty" form:"logo"` // 文件
	//[ 4] href                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Href string `gorm:"column:href;type:varchar(255);size:255;" json:"href,omitempty" form:"href" binding:"required"` // 跳转的路径
	//[ 5] sort                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Sort int32 `gorm:"column:sort;type:smallint;default:0;" json:"sort,omitempty" form:"sort"` // 排序
	//[ 6] description                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Description null.String `gorm:"column:description;type:varchar(255);size:255;" json:"description,omitempty" form:"description"` // 描述

	*models.BaseVO `gorm:"-"`
}

func Links2Interfaces(links []FriendshipLink) []interface{} {
	ints := make([]interface{}, len(links))
	for i, link := range links {
		ints[i] = link
	}
	return ints
}

var friendship_linkTableInfo = &TableInfo{
	Name: "friendship_link",
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
			Name:               "title",
			Comment:            `标题`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Title",
			GoFieldType:        "null.String",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "logo",
			Comment:            `文件`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Logo",
			GoFieldType:        "string",
			JSONFieldName:      "logo",
			ProtobufFieldName:  "logo",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "href",
			Comment:            `跳转的路径`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Href",
			GoFieldType:        "string",
			JSONFieldName:      "href",
			ProtobufFieldName:  "href",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
	},
}

// TableName sets the insert table name for this struct type
func (f *FriendshipLink) TableName() string {
	return "friendship_link"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FriendshipLink) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FriendshipLink) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FriendshipLink) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FriendshipLink) TableInfo() *TableInfo {
	return friendship_linkTableInfo
}
