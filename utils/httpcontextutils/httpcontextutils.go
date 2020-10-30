package httpcontextutils

import (
	"fmt"
	"net/http"

	"github.com/mssola/user_agent"
)

func GetOsName(request *http.Request) string {
	UAstr := request.Header.Get("User-Agent")
	if UAstr == "" {
		return "not found osName"
	}
	userAgent := user_agent.New(UAstr)
	return fmt.Sprintf("%s %s", userAgent.OSInfo().Name, userAgent.OSInfo().Version)
}

func GetBrowserVersion(request *http.Request) string {
	UAstr := request.Header.Get("User-Agent")
	if UAstr == "" {
		return "not found version"
	}
	userAgent := user_agent.New(UAstr)
	_, browserVersion := userAgent.Browser()
	return browserVersion

}

func GetBrowserName(request *http.Request) string {
	UAstr := request.Header.Get("User-Agent")
	if UAstr == "" {
		return "not found name"
	}
	userAgent := user_agent.New(UAstr)
	browserName, _ := userAgent.Browser()
	return browserName
}
