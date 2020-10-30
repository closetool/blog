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
	initial.InitLog()
	viper.Set("db_location", "root:%s@/test?charset=utf8")
	viper.Set("db_password", "123456")
	db.DbInit()
	viper.Set("github_auth_url", "https://github.com/login/oauth/authorize?scope=public_repo,read:user&client_id=09087b58751fd0859bce")
	//config.LoadConfigurationFromBranch(
	//	viper.GetString("config_server_url"),
	//	"userservice",
	//	viper.GetString("profile"),
	//	viper.GetString("branch"),
	//)

	r = gin.New()
	//r.Use(middlewares.Recover())
	group := r.Group("/auth")
	routeutils.RegisterRoute(Routes, group)
}

func TestRegisterAdminByGithub(t *testing.T) {
	Convey("Given a request to /auth/admin/v1/register", t, func() {
		req := httptest.NewRequest("POST", "/auth/admin/v1/register", strings.NewReader(`{ "email":"c299999999@qq.com","password":"abc123" }`))
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

func TestGetUserInfo(t *testing.T) {
	Convey("Given a request to server", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("GET", "/auth/user/v1/get", nil)
		req.Header.Set(constants.AuthHeader, token)
		resp := httptest.NewRecorder()
		Convey("When pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then response body should have message", func() {
				logrus.Debugf("response = %s", resp.Body.String())
				So(resp.Result().StatusCode, ShouldEqual, 200)
				So(resp.Body.String(), ShouldContainSubstring, "00000")
			})
		})
	})
}

func generateAdminToken() (string, error) {
	admin := &po.AuthUser{}
	db.DB.Where("role_id=?", 2).Get(admin)
	logrus.Debugf("admin = %#v", admin)
	token, err, _ := utils.GenerateToken(admin)
	logrus.Debugf("token = %v", token)
	return token, err
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

func TestGetUserList(t *testing.T) {
	Convey("Given a request to /auth/user/v1/list", t, func() {
		req := httptest.NewRequest("GET", "/auth/user/v1/list?keywords=close",
			strings.NewReader(`{"name":"closetool"}`))
		token, _ := generateAdminToken()
		req.Header.Add(constants.AuthHeader, token)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
				So(resp.Body.String(), ShouldContainSubstring, "closetool")
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
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
			})
		})
	})
}

func TestOauthByGithub(t *testing.T) {
	Convey("Given a request to /auth/github/v1/get", t, func() {
		req := httptest.NewRequest("GET", "/auth/github/v1/get", nil)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
				So(resp.Body.String(), ShouldContainSubstring,
					"https://github.com/login/oauth/authorize?scope=public_repo,read:user&client_id=09087b58751fd0859bce")
			})
		})
	})
}

func TestDeleteUserFail(t *testing.T) {
	Convey("Generate admin token", t, func() {
		token, _ := generateAdminToken()

		Convey("Given a request to server", func() {
			req := httptest.NewRequest("DELETE", "/auth/user/v1/10086", nil)
			req.Header.Set(constants.AuthHeader, token)
			resp := httptest.NewRecorder()
			Convey("When pass request to server", func() {
				r.ServeHTTP(resp, req)
				Convey("Then response body should have message", func() {
					logrus.Debugf("response = %s", resp.Body.String())
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
			logrus.Panicf("generate token failed: %v", err)
		}
		logrus.Debugf("token = %v", token)

		Convey("Given a request to server", func() {
			req := httptest.NewRequest("DELETE", "/auth/user/v1/2", nil)
			req.Header.Set(constants.AuthHeader, token)
			resp := httptest.NewRecorder()
			Convey("When pass request to server", func() {
				r.ServeHTTP(resp, req)
				Convey("Then response body should have message", func() {
					logrus.Debugf("response = %s", resp.Body.String())
					So(resp.Result().StatusCode, ShouldEqual, 200)
					So(resp.Body.String(), ShouldContainSubstring, "00000")
				})
			})
		})
	})
}

func TestSaveUserByGithub(t *testing.T) {
	Convey("Given a request to /auth/user/v1/login", t, func() {
		req := httptest.NewRequest("POST", "/auth/user/v1/login", strings.NewReader(`{ "socialId": "123456", "avatar": "https://localhost/images/01.jpg", "name": "closetool" }`))
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

func TestLogin(t *testing.T) {
	Convey("Given a request to /auth/admin/v1/login", t, func() {
		req := httptest.NewRequest("POST", "/auth/admin/v1/login", strings.NewReader(`{ "email":"c299999999@qq.com","password":"abc123" }`))
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "c299999999@qq.com")
			})
		})
	})
}

func TestUpdatePassword(t *testing.T) {
	Convey("Given a request to /auth/password/v1/update", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("PUT", "/auth/password/v1/update",
			strings.NewReader(`{ "passwordOld":"abc123","password":"admin" }`))
		req.Header.Set(constants.AuthHeader, token)
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

func TestUpdateAdmin(t *testing.T) {
	Convey("Given a request to /auth/admin/v1/update", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("PUT", "/auth/admin/v1/update",
			strings.NewReader(`{"name":"closetool_admin",
			"email":"c299999999@qq.com",
			"avatar":"https://avatars3.githubusercontent.com/u/52988625?v=4",
			"introduction":"sql boy"}`))

		req.Header.Set(constants.AuthHeader, token)
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

func TestUpdateUser(t *testing.T) {
	Convey("Given a request to /auth/user/v1/update", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("PUT", "/auth/user/v1/update",
			strings.NewReader(`{"Id":4,"name":"closetool_user",
			"email":"4closetool3@gmail.com","avatar":"","introduction":"closetool's introduction"}`))

		req.Header.Set(constants.AuthHeader, token)
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

func TestSaveSocial(t *testing.T) {
	Convey("Given a request to /auth/social/v1/add", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("POST", "/auth/social/v1/add",
			strings.NewReader(`{"code":"github","content":"https://github.com/closetool",
			"showType":3,"isEnabled":1,"isHome":1}`))

		req.Header.Set(constants.AuthHeader, token)
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

func TestEditSocial(t *testing.T) {
	Convey("Given a request to /auth/social/v1/update", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("PUT", "/auth/social/v1/update",
			strings.NewReader(`{"id":1,"code":"github","content":"https://github.com/closetool",
			"showType":1,"isEnabled":1,"isHome":1}`))

		req.Header.Set(constants.AuthHeader, token)
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

func TestGetSocial(t *testing.T) {
	Convey("Given a request to /auth/social/v1/1", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("GET", "/auth/social/v1/1", nil)

		req.Header.Set(constants.AuthHeader, token)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
				So(resp.Body.String(), ShouldContainSubstring, "github.com/closetool")
			})
		})
	})
}

func TestGetSocialList(t *testing.T) {
	Convey("Given a request to /auth/list/v1/social", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("GET", "/auth/list/v1/social?page=2&size=1", nil)

		req.Header.Set(constants.AuthHeader, token)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
				So(resp.Body.String(), ShouldContainSubstring, "github.com/closetool")
			})
		})
	})
}

func TestGetSocialInfo(t *testing.T) {
	Convey("Given a request to /auth/info/v1/social", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("GET", "/auth/info/v1/social", nil)

		req.Header.Set(constants.AuthHeader, token)
		resp := httptest.NewRecorder()
		Convey("Pass request to server", func() {
			r.ServeHTTP(resp, req)
			Convey("Then reponse body should be", func() {
				logrus.Debugf("response = %v", resp.Body.String())
				So(resp.Body.String(), ShouldContainSubstring, "00000")
				So(resp.Body.String(), ShouldContainSubstring, "github.com/closetool")
			})
		})
	})
}

func TestDelSocial(t *testing.T) {
	Convey("Given a request to /auth/social/v1/1", t, func() {
		token, _ := generateAdminToken()
		req := httptest.NewRequest("DELETE", "/auth/social/v1/1", nil)

		req.Header.Set(constants.AuthHeader, token)
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
