package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type springCloudConfig struct {
	AppName         string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Branch          string           `json:"label"`
	Varsion         string           `json:"version"`
	PropertySources []propertySource `json:"propertySources"`
}

type propertySource struct {
	URL    string            `json:"name"`
	Source map[string]interface{} `json:"source"`
}

// LoadConfigurationFromBranch loads config from address pointed out
// http://localhost:8888/mysql/test/branch
func LoadConfigurationFromBranch(configServerURL, appName, profile, branch string) {
	url := fmt.Sprintf("%s/%s/%s/%s", configServerURL, appName, profile, branch)
	body, err := fetchConfiguration(url)
	if err != nil {
		logrus.Panicf("Could not load configuration. Terminating. Error: %v", err)
	}
	parseConfiguration(body)
}

func fetchConfiguration(url string) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in failing: %v", r)
		}
	}()
	logrus.Infof("Getting config from %v", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		logrus.Panicf("Could not load configuration from %v, can not start. Terminating. Error: %v", url, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Panicf("Error reading configuration: %v", err)
	}
	return body, err
}

func parseConfiguration(body []byte) {
	var configs springCloudConfig

	err := json.Unmarshal(body, &configs)
	if err != nil {
		logrus.Panicf("Cannot parse configuration. Error: %v", err)
	}

	for k, v := range configs.PropertySources[0].Source {
		viper.Set(k, v)
		logrus.Infof("Loading config property %v => %v", k, v)
	}
	if viper.IsSet("service_name") {
		logrus.Infof("Successfully loaded configuration for service %s", viper.GetString("service_name"))
	}

}
