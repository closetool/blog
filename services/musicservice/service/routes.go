package service

import (
	"net/http"

	"github.com/closetool/blog/system/models"
	"github.com/gin-gonic/gin"
)

var Routes = []models.Route{
	{Method: "GET", Pattern: "/music/v1/list", MiddleWare: nil, HandlerFunc: GetPlayList},
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]bool{"health": true})
	}},
}
