package models

import (
	"database/sql/driver"
	"encoding/json"
)

type BStruct struct {
	Path  string `json:"path"`
	Label string `json:"label"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Url   string `json:"url"`
}

type BList []*BStruct

func (b BList) Value() (driver.Value, error) {
	d, err := json.Marshal(b)
	return string(d), err
}

// 注意，这里的接收器是指针类型，否则无法把数据从数据库读到结构体
func (b *BList) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), b)
}
