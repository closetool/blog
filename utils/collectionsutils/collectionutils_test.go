package collectionsutils

import (
	"fmt"
	"testing"

	"github.com/closetool/blog/system/initial"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func init() {
	initial.InitConfig("collectionutils")
	viper.Set("log_level", fmt.Sprintf("%d", logrus.DebugLevel))
	initial.InitLog()
	//config.LoadConfigurationFromBranch(
	//	viper.GetString("config_server_url"),
	//	"userservice",
	//	viper.GetString("profile"),
	//	viper.GetString("branch"),
	//)
}

func TestRandomString(t *testing.T) {
	Convey("Random String", t, func() {
		res := RandomString(32)
		Convey("result should have 32 bit", func() {
			logrus.Debugf("res = %v\n", string(res))
			So(res, ShouldHaveLength, 32)
		})
	})
}
