package service

import (
	authmid "github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
	"github.com/gin-gonic/gin"
)

var ConfigRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: health},
	{Method: "PUT", Pattern: "/config/v1/update", MiddleWare: gin.HandlersChain{authmid.AdminToken}, HandlerFunc: transaction.GormTx(updateConfig)},
	{Method: "GET", Pattern: "/config/v1/list", MiddleWare: gin.HandlersChain{authmid.AdminToken}, HandlerFunc: getConfigList},
	{Method: "GET", Pattern: "/config-base/v1/list", MiddleWare: nil, HandlerFunc: getConfigBaseList},
}
