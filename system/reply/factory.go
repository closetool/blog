package reply

import (
	"fmt"
)

func CreateWithError() *Reply {
	return CreateWithErrorX(Error)
}

func CreateWithErrorX(errCode int) *Reply {
	reply := createWithErrorFlag()
	reply.ReplyCode = handleErrCode(errCode)
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
		ReplyCode: handleErrCode(Success),
		Message:   getMessage(Success),
	}
}

func handleErrCode(errCode int) string {
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
