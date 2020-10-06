package service

import (
	"github.com/closetool/blog/system/models"
)

var Routes = []models.Route{
	{Method: "GET", Pattern: "/master/v1/get", MiddleWare: nil, HandlerFunc: GetMasterUserInfo},
}
