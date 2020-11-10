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


CREATE TABLE `posts_attribute` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `content` longtext NOT NULL COMMENT '内容',
  `posts_id` bigint(20) NOT NULL COMMENT '文章表主键',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT

JSON Sample
-------------------------------------
{    "id": 94,    "content": "mkYMUFunwypIbfkakMeYEZDEN",    "posts_id": 35}



*/

// PostsAttribute struct is a row record of the posts_attribute table in the test database
type PostsAttribute struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"` // 主键
	//[ 1] content                                        text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Content string `gorm:"column:content;type:text;size:4294967295;" json:"content,omitempty" form:"content"` // 内容
	//[ 2] posts_id                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PostsID int64 `gorm:"column:posts_id;type:bigint;" json:"postsId,omitempty" form:"postsId"` // 文章表主键
}

var posts_attributeTableInfo = &TableInfo{
	Name: "posts_attribute",
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
			Name:               "content",
			Comment:            `内容`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(4294967295)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       4294967295,
			GoFieldName:        "Content",
			GoFieldType:        "string",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "posts_id",
			Comment:            `文章表主键`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PostsID",
			GoFieldType:        "int64",
			JSONFieldName:      "posts_id",
			ProtobufFieldName:  "posts_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PostsAttribute) TableName() string {
	return "posts_attribute"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PostsAttribute) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PostsAttribute) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PostsAttribute) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PostsAttribute) TableInfo() *TableInfo {
	return posts_attributeTableInfo
}
