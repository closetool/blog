package pageutils

import (
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/sirupsen/logrus"
)

var defaultPage = reply.PageInfo{
	Page:  1,
	Size:  10,
	Total: 0,
}

func CheckAndInitPage(base *models.BaseVO) *reply.PageInfo {
	if base == nil {
		logrus.Infoln("page info is null")
		return &reply.PageInfo{
			Page:  defaultPage.Page,
			Size:  defaultPage.Size,
			Total: defaultPage.Total,
		}
	}

	rpl := &reply.PageInfo{
		Page:  base.Page,
		Size:  base.Size,
		Total: defaultPage.Total,
	}
	if rpl.Page == 0 {
		rpl.Page = 1
	}
	if rpl.Size == 0 {
		rpl.Size = 10
	}
	if rpl.Size > 20 {
		rpl.Size = 20
	}

	return rpl
}

func StartAndEnd(page *reply.PageInfo) (int, int) {
	if page == nil {
		return 10, 0
	}
	return int(page.Size), int((page.Page - 1) * page.Size)
}
