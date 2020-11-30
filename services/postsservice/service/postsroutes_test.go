package service

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
)

var PostsRoutesTest = []models.Route{
	{Method: "GET", Pattern: "/list/v1/posts", MiddleWare: nil, HandlerFunc: getPostsList},
	{Method: "GET", Pattern: "/weight/v1/list", MiddleWare: nil, HandlerFunc: getWeightList},
	{Method: "GET", Pattern: "/archive/v1/list", MiddleWare: nil, HandlerFunc: getArchiveTotalByDateList},
	{Method: "GET", Pattern: "/hot/v1/list", MiddleWare: nil, HandlerFunc: getHotPostsList},
	{Method: "POST", Pattern: "/posts/v1/add", MiddleWare: nil, HandlerFunc: transaction.GormTx(savePosts)},
	{Method: "GET", Pattern: "/posts/v1/:id", MiddleWare: nil, HandlerFunc: getPosts},
	{Method: "DELETE", Pattern: "/posts/v1/:id", MiddleWare: nil, HandlerFunc: transaction.GormTx(deletePosts)},
	{Method: "PUT", Pattern: "/posts/v1/update", MiddleWare: nil, HandlerFunc: transaction.GormTx(updatePosts)},
	{Method: "PUT", Pattern: "/status/v1/update", MiddleWare: nil, HandlerFunc: transaction.GormTx(updatePostsStatus)},
}
