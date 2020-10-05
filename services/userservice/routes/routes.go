package routes

import (
	"github.com/closetool/blog/services/userservice/handlers"
	"github.com/closetool/blog/system/models"
)

var Routes = []models.Route{
	{Method:"GET",Pattern: "/master/v1/get",MiddleWare: nil,HandlerFunc:handlers.GetMasterUserInfo},
}