package service

import (
	"strconv"

	"github.com/closetool/blog/services/categoryservice/models/po"
	"github.com/closetool/blog/services/categoryservice/models/vo"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/pageutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getTagsList(c *gin.Context) {
	tagsVO := &vo.Tags{}
	c.ShouldBindQuery(tagsVO)
	logrus.Debugf("tagsVO = %#v", tagsVO)
	tagsList := make([]interface{}, 0)
	if tagsVO.BaseVO == nil || (tagsVO.Page == 0 && tagsVO.Size == 0) {
		records := make([]*po.Tags, 0)
		if err := db.DB.Find(&records); err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
			return
		}
		for _, record := range records {
			tagsList = append(tagsList, &vo.Tags{
				Id:   record.Id,
				Name: record.Name,
			})
		}
		reply.CreateJSONModels(c, tagsList)
		return
	}

	session := db.DB.NewSession()
	if tagsVO.Keywords != "" {
		session = session.Where("name like ?", "%"+tagsVO.Keywords+"%")
	}
	if tagsVO.Name != "" {
		session = session.Where("name = ?", tagsVO.Name)
	}

	page := pageutils.CheckAndInitPage(tagsVO.BaseVO)
	tagsPOs := make([]*po.Tags, 0)

	var err error
	if page.Total, err = session.Limit(pageutils.StartAndEnd(page)).FindAndCount(&tagsPOs); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	tagsVOs := make([]interface{}, 0)
	for _, tagsPO := range tagsPOs {
		tagsVOs = append(tagsVOs, vo.Tags{
			Name: tagsPO.Name,
			Id:   tagsPO.Id,
		})
	}
	reply.CreateJSONPaging(c, tagsVOs, page)
}

func getTagsAndArticleQuantityList(c *gin.Context) {
	records := make([]*po.Tags, 0)
	if err := db.DB.Find(&records); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}

	tagsVOList := make([]interface{}, 0)

	for _, record := range records {
		//TODO:修改为向postsservice，发送请求
		res, err := db.DB.SQL(`SELECT Count(*) as total 
		FROM closetool_posts_tags 
		WHERE tags_id = ?`, record.Id).QueryInterface()
		if err != nil {
			reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		}
		total, ok := res[0]["total"].(int64)
		logrus.Debugf("total is int64: %v", ok)
		if !ok {
			reply.CreateJSONError(c, reply.Error)
			return
		}
		tagsVOList = append(tagsVOList, &vo.Tags{
			Name:       record.Name,
			Id:         record.Id,
			PostsTotal: total,
		})
	}

	reply.CreateJSONModels(c, tagsVOList)
}

func getTags(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return
	}
	logrus.Debugf("id = %v", id)

	tagsPO := &po.Tags{}
	if _, err := db.DB.ID(id).Get(tagsPO); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return
	}
	tagsVO := vo.Tags{
		Name: tagsPO.Name,
		Id:   tagsPO.Id,
	}
	reply.CreateJSONModel(c, tagsVO)
}

func saveTags(c *gin.Context) error {
	tagsVO := &vo.Tags{}
	err := c.ShouldBindJSON(tagsVO)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	//session, ok := c.Get("session")
	//if !ok {
	//	reply.CreateJSONError(c, reply.AccountNotExist)
	//}
	//user, ok := session.(uservo.AuthUser)
	//if !ok {
	//	reply.CreateJSONError(c, reply.AccountNotExist)
	//}

	db.DB.InsertOne(&po.Tags{
		Name: tagsVO.Name,
		//CreateBy: user.Id,
		//UpdateBy: user.Id,
	})
	reply.CreateJSONsuccess(c)
	return nil
}

func updateTags(c *gin.Context) error {
	tagsVO := &vo.Tags{}
	err := c.ShouldBindJSON(tagsVO)
	if err != nil || tagsVO.Id == 0 {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}

	if count, err := db.DB.ID(tagsVO.Id).Cols("name").Update(&po.Tags{Name: tagsVO.Name}); err != nil {
		reply.CreateJSONError(c, reply.DatabaseSqlParseError)
		return err
	} else if count == 0 {
		reply.CreateJSONError(c, reply.DataNoExist)
		return nil
	}
	reply.CreateJSONsuccess(c)
	return nil
}

func deleteTags(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		reply.CreateJSONError(c, reply.ParamError)
		return err
	}
	logrus.Debugf("id = %v", id)

	db.DB.ID(id).Delete(&po.Tags{})
	db.DB.Where("tags_id = ?", id).Delete(&po.CategoryTags{})
	//TODO:向消息总线发送任务，请求postsservice删除相关poststags
	//db.DB.Where("tags_id = ?", id).Delete()
	reply.CreateJSONsuccess(c)
	return nil
}
