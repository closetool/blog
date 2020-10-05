package dbutils

import (
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

func InitTables(db *xorm.Engine,beans ...interface{}) error {
	var err error
	_,err = db.Transaction(func(session *xorm.Session) (interface{}, error) {
		for _,bean := range beans{
			err = db.Sync2(bean)
			if err != nil {
				logrus.Errorf("An error occurred when sync db: %v",err)
			}
		}
		return nil,err
	})
	return err
}
