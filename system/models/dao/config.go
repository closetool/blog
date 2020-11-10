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

// GetAllConfig is a function to get a slice of record(s) from config table in the test database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllConfig(DB *gorm.DB, page, pagesize int, order string) (results []*model.Config, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Config{})
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

// GetConfig is a function to get a single record from the config table in the test database
// error - ErrNotFound, db Find error
func GetConfig(DB *gorm.DB, argId int64) (record *model.Config, err error) {
	record = &model.Config{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddConfig is a function to add a single record to config table in the test database
// error - ErrInsertFailed, db save call failed
func AddConfig(DB *gorm.DB, record *model.Config) (result *model.Config, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateConfig is a function to update a single record from config table in the test database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateConfig(DB *gorm.DB, argId int64, updated *model.Config) (result *model.Config, RowsAffected int64, err error) {

	result = &model.Config{}
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

// DeleteConfig is a function to delete a single record from config table in the test database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteConfig(DB *gorm.DB, argId int64) (rowsAffected int64, err error) {

	record := &model.Config{}
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
