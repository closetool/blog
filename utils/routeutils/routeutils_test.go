package routeutils

import (
	"github.com/closetool/blog/system/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestRegisterRoute(t *testing.T) {
	Convey("Given routes", t, func() {
		routes := []models.Route{
			{Method: "GET", Pattern: "/:name", MiddleWare: nil, HandlerFunc: func(c *gin.Context) {
				c.String(http.StatusOK, "hello, %s", c.Param("name"))
			}},
		}

		r := gin.New()
		g := r.Group("/test")
		Convey("Pass routes to func", func() {
			RegisterRoute(routes, g)

			Convey("Then should equal", func() {
				req := httptest.NewRequest("GET", "/test/world", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)
				body, _ := ioutil.ReadAll(resp.Body)
				So(string(body), ShouldEqual, "hello, world")
			})
		})
	})
}
