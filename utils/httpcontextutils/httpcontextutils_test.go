package httpcontextutils

import (
	"log"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOsName(t *testing.T) {
	Convey("create a request with user-agent", t, func() {
		req := httptest.NewRequest("GET", "http://closetool.site", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
		Convey("pass it to method GetOsName", func() {
			osName := GetOsName(req)
			log.Printf("osName = %v", osName)
			So(osName, ShouldContainSubstring, "Windows 10")
		})
	})
}
func TestGetBrowserVersion(t *testing.T) {
	Convey("create a request with user-agent", t, func() {
		req := httptest.NewRequest("GET", "http://closetool.site", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
		Convey("pass it to method GetOsName", func() {
			osName := GetBrowserVersion(req)
			log.Printf("osName = %v", osName)
			So(osName, ShouldContainSubstring, "84.0.4147")
		})
	})
}
