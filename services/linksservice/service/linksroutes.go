package service

import (
	authmid "github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
	"github.com/gin-gonic/gin"
)

var LinksRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: health},
	{Method: "GET", Pattern: "/list/v1/link", MiddleWare: nil, HandlerFunc: getFriendshipLinkList},
	{Method: "GET", Pattern: "/link/v2/list", MiddleWare: nil, HandlerFunc: getFriendshipLinkMap},
	{Method: "POST", Pattern: "/link/v1/add", MiddleWare: gin.HandlersChain{authmid.AdminToken}, HandlerFunc: transaction.GormTx(saveFriendshipLink)},
	{Method: "PUT", Pattern: "/link/v1/update", MiddleWare: gin.HandlersChain{authmid.AdminToken}, HandlerFunc: transaction.GormTx(updateFriendshipLink)},
	{Method: "GET", Pattern: "/link/v1/:id", MiddleWare: nil, HandlerFunc: getFriendshipLink},
	{Method: "DELETE", Pattern: "/link/v1/:id", MiddleWare: nil, HandlerFunc: transaction.GormTx(deleteFriendshipLink)},
}
