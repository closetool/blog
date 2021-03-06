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
	{Method: "DELETE", Pattern: "/user/v1/:id", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(deleteUser)},
	{Method: "PUT", Pattern: "/status/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(saveAuthUserStatus)},
	{Method: "GET", Pattern: "/master/v1/get", MiddleWare: nil, HandlerFunc: getMasterUserInfo},
	{Method: "GET", Pattern: "/user/v1/list", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: getUserList},
	{Method: "GET", Pattern: "/github/v1/get", MiddleWare: nil, HandlerFunc: oathLoginByGithub},
	{Method: "POST", Pattern: "/user/v1/login", MiddleWare: nil, HandlerFunc: transaction.GormTx(saveUserByGithub)},
	{Method: "POST", Pattern: "/admin/v1/register", MiddleWare: nil, HandlerFunc: transaction.GormTx(registerAdminByGithub)},
	{Method: "POST", Pattern: "/admin/v1/login", MiddleWare: nil, HandlerFunc: transaction.GormTx(login)},
	{Method: "PUT", Pattern: "/password/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(updatePassword)},
	{Method: "PUT", Pattern: "/admin/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(updateAdmin)},
	{Method: "PUT", Pattern: "/user/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(updateUser)},
	{Method: "POST", Pattern: "/auth/v1/logout", MiddleWare: gin.HandlersChain{middlewares.UserToken}, HandlerFunc: transaction.GormTx(logout)},
	//TODO
	{Method: "GET", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},
	{Method: "POST", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},
	{Method: "DELETE", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},
	{Method: "PUT", Pattern: "/auth/v1/avatar", MiddleWare: nil, HandlerFunc: getAvatar},

	{Method: "POST", Pattern: "/social/v1/add", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(saveSocial)},
	{Method: "PUT", Pattern: "/social/v1/update", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: transaction.GormTx(editSocial)},
	{Method: "GET", Pattern: "/social/v1/:id", MiddleWare: nil, HandlerFunc: getSocial},
	{Method: "DELETE", Pattern: "/social/v1/:id", MiddleWare: nil, HandlerFunc: transaction.GormTx(delSocial)},
	//FIXME
	//{Method: "GET", Pattern: "/social/v1/list", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: getSocialList},
	{Method: "GET", Pattern: "/list/v1/social", MiddleWare: gin.HandlersChain{middlewares.AdminToken}, HandlerFunc: getSocialList},
	//TODO
	{Method: "GET", Pattern: "/socials/v1/social", MiddleWare: nil, HandlerFunc: getSocialEnableList},
	{Method: "GET", Pattern: "/info/v1/social", MiddleWare: nil, HandlerFunc: getSocialInfo},
	{Method: "POST", Pattern: "/email/v1/send", MiddleWare: nil, HandlerFunc: sendEmail},
}
