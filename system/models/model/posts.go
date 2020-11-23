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


CREATE TABLE `posts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `author_id` bigint(255) DEFAULT NULL COMMENT '文章创建人',
  `title` varchar(64) NOT NULL COMMENT '文章标题',
  `thumbnail` varchar(255) DEFAULT NULL COMMENT '封面图',
  `comments` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
  `is_comment` smallint(6) DEFAULT '1' COMMENT '是否打开评论 (0 不打开 1 打开 )',
  `category_id` bigint(20) DEFAULT NULL COMMENT '分类主键',
  `sync_status` smallint(6) NOT NULL DEFAULT '0' COMMENT '同步到byteblogs状态',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1 草稿 2 发布',
  `summary` varchar(255) NOT NULL COMMENT '摘要',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '文章权重',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT

JSON Sample
-------------------------------------
{    "comments": 58,    "sync_status": 34,    "create_time": "2210-01-27T06:58:53.751289631+08:00",    "is_comment": 83,    "weight": 97,    "author_id": 10,    "title": "gRpPKTJAueBgcxZaeRrwmWTWY",    "category_id": 97,    "status": 93,    "summary": "ApoSUhsQQDpxuLRyBvFNeZyEo",    "update_time": "2089-11-13T01:14:23.147893334+08:00",    "id": 88,    "thumbnail": "bqjBgXEDcleSDnpKeHieLrTEH",    "views": 95}



*/

// Posts struct is a row record of the posts table in the test database
type Posts struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primaryKey;autoIncrement;column:id;" json:"id,omitempty" form:"id"` // 主键
	//[ 1] author_id                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AuthorID null.Int `gorm:"column:author_id;type:bigint;" json:"authorId,omitempty" form:"authorId"` // 文章创建人
	//[ 2] title                                          varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Title string `gorm:"column:title;type:varchar(64);size:64;" json:"title,omitempty" form:"title"` // 文章标题
	//[ 3] thumbnail                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Thumbnail null.String `gorm:"column:thumbnail;type:varchar(255);size:255;" json:"thumbnail,omitempty" form:"thumbnail"` // 封面图
	//[ 4] comments                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Comments int32 `gorm:"column:comments;type:int;default:0;" json:"comments,omitempty" form:"comments"` // 评论数
	//[ 5] is_comment                                     smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [1]
	IsComment null.Int `gorm:"column:is_comment;type:smallint;default:1;" json:"isComment,omitempty" form:"isComment"` // 是否打开评论 (0 不打开 1 打开 )
	//[ 6] category_id                                    bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	CategoryID null.Int `gorm:"column:category_id;type:bigint;" json:"categoryId,omitempty" form:"categoryId"` // 分类主键
	//[ 7] sync_status                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	SyncStatus int32 `gorm:"column:sync_status;type:smallint;default:0;" json:"syncStatus,omitempty" form:"syncStatus"` // 同步到byteblogs状态
	//[ 8] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Status int32 `gorm:"column:status;type:int;default:1;" json:"status,omitempty" form:"status"` // 状态 1 草稿 2 发布
	//[ 9] summary                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Summary string `gorm:"column:summary;type:varchar(255);size:255;" json:"summary,omitempty" form:"summary"` // 摘要
	//[10] views                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Views int32 `gorm:"column:views;type:int;default:0;" json:"views,omitempty" form:"views"` // 浏览次数
	//[11] weight                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Weight int32 `gorm:"column:weight;type:int;default:0;" json:"weight,omitempty" form:"weight"` // 文章权重
	//[12] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateTime int64 `gorm:"column:create_time;type:datetime;autoCreateTime:milli;" json:"createTime,omitempty" form:"createTime"` // 创建时间
	//[13] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime int64 `gorm:"column:update_time;type:datetime;autoUpdateTime:milli;" json:"updateTime,omitempty" form:"updateTime"` // 更新时间

	PostsAttribute PostsAttribute `gorm:"foreignKey:PostsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PostsTags []PostsTags `gorm:"foreignKey:PostsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PostsComments []PostsComments `gorm:"foreignKey:PostsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	*models.BaseVO `gorm:"-"`

	Content string `gorm:"-" json:"content" form:"content"`

	CommentsTotal int64 `gorm:"-" json:"commentsTotal" form:"commentsTotal"`

	ViewsTotal int64 `gorm:"-" json:"viewsTotal" form:"viewsTotal"`

	DraftTotal int64 `gorm:"-" json:"draftTotal" form:"draftTotal"`

	SyncTotal int64 `gorm:"-" json:"syncTotal" form:"syncTotal"`

	TodayPublishTotal int64 `gorm:"-" json:"todayPublishTotal" form:"todayPublishTotal"`

	TagList []string `gorm:"-" json:"tagList" form:"tagList"`

	SocialID string `gorm:"-" json:"socialId" form:"socialId"`

	Year int64 `gorm:"-" json:"year" form:"year"`

	TagsName string `gorm:"-" json:"tagsName" form:"tagsName"`

	IsWeight int64 `gorm:"-" json:"isWeight" form:"isWeight"`

	CategoryName string `gorm:"-" json:"categoryName" form:"categoryName"`

	Author string `gorm:"-" json:"author" form:"author"`

	ArchivePosts []Posts `gorm:"-" json:"archivePosts" form:"archivePosts"`

	PostsTagsID int64 `gorm:"-" json:"postsTagsId" form:"postsTagsId"`

	TagsList []Tags `gorm:"-" json:"tagsList" form:"tagsList"`

	ArchiveDate *models.JSONTime `gorm:"-" json:"archiveDate" form:"archiveDate"`

	ArchiveTotal int64 `gorm:"-" json:"archiveTotal" form:"archiveTotal"`
}

func Posts2Interfaces(posts []Posts) []interface{} {
	ints := make([]interface{}, len(posts))
	for i, post := range posts {
		ints[i] = post
	}
	return ints
}

var postsTableInfo = &TableInfo{
	Name: "posts",
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
			Name:               "author_id",
			Comment:            `文章创建人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AuthorID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "author_id",
			ProtobufFieldName:  "author_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
			Comment:            `文章标题`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "thumbnail",
			Comment:            `封面图`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Thumbnail",
			GoFieldType:        "null.String",
			JSONFieldName:      "thumbnail",
			ProtobufFieldName:  "thumbnail",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "comments",
			Comment:            `评论数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Comments",
			GoFieldType:        "int32",
			JSONFieldName:      "comments",
			ProtobufFieldName:  "comments",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_comment",
			Comment:            `是否打开评论 (0 不打开 1 打开 )`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsComment",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_comment",
			ProtobufFieldName:  "is_comment",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "category_id",
			Comment:            `分类主键`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "int64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sync_status",
			Comment:            `同步到byteblogs状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "SyncStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "sync_status",
			ProtobufFieldName:  "sync_status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "status",
			Comment:            `状态 1 草稿 2 发布`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "summary",
			Comment:            `摘要`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Summary",
			GoFieldType:        "string",
			JSONFieldName:      "summary",
			ProtobufFieldName:  "summary",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "views",
			Comment:            `浏览次数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Views",
			GoFieldType:        "int32",
			JSONFieldName:      "views",
			ProtobufFieldName:  "views",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "weight",
			Comment:            `文章权重`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Weight",
			GoFieldType:        "int32",
			JSONFieldName:      "weight",
			ProtobufFieldName:  "weight",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Posts) TableName() string {
	return "posts"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Posts) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Posts) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Posts) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Posts) TableInfo() *TableInfo {
	return postsTableInfo
}
