package drone

import (
	`github.com/storezhang/simaqian`
)

type plugin interface {
	// 加载配置
	load() (config *Config, err error)

	// 插件运行
	run(logger simaqian.Logger) (err error)
}
