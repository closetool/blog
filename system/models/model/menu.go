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


CREATE TABLE `menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单Id',
  `title` varchar(32) NOT NULL COMMENT '名称',
  `icon` varchar(255) NOT NULL COMMENT 'icon图标',
  `url` varchar(255) NOT NULL COMMENT '跳转路径',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='菜单表'

JSON Sample
-------------------------------------
{    "url": "ICkIaUuJxlKwkVlArqYvKfxMa",    "sort": 47,    "id": 58,    "parent_id": 67,    "title": "dhnuVgDhfKWPuaTpRREgZtRJS",    "icon": "DAWhVIuNRKwUxTkHfDjKLdZBH"}



*/

// Menu struct is a row record of the menu table in the test database
type Menu struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty,omitempty" form:"id"`
	//[ 1] parent_id                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ParentID null.Int `gorm:"column:parent_id;type:bigint;default:null;" json:"parentId,omitempty,omitempty" form:"parentId"` // 父菜单Id
	//[ 2] title                                          varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Title string `gorm:"column:title;type:varchar(32);size:32;" json:"title,omitempty,omitempty" form:"title"` // 名称
	//[ 3] icon                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Icon string `gorm:"column:icon;type:varchar(255);size:255;" json:"icon,omitempty,omitempty" form:"icon"` // icon图标
	//[ 4] url                                            varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	URL string `gorm:"column:url;type:varchar(255);size:255;" json:"url,omitempty,omitempty" form:"url"` // 跳转路径
	//[ 5] sort                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Sort int32 `gorm:"column:sort;type:smallint;default:0;" json:"sort,omitempty,omitempty" form:"sort"` // 排序

	Children []Menu `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"child,omitempty" form:"child"`

	*models.BaseVO `gorm:"-"`
}

func MenuToInterface(menus []Menu) []interface{} {
	menusInterface := make([]interface{}, len(menus))
	for i, menu := range menus {
		menusInterface[i] = menu
	}
	return menusInterface
}

var menuTableInfo = &TableInfo{
	Name: "menu",
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
			Name:               "parent_id",
			Comment:            `父菜单Id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ParentID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "parent_id",
			ProtobufFieldName:  "parent_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
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
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "icon",
			Comment:            `icon图标`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Icon",
			GoFieldType:        "string",
			JSONFieldName:      "icon",
			ProtobufFieldName:  "icon",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "url",
			Comment:            `跳转路径`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "URL",
			GoFieldType:        "string",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
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
	},
}

// TableName sets the insert table name for this struct type
func (m *Menu) TableName() string {
	return "menu"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Menu) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Menu) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Menu) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Menu) TableInfo() *TableInfo {
	return menuTableInfo
}
