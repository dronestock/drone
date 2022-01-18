package drone

import (
	`fmt`
	`time`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/mengpo`
	`github.com/storezhang/simaqian`
	`github.com/storezhang/validatorx`
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
	_configuration := plugin.configuration()
	if err = mengpo.Set(_configuration); nil != err {
		logger.Error(`加载配置出错`, _configuration.fields().Connect(field.Error(err))...)
	} else {
		logger.Info(`加载配置成功`, _configuration.fields()...)
	}
	if nil != err {
		return
	}

	// 数据验证
	if err = validatorx.Struct(_configuration); nil != err {
		logger.Error(`配置验证未通过`, _configuration.fields().Connect(field.Error(err))...)
	} else {
		logger.Info(`配置验证通过，继续执行`)
	}

	config := _configuration.config()
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
