package service

import (
	"net/http"

	"github.com/closetool/blog/services/musicservice/utils"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/reply"
	"github.com/gin-gonic/gin"
)

func GetPlayList(c *gin.Context) {
	playlist, err := utils.GetPlaylist()
	if err != nil {
		c.JSON(http.StatusOK, reply.CreateWithErrorX(reply.Error))
	} else {
		log.Logger.Debugf("parsed json info: %s", playlist)
		rp := reply.CreateWithModel(playlist)
		c.JSON(http.StatusOK, rp)
	}
}
