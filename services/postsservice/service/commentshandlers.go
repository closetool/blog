package service

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/closetool/blog/utils/sessionutils"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func savePostsComments(c *gin.Context, tx *gorm.DB) {
	user, err := sessionutils.GetSession(c)
	if err != nil {
		logrus.Errorf("account have no privilege: %v", err)
		panic(reply.AccessNoPrivilege)
	}

	comment := model.PostsComments{}
	err = c.ShouldBindJSON(&comment)
	if err != nil {
		logrus.Errorf("can not bind param: %v", err)
		panic(reply.ParamError)
	}

	if comment.Content == "" && comment.PostsID == 0 {
		panic(reply.ParamError)
	}

	comment.AuthorID = user.ID

	treePath := ""
	if comment.ParentID == 0 {
		if _, _, err := dao.AddPostsComments(tx, &comment); err != nil {
			logrus.Errorf("insert comment failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
		treePath = fmt.Sprintf("%d%s", comment.ID, constants.TreePath)
	} else {
		parentComment := &model.PostsComments{}
		parentComment, err := dao.GetPostsComments(tx, comment.ParentID)
		if err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				logrus.Errorf("there is no parent comment: %v", err)
				panic(reply.DataNoExist)
			default:
				logrus.Errorf("can not get parent comment: %v", err)
				panic(reply.DatabaseSqlParseError)
			}
		}
		if parentComment.ID == 0 {
			logrus.Errorf("get parent comment failed")
			panic(reply.DatabaseSqlParseError)
		}

		comment.ParentID = parentComment.ID

		if _, _, err := dao.AddPostsComments(tx, &comment); err != nil {
			logrus.Errorf("insert comment failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}

		treePath = fmt.Sprintf("%s%d%s", parentComment.TreePath, comment.ID, constants.TreePath)
	}

	comment.TreePath = null.StringFrom(treePath)

	if err := tx.Where("id = ?", comment.ID).Updates(&comment).Error; err != nil {
		logrus.Errorf("can not update comment")
		panic(reply.DatabaseSqlParseError)
	}

	err = incrementComments(tx, comment.PostsID)
	if err != nil {
		logrus.Errorf("couldn't update post's comment count: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func replyComments(c *gin.Context, tx *gorm.DB) {
	comment := model.PostsComments{}
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		panic(reply.ParamError)
	}
	rpl, err := messaging.Client.PublishOnQueueWaitReply([]byte{}, "auth.selectAdmin")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		panic(reply.Error)
	}

	admin := model.AuthUser{}
	jsoniter.Get(rpl, "model").ToVal(&admin)
	if admin.ID == 0 {
		panic(reply.AccountNotExist)
	}

	commentPO := model.PostsComments{}
	if err := tx.Where("id = ?", comment.ParentID).First(&commentPO).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logrus.Errorf("no parent comment: %v", err)
			panic(reply.DataNoExist)
		default:
			logrus.Errorf("")
			panic(reply.DatabaseSqlParseError)
		}
	}

	commentPO.ParentID = comment.ParentID
	commentPO.Content = comment.Content
	commentPO.AuthorID = admin.ID

	dao.AddPostsComments(tx, &commentPO)

	commentPO.TreePath = null.StringFrom(fmt.Sprintf("%s%d%s", commentPO.TreePath, commentPO.ID, constants.TreePath))
	if _, _, err := dao.UpdatePostsComments(tx, commentPO.ID, &commentPO); err != nil {
		switch err {
		case dao.ErrNotFound:
			logrus.Errorf("there no comment %d: %v", commentPO.ID, err)
			panic(reply.DataNoExist)
		case dao.ErrUpdateFailed:
			logrus.Errorf("can not update comments: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	}

	err = incrementComments(tx, commentPO.PostsID)
	if err != nil {
		logrus.Errorf("increase comments' count failed: %v", err)
		panic(reply.Error)
	}
	reply.CreateJSONsuccess(c)
}

func incrementComments(tx *gorm.DB, posts_id int64) error {
	err := tx.Exec("update posts set comments = comments+1 where id = ?", posts_id).Error
	return err
}

func deletePostsComments(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	if _, err := dao.DeletePostsComments(tx, id); err != nil {
		logrus.Errorf("delete comment failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONsuccess(c)
}

func getPostsComments(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	comment := model.PostsComments{}

	if err := db.Gorm.Model(&model.PostsComments{}).
		Where("id = ?", id).
		Preload("ParentComments").
		Preload("Posts").
		First(&comment).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logrus.Errorf("this is no comment: %v", err)
			panic(reply.DataNoExist)
		default:
			logrus.Errorf("can not get comment: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	}

	userIDs := []int64{comment.AuthorID}
	if comment.ParentComments != nil {
		userIDs = append(userIDs, comment.ParentComments.AuthorID)
	}

	rpl, err := messaging.SendRequest("auth.getUserById", userIDs)
	if err != nil || rpl.Success == 0 {
		logrus.Errorf("amqp request failed: %v", err)
		panic(reply.Error)
	}
	users := rpl.Model.(map[int64]model.AuthUser)

	comment.AuthorName = users[comment.AuthorID].Name.String
	comment.AuthorAvatar = users[comment.AuthorID].Avatar.String
	if comment.ParentComments != nil {
		comment.ParentUserName = users[comment.ParentComments.AuthorID].Name.String
	}
	comment.Title = comment.Posts.Title

	reply.CreateJSONModel(c, comment)
}

func getPostsCommentsByPostsIdList(c *gin.Context) {
	comment := model.PostsComments{}
	err := c.ShouldBindQuery(&comment)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	page := pageutils.CheckAndInitPage(comment.BaseVO)
	comments := []model.PostsComments{}

	if err := db.Gorm.Model(&comment).
		Preload("ParentComments").
		Where("posts_id = ?", comment.PostsID).
		Count(&page.Total).
		Scopes(dao.Paginate(page)).
		Find(&comments).Error; err != nil {
		logrus.Errorf("can not get comments")
		panic(reply.DatabaseSqlParseError)
	}

	userIds := []int64{}
	for _, c := range comments {
		userIds = append(userIds, c.AuthorID)
		if c.ParentComments != nil {
			userIds = append(userIds, c.ParentComments.AuthorID)
		}
	}

	rpl, err := messaging.SendRequest("auth.getUserById", userIds)
	if err != nil || rpl.Success != 1 {
		logrus.Errorf("amqp request failed: %v", err)
		panic(reply.Error)
	}
	users := rpl.Model.(map[int64]model.AuthUser)

	for i := range comments {
		comments[i].AuthorName = users[comments[i].AuthorID].Name.String
		comments[i].AuthorAvatar = users[comments[i].AuthorID].Avatar.String
		if comments[i].ParentComments != nil {
			comments[i].ParentUserName = users[comments[i].ParentComments.AuthorID].Name.String
		}
	}

	reply.CreateJSONPaging(c, model.Comments2Ints(comments), page)
}

//太麻烦了 复制粘贴 有缘再改吧
func getPostsCommentsList(c *gin.Context) {
	comment := model.PostsComments{}
	err := c.ShouldBindQuery(&comment)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	page := pageutils.CheckAndInitPage(comment.BaseVO)
	comments := []model.PostsComments{}

	if err := db.Gorm.Model(&comment).
		Preload("ParentComments").
		Preload("Posts").
		Scopes(dao.CommentsCond(&comment)).
		Count(&page.Total).
		Scopes(dao.Paginate(page)).
		Find(&comments).Error; err != nil {
		logrus.Errorf("can not get comments")
		panic(reply.DatabaseSqlParseError)
	}

	userIds := []int64{}
	for _, c := range comments {
		userIds = append(userIds, c.AuthorID)
		if c.ParentComments != nil {
			userIds = append(userIds, c.ParentComments.AuthorID)
		}
	}

	rpl, err := messaging.SendRequest("auth.getUserById", userIds)
	if err != nil || rpl.Success != 1 {
		logrus.Errorf("amqp request failed: %v", err)
		panic(reply.Error)
	}
	users := rpl.Model.(map[int64]model.AuthUser)

	for i := range comments {
		comments[i].AuthorName = users[comments[i].AuthorID].Name.String
		comments[i].AuthorAvatar = users[comments[i].AuthorID].Avatar.String
		if comments[i].ParentComments != nil {
			comments[i].ParentUserName = users[comments[i].ParentComments.AuthorID].Name.String
		}
	}

	reply.CreateJSONPaging(c, model.Comments2Ints(comments), page)
}
