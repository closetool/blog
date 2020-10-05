package reply

import (
	"fmt"
	"github.com/closetool/blog/system/blogerrors"
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
		Success: 1,
	}
}

func handleErrCode(errCode int) string {
	return fmt.Sprintf("%05d",errCode)
}

func getMessage(errCode int) string{
	return blogerrors.Errors[errCode]
}