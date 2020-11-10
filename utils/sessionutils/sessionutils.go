package sessionutils

import (
	"fmt"

	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) (*vo.AuthUser, error) {
	session, exist := c.Get("session")
	if !exist {
		return nil, fmt.Errorf("context has no session")
	}
	user, ok := session.(vo.AuthUser)
	if !ok {
		return nil, fmt.Errorf("can not convert session into vo.AuthUser")
	}
	return &user, nil
}
