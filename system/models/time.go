package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	t.Time = t1
	return err
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix()*1000, 10)), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	switch value := v.(type) {
	case int64:
		t.Time = time.Unix(value/1000, 0)
	case time.Time:
		t.Time = value
	default:
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	return nil
}
