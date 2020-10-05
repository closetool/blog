package routes

import (
	"github.com/closetool/blog/services/musicservice/handlers"
	"github.com/closetool/blog/system/models"
)

var Routes = []models.Route{
	{Method:"GET",Pattern: "/music/v1/list",MiddleWare: nil,HandlerFunc:handlers.GetPlayList},
}