package service

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
)

var TagsRoutes = []models.Route{
	{Method: "GET", Pattern: "/list/v1/tags", MiddleWare: nil, HandlerFunc: getTagsList},
	{Method: "GET", Pattern: "/tags-article-quantity/v1/list", MiddleWare: nil, HandlerFunc: getTagsAndArticleQuantityList},
	{Method: "GET", Pattern: "/tags/v1/:id", MiddleWare: nil, HandlerFunc: getTags},
	{Method: "POST", Pattern: "/tags/v1/add", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: transaction.GormTx(saveTags)},
	{Method: "PUT", Pattern: "/tags/v1/update", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: transaction.GormTx(updateTags)},
	{Method: "DELETE", Pattern: "/tags/v1/:id", MiddleWare: /*gin.HandlersChain{middlewares.AdminToken}*/ nil, HandlerFunc: transaction.GormTx(deleteTags)},
}
