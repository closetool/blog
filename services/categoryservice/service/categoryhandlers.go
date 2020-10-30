package service

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/closetool/blog/services/categoryservice/models/po"
	"github.com/closetool/blog/services/categoryservice/models/vo"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]bool{"health": true})
}

func saveCategory(c *gin.Context) error {
	categoryVO := &vo.Category{}
	err := c.ShouldBindJSON(categoryVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}
	categoryPO := &po.Category{
		Name: categoryVO.Name,
	}

	if ok, err := db.DB.Where("name = ?", categoryPO.Name).Get(categoryPO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	} else if !ok {
		if _, err := db.DB.InsertOne(categoryPO); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return err
		}
	}

	for _, tag := range categoryVO.TagsList {
		tagPO := &po.Tags{
			Name: tag.Name,
		}
		if ok, err := db.DB.Where("name = ?", tagPO.Name).Get(tagPO); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return err
		} else if !ok {
			if _, err := db.DB.InsertOne(tagPO); err != nil {
				reply.CreateJSONError(c, reply.DatabaseSqlParseError)
				return err
			}
		}

		categoryTagsPO := &po.CategoryTags{
			TagsId:     tagPO.Id,
			CategoryId: categoryPO.Id,
		}

		if ok, err := db.DB.Where("tags_id = ? and category_id = ?", categoryTagsPO.TagsId, categoryTagsPO.CategoryId).Get(categoryTagsPO); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return err
		} else if !ok {
			if _, err := db.DB.InsertOne(categoryTagsPO); err != nil {
				reply.CreateJSONError(c, reply.DatabaseSqlParseError)
				return err
			}
		}
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func statisticsList(c *gin.Context) {
	categoryVO := &vo.Category{}
	c.ShouldBindQuery(categoryVO)
	page := pageutils.CheckAndInitPage(categoryVO.BaseVO)
	rpl, err := messaging.Client.PublishOnQueueWaitReply(nil, "posts.getCategoryIDAndCount")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
	}

	IDAndCount := make(map[int64]int64)

	jsoniter.Get(rpl, "model").ToVal(&IDAndCount)
	logrus.Debugln(IDAndCount)

	categoryPOs := make([]po.Category, 0)
	err = db.DB.Cols("id", "name").Limit(pageutils.StartAndEnd(page)).Find(&categoryPOs)
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	categoryVOList := make([]interface{}, 0)
	for _, categoryPO := range categoryPOs {
		category := &vo.Category{
			Id:    categoryPO.Id,
			Name:  categoryPO.Name,
			Total: IDAndCount[categoryPO.Id],
		}
		categoryVOList = append(categoryVOList, category)
	}
	reply.CreateJSONModels(c, categoryVOList)
}

//func statisticsList(c *gin.Context) {
//	categoryVO := &vo.Category{}
//	c.ShouldBindQuery(categoryVO)
//	page := pageutils.CheckAndInitPage(categoryVO.BaseVO)
//	count, offset := pageutils.StartAndEnd(page)
//	//TODO:发送请求给postsservice
//	results, err := db.DB.SQL(`SELECT id,
//	(SELECT COUNT( 1 ) FROM closetool_posts WHERE category_id = category.id AND status=2) as total,
//	name FROM closetool_category AS category limit ?,?`, offset, count).QueryString()
//	if err != nil {
//		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
//		return
//	}
//
//	categoryVOList := make([]interface{}, 0)
//	for _, result := range results {
//		total, err := strconv.ParseInt(result["total"], 10, 64)
//		if err != nil {
//			reply.CreateJSONError(c, reply.Error)
//			return
//		}
//		category := &vo.Category{
//			Name:  result["name"],
//			Total: total,
//		}
//		categoryVOList = append(categoryVOList, category)
//	}
//	reply.CreateJSONModels(c, categoryVOList)
//}

func updateCategory(c *gin.Context) error {
	categoryVO := &vo.Category{}
	err := c.ShouldBindJSON(categoryVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	if categoryVO.Id == 0 {
		reply.CreateJSONError(c, reply.ParamError)
		return nil
	}

	if count, err := db.DB.ID(categoryVO.Id).Count(&po.Category{}); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	} else if count == 0 {
		reply.CreateJSONError(c, reply.DataNoExist)
		return nil
	}

	categoryPO := &po.Category{
		Name: categoryVO.Name,
		Id:   categoryVO.Id,
	}
	if _, err := db.DB.ID(categoryPO.Id).Cols("name").Update(categoryPO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}

	if _, err := db.DB.Where("category_id = ?", categoryVO.Id).Delete(&po.CategoryTags{}); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}

	for _, tag := range categoryVO.TagsList {
		tagPO := &po.Tags{
			Name: tag.Name,
		}
		if ok, err := db.DB.Where("name = ?", tagPO.Name).Get(tagPO); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return err
		} else if !ok {
			if _, err := db.DB.InsertOne(tagPO); err != nil {
				reply.CreateJSONError(c, reply.DatabaseSqlParseError)
				return err
			}
		}

		categoryTagsPO := &po.CategoryTags{
			TagsId:     tagPO.Id,
			CategoryId: categoryPO.Id,
		}

		if ok, err := db.DB.Where("tags_id = ? and category_id = ?", categoryTagsPO.TagsId, categoryTagsPO.CategoryId).Get(categoryTagsPO); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return err
		} else if !ok {
			if _, err := db.DB.InsertOne(categoryTagsPO); err != nil {
				reply.CreateJSONError(c, reply.DatabaseSqlParseError)
				return err
			}
		}
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func getCategoryTags(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}
	categoryPO := &po.Category{}
	if ok, err := db.DB.ID(id).Get(categoryPO); !ok || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	categoryTags := make([]*po.CategoryTags, 0)
	if _, err := db.DB.Where("category_id = ?", id).FindAndCount(&categoryTags); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	tagsVOList := make([]*vo.Tags, 0)
	for _, categoryTag := range categoryTags {
		tag := &po.Tags{}
		if _, err := db.DB.ID(categoryTag.TagsId).Get(tag); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return
		}
		tagsVOList = append(tagsVOList, &vo.Tags{Name: tag.Name, Id: tag.Id})
	}

	category := &vo.Category{
		Id:       id,
		Name:     categoryPO.Name,
		TagsList: tagsVOList,
	}
	reply.CreateJSONModel(c, category)
}

func getCategoryTagsList(c *gin.Context) {
	categoryVO := &vo.Category{}
	c.ShouldBindJSON(categoryVO)
	page := pageutils.CheckAndInitPage(categoryVO.BaseVO)
	session := db.DB.NewSession()
	if categoryVO.BaseVO != nil && categoryVO.Keywords != "" {
		session = session.Where("keywords like ?", "%"+categoryVO.Keywords+"%")
	}
	if categoryVO.Name != "" {
		session = session.Where("name = ?", categoryVO.Name)
	}
	session = session.Limit(pageutils.StartAndEnd(page))
	session = session.Desc("id")
	categoryList := make([]*po.Category, 0)
	if count, err := session.FindAndCount(&categoryList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	} else {
		page.Total = count
	}

	categoryVOList := make([]interface{}, 0)
	for _, categoryPO := range categoryList {
		categoryTagsList := make([]*po.CategoryTags, 0)
		if err := db.DB.Where("category_id = ?", categoryPO.Id).Find(&categoryTagsList); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return
		}
		TagsVOList := make([]*vo.Tags, 0)
		for _, categoryTags := range categoryTagsList {
			tagsPO := &po.Tags{}
			if _, err := db.DB.ID(categoryTags.TagsId).Get(tagsPO); err != nil {
				reply.CreateJSONError(c, reply.DatabaseSqlParseError)
				return
			}
			TagsVOList = append(TagsVOList, &vo.Tags{
				Name: tagsPO.Name,
			})
		}
		categoryVOList = append(categoryVOList, vo.Category{
			Id:       categoryPO.Id,
			Name:     categoryPO.Name,
			TagsList: TagsVOList,
		})
	}
	reply.CreateJSONPaging(c, categoryVOList, page)
}

func getCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}

	logrus.Debugf("Id = %d", id)

	category := &po.Category{}
	if ok, err := db.DB.ID(id).Get(category); !ok || err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	logrus.Debugf("category = %#v", category)

	categoryTagsList := make([]*po.CategoryTags, 0)
	if err := db.DB.Where("category_id = ?", category.Id).Find(&categoryTagsList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	TagsVOList := make([]*vo.Tags, 0)
	for _, categoryTags := range categoryTagsList {
		tagsPO := &po.Tags{}
		if _, err := db.DB.ID(categoryTags.TagsId).Get(tagsPO); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return
		}
		TagsVOList = append(TagsVOList, &vo.Tags{
			Name: tagsPO.Name,
		})
	}
	categoryVO := &vo.Category{
		Id:       category.Id,
		Name:     category.Name,
		TagsList: TagsVOList,
	}

	reply.CreateJSONModel(c, categoryVO)
}

func getCategoryList(c *gin.Context) {

	rpl, err := messaging.Client.PublishOnQueueWaitReply(nil, "posts.tags.getTagsIDAndCount")
	if err != nil || !bytes.Contains(rpl, []byte("00000")) {
		reply.CreateJSONError(c, reply.Error)
	}

	IdAndCount := make(map[int64]int64)
	jsoniter.Get(rpl, "model").ToVal(&IdAndCount)

	categoryList := make([]*po.Category, 0)
	if err := db.DB.Desc("id").Find(&categoryList); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	categoryTags := make([]*po.CategoryTags, 0)
	if err := db.DB.Find(&categoryTags); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	categoryTagsMap := make(map[int64][]*po.CategoryTags)

	for _, categoryTag := range categoryTags {
		temp := categoryTagsMap[categoryTag.CategoryId]
		if temp == nil {
			categoryTagsMap[categoryTag.CategoryId] = []*po.CategoryTags{categoryTag}
		} else {
			temp = append(temp, categoryTag)
			categoryTagsMap[categoryTag.CategoryId] = temp
		}
	}

	categoryVOList := make([]interface{}, 0)
	for _, category := range categoryList {
		total := int64(0)
		for _, categoryTag := range categoryTagsMap[category.Id] {
			total += IdAndCount[categoryTag.TagsId]
		}
		categoryVOList = append(categoryVOList, vo.Category{
			Name:  category.Name,
			Id:    category.Id,
			Total: total,
		})
	}

	reply.CreateJSONModels(c, categoryVOList)
}

//func getCategoryList(c *gin.Context) {
//	categoryList := make([]*po.Category, 0)
//	if err := db.DB.Desc("id").Find(&categoryList); err != nil {
//		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
//		return
//	}
//
//	//TODO:向postsservice发送请求
//	res, err := db.DB.SQL(`SELECT categoryTags.category_id id,COUNT(*) total
//	FROM closetool_category_tags categoryTags
//	LEFT JOIN
//	closetool_posts_tags postsTags ON postsTags.tags_id = categoryTags.tags_id
//	WHERE postsTags.posts_id IS NOT NULL
//	GROUP BY categoryTags.category_id`).QueryString()
//
//	if err != nil {
//		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
//		return
//	}
//	totals := make(map[int64]int64)
//	for _, m := range res {
//		id, err := strconv.ParseInt(m["id"], 10, 64)
//		if err != nil {
//			reply.CreateJSONError(c, reply.Error)
//			return
//		}
//
//		total, err := strconv.ParseInt(m["total"], 10, 64)
//		if err != nil {
//			reply.CreateJSONError(c, reply.Error)
//			return
//		}
//
//		totals[id] = total
//	}
//
//	categoryVOList := make([]interface{}, 0)
//	for _, category := range categoryList {
//		categoryVOList = append(categoryVOList, vo.Category{
//			Name:  category.Name,
//			Id:    category.Id,
//			Total: totals[category.Id],
//		})
//	}
//
//	reply.CreateJSONModels(c, categoryVOList)
//}

func deleteCategory(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	_, err = db.DB.ID(id).Delete(&po.Category{})
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}
	_, err = db.DB.Where("category_id = ?", id).Delete(&po.CategoryTags{})
	if err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	}
	reply.CreateJSONsuccess(c)
	return nil
}
