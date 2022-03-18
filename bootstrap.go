package drone

import (
	`fmt`
	`os`
	`sync`
	`time`

	`github.com/goexl/gox`
	`github.com/goexl/gox/field`
	`github.com/goexl/mengpo`
	`github.com/goexl/simaqian`
	`github.com/goexl/xiren`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
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

	// 处理别名
	if err = parseAliases(_options.aliases...); nil != err {
		return
	}

	// 加载配置
	configuration := _plugin.Config()
	err = mengpo.Set(configuration, mengpo.Before(toSlice))
	fields := configuration.Fields().Connects(configuration.Base().Fields())
	if nil != err {
		logger.Error(`加载配置出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`加载配置成功`, fields...)
	}
	if nil != err {
		return
	}

	// 数据验证
	if err = xiren.Struct(configuration); nil != err {
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

	base := configuration.Base()
	// 设置日志级别
	if base.Debug {
		logger.Sets(simaqian.Level(simaqian.LevelDebug))
	}
	base.Logger = logger

	// 执行插件
	wg := new(sync.WaitGroup)
	for _, step := range _plugin.Steps() {
		if err = execStep(step, wg, base); nil != err && step.options._break {
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

func execStep(step *Step, wg *sync.WaitGroup, base *PluginBase) (err error) {
	if step.options.async {
		err = execStepAsync(step, wg, base)
	} else {
		err = execStepSync(step, base)
	}

	return
}

func execStepSync(step *Step, base *PluginBase) error {
	return execDo(step.do, step.options, base)
}

func execStepAsync(step *Step, wg *sync.WaitGroup, base *PluginBase) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = execDo(step.do, step.options, base); nil != err {
			panic(err)
		}
	}()

	return
}

func execDo(do do, options *stepOptions, base *PluginBase) (err error) {
	fields := gox.Fields{
		field.String(`name`, options.name),
		field.Bool(`async`, options.async),
		field.Bool(`retry`, options.retry),
		field.Bool(`break`, options._break),
	}
	base.Info(`步骤执行开始`, fields...)

	undo := false
	for count := 0; count < base.Counts; count++ {
		if undo, err = do(); (nil == err) || (0 == count && !base.Retry && !options.retry) || undo {
			break
		} else {
			time.Sleep(base.Backoff)
			base.Info(`步骤执行遇到错误`, field.String(`name`, options.name), field.Int(`count`, count), field.Error(err))
		}
	}

	if nil != err {
		if base.Retry {
			base.Error(`步骤执行尝试所有重试后出错`, fields.Connect(field.Error(err))...)
		} else {
			base.Error(`步骤执行出错`, fields...)
		}
	} else {
		if undo {
			base.Info(`步骤未执行`, fields...)
		} else {
			base.Info(`步骤执行成功`, fields...)
		}
	}

	return
}
