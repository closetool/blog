package reply

type Reply struct {
	Success   int           `json:"success"`
	ReplyCode string        `json:"resultCode"`
	Message   string        `json:"message"`
	Model     interface{}   `json:"model"`
	Models    []interface{} `json:"models"`
	Pageinfo  PageInfo      `json:"pageInfo"`
	Extra     []interface{} `json:"extra"`
}

type PageInfo struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
