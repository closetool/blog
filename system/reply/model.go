package reply

import "github.com/closetool/blog/system/models"

type Reply struct {
	Success   int           `json:"success"`
	ReplyCode string        `json:"resultCode"`
	Message   string        `json:"message"`
	Model     interface{}   `json:"model,omitempty"`
	Models    []interface{} `json:"models,omitempty"`
	Pageinfo  PageInfo      `json:"pageInfo"`
	Extra     []interface{} `json:"extra,omitempty"`
}

type PageInfo struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}

func GetPageInfo(basevo *models.BaseVO) PageInfo {
	return PageInfo{
		Page: basevo.Page,
		Size: basevo.Size,
	}
}
