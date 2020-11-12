package service

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
)

var CategoryRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: health},
	{Method: "POST", Pattern: "/category/v1/add", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: transaction.GormTx(saveCategory)},
	{Method: "GET", Pattern: "/statistics/v1/list", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: statisticsList},
	{Method: "PUT", Pattern: "/category/v1/update", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: transaction.GormTx(updateCategory)},
	{Method: "GET", Pattern: "/category-tags/v1/:id", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: getCategoryTags},
	{Method: "GET", Pattern: "/list/v1/category-tags", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: getCategoryTagsList},
	{Method: "GET", Pattern: "/category/v1/:id", MiddleWare: nil, HandlerFunc: getCategoryTags},
	{Method: "GET", Pattern: "/list/v1/category", MiddleWare: nil, HandlerFunc: getCategoryList},
	{Method: "DELETE", Pattern: "/category/v1/:id", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: transaction.GormTx(deleteCategory)},
}
