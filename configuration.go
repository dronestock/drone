package drone

import (
	`github.com/storezhang/gox`
)

// Configuration 配置
type Configuration interface {
	// Fields 导出所有字段
	Fields() gox.Fields

	// Basic 基础配置
	Basic() *Config
}
