package drone

import (
	"encoding/json"
)

type cardOutput struct {
	// 地址
	Schema string `json:"schema"`
	// 数据
	Data json.RawMessage `json:"data"`
}
