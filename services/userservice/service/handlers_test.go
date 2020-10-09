package service

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/utils"
	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/db"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/log"
	"github.com/closetool/blog/system/middlewares"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

var r *gin.Engine

func init() {
	initial.InitConfig("userservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	log.InitLog()
	viper.Set("db_location", "root:%s@/test?charset=utf8")
	viper.Set("db_password", "123456")
	db.DbInit()

	r = gin.New()
	r.Use(middlewares.Recover())
	group := r.Group("/auth")
	routeutils.RegisterRoute(Routes, group)
}

func TestGetUserInfo(t *testing.T) {
	Convey("Given a request to server", t, func() {
		req := httptest.NewRequest("GET", "/auth/user/v1/get", nil)
		req.Header.Set(constants.AuthHeader, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwibmFtZSI6ImNsb3NldG9vbCIsImV4cCI6MTYwMjc0NDQxNH0.Rlg4j193eUmymCmcBeemvbI9k-fPRDMQMfTKFYrxE3Y")
		resp := httptest.NewRecorder()
		Convey("When pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then response body should have message", func() {
				log.Logger.Debugf("response = %s\n", resp.Body.String())
				So(resp.Result().StatusCode, ShouldEqual, 200)
				So(resp.Body.String(), ShouldContainSubstring, "00000")
			})
		})
	})
}

func generateAdminToken() (string, error) {
	admin := &po.AuthUser{}
	db.DB.Where("role_id=?", 2).Get(admin)
	log.Logger.Debugf("admin = %#v\n", admin)
	return utils.GenerateToken(admin)
}

func TestSaveAuthUserStatus(t *testing.T) {
	Convey("Generate admin token", t, func() {
		token, _ := generateAdminToken()
		Convey("Given a request to /auth/status/v1/update", func() {
			req := httptest.NewRequest("PUT", "/auth/status/v1/update", strings.NewReader(`{"id":2,"status":1}`))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			req.Header.Set(constants.AuthHeader, token)
			resp := httptest.NewRecorder()
			Convey("Pass request to server", func() {
				r.ServeHTTP(resp, req)
				Convey("Then reponse body should have", func() {
					So(resp.Body.String(), ShouldContainSubstring, "00000")
				})
			})
		})
	})
}

func TestGetMasterUserInfo(t *testing.T) {
	Convey("Given a request to /auth/master/v1/get", t, func() {
		req := httptest.NewRequest("GET", "/auth/master/v1/get", nil)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should have", func() {
				log.Logger.Debugf("response = %v\n", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
			})
		})
	})
}

func TestDeleteUserFail(t *testing.T) {
	Convey("Generate admin token", t, func() {
		token, _ := generateAdminToken()

		Convey("Given a request to server", func() {
			req := httptest.NewRequest("DELETE", "/auth/user/v1/4", nil)
			req.Header.Set(constants.AuthHeader, token)
			resp := httptest.NewRecorder()
			Convey("When pass request to server", func() {
				r.ServeHTTP(resp, req)
				Convey("Then response body should have message", func() {
					log.Logger.Debugf("response = %s\n", resp.Body.String())
					So(resp.Result().StatusCode, ShouldEqual, 200)
					So(resp.Body.String(), ShouldContainSubstring, "00011")
				})
			})
		})
	})
}

func TestDeleteUserSucceed(t *testing.T) {
	Convey("Generate admin token", t, func() {
		token, err := generateAdminToken()
		if err != nil {
			log.Logger.Panicf("generate token failed: %v\n", err)
		}
		log.Logger.Debugf("token = %v\n", token)

		Convey("Given a request to server", func() {
			req := httptest.NewRequest("DELETE", "/auth/user/v1/2", nil)
			req.Header.Set(constants.AuthHeader, token)
			resp := httptest.NewRecorder()
			Convey("When pass request to server", func() {
				r.ServeHTTP(resp, req)
				Convey("Then response body should have message", func() {
					log.Logger.Debugf("response = %s\n", resp.Body.String())
					So(resp.Result().StatusCode, ShouldEqual, 200)
					So(resp.Body.String(), ShouldContainSubstring, "00000")
				})
			})
		})
	})
}
