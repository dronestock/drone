package drone

import (
	`github.com/storezhang/simaqian`
)

// Bootstrap 启动插件
func Bootstrap(plugin plugin, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	var logger simaqian.Logger
	if logger, err = simaqian.New(); nil != err {
		return
	}

	// 解析数组环境变量

	// 加载配置
	if err = plugin.load(); nil != err {
		return
	}

	// 执行插件
	if err = plugin.run(logger); nil != err {
		return
	}

	return
}
