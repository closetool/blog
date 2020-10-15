package transaction

import (
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
)

func Wrapper(fn func(*gin.Context) error) func(*gin.Context) {
	return func(g *gin.Context) {
		session := db.DB.NewSession()
		defer session.Close()

		err := session.Begin()
		if err != nil {
			session.Rollback()
			reply.CreateJSONError(g, reply.Error)
			return
		}
		err = fn(g)
		if err != nil {
			session.Rollback()
			log.Logger.Errorf("transaction rollbacked: %v\n", err)
			return
		}
		session.Commit()
		log.Logger.Debugf("transaction committed\n")
	}
}
