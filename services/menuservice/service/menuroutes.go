package service

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
)

var MenuRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: health},
	{Method: "POST", Pattern: "/menu/v1/add", MiddleWare: nil, HandlerFunc: transaction.GormTx(saveMenu)},
	{Method: "GET", Pattern: "/menu/v1/:id", MiddleWare: nil, HandlerFunc: getMenu},
	{Method: "GET", Pattern: "/list/v1/menu", MiddleWare: nil, HandlerFunc: getMenuList},
	{Method: "GET", Pattern: "/front/v1/list", MiddleWare: nil, HandlerFunc: getFrontMenuList},
	{Method: "PUT", Pattern: "/menu/v1/update", MiddleWare: nil, HandlerFunc: transaction.GormTx(updateMenu)},
	{Method: "DELETE", Pattern: "/menu/v1/:id", MiddleWare: nil, HandlerFunc: transaction.GormTx(deleteMenu)},
}
