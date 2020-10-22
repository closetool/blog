package utils

import (
	"fmt"
	"testing"

	"github.com/closetool/blog/services/musicservice/models"
	"github.com/closetool/blog/system/initial"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func init() {
	initial.InitConfig("musicservice")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	initial.InitLog()
}
func TestParsePlaylist(t *testing.T) {
	Convey("Given function json string to parsing", t, func() {
		info := []byte(`{
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
		}`)
		Convey("When function working on this json string", func() {
			res := parsePlaylist(info)
			Convey("Then res should is a music object", func() {
				So(len(res), ShouldEqual, 1)
				//So(res[0].Name, ShouldEqual, "屋顶")
				//So(res[0].URL, ShouldEqual, "https://music.163.com/song/media/outer/url?id=298317.mp3")
				//So(res[0].Artists, ShouldEqual, "温岚/周杰伦")
				//So(res[0].Cover, ShouldEqual, "http://p3.music.126.net/vu3Cdo_dPq8HKOPI6_YXfA==/74766790689775.jpg")
				//So(res[0].Lrc, ShouldEqual, "")
				So(res[0], ShouldResemble, &models.Music{
					Name:    "屋顶",
					URL:     "https://music.163.com/song/media/outer/url?id=298317.mp3",
					Artists: "温岚/周杰伦",
					Cover:   "http://p3.music.126.net/vu3Cdo_dPq8HKOPI6_YXfA==/74766790689775.jpg",
					Lrc:     "",
				})
			})
		})
	})
}
