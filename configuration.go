package drone

import (
	`github.com/storezhang/gox`
)

type configuration interface {
	// 导出所有字段
	fields() gox.Fields

	// 基础配置
	config() *Config
}
