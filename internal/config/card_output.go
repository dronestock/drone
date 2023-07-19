package config

import (
	"encoding/json"
)

type CardOutput struct {
	// 地址
	Schema string `json:"schema"`
	// 数据
	Data json.RawMessage `json:"data"`
}
