package service

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/closetool/blog/services/postsservice/models/po"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/models/dao"
	"github.com/closetool/blog/system/models/model"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/closetool/blog/utils/previewtextutils"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Health(c *gin.Context) {
	if db.Gorm == nil {
		c.JSON(http.StatusOK, map[string]bool{"health": false})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

type PostsAndTags struct {
	po.Posts     `xorm:"extends"`
	po.PostsTags `xorm:"extends"`
}

func (p PostsAndTags) TableName() string {
	return "closetool_posts"
}

func getPostsListWeight(c *gin.Context, IsWeight int64) {
	post := model.Posts{}
	c.ShouldBindQuery(&post)
	logrus.Debugf("%#v", post)
	logrus.Debugf("%#v", post.BaseVO)
	page := pageutils.CheckAndInitPage(post.BaseVO)
	logrus.Debugf("%#v", page)

	post.IsWeight = IsWeight

	posts := make([]model.Posts, 0)

	db.Gorm.Model(&post).Preload("PostsTags").Scopes(dao.PostsCond(&post)).Count(&page.Total).Scopes(dao.Paginate(page)).Find(&posts)

	var categoryNames, userNames = make(map[int64]string), make(map[int64]string)
	categoryIds := make([]int64, len(posts))
	userIds := make([]int64, len(posts))
	for i, post := range posts {
		categoryIds[i] = post.CategoryID.Int64
		userIds[i] = post.AuthorID.Int64
	}

	temp, err := jsoniter.Marshal(categoryIds)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	rpl, err := messaging.Client.PublishOnQueueWaitReply(temp, "category.getCategoryNameById")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		logrus.Debugln(string(rpl))
		reply.CreateJSONError(c, reply.Error)
		return
	}
	jsoniter.Get(rpl, "model").ToVal(&categoryNames)

	temp, err = jsoniter.Marshal(userIds)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	rpl, err = messaging.Client.PublishOnQueueWaitReply(temp, "auth.getUserNameById")
	logrus.Debugln(string(rpl))
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
	}
	jsoniter.Get(rpl, "model").ToVal(&userNames)

	for i, post := range posts {

		posts[i].Author = userNames[post.AuthorID.Int64]
		posts[i].CategoryName = categoryNames[post.CategoryID.Int64]

		tagsList := make([]model.Tags, 0)
		tagsIds := make([]int64, 0)
		for _, postsTags := range post.PostsTags {
			tagsIds = append(tagsIds, postsTags.TagsID)
		}

		if len(tagsIds) != 0 {
			bts, err := jsoniter.Marshal(tagsIds)
			if err != nil {
				reply.CreateJSONError(c, reply.Error)
				return
			}
			rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, "tags.getTagsByIds")
			if err != nil || !bytes.Contains(rpl, []byte("00000")) {
				reply.CreateJSONError(c, reply.Error)
				return
			}
			jsoniter.Get(rpl, "models").ToVal(&tagsList)
			posts[i].TagsList = tagsList
		}
	}

	ints := model.Posts2Interfaces(posts)

	reply.CreateJSONPaging(c, ints, page)
}

func getPostsList(c *gin.Context) {
	getPostsListWeight(c, 0)
}

func getWeightList(c *gin.Context) {
	getPostsListWeight(c, 1)
}

func getArchiveTotalByDateList(c *gin.Context) {

	posts := make([]model.Posts, 0)

	if rows, err := db.Gorm.Raw(`SELECT
			FROM_UNIXTIME( create_time/1000, "%Y-%m-01 00:00:00" ) archiveDate,
			COUNT(*) articleTotal
			FROM
			posts	
			GROUP BY FROM_UNIXTIME( create_time/1000, "%Y-%m-01 00:00:00" )`).Rows(); err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("select archive data from db failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var (
				archiveDate  time.Time
				articleTotal int64
			)
			rows.Scan(&archiveDate, &articleTotal)
			posts = append(posts, model.Posts{
				ArchiveDate:  &models.JSONTime{archiveDate},
				ArchiveTotal: articleTotal,
			})
		}
	}

	for i, post := range posts {
		archivePosts := make([]model.Posts, 0)
		if err := db.Gorm.
			Where(`DATE_FORMAT( create_time,"%Y-%m-01 00:00:00")=DATE_FORMAT(?, "%Y-%m-01 00:00:00" )`, post.ArchiveDate).
			Find(&archivePosts).Error; err != nil && err != gorm.ErrRecordNotFound {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return
		}

		posts[i].ArchivePosts = archivePosts
	}
	reply.CreateJSONModels(c, model.Posts2Interfaces(posts))
}

func getHotPostsList(c *gin.Context) {
	post := &model.Posts{}
	if err := c.ShouldBindQuery(post); err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}
	page := pageutils.CheckAndInitPage(post.BaseVO)

	rpl, err := messaging.Client.PublishOnQueueWaitReply([]byte(constants.PostsDetail), "logs.getParamGroupByCode")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	if !bytes.Contains(rpl, []byte("models")) {
		reply.CreateJSONsuccess(c)
		return
	}

	logs := make([]model.AuthUserLog, 0)
	jsoniter.Get(rpl, "models").ToVal(&logs)
	ids := make([]interface{}, 0)
	for _, log := range logs {
		id := jsoniter.Get([]byte(log.Parameter.String), "id").ToInt64()
		ids = append(ids, id)
	}

	hotPosts := make([]model.Posts, 0)
	if err := db.Gorm.Where("id in ?", ids).Count(&page.Total).Scopes(dao.Paginate(page)).Find(&hotPosts).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Errorf("find posts by id failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	}

	reply.CreateJSONPaging(c, model.Posts2Interfaces(hotPosts), page)
}

func savePosts(c *gin.Context, tx *gorm.DB) {
	post := model.Posts{
		Status: 2,
	}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	s, exist := c.Get("session")
	user, ok := s.(model.AuthUser)
	if !exist || !ok {
		panic(reply.AccessNoPrivilege)
	}

	html := markdown.ToHTML([]byte(post.Content), nil, nil)
	post.Summary = previewtextutils.GetText(string(html), 126)
	post.AuthorID.Scan(user.ID)
	post.PostsAttribute.Content = post.Content

	if err := tx.Create(&post).Error; err != nil {
		logrus.Errorf("failed to insert into db: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	if post.TagsList == nil {
		reply.CreateJSONsuccess(c)
	}

	bts, err := jsoniter.Marshal(post.TagsList)
	if err != nil {
		panic(reply.Error)
	}

	var rpl []byte
	if rpl, err = messaging.Client.PublishOnQueueWaitReply(bts, "tags.addTags"); err != nil {
		panic(reply.Error)
	}
	if !bytes.Contains(rpl, []byte("00000")) || !bytes.Contains(rpl, []byte("models")) {
		panic(reply.Error)
	}
	ids := make([]int64, 0)
	jsoniter.Get(rpl, "models").ToVal(&ids)
	post.PostsTags = make([]model.PostsTags, 0)
	for _, id := range ids {
		postsTag := model.PostsTags{TagsID: id, PostsID: post.ID}
		post.PostsTags = append(post.PostsTags, postsTag)
	}

	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&post).Error; err != nil {
		logrus.Errorf("update post failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONsuccess(c)
}

func getPosts(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}
	post := model.Posts{}
	if err := db.Gorm.Preload("PostsTags").Preload("PostsAttribute").First(&post, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			panic(reply.DataNoExist)
		default:
			logrus.Errorf("get post from db failed: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	}

	post.Content = post.PostsAttribute.Content

	bts, err := jsoniter.Marshal([]int64{post.CategoryID.Int64})
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	var rpl []byte
	if rpl, err = messaging.Client.PublishOnQueueWaitReply(bts, "category.getCategoryNameById"); err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	if !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}

	categoryNames := map[int64]string{}
	jsoniter.Get(rpl, "model").ToVal(&categoryNames)
	post.CategoryName = categoryNames[post.CategoryID.Int64]

	ids := []int64{}
	for _, postsTag := range post.PostsTags {
		ids = append(ids, postsTag.ID)
	}

	bts, err = jsoniter.Marshal(ids)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	if rpl, err = messaging.Client.PublishOnQueueWaitReply(bts, "tags.getTagsByIds"); err != nil {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	if !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}

	tagsList := []model.Tags{}
	jsoniter.Get(rpl, "models").ToVal(tagsList)
	post.TagsList = tagsList

	post.Views++
	if err := db.Gorm.Select("views").Updates(&post).Error; err != nil {
		logrus.Errorf("update post failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}
	reply.CreateJSONModel(c, post)
}

func deletePosts(c *gin.Context, tx *gorm.DB) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	if _, err := dao.DeletePosts(tx, id); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			panic(reply.DataNoExist)
		default:
			logrus.Errorf("can not delete post from db: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	}

	reply.CreateJSONsuccess(c)
}

func updatePosts(c *gin.Context, tx *gorm.DB) {
	s, exist := c.Get("session")
	user, ok := s.(model.AuthUser)
	if !exist || !ok {
		panic(reply.AccessNoPrivilege)
	}

	post := model.Posts{}
	err := c.ShouldBindJSON(&post)
	if err != nil || post.ID == 0 {
		panic(reply.ParamError)
	}

	html := markdown.ToHTML([]byte(post.Content), nil, nil)
	post.Summary = previewtextutils.GetText(string(html), 126)
	post.AuthorID.Scan(user.ID)
	if err := tx.Updates(&post).Error; err != nil {
		logrus.Errorf("can not update post: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	//tx.Model(&post).Association("PostsTags").Replace(&post.PostsTags)

	post.PostsAttribute.Content = post.Content
	if err := tx.Model(&post).Association("PostsAttribute").Replace(&post.PostsAttribute).Error; err != nil {
		logrus.Errorf("can not update posts attribute: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	tagsList := post.TagsList

	bts, _ := jsoniter.Marshal(tagsList)
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, "tags.addTags")
	if err != nil {
		logrus.Errorf("add tags failed: %v", err)
		panic(reply.Error)
	}
	if !bytes.Contains(rpl, []byte("00000")) {
		logrus.Errorf("add tags failed: %v", string(rpl))
		panic(reply.Error)
	}

	ids := make([]int64, 0)
	jsoniter.Get(rpl, "models").ToVal(&ids)

	post.PostsTags = []model.PostsTags{}

	for _, id := range ids {
		post.PostsTags = append(post.PostsTags, model.PostsTags{PostsID: post.ID, TagsID: id})
	}

	if err := tx.Model(&post).Association("PostsAttribute").Replace(&post.PostsTags).Error; err != nil {
		logrus.Errorf("can not update posts tags: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	reply.CreateJSONsuccess(c)
}

func updatePostsStatus(c *gin.Context, tx *gorm.DB) {
	post := model.Posts{}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		logrus.Errorf("binding param failed: %v", err)
		panic(reply.ParamError)
	}

	if post.ID == 0 {
		logrus.Errorf("param err")
		panic(reply.ParamError)
	}

	if err := tx.Model(&post).Where("id = ?", post.ID).Update("status", post.Status).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logrus.Errorf("there's no data: %v", post.ID)
			panic(reply.DataNoExist)
		default:
			logrus.Errorf("can not update post's status: %v", err)
			panic(reply.DatabaseSqlParseError)
		}
	}

	reply.CreateJSONsuccess(c)
}

func getArchiveGrouopYearList(c *gin.Context) {
	posts := []model.Posts{}
	rows, err := db.Gorm.Raw(`select
	id,
	title,
	create_time,
	FROM_UNIXTIME(create_time,"%Y") year 
	FROM posts
	order by
	FROM_UNIXTIME(create_time,"%Y") DESC`).Rows()
	if err != nil {
		logrus.Errorf("select posts from db failed: %v", err)
		panic(reply.DatabaseSqlParseError)
	}

	for rows.Next() {
		var (
			id         int64
			title      string
			createTime int64
			year       int64
		)
		rows.Scan(&id, &title, &createTime, &year)
		posts = append(posts, model.Posts{
			ID:         id,
			Title:      title,
			CreateTime: createTime,
			Year:       year,
		})
	}

	reply.CreateJSONModels(c, model.Posts2Interfaces(posts))
}
