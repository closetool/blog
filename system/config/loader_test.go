package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	Convey("Given JSON response body", t, func() {
		var body = `{"name":"userservice","profiles":["test"],"label":"blog","version":"bf30f2d8c0d62057e5ec1cae88d59118d3d28808","propertySources":[{"name":"https://github.com/closetool/go-micro-service-config.git/userservice-test.yml","source":{"username":"root","url":"tcp(localhost:3306)/blog?charset=utf8&serverTimezone=Hongkong","service_name":"userservice","log_file_path":"./","log_file_name":"log","log_level":4,"port":3456,"password":"123456"}}]}`
		Convey("When parsed", func() {
			parseConfiguration([]byte(body))
			Convey("Then viper should have parsed info", func() {
				So(viper.Get("service_name"), ShouldEqual, "userservice")
			})
		})
	})
}
