package dao

import (
	"time"

	"github.com/closetool/blog/system/models/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllAuthUserLog is a function to get a slice of record(s) from auth_user_log table in the test database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAuthUserLog(DB *gorm.DB, page, pagesize int, order string) (results []*model.AuthUserLog, totalRows int64, err error) {

	resultOrm := DB.Model(&model.AuthUserLog{})
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

// GetAuthUserLog is a function to get a single record from the auth_user_log table in the test database
// error - ErrNotFound, db Find error
func GetAuthUserLog(DB *gorm.DB, argId int64) (record *model.AuthUserLog, err error) {
	record = &model.AuthUserLog{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAuthUserLog is a function to add a single record to auth_user_log table in the test database
// error - ErrInsertFailed, db save call failed
func AddAuthUserLog(DB *gorm.DB, record *model.AuthUserLog) (result *model.AuthUserLog, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAuthUserLog is a function to update a single record from auth_user_log table in the test database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAuthUserLog(DB *gorm.DB, argId int64, updated *model.AuthUserLog) (result *model.AuthUserLog, RowsAffected int64, err error) {

	result = &model.AuthUserLog{}
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

// DeleteAuthUserLog is a function to delete a single record from auth_user_log table in the test database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAuthUserLog(DB *gorm.DB, argId int64) (rowsAffected int64, err error) {

	record := &model.AuthUserLog{}
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
