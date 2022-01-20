package drone

import (
	`fmt`
	`os`
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
func Bootstrap(constructor constructor, opts ...option) (err error) {
	_plugin := constructor()
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
	configuration := _plugin.Configuration()
	fields := configuration.Fields().Connects(configuration.Basic().Fields())
	if err = mengpo.Set(configuration); nil != err {
		logger.Error(`加载配置出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`加载配置成功`, fields...)
	}
	if nil != err {
		return
	}

	// 数据验证
	if err = validatorx.Struct(configuration); nil != err {
		logger.Error(`配置验证未通过`, configuration.Fields().Connect(field.Error(err))...)
	} else {
		logger.Info(`配置验证通过，继续执行`)
	}
	if nil != err {
		return
	}

	// 设置配置信息
	if unset, setErr := configuration.Setup(); nil != setErr {
		logger.Error(`设置配置信息出错`, configuration.Fields().Connect(field.Error(err))...)
		err = setErr
	} else if !unset {
		logger.Info(`设置配置信息完成，继续执行`)
	}
	if nil != err {
		return
	}

	config := configuration.Basic()
	// 设置日志级别
	if config.Debug {
		logger.Sets(simaqian.Level(simaqian.LevelDebug))
	}

	// 执行插件
	wg := new(sync.WaitGroup)
	for _, step := range _plugin.Steps() {
		if err = execStep(step, wg, config, logger); nil != err && step.options._break {
			return
		}
	}
	wg.Wait()

	// 记录日志
	if nil != err {
		logger.Error(fmt.Sprintf(`%s插件执行出错，请检查`, _options.name))
	} else {
		logger.Info(fmt.Sprintf(`%s插件执行成功，恭喜`, _options.name))

		// 退出程序，解决最外层panic报错的问题
		// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
		// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
		os.Exit(0)
	}

	return
}

func execStep(step *Step, wg *sync.WaitGroup, config *Config, logger simaqian.Logger) (err error) {
	if step.options.async {
		err = execStepAsync(step, wg, config, logger)
	} else {
		err = execStepSync(step, config, logger)
	}

	return
}

func execStepSync(step *Step, config *Config, logger simaqian.Logger) error {
	return execDo(step.do, step.options, config, logger)
}

func execStepAsync(step *Step, wg *sync.WaitGroup, config *Config, logger simaqian.Logger) (err error) {
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
	undo := false
	for count := 0; count < config.Counts; count++ {
		if undo, err = do(logger); (nil == err) || (0 == count && !config.Retry) || undo {
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
		if undo {
			logger.Info(`步骤未执行`, field.String(`name`, options.name))
		} else {
			logger.Info(`步骤执行成功`, field.String(`name`, options.name))
		}
	}

	return
}
