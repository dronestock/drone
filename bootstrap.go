package drone

import (
	`fmt`
	`sync`
	`time`

	`github.com/storezhang/gox`
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
	wg := new(sync.WaitGroup)
	for _, _step := range plugin.steps() {
		err = execStep(_step, wg, config, logger)
	}
	wg.Wait()

	// 记录日志
	if nil != err {
		logger.Error(fmt.Sprintf(`%s插件执行出错，请检查`, _options.name))
	} else {
		logger.Info(fmt.Sprintf(`%s插件执行成功，恭喜`, _options.name))
	}

	return
}

func execStep(step *step, wg *sync.WaitGroup, config *Config, logger simaqian.Logger) (err error) {
	if step.options.parallelism {
		err = execStepAsync(step, wg, config, logger)
	} else {
		err = execStepSync(step, config, logger)
	}

	return
}

func execStepSync(step *step, config *Config, logger simaqian.Logger) error {
	return execDo(step.do, step.options, config, logger)
}

func execStepAsync(step *step, wg *sync.WaitGroup, config *Config, logger simaqian.Logger) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = execDo(step.do, step.options, config, logger); nil != err {
			panic(err)
		}
	}()

	return
}

func execDo(do do, options *stepOptions, config *Config, logger simaqian.Logger) (err error) {
	for count := 0; count < config.Counts; count++ {
		if err = do(logger); (nil == err) || (0 == count && !config.Retry) {
			break
		} else {
			time.Sleep(config.Backoff)
			logger.Info(`步骤执行遇到错误`, field.String(`name`, options.name), field.Int(`count`, count), field.Error(err))
		}
	}

	if nil != err {
		fields := gox.Fields{
			field.String(`name`, options.name),
			field.Error(err),
		}
		if config.Retry {
			logger.Error(`步骤执行尝试所有重试后出错`, fields...)
		} else {
			logger.Error(`步骤执行出错`, fields...)
		}
	} else {
		logger.Info(`步骤执行成功`, field.String(`name`, options.name))
	}

	return
}
