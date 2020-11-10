package transaction

import (
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

func Wrapper(fn func(*gin.Context, *xorm.Session) error) func(*gin.Context) {
	return func(g *gin.Context) {
		session := db.DB.NewSession()
		defer session.Close()

		err := session.Begin()
		if err != nil {
			session.Rollback()
			reply.CreateJSONError(g, reply.Error)
			return
		}
		err = fn(g, session)
		if err != nil {
			session.Rollback()
			logrus.Errorf("transaction rollbacked: %v\n", err)
			return
		}
		session.Commit()
		logrus.Debugf("transaction committed\n")
	}
}

func GormTx(fn func(*gin.Context, *gorm.DB)) func(*gin.Context) {
	return func(g *gin.Context) {
		tx := db.Gorm.Begin()
		defer func() {
			if o := recover(); o != nil {
				tx.Rollback()
				logrus.Errorf("transaction rollbacked: %v\n", o)
				panic(o)
			}
		}()

		fn(g, tx)

		tx.Commit()
		logrus.Debugf("transaction committed\n")
	}
}
