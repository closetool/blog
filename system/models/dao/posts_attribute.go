package dao

import (
	"time"

	"github.com/closetool/blog/system/models/model"
	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllPostsAttribute is a function to get a slice of record(s) from posts_attribute table in the test database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllPostsAttribute(DB *gorm.DB, page, pagesize int, order string) (results []*model.PostsAttribute, totalRows int64, err error) {

	resultOrm := DB.Model(&model.PostsAttribute{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetPostsAttribute is a function to get a single record from the posts_attribute table in the test database
// error - ErrNotFound, db Find error
func GetPostsAttribute(DB *gorm.DB, argId int64) (record *model.PostsAttribute, err error) {
	record = &model.PostsAttribute{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddPostsAttribute is a function to add a single record to posts_attribute table in the test database
// error - ErrInsertFailed, db save call failed
func AddPostsAttribute(DB *gorm.DB, record *model.PostsAttribute) (result *model.PostsAttribute, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdatePostsAttribute is a function to update a single record from posts_attribute table in the test database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdatePostsAttribute(DB *gorm.DB, argId int64, updated *model.PostsAttribute) (result *model.PostsAttribute, RowsAffected int64, err error) {

	result = &model.PostsAttribute{}
	db := DB.First(result, argId)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeletePostsAttribute is a function to delete a single record from posts_attribute table in the test database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeletePostsAttribute(DB *gorm.DB, argId int64) (rowsAffected int64, err error) {

	record := &model.PostsAttribute{}
	db := DB.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
