package service

import (
	"github.com/closetool/blog/services/userservice/middlewares"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/transaction"
	"github.com/gin-gonic/gin"
)

var Routes = []models.Route{
	{Method: "GET", Pattern: "/health", MiddleWare: nil, HandlerFunc: Health},
	{Method: "GET", Pattern: "/user/v1/get", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: getUserInfo},
	{Method: "DELETE", Pattern: "/user/v1/:id", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.Wrapper(deleteUser)},
	{Method: "PUT", Pattern: "/status/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.Wrapper(saveAuthUserStatus)},
	{Method: "GET", Pattern: "/master/v1/get", MiddleWare: nil, HandlerFunc: getMasterUserInfo},
	{Method: "GET", Pattern: "/user/v1/list", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: getUserList},
	{Method: "GET", Pattern: "/github/v1/get", MiddleWare: nil, HandlerFunc: oathLoginByGithub},
	{Method: "POST", Pattern: "/user/v1/login", MiddleWare: nil, HandlerFunc: transaction.Wrapper(saveUserByGithub)},
	{Method: "POST", Pattern: "/admin/v1/register", MiddleWare: nil, HandlerFunc: transaction.Wrapper(registerAdminByGithub)},
	{Method: "POST", Pattern: "/admin/v1/login", MiddleWare: nil, HandlerFunc: login},
	{Method: "PUT", Pattern: "/password/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.Wrapper(updatePassword)},
	{Method: "PUT", Pattern: "/admin/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.Wrapper(updateAdmin)},
	{Method: "PUT", Pattern: "/user/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.Wrapper(updateUser)},
	{Method: "POST", Pattern: "/auth/v1/logout", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: transaction.Wrapper(logout)},
	//TODO
	{Method: "GET", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},
	{Method: "POST", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},
	{Method: "DELETE", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},
	{Method: "PUT", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},

	{Method: "POST", Pattern: "/social/v1/add", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.Wrapper(saveSocial)},
}
