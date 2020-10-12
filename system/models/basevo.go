package models

type BaseVO struct {
	Keywords string `json:"keywords,omitempty" form:"keywords"`
	Page     int64  `json:"page,omitempty" form:"page"`
	Size     int64  `json:"size,omitempty" form:"size"`
}
