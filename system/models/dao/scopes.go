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
			db = db.Where("title like ?", surround(m.Keywords))
		}
		if m.Title != "" {
			db = db.Where("title = ?", m.Title)
		}
		if m.ParentID.Valid {
			db = db.Where("parent_id = ?", m.ParentID)
		}
		if m.Icon != "" {
			db = db.Where("icon = ?", m.Icon)
		}
		if m.URL != "" {
			db = db.Where("url like ?", "%"+m.URL+"%")
		}
		db = db.Where("sort = ?", m.Sort)
		return db
	}
}

func LinkCond(l *model.FriendshipLink) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if l.BaseVO != nil && l.Keywords != "" {
			db = db.Where("name like ?", surround(l.Keywords))
		}
		if l.Href != "" {
			db = db.Where("href like ?", surround(l.Href))
		}
		if l.Name != "" {
			db = db.Where("name = ?", l.Name)
		}
		return db
	}
}

func CategoryCond(c *model.Category) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.BaseVO != nil && c.Keywords != "" {
			db = db.Where("name like ?", surround(c.Keywords))
		}
		if c.Name != "" {
			db = db.Where("name = ?", c.Name)
		}
		return db
	}
}

func TagsCond(t *model.Tags) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if t.BaseVO != nil && t.Keywords != "" {
			db = db.Where("name like ?", surround(t.Keywords))
		}
		if t.Name != "" {
			db = db.Where("name = ?", t.Name)
		}
		return db
	}
}

func UserCond(u *model.AuthUser) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if u.BaseVO != nil && u.Keywords != "" {
			db = db.Where("name like ?", surround(u.Keywords))
		}
		if u.Name.Valid {
			db = db.Where("name = ?", u.Name)
		}
		if u.Status.Valid {
			db = db.Where("status = ?", u.Status)
		}
		return db
	}
}

func SocialCond(s *model.AuthUserSocial) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.BaseVO != nil && s.Keywords != "" {
			db = db.Where("code like ?", surround(s.Keywords))
		}
		if s.Code != "" {
			db = db.Where("code = ?", s.Code)
		}
		if s.Content.Valid {
			db = db.Where("content = ?", s.Content)
		}
		if s.ShowType != 0 {
			db = db.Where("show_type = ?", s.ShowType)
		}
		if s.Remark.Valid {
			db = db.Where("remark = ?", s.Remark)
		}
		if s.IsEnabled != -1 {
			db = db.Where("is_enabled = ?", s.IsEnabled)
		}
		if s.IsHome.Valid {
			db = db.Where("is_home = ?", s.IsHome)
		}
		db = db.Order("id")
		return db
	}
}

func PostsCond(p *model.Posts) func(*gorm.DB) *gorm.DB {
	return func(DB *gorm.DB) *gorm.DB {
		if p.BaseVO != nil && p.Keywords != "" {
			DB = DB.Where("title like ?", surround(p.Keywords))
		}
		if p.ID != 0 {
			DB = DB.Where("posts.id = ?", p.ID)
		}
		if p.CreateTime != 0 {
			DB = DB.Where("posts.create_time = ?", p.CreateTime)
		}
		if p.CategoryID.Valid {
			DB = DB.Where("category_id = ?", p.CategoryID)
		}
		if p.PostsTagsID != 0 {
			DB = DB.Where("posts_tags.tags_id = ?", p.PostsTagsID)
		}
		if p.Title != "" {
			DB = DB.Where("title = ?", p.Title)
		}
		if p.Status != 0 {
			DB = DB.Where("status = ?", p.Status)
		}
		if p.IsWeight != 0 {
			DB = DB.Order("weight")
		} else {
			DB = DB.Order("posts.id")
		}
		return DB
	}
}

func CommentsCond(c *model.PostsComments) func(*gorm.DB) *gorm.DB {
	return func(DB *gorm.DB) *gorm.DB {
		if c.BaseVO != nil && c.Keywords != "" {
			DB = DB.Where("posts_comments.content = ? or posts.title like ?", c.Keywords, c.Keywords)
		}

		if c.ID != 0 {
			DB = DB.Where("posts_comments.id = ?", c.ID)
		}
		return DB
	}
}

func LogsCond(l *model.AuthUserLog) func(*gorm.DB) *gorm.DB {
	return func(DB *gorm.DB) *gorm.DB {
		if l.UserID != "" {
			DB = DB.Where("user_id = ?", l.UserID)
		}
		if l.IP != "" {
			DB = DB.Where("ip = ?", l.IP)
		}
		if l.URL != "" {
			DB = DB.Where("url like ?", surround(l.URL))
		}
		if l.Parameter.Valid {
			DB = DB.Where("parameter like ?", surround(l.Parameter.String))
		}
		if l.Device.Valid {
			DB = DB.Where("device like ?", surround(l.Device.String))
		}
		if l.Description.Valid {
			DB = DB.Where("description like ?", surround(l.Description.String))
		}
		if l.Code.Valid {
			DB = DB.Where("code = ?", surround(l.Code.String))
		}
		if l.BrowserName.Valid {
			DB = DB.Where("browset_name like ?", surround(l.BrowserName.String))
		}
		if l.BrowserVersion.Valid {
			DB = DB.Where("browser_version = ?", surround(l.BrowserName.String))
		}
		if l.CreateTime != 0 {
			DB = DB.Where("FROM_UNIXTIME( ? ,'%Y-%m-%d')=FROM_UNIXTIME(create_time, '%Y-%m-%d')", l.CreateTime)
		}
		return DB
	}
}
