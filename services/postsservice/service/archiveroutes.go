package service

import "github.com/closetool/blog/system/models"

var ArchiveRoutes = []models.Route{
	{Method: "GET", Pattern: "/archive/v1/list", MiddleWare: nil, HandlerFunc: getArchiveTotalByDateList},
	{Method: "GET", Pattern: "/year/v1/list", MiddleWare: nil, HandlerFunc: getArchiveGrouopYearList},
}
