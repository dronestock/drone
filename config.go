package drone

import (
	"github.com/goexl/gox"
)

// Config 配置
type Config interface {
	// Setup 设置配置信息
	Setup() (err error)

	// Fields 导出所有字段
	Fields() gox.Fields[any]

	base() *Base
}
