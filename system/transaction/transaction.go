package transaction

import (
	"github.com/closetool/blog/system/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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
