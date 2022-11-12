package drone

import (
	"time"
)

// Plugin 插件接口，任何插件都需要实现这个接口
type Plugin interface {
	// Scheme 卡片模板
	Scheme() (scheme string)

	// Card 卡片数据
	Card() (card any, err error)

	// Interval 卡片数据写入间隔
	Interval() time.Duration

	// Config 加载配置
	Config() (config Config)

	// Steps 插件运行步骤
	Steps() Steps
}
