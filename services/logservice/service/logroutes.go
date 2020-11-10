package service

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
)

var LogRoutes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: Health},
	{Method: "GET", Pattern: "/logs/v1/:id", MiddleWare: nil, HandlerFunc: getLogs},
	{Method: "DELETE", Pattern: "/logs/v1/:id", MiddleWare: nil, HandlerFunc: transaction.Wrapper(deleteLogs)},
	{Method: "GET", Pattern: "/list/v1/logs", MiddleWare: nil, HandlerFunc: getLogsList},
}
