package plugin

import (
	"github.com/dronestock/drone/internal/expr"
	"github.com/dronestock/drone/internal/step"
)

// Plugin 插件接口，任何插件都需要实现这个接口
type Plugin interface {
	// Scheme 卡片模板
	Scheme() (scheme string)

	// Carding 卡片数据
	Carding() (card any, err error)

	// Config 加载配置
	Config() (config Config)

	// Steps 插件运行步骤
	Steps() step.Steps

	// Expressions 表达式
	Expressions() expr.Expressions
}
