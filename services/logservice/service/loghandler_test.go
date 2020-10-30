package service

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

var r *gin.Engine

func init() {
	initial.InitConfig("logservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	initial.InitLog()
	viper.Set("db_location", "root:%s@/test?charset=utf8")
	viper.Set("db_password", "123456")
	db.DbInit()

	r = gin.New()
	//r.Use(middlewares.Recover())
	group := r.Group("/logs")
	routeutils.RegisterRoute(LogRoutes, group)
}

func TestGetLogs(t *testing.T) {
	Convey("Given a request to /logs/logs/v1/1", t, func() {
		req := httptest.NewRequest("GET", "/logs/logs/v1/1", nil)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
			})
		})
	})
}

func TestGetLogsList(t *testing.T) {
	Convey("Given a request to /logs/list/v1/logs", t, func() {
		req := httptest.NewRequest("GET", "/logs/list/v1/logs?size=10&page=1&userId=3", nil)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
			})
		})
	})
}
