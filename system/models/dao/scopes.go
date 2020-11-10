package dao

import (
	"fmt"

	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"gorm.io/gorm"
)

func Paginate(page *reply.PageInfo) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageSize, offset := pageutils.StartAndEnd(page)
		return db.Offset(offset).Limit(pageSize)
	}
}

func surround(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

func MenuCond(m *model.Menu) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if m.BaseVO != nil && m.Keywords != "" {
			db.Where("title like ?", surround(m.Keywords))
		}
		if m.Title != "" {
			db.Where("title = ?", m.Title)
		}
		if m.ParentID.Valid {
			db.Where("parent_id = ?", m.ParentID)
		}
		if m.Icon != "" {
			db.Where("icon = ?", m.Icon)
		}
		if m.URL != "" {
			db.Where("url like ?", "%"+m.URL+"%")
		}
		db.Where("sort = ?", m.Sort)
		return db
	}
}

func LinkCond(l *model.FriendshipLink) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if l.BaseVO != nil && l.Keywords != "" {
			db.Where("name like ?", surround(l.Keywords))
		}
		if l.Href != "" {
			db.Where("href like ?", surround(l.Href))
		}
		if l.Name != "" {
			db.Where("name = ?", l.Name)
		}
		return db
	}
}
