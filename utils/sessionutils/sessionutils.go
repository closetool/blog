package sessionutils

import (
	"fmt"

	"github.com/closetool/blog/system/models/model"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) (*model.AuthUser, error) {
	session, exist := c.Get("session")
	if !exist {
		return nil, fmt.Errorf("context has no session")
	}
	user, ok := session.(model.AuthUser)
	if !ok {
		return nil, fmt.Errorf("can not convert session into vo.AuthUser")
	}
	return &user, nil
}
