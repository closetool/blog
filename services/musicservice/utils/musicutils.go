package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/closetool/blog/system/constants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetPalyList() {
	playlistUrl := fmt.Sprintf("%s%s", constants.MUSIC_PREFIX_URL, viper.Get("music_playlist_id"))
	resp, err := http.Get(playlistUrl)
	if err != nil {
		logrus.Errorf("get music play list failed: %v", err)
		return
	}

	playlistInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("read response body of playlist failed: %v", err)
		return
	}

	parsePlaylist(playlistInfo)
}

func parsePlaylist(info []byte) {
	tracks := jsoniter.Get(info,"result","tracks")

}
