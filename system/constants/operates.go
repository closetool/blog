package constants

const (
	PostsDefault = "000"
	PostsList    = "001"
	PostsDetail  = "002"
)

var OperationNames = map[string]string{
	PostsDefault: "默认类型",
	PostsList:    "查询文章列表",
	PostsDetail:  "查询文章详情",
}
