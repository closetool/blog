package initial

import (
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
)

func InitServer(routes []models.Route, groupRoute string) *gin.Engine {
	r := gin.New()
	r.Use(middlewares.LogToFile())
	r.Use(middlewares.Recover())
	group := r.Group(groupRoute)

	routeutils.RegisterRoute(routes, group)
	log.Logger.Infof("server initialized")
	return r
}
