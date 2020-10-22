package service

import (
	"fmt"
	"net/http/httptest"
	"strings"
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
	initial.InitConfig("categoryservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	initial.InitLog()
	viper.Set("db_location", "root:%s@/test?charset=utf8")
	viper.Set("db_password", "123456")
	db.DbInit()
	viper.Set("github_auth_url", "https://github.com/login/oauth/authorize?scope=public_repo,read:user&client_id=09087b58751fd0859bce")

	r = gin.New()
	//r.Use(middlewares.Recover())
	group := r.Group("/category")
	routeutils.RegisterRoute(CategoryRoutesTest, group)
}

func TestSaveCategoryWithNoName(t *testing.T) {
	Convey("Given a request to /category/category/v1/add", t, func() {
		req := httptest.NewRequest("POST", "/category/category/v1/add", strings.NewReader(`{  }`))
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00003")
			})
		})
	})
}

func TestSaveCategoryWithName(t *testing.T) {
	Convey("Given a request to /category/category/v1/add", t, func() {
		req := httptest.NewRequest("POST", "/category/category/v1/add", strings.NewReader(`{"name":"test","tagsList":[{"name":"test2"}]}`))
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

func TestStatisticsList(t *testing.T) {
	Convey("Given a request to /category/statistics/v1/list", t, func() {
		req := httptest.NewRequest("GET", "/category/statistics/v1/list?size=1&page=3", nil)
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

func TestUpdateCategory(t *testing.T) {
	Convey("Given a request to /category/category/v1/update", t, func() {
		req := httptest.NewRequest("PUT", "/category/category/v1/update",
			strings.NewReader(`{"id":4,"name":"test","tagsList":[{"name":"test2"},{"name":"testall"},{"name":"test1"},{"name":"test3"}]}`))
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

func TestGetCategoryTags(t *testing.T) {
	Convey("Given a request to /category/category-tags/v1/4", t, func() {
		req := httptest.NewRequest("GET", "/category/category-tags/v1/4", nil)
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

func TestGetCategoryTagsList(t *testing.T) {
	Convey("Given a request to /category/list/v1/category-tags", t, func() {
		req := httptest.NewRequest("GET", "/category/list/v1/category-tags", nil)
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

func TestGetCategory(t *testing.T) {
	Convey("Given a request to /category/category/v1/1", t, func() {
		req := httptest.NewRequest("GET", "/category/category/v1/1", nil)
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

func TestGetCategoryList(t *testing.T) {
	Convey("Given a request to /category/list/v1/category", t, func() {
		req := httptest.NewRequest("GET", "/category/list/v1/category", nil)
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
