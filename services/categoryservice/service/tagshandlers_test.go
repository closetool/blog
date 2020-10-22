package service

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func init() {
	initial.InitConfig("categoryservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	initial.InitLog()
	viper.Set("db_location", "root:%s@/test?charset=utf8")
	viper.Set("db_password", "123456")
	db.DbInit()

	//r.Use(middlewares.Recover())
	group := r.Group("/tags")
	routeutils.RegisterRoute(TagsRoutesTest, group)
}

func TestGetTagsListNoPaging(t *testing.T) {
	Convey("Given a request to /tags/list/v1/tags", t, func() {
		req := httptest.NewRequest("GET", "/tags/list/v1/tags", nil)
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

func TestGetTagsList(t *testing.T) {
	Convey("Given a request to /tags/list/v1/tags", t, func() {
		req := httptest.NewRequest("GET", "/tags/list/v1/tags?page=2&size=2", nil)
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

func TestGetTagsAndArticleQuantityList(t *testing.T) {
	Convey("Given a request to /tags/tags-article-quantity/v1/list", t, func() {
		req := httptest.NewRequest("GET", "/tags/tags-article-quantity/v1/list", nil)
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

func TestGetTags(t *testing.T) {
	Convey("Given a request to /tags/tags/v1/1", t, func() {
		req := httptest.NewRequest("GET", "/tags/tags/v1/1", nil)
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

func TestSaveTags(t *testing.T) {
	Convey("Given a request to /tags/tags/v1/add", t, func() {
		req := httptest.NewRequest("POST", "/tags/tags/v1/add",
			strings.NewReader(`{"name":"closetool"}`))
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

func TestUpdateTags(t *testing.T) {
	Convey("Given a request to /tags/tags/v1/update", t, func() {
		req := httptest.NewRequest("PUT", "/tags/tags/v1/update",
			strings.NewReader(`{"id":5,"name":"closetool2"}`))
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

func TestDeleteTags(t *testing.T) {
	Convey("Given a request to /tags/tags/v1/5", t, func() {
		req := httptest.NewRequest("DELETE", "/tags/tags/v1/5", nil)
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
