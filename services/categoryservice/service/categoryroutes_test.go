package service

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
)

var CategoryRoutesTest = []models.Route{
	{Method: "POST", Pattern: "/category/v1/add", MiddleWare: nil, HandlerFunc: transaction.Wrapper(saveCategory)},
	{Method: "GET", Pattern: "/statistics/v1/list", MiddleWare: nil, HandlerFunc: statisticsList},
	{Method: "PUT", Pattern: "/category/v1/update", MiddleWare: nil, HandlerFunc: transaction.Wrapper(updateCategory)},
	{Method: "GET", Pattern: "/category-tags/v1/:id", MiddleWare: nil, HandlerFunc: getCategoryTags},
	{Method: "GET", Pattern: "/list/v1/category-tags", MiddleWare: nil, HandlerFunc: getCategoryTagsList},
	{Method: "GET", Pattern: "/category/v1/:id", MiddleWare: nil, HandlerFunc: getCategory},
	{Method: "GET", Pattern: "/list/v1/category", MiddleWare: nil, HandlerFunc: getCategoryList},
	{Method: "DELETE", Pattern: "/category/v1/:id", MiddleWare: nil, HandlerFunc: transaction.Wrapper(deleteCategory)},
}
