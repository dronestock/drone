package drone

import (
	`github.com/storezhang/gox`
)

// Configuration 配置
type Configuration interface {
	// Setup 设置配置信息
	Setup() (unset bool, err error)

	// Fields 导出所有字段
	Fields() gox.Fields

	// Basic 基础配置
	Basic() *Config
}
