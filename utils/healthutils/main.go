package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	if len(os.Args) < 2 {
		logrus.Errorln("param not enough")
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/health", os.Args[1])
	logrus.Infof("get health info through %s", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		logrus.Errorf("error: %v", err)
		os.Exit(1)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			os.Exit(1)
		}
		rpl := new(reply)
		json.Unmarshal(body, rpl)
		logrus.Debugf("reply = %#v", rpl)
		if !rpl.Health {
			os.Exit(1)
		}
		os.Exit(0)
	}
}

type reply struct {
	Health bool `json:"health"`
}
