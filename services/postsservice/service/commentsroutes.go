package service

import (
	authmidware "github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
	"github.com/gin-gonic/gin"
)

var CommentsRoutes = []models.Route{
	{Method: "POST", Pattern: "/comments/v1/add", MiddleWare: gin.HandlersChain{authmidware.UserToken}, HandlerFunc: transaction.GormTx(savePostsComments)},
	{Method: "POST", Pattern: "/admin/v1/reply", MiddleWare: gin.HandlersChain{authmidware.UserToken}, HandlerFunc: transaction.GormTx(replyComments)},
	{Method: "DELETE", Pattern: "/comments/v1/:id", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: transaction.GormTx(deletePostsComments)},
	{Method: "GET", Pattern: "/comments/v1/:id", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: getPostsComments},
	{Method: "GET", Pattern: "/comments-posts/v1/list", MiddleWare: nil, HandlerFunc: getPostsCommentsByPostsIdList},
	{Method: "GET", Pattern: "/get/v1/comments", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: getPostsCommentsList},
}
