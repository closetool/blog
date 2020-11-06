package service

import (
	"github.com/closetool/blog/services/postsservice/middlewares"
	"github.com/closetool/blog/system/constants"
	authmidware "github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
	"github.com/gin-gonic/gin"
)

var PostsRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: Health},
	{Method: "GET", Pattern: "/list/v1/posts", MiddleWare: gin.HandlersChain{middlewares.AuthUserLogMiddleware(constants.PostsList, "文章列表")}, HandlerFunc: getPostsList},
	{Method: "GET", Pattern: "/weight/v1/list", MiddleWare: gin.HandlersChain{middlewares.AuthUserLogMiddleware(constants.PostsList, "权重列表")}, HandlerFunc: getWeightList},
	{Method: "GET", Pattern: "/archive/v1/list", MiddleWare: nil, HandlerFunc: getArchiveTotalByDateList},
	{Method: "GET", Pattern: "/hot/v1/list", MiddleWare: nil, HandlerFunc: getHotPostsList},
	{Method: "POST", Pattern: "/posts/v1/add", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: transaction.Wrapper(savePosts)},
	{Method: "GET", Pattern: "/posts/v1/:id", MiddleWare: gin.HandlersChain{middlewares.AuthUserLogMiddleware(constants.PostsDetail, "获取文章")}, HandlerFunc: getPosts},
	{Method: "DELETE", Pattern: "/posts/v1/:id", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: transaction.Wrapper(deletePosts)},
	{Method: "PUT", Pattern: "/posts/v1/update", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: transaction.Wrapper(updatePosts)},
	{Method: "PUT", Pattern: "/status/v1/update", MiddleWare: gin.HandlersChain{authmidware.AdminToken}, HandlerFunc: transaction.Wrapper(updatePostsStatus)},
}
