package po

type PostsAttribute struct {
	Id      int64  `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Content string `xorm:"not null comment('内容') LONGTEXT"`
	PostsId int64  `xorm:"not null comment('文章表主键') BIGINT(20)"`
}
