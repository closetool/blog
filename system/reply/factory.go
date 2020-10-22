package reply

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWithError() *Reply {
	return CreateWithErrorX(Error)
}

func CreateWithErrorX(errCode int) *Reply {
	reply := createWithErrorFlag()
	reply.ReplyCode = HandleErrCode(errCode)
	reply.Message = getMessage(errCode)
	return reply
}

func createWithErrorFlag() *Reply {
	return &Reply{
		Success: 0,
	}
}

func CreateWithSuccess() *Reply {
	return &Reply{
		Success:   1,
		ReplyCode: HandleErrCode(Success),
		Message:   getMessage(Success),
	}
}

func HandleErrCode(errCode int) string {
	return fmt.Sprintf("%05d", errCode)
}

func getMessage(errCode int) string {
	return Errors[errCode]
}

func CreateWithModel(model interface{}) *Reply {
	reply := CreateWithSuccess()
	reply.Model = model
	return reply
}

func CreateWithModels(models []interface{}) *Reply {
	reply := CreateWithSuccess()
	reply.Models = models
	return reply
}

func CreateWithPaging(models []interface{}, page *PageInfo) *Reply {
	reply := CreateWithModels(models)
	reply.Pageinfo.Page = page.Page
	reply.Pageinfo.Size = page.Size
	reply.Pageinfo.Total = page.Total
	return reply
}

func CreateJSONError(c *gin.Context, errCode int) {
	c.JSON(http.StatusOK, CreateWithErrorX(errCode))
}

func CreateJSONModel(c *gin.Context, model interface{}) {
	c.JSON(http.StatusOK, CreateWithModel(model))
}

func CreateJSONModels(c *gin.Context, models []interface{}) {
	c.JSON(http.StatusOK, CreateWithModels(models))
}

func CreateJSONPaging(c *gin.Context, models []interface{}, page *PageInfo) {
	c.JSON(http.StatusOK, CreateWithPaging(models, page))
}

func CreateJSONsuccess(c *gin.Context) {
	c.JSON(http.StatusOK, CreateWithSuccess())
}
