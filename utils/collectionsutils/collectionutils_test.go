package collectionsutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomString(t *testing.T) {
	Convey("Random String", t, func() {
		res := RandomString(32)
		Convey("result should have 32 bit", func() {
			So(res, ShouldHaveLength, 32)
		})
	})
}
