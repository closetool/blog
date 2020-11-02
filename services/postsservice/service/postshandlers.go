package service

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	categoryvo "github.com/closetool/blog/services/categoryservice/models/vo"
	logsvo "github.com/closetool/blog/services/logservice/models/vo"
	"github.com/closetool/blog/services/postsservice/models/po"
	"github.com/closetool/blog/services/postsservice/models/vo"
	uservo "github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/models"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/closetool/blog/utils/previewtextutils"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func Health(c *gin.Context) {
	if db.DB == nil {
		c.JSON(http.StatusOK, map[string]bool{"health": false})
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
	postsVO := vo.Posts{Status: -1}
	c.ShouldBindQuery(&postsVO)
	logrus.Debugf("%#v", postsVO)
	logrus.Debugf("%#v", postsVO.BaseVO)
	page := pageutils.CheckAndInitPage(postsVO.BaseVO)
	logrus.Debugf("%#v", page)

	postsVO.IsWeight = IsWeight

	session := db.DB.NewSession()

	if postsVO.BaseVO != nil && postsVO.Keywords != "" {
		session = session.Where("title like ?", "%"+postsVO.Keywords+"%")
	}
	if postsVO.Id != 0 {
		session = session.Where("closetool_posts.id = ?", postsVO.Id)
	}
	if postsVO.CreateTime != nil {
		session = session.Where("closetool_posts.create_time = ?", postsVO.CreateTime)
	}
	if postsVO.CategoryId != 0 {
		session = session.Where("category_id = ?", postsVO.CategoryId)
	}
	if postsVO.PostsTagsId != 0 {
		session = session.Where("closetool_posts_tags.tags_id = ?", postsVO.PostsTagsId)
	}
	if postsVO.Title != "" {
		session = session.Where("title = ?", postsVO.Title)
	}
	if postsVO.Status != -1 {
		session = session.Where("status = ?", postsVO.Status)
	}
	if postsVO.IsWeight != 0 {
		session = session.Desc("weight")
	} else {
		session = session.Desc("closetool_posts.id")
	}

	session = session.Limit(pageutils.StartAndEnd(page))
	session = session.
		Table("closetool_posts").
		Join("LEFT OUTER", "closetool_posts_tags", "closetool_posts.id=closetool_posts_tags.posts_id")
	var err error
	//page.Total, err = session.Count(&PostsAndTags{})
	//logrus.Debugln(page.Total, err)

	POs := make([]PostsAndTags, 0)
	if page.Total, err = session.
		//Distinct("closetool_posts.id").
		FindAndCount(&POs); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		logrus.Debug(err)
		return
	}

	logrus.Debugf("%d %#v", page.Total, POs)

	var categoryNames, userNames = make(map[int64]string), make(map[int64]string)
	categoryIds := make([]int64, len(POs))
	userIds := make([]int64, len(POs))
	for i, PO := range POs {
		categoryIds[i] = PO.CategoryId
		userIds[i] = PO.AuthorId
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

	postsVOs := make([]interface{}, 0)
	for _, PO := range POs {
		temp := vo.Posts{
			Id:           PO.Posts.Id,
			Title:        PO.Title,
			Status:       PO.Status,
			Summary:      PO.Summary,
			Thumbnail:    PO.Thumbnail,
			Author:       userNames[PO.AuthorId],
			Views:        PO.Views,
			Comments:     PO.Comments,
			CategoryId:   PO.CategoryId,
			CategoryName: categoryNames[PO.CategoryId],
			Weight:       PO.Weight,
			CreateTime:   &models.JSONTime{PO.Posts.CreateTime},
		}
		postsTagsList := make([]po.PostsTags, 0)
		db.DB.Where("posts_id = ?", temp.Id).Find(&postsTagsList)
		tagsList := make([]*categoryvo.Tags, 0)
		tagsIds := make([]int64, 0)
		for _, postsTags := range postsTagsList {
			tagsIds = append(tagsIds, postsTags.TagsId)
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
			temp.TagsList = tagsList
		}
		postsVOs = append(postsVOs, temp)
	}
	reply.CreateJSONPaging(c, postsVOs, page)
}

func getPostsList(c *gin.Context) {
	getPostsListWeight(c, 0)
}

func getWeightList(c *gin.Context) {
	getPostsListWeight(c, 1)
}

func getArchiveTotalByDateList(c *gin.Context) {

	postsVOList := make([]vo.Posts, 0)
	var (
		err error
		m   []map[string]string
	)
	if m, err = db.DB.SQL(`SELECT
			DATE_FORMAT( create_time, "%Y-%m-01 00:00:00" ) archiveDate,
			COUNT(*) articleTotal
			FROM
			closetool_posts
			GROUP BY DATE_FORMAT( create_time, "%Y-%m-01 00:00:00" )`).QueryString(); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	for _, data := range m {
		tm, err := time.Parse("2006-01-02 15:04:05", data["archiveDate"])
		if err != nil {
			reply.CreateJSONError(c, reply.Error)
			return
		}
		count, err := strconv.ParseInt(data["articleTotal"], 10, 64)
		if err != nil {
			reply.CreateJSONError(c, reply.Error)
			return
		}
		postsVOList = append(postsVOList, vo.Posts{
			ArticleTotal: count,
			ArchiveDate:  &models.JSONTime{tm},
		})
	}

	for i, postsVO := range postsVOList {
		if err != nil {
			reply.CreateJSONError(c, reply.Error)
			return
		}
		postsPOs := make([]po.Posts, 0)
		if err := db.DB.
			Where(`DATE_FORMAT( create_time,"%Y-%m-01 00:00:00")=DATE_FORMAT(?, "%Y-%m-01 00:00:00" )`, postsVO.ArchiveDate).
			Find(&postsPOs); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return
		}

		postsVOs := make([]*vo.Posts, 0)
		for _, postsPO := range postsPOs {
			postsVOs = append(postsVOs, &vo.Posts{
				Id:         postsPO.Id,
				AuthorId:   postsPO.AuthorId,
				Title:      postsPO.Title,
				Thumbnail:  postsPO.Thumbnail,
				Comments:   postsPO.Comments,
				IsComment:  postsPO.IsComment,
				CategoryId: postsPO.CategoryId,
				SyncStatus: postsPO.SyncStatus,
				Status:     postsPO.Status,
				Summary:    postsPO.Summary,
				Views:      postsPO.Views,
				Weight:     postsPO.Weight,
				CreateTime: &models.JSONTime{postsPO.CreateTime},
				UpdateTime: &models.JSONTime{postsPO.UpdateTime},
			})
		}
		postsVOList[i].ArchivePosts = postsVOs
	}
	temp := make([]interface{}, 0)
	for _, postsVO := range postsVOList {
		temp = append(temp, postsVO)
	}
	reply.CreateJSONModels(c, temp)
}

//TODO:先实现logservice
func getHotPostsList(c *gin.Context) {
	postsVO := &vo.Posts{}
	c.ShouldBindQuery(postsVO)
	page := pageutils.CheckAndInitPage(postsVO.BaseVO)

	rpl, err := messaging.Client.PublishOnQueueWaitReply([]byte(constants.PostsDetail), "logs.getParamGroupByCode")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
		return
	}
	if !bytes.Contains(rpl, []byte("models")) {
		reply.CreateJSONsuccess(c)
		return
	}
	logsVOs := make([]logsvo.AuthUserLog, 0)
	jsoniter.Get(rpl, "models").ToVal(&logsVOs)
	ids := make([]interface{}, 0)
	for _, logVO := range logsVOs {
		id := jsoniter.Get([]byte(logVO.Parameter), "id").ToInt64()
		ids = append(ids, id)
	}
	postsPOs := make([]po.Posts, 0)
	if page.Total, err = db.DB.In("id", ids...).Limit(pageutils.StartAndEnd(page)).FindAndCount(&postsPOs); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	postsVOs := make([]interface{}, 0)
	for _, postsPO := range postsPOs {
		postsVOs = append(postsVOs, &vo.Posts{
			Id:         postsPO.Id,
			AuthorId:   postsPO.AuthorId,
			Title:      postsPO.Title,
			Thumbnail:  postsPO.Thumbnail,
			Comments:   postsPO.Comments,
			IsComment:  postsPO.IsComment,
			CategoryId: postsPO.CategoryId,
			SyncStatus: postsPO.SyncStatus,
			Status:     postsPO.Status,
			Summary:    postsPO.Summary,
			Views:      postsPO.Views,
			Weight:     postsPO.Weight,
			CreateTime: &models.JSONTime{postsPO.CreateTime},
			UpdateTime: &models.JSONTime{postsPO.UpdateTime},
		})
	}

	reply.CreateJSONPaging(c, postsVOs, page)
}

func savePosts(c *gin.Context) error {
	postsVO := &vo.Posts{
		Status:    2,
		IsComment: 1,
	}
	err := c.ShouldBindJSON(postsVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	session, exist := c.Get("session")
	user, ok := session.(uservo.AuthUser)
	if !exist || !ok {
		reply.CreateJSONError(c, reply.AccessNoPrivilege)
		return fmt.Errorf("获取session失败")
	}

	html := markdown.ToHTML([]byte(postsVO.Content), nil, nil)
	postsPO := &po.Posts{
		Title:      postsVO.Title,
		Thumbnail:  postsVO.Thumbnail,
		Status:     postsVO.Status,
		Summary:    previewtextutils.GetText(string(html), 126),
		IsComment:  postsVO.IsComment,
		AuthorId:   user.Id,
		CategoryId: postsVO.CategoryId,
		Weight:     postsVO.Weight,
	}
	if _, err := db.DB.InsertOne(postsPO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}
	postsAttributePO := &po.PostsAttribute{
		Content: postsVO.Content,
		PostsId: postsPO.Id,
	}
	if _, err := db.DB.InsertOne(postsAttributePO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}
	if postsVO.TagsList == nil {
		reply.CreateJSONsuccess(c)
		return nil
	}

	bts, err := jsoniter.Marshal(postsVO.TagsList)
	if err != nil {
		reply.CreateJSONError(c, reply.Error)
		return err
	}

	var rpl []byte
	if rpl, err = messaging.Client.PublishOnQueueWaitReply(bts, "tags.addTags"); err != nil {
		reply.CreateJSONError(c, reply.Error)
		return err
	}
	if !bytes.Contains(rpl, []byte("00000")) || !bytes.Contains(rpl, []byte("models")) {
		reply.CreateJSONError(c, reply.Error)
		return fmt.Errorf("create tags failed")
	}
	ids := make([]int64, 0)
	jsoniter.Get(rpl, "models").ToVal(&ids)
	for _, id := range ids {
		postsTagsPO := &po.PostsTags{TagsId: id, PostsId: postsPO.Id}
		if _, err := db.DB.InsertOne(postsTagsPO); err != nil {
			reply.CreateJSONError(c, reply.Error)
			return err
		}
	}
	reply.CreateJSONsuccess(c)
	return nil
}
