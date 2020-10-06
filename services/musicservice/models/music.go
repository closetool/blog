package models

import (
	jsoniter "github.com/json-iterator/go"
)

type Music struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Artists string `json:"artist"`
	Cover   string `json:"cover"`
	Lrc     string `json:"lrc"`
}

type PlayList []*Music

func (p PlayList) String() string {
	rtn, _ := jsoniter.MarshalToString(p)
	return rtn
}
