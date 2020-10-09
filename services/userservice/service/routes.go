package service

import (
	"github.com/closetool/blog/services/userservice/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/gin-gonic/gin"
)

var Routes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: Health},
	{Method: "GET", Pattern: "/user/v1/get", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: getUserInfo},
	{Method: "DELETE", Pattern: "/user/v1/:id", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: deleteUser},
	{Method: "PUT", Pattern: "/status/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: saveAuthUserStatus},
	{Method: "GET", Pattern: "/master/v1/get", MiddleWare: nil, HandlerFunc: getMasterUserInfo},
	{Method: "GET", Pattern: "/user/v1/list", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: getUserList},
}
