package routeutils

import (
	"github.com/closetool/blog/system/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(Routes []models.Route,group *gin.RouterGroup){
	for _, r := range Routes{
		if r.MiddleWare == nil {
			r.MiddleWare = make([]gin.HandlerFunc,0)
		}
		r.MiddleWare = append(r.MiddleWare, r.HandlerFunc)
		group.Handle(r.Method, r.Pattern, r.MiddleWare...)
	}
}
