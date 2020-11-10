package service

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/closetool/blog/services/postsservice/models/po"
	"github.com/closetool/blog/services/postsservice/models/vo"
	uservo "github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/closetool/blog/utils/sessionutils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"xorm.io/xorm"
)

func savePostsComments(c *gin.Context, session *xorm.Session) error {
	user, err := sessionutils.GetSession(c)
	if err != nil {
		reply.CreateJSONError(c, reply.AccountNotExist)
		return err
	}

	postsCommentsVO := &vo.PostsComments{}
	err = c.ShouldBind(postsCommentsVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}
	if postsCommentsVO.Content == "" {
		reply.CreateJSONError(c, reply.ParamError)
		return fmt.Errorf("param error")
	}

	commentsPO := &po.PostsComments{
		AuthorId: user.Id,
		Content:  postsCommentsVO.Content,
		PostsId:  postsCommentsVO.PostsId,
	}

	treePath := ""
	if postsCommentsVO.ParentId == 0 {
		session.InsertOne(commentsPO)
		treePath = fmt.Sprintf("%d%s", commentsPO.Id, constants.TreePath)
	} else {
		parentComments := &po.PostsComments{}
		session.ID(postsCommentsVO.ParentId).Get(parentComments)
		if parentComments.Id == 0 {
			reply.CreateJSONError(c, reply.DataNoExist)
			return fmt.Errorf("no parent comments")
		}

		commentsPO.ParentId = parentComments.Id
		session.InsertOne(commentsPO)

		treePath = fmt.Sprintf("%s%d%s", parentComments.TreePath, commentsPO.Id, constants.TreePath)
	}

	commentsPO.TreePath = treePath
	session.ID(commentsPO.Id).Update(commentsPO)
	err = incrementComments(session, commentsPO.PostsId)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return err
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func replyComments(c *gin.Context, session *xorm.Session) error {
	commentsVO := &vo.PostsComments{}
	err := c.ShouldBindJSON(commentsVO)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return err
	}
	rpl, err := messaging.Client.PublishOnQueueWaitReply([]byte{}, "auth.selectAdmin")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return err
	}

	user := uservo.AuthUser{}
	jsoniter.Get(rpl, "model").ToVal(&user)
	if user.Id == 0 {
		reply.CreateJSONError(c, reply.AccountNotExist)
		return fmt.Errorf("can not get admin")
	}
	commentsPO := po.PostsComments{}
	if _, err := session.ID(commentsVO.ParentId).Get(&commentsPO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}

	commentsPO.ParentId = commentsVO.ParentId
	commentsPO.Content = commentsVO.Content
	commentsPO.AuthorId = user.Id
	commentsPO.CreateTime = time.Time{}

	session.InsertOne(commentsPO)
	commentsPO.TreePath = fmt.Sprintf("%s%d%s", commentsPO.TreePath, commentsPO.Id, constants.TreePath)
	if _, err := session.ID(commentsPO.Id).Update(commentsPO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}

	err = incrementComments(session, commentsPO.PostsId)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return err
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func incrementComments(session *xorm.Session, posts_id int64) error {
	_, err := session.Exec("update closetool_posts set comments = comments+1 where id = ?", posts_id)
	return err
}

func deletePostsComments(c *gin.Context, session *xorm.Session) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	if _, err := session.ID(id).Delete(&po.PostsComments{}); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func getPostsComments(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}

	comments := po.PostsComments{}
	if _, err = db.DB.ID(id).Get(&comments); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	parentIds := []interface{}{}
	userIds := []int64{}

	parentIds = append(parentIds, comments.ParentId)
	userIds = append(userIds, comments.AuthorId)

	parentCommentsList := map[int64]po.PostsComments{}
	if err := db.DB.In("id", parentIds...).Find(&parentCommentsList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	for _, parentComments := range parentCommentsList {
		userIds = append(userIds, parentComments.AuthorId)
	}

	bts, err := jsoniter.Marshal(userIds)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, "auth.getUserById")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	authors := map[int64]uservo.AuthUser{}
	jsoniter.Get(rpl, "model").ToVal(authors)

	posts := po.Posts{}
	if _, err := db.DB.ID(comments.PostsId).Get(&posts); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	author := authors[comments.AuthorId]
	parentAuthor := authors[parentCommentsList[comments.ParentId].AuthorId]
	commentsVO := vo.PostsComments{
		Id:             comments.Id,
		Content:        comments.Content,
		CreateTime:     &models.JSONTime{comments.CreateTime},
		AuthorName:     author.Name,
		AuthorAvatar:   author.Avatar,
		ParentUserName: parentAuthor.Name,
		Title:          posts.Title,
	}
	reply.CreateJSONModel(c, commentsVO)
}

func getPostsCommentsByPostsId(c *gin.Context) {
	commentsVO := vo.PostsComments{}
	err := c.ShouldBindJSON(&commentsVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}

	page := pageutils.CheckAndInitPage(commentsVO.BaseVO)
	commentsList := []po.PostsComments{}
	if page.Total, err = db.DB.Desc("create_time").Where("posts_id = ?", commentsVO.PostsId).FindAndCount(&commentsList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	parentIds := []interface{}{}
	userIds := []int64{}
	for _, comments := range commentsList {
		parentIds = append(parentIds, comments.ParentId)
		userIds = append(userIds, comments.AuthorId)
	}

	parentCommentsList := map[int64]po.PostsComments{}
	if err := db.DB.In("id", parentIds...).Find(&parentCommentsList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	for _, parentComments := range parentCommentsList {
		userIds = append(userIds, parentComments.AuthorId)
	}

	bts, err := jsoniter.Marshal(userIds)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, "auth.getUserById")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	authors := map[int64]uservo.AuthUser{}
	jsoniter.Get(rpl, "model").ToVal(authors)
	commentsVOList := []interface{}{}
	for _, comments := range commentsList {
		author := authors[comments.AuthorId]
		parentAuthor := authors[parentCommentsList[comments.ParentId].AuthorId]
		commentsVO := vo.PostsComments{
			Id:             comments.Id,
			Content:        comments.Content,
			CreateTime:     &models.JSONTime{comments.CreateTime},
			AuthorName:     author.Name,
			AuthorAvatar:   author.Avatar,
			ParentUserName: parentAuthor.Name,
		}
		commentsVOList = append(commentsVOList, commentsVO)
	}
	reply.CreateJSONPaging(c, commentsVOList, page)
}

//太麻烦了 复制粘贴 有缘再改吧
func getPostsCommentsList(c *gin.Context) {
	commentsVO := vo.PostsComments{}
	err := c.ShouldBindJSON(&commentsVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}

	page := pageutils.CheckAndInitPage(commentsVO.BaseVO)
	commentsList := []po.PostsComments{}
	if commentsVO.BaseVO != nil && commentsVO.Keywords != "" {
		db.DB.Where("closetool_posts_comments.content like", "%"+commentsVO.Keywords+"%")
	}
	if commentsVO.Id != 0 {
		db.DB.Where("closetool_posts_comments.id = ?", commentsVO.Id)
	}
	if page.Total, err = db.DB.Desc("create_time").Where("posts_id = ?", commentsVO.PostsId).FindAndCount(&commentsList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	parentIds := []interface{}{}
	userIds := []int64{}
	for _, comments := range commentsList {
		parentIds = append(parentIds, comments.ParentId)
		userIds = append(userIds, comments.AuthorId)
	}

	parentCommentsList := map[int64]po.PostsComments{}
	if err := db.DB.In("id", parentIds...).Find(&parentCommentsList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	for _, parentComments := range parentCommentsList {
		userIds = append(userIds, parentComments.AuthorId)
	}

	bts, err := jsoniter.Marshal(userIds)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, "auth.getUserById")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	authors := map[int64]uservo.AuthUser{}
	jsoniter.Get(rpl, "model").ToVal(authors)
	commentsVOList := []interface{}{}
	for _, comments := range commentsList {
		author := authors[comments.AuthorId]
		parentAuthor := authors[parentCommentsList[comments.ParentId].AuthorId]
		commentsVO := vo.PostsComments{
			Id:             comments.Id,
			Content:        comments.Content,
			CreateTime:     &models.JSONTime{comments.CreateTime},
			AuthorName:     author.Name,
			AuthorAvatar:   author.Avatar,
			ParentUserName: parentAuthor.Name,
		}
		commentsVOList = append(commentsVOList, commentsVO)
	}
	reply.CreateJSONPaging(c, commentsVOList, page)
}
