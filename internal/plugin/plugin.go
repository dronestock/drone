package plugin

import (
	"github.com/dronestock/drone/internal/core"
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

	// Setup 设置配置信息
	Setup() (err error)

	// Before 插件运行前执行
	Before() (err error)

	// Steps 插件运行步骤
	Steps() step.Steps

	// After 插件运行后执行
	After() (err error)

	// Expressions 表达式
	Expressions() core.Expressions
}
