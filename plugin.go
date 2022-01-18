package drone

import (
	`github.com/storezhang/simaqian`
)

type plugin interface {
	// 加载配置
	configuration() (configuration configuration)

	// 插件运行
	run(logger simaqian.Logger) (err error)
}
