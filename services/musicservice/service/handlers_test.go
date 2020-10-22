package service

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/closetool/blog/services/musicservice/models"
	"github.com/closetool/blog/system/initial"
	"github.com/closetool/blog/utils/routeutils"
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func init() {
	initial.InitConfig("musicservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	initial.InitLog()
}

func TestGetPlaylist(t *testing.T) {
	Convey("Set http mock", t, func() {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://music.163.com/api/playlist/detail?id=165447515",
			httpmock.NewStringResponder(200, `{
			"result": {
				"tracks": [{
					"name": "屋顶",
					"id": 298317,
					"artists": [{
						"name": "温岚"
					}, {
						"name": "周杰伦"
					}],
					"album": {
						"blurPicUrl": "http://p3.music.126.net/vu3Cdo_dPq8HKOPI6_YXfA==/74766790689775.jpg"
					}
				}]
			}
		}`))

		Convey("Given a HTTP request to /music/music/v1/list", func() {
			viper.Set("music_playlist_id", "165447515")
			req := httptest.NewRequest("GET", "/music/music/v1/list", nil)
			resp := httptest.NewRecorder()
			r := gin.New()
			g := r.Group("/music")
			routeutils.RegisterRoute(Routes, g)
			Convey("When pass req to server", func() {
				r.ServeHTTP(resp, req)
				Convey("Response should contains reply's json string", func() {
					res, _ := ioutil.ReadAll(resp.Body)
					So(jsoniter.Get(res, "success").ToInt(),
						ShouldEqual, 1)
					parsedMusic := &models.Music{}
					jsoniter.Get(res, "model", 0).ToVal(parsedMusic)
					music := models.Music{
						Name:    "屋顶",
						URL:     "https://music.163.com/song/media/outer/url?id=298317.mp3",
						Artists: "温岚/周杰伦",
						Cover:   "http://p3.music.126.net/vu3Cdo_dPq8HKOPI6_YXfA==/74766790689775.jpg",
						Lrc:     "",
					}
					So(*parsedMusic, ShouldResemble, music)
				})
			})
		})
	})
}
