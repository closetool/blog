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

// GetAllFriendshipLink is a function to get a slice of record(s) from friendship_link table in the test database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFriendshipLink(DB *gorm.DB, page, pagesize int, order string) (results []*model.FriendshipLink, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FriendshipLink{})
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

// GetFriendshipLink is a function to get a single record from the friendship_link table in the test database
// error - ErrNotFound, db Find error
func GetFriendshipLink(DB *gorm.DB, argId int64) (record *model.FriendshipLink, err error) {
	record = &model.FriendshipLink{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFriendshipLink is a function to add a single record to friendship_link table in the test database
// error - ErrInsertFailed, db save call failed
func AddFriendshipLink(DB *gorm.DB, record *model.FriendshipLink) (result *model.FriendshipLink, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFriendshipLink is a function to update a single record from friendship_link table in the test database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFriendshipLink(DB *gorm.DB, argId int64, updated *model.FriendshipLink) (result *model.FriendshipLink, RowsAffected int64, err error) {

	result = &model.FriendshipLink{}
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

// DeleteFriendshipLink is a function to delete a single record from friendship_link table in the test database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFriendshipLink(DB *gorm.DB, argId int64) (rowsAffected int64, err error) {

	record := &model.FriendshipLink{}
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
