package plugin

import (
	"github.com/goexl/gox"
)

// Config 配置
type Config interface {
	// Fields 导出所有字段
	Fields() gox.Fields[any]

	base() *Base
}
