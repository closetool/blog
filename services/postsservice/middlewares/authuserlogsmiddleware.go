package middlewares

import (
	"time"

	authuservo "github.com/closetool/blog/services/logservice/model/vo"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/utils/httpcontextutils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func AuthUserLogMiddleware(code string, module string) func(c *gin.Context) {
	return func(c *gin.Context) {
		log := &authuservo.AuthUserLog{
			Code:           code,
			Description:    module,
			Ip:             c.ClientIP(),
			Url:            c.Request.URL.Path,
			Device:         httpcontextutils.GetOsName(c.Request),
			BrowserName:    httpcontextutils.GetBrowserName(c.Request),
			BrowserVersion: httpcontextutils.GetBrowserVersion(c.Request),
		}
		m := make(map[string]string)
		if code == constants.PostsDetail {
			for _, p := range c.Params {
				m[p.Key] = m[p.Value]
			}
		} else {
			for k, v := range c.Request.URL.Query() {
				if len(v) > 0 {
					m[k] = v[0]
				}
			}
		}
		log.Parameter, _ = jsoniter.MarshalToString(m)
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		dur := endTime.Sub(startTime)
		log.RunTime = dur.Milliseconds()
		bts, _ := jsoniter.Marshal(log)
		messaging.Client.PublishOnQueue(bts, "logs.saveLogs")
	}
}
