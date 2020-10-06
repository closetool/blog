package reply

import (
	"fmt"
)

func CreateWithError(errCode int) *Reply {
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

func createWithSuccessFlag() *Reply {
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
	reply := createWithSuccessFlag()
	reply.Model = model
	return reply
}
