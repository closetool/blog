package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/closetool/blog/services/userservice/models/po"
	"github.com/closetool/blog/services/userservice/models/vo"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/system/log"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func init() {
	initial.InitConfig("userservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	log.InitLog()
}
func TestGenerateToken(t *testing.T) {
	Convey("Given a user info to generate token", t, func() {
		user := &po.AuthUser{
			Id:       2,
			Name:     "closetool",
			Password: "6193097d60899bcfd4f00a0896c9b610",
		}

		Convey("When pass it into function", func() {
			expected := &vo.AuthUser{
				Id:   2,
				Name: "closetool",
			}
			token, _ := GenerateToken(user)
			Convey("Then token should equal", func() {
				res := strings.Split(token, ".")
				body, _ := base64.RawStdEncoding.DecodeString(res[1])
				parsedUser := &vo.AuthUser{}
				json.Unmarshal(body, parsedUser)
				parsedUser.StandardClaims = nil

				So(parsedUser, ShouldResemble, expected)
			})
		})
	})
}

func TestVerifyToken(t *testing.T) {
	Convey("Generate a token", t, func() {
		user := &po.AuthUser{
			Id:       2,
			Name:     "closetool",
			Password: "6193097d60899bcfd4f00a0896c9b610",
		}
		token, err := GenerateToken(user)
		if err != nil {
			log.Logger.Errorf("generate token failed: %v\n", err)
			t.Fail()
		}

		log.Logger.Debugf("token = %v\n", token)
		Convey("Given a token string to verifing", func() {
			Convey("When pass it to function", func() {
				res := VerifyToken(token, user.Password)
				Convey("Then claim should equal", func() {
					So(res, ShouldBeTrue)
				})
			})
		})
	})
}
