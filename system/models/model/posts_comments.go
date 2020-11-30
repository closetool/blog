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


CREATE TABLE `posts_comments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author_id` bigint(20) NOT NULL,
  `content` varchar(255) NOT NULL,
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `posts_id` bigint(20) NOT NULL,
  `tree_path` varchar(128) DEFAULT NULL COMMENT '层级结构',
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='评论表'

JSON Sample
-------------------------------------
{    "parent_id": 62,    "status": 81,    "posts_id": 22,    "tree_path": "dAHrPjqOCetctNwTIAxwBtMwL",    "create_time": "2137-08-15T18:33:52.477096064+08:00",    "id": 66,    "author_id": 2,    "content": "esMPWNhlLubFUgtEgaMKNwWKy"}



*/

// PostsComments struct is a row record of the posts_comments table in the test database
type PostsComments struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"`
	//[ 1] author_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AuthorID int64 `gorm:"column:author_id;type:bigint;" json:"authorId,omitempty" form:"authorId"`
	//[ 2] content                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Content string `gorm:"column:content;type:varchar(255);size:255;" json:"content,omitempty" form:"content"`
	//[ 3] parent_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ParentID int64 `gorm:"column:parent_id;type:bigint;default:0;" json:"parentId,omitempty" form:"parentId"`
	//[ 4] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status int32 `gorm:"column:status;type:int;default:0;" json:"status,omitempty" form:"status"`
	//[ 5] posts_id                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PostsID int64 `gorm:"column:posts_id;type:bigint;" json:"postsId,omitempty" form:"postsId"`
	//[ 6] tree_path                                      varchar(128)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	TreePath null.String `gorm:"column:tree_path;type:varchar(128);size:128;" json:"treePath,omitempty" form:"treePath"` // 层级结构
	//[ 7] create_time                                    datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime int64 `gorm:"column:create_time;type:datetime;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"`

	Title string `gorm:"-" json:"title" form:"title"`

	AuthorAvatar string `gorm:"-" json:"authorAvatar" form:"authorAvatar"`

	AuthorName string `gorm:"-" json:"authorName" form:"authorName"`

	ParentUserName string `gorm:"-" json:"parentUserName" form:"parentUserName"`

	ParentComments *PostsComments `gorm:"foreignKey:ParentID"`

	Posts Posts `gorm:"foreignKey:PostsID;references:Posts.ID"`

	*models.BaseVO
}

func Comments2Ints(comments []PostsComments) []interface{} {
	ints := []interface{}{}
	for i := range comments {
		ints = append(ints, comments[i])
	}
	return ints
}

var posts_commentsTableInfo = &TableInfo{
	Name: "posts_comments",
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
			Name:               "author_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AuthorID",
			GoFieldType:        "int64",
			JSONFieldName:      "author_id",
			ProtobufFieldName:  "author_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "content",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Content",
			GoFieldType:        "string",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "parent_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ParentID",
			GoFieldType:        "int64",
			JSONFieldName:      "parent_id",
			ProtobufFieldName:  "parent_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "status",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "posts_id",
			Comment:            ``,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "tree_path",
			Comment:            `层级结构`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "TreePath",
			GoFieldType:        "null.String",
			JSONFieldName:      "tree_path",
			ProtobufFieldName:  "tree_path",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "create_time",
			Comment:            ``,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PostsComments) TableName() string {
	return "posts_comments"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PostsComments) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PostsComments) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PostsComments) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PostsComments) TableInfo() *TableInfo {
	return posts_commentsTableInfo
}
