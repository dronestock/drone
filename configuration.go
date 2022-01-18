package drone

import (
	`github.com/storezhang/gox`
)

// Configuration 配置
type Configuration interface {
	// 导出所有字段
	fields() gox.Fields

	// 基础配置
	config() *Config
}
