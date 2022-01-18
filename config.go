package drone

import (
	`time`
)

// Config 插件基础配置
type Config struct {
	// 是否启用默认配置
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示详细信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`
	// 是否显示调试信息
	Debug bool `default:"${PLUGIN_DEBUG=${DEBUG=false}}"`

	// 是否重试
	Retry bool `default:"${PLUGIN_RETRY=${RETRY=true}}"`
	// 重试次数
	Times int `default:"${PLUGIN_TIMES=${TIMES=true}}"`
	// 重试间隔
	Backoff time.Duration `default:"${PLUGIN_BACKOFF=${BACKOFF=5s}}"`
}
