package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/closetool/blog/services/musicservice/models"
	"github.com/closetool/blog/system/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

const (
	MusicPrefixURL = "https://music.163.com/api/playlist/detail?id="
	PlayURL        = "https://music.163.com/song/media/outer/url?id="
)

func GetPlaylist() (models.PlayList, error) {
	playlistURL := fmt.Sprintf("%s%s", MusicPrefixURL, viper.GetString("music_playlist_id"))
	resp, err := http.Get(playlistURL)
	if err != nil || resp.StatusCode != 200 {
		log.Logger.Errorf("get music play list failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	playlistInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Errorf("read response body from playlist failed: %v", err)
		return nil, err
	}

	return parsePlaylist(playlistInfo), nil
}

func parsePlaylist(info []byte) models.PlayList {
	tracks := jsoniter.Get(info, "result", "tracks")
	musics := make([]*models.Music, tracks.Size())

	for i := 0; i < tracks.Size(); i++ {
		track := tracks.Get(i)

		music := new(models.Music)

		music.Name = track.Get("name").ToString()
		songURL := fmt.Sprintf("%s%s%s", PlayURL, track.Get("id").ToString(), ".mp3")
		music.URL = songURL
		music.Artists = getAllArtists(track)
		music.Cover = track.Get("album").Get("blurPicUrl").ToString()
		musics[i] = music
	}
	return musics
}

func getRedirectURL(URL string) string {
	return ""
}

func getAllArtists(js jsoniter.Any) string {
	artistsArr := make([]string, 0)
	artists := js.Get("artists")
	for i := 0; i < artists.Size(); i++ {
		artistsArr = append(artistsArr, artists.Get(i).Get("name").ToString())
	}
	return strings.Join(artistsArr, `/`)
}
