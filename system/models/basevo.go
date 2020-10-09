package models

type BaseVO struct {
	Keywords string `json:"keywords,omitempty"`
	Page     int64  `json:"page,omitempty"`
	Size     int64  `json:"size,omitempty"`
}
