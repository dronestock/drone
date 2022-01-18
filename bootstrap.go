package drone

import (
	`fmt`
	`time`

	`github.com/storezhang/simaqian`
)

var _ = Bootstrap

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
	if err = parseConfigs(_options.configs...); nil != err {
		return
	}

	// 加载配置
	var config *Config
	if config, err = plugin.load(); nil != err {
		return
	}

	// 设置日志级别
	if config.Debug {
		logger.Sets(simaqian.Level(simaqian.LevelDebug))
	}

	// 执行插件
	for count := 0; count < config.Counts; count++ {
		if err = plugin.run(logger); (nil == err) || (0 == count && !config.Retry) {
			break
		} else {
			time.Sleep(config.Backoff)
		}
	}

	// 记录日志
	if nil != err {
		logger.Error(fmt.Sprintf(`%s插件执行出错，请检查`, _options.name))
	} else {
		logger.Info(fmt.Sprintf(`%s插件执行成功，恭喜`, _options.name))
	}

	return
}
