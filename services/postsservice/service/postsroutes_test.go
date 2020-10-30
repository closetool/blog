package service

import "github.com/closetool/blog/system/models"

var PostsRoutesTest = []models.Route{
	{Method: "GET", Pattern: "/posts/v1/list", MiddleWare: nil, HandlerFunc: getPostsList},
	{Method: "GET", Pattern: "/weight/v1/list", MiddleWare: nil, HandlerFunc: getWeightList},
	{Method: "GET", Pattern: "/archive/v1/list", MiddleWare: nil, HandlerFunc: getArchiveTotalByDateList},
	{Method: "GET", Pattern: "/hot/v1/list", MiddleWare: nil, HandlerFunc: getHotPostsList},
}
