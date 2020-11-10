package service

import "github.com/closetool/blog/system/models"

var FileRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: health},
	{Method: "POST", Pattern: "/file/v1/upload", MiddleWare: nil, HandlerFunc: uploadFile},
}
