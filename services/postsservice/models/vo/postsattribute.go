package vo

type PostsAttribute struct {
	Id      int64  `form:"id" json:"id,omitempty"`
	Content string `form:"content" json:"content,omitempty"`
	PostsId int64  `form:"postsId" json:"postsId,omitempty"`
}
