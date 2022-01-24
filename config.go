package drone

import (
	`github.com/storezhang/gox`
)

// Config 配置
type Config interface {
	// Setup 设置配置信息
	Setup() (unset bool, err error)

	// Fields 导出所有字段
	Fields() gox.Fields

	// Base 基础配置
	Base() *PluginBase
}
