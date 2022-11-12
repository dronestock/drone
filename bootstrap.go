package drone

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/mengpo"
	"github.com/goexl/simaqian"
	"github.com/goexl/xiren"
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
	if logger, err = simaqian.New(simaqian.Output(simaqian.Stdout())); nil != err {
		return
	}

	// 处理别名
	if err = parseAliases(_options.aliases...); nil != err {
		return
	}

	// 加载配置
	configuration := _plugin.Config()
	err = mengpo.Set(configuration, mengpo.EnvGetter(envGetter), mengpo.Processor(new(sliceProcessor)))
	fields := configuration.Fields().Connects(configuration.BaseConfig().Fields())
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

	base := configuration.BaseConfig()
	base.Logger = logger
	// 设置日志级别
	base.Logger.Sets(simaqian.Levels(base.Level))

	// 开始卡片信息写入
	go startCard(_plugin, base)

	// 执行插件
	wg := new(sync.WaitGroup)
	for _, step := range _plugin.Steps() {
		if err = execStep(step, wg, base); nil != err && step.options._break {
			return
		}
	}
	wg.Wait()

	// 写入最终数据到卡片中
	_ = writeCard(_plugin, base)

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

func execStep(step *Step, wg *sync.WaitGroup, base *Base) (err error) {
	if step.options.async {
		err = execStepAsync(step, wg, base)
	} else {
		err = execStepSync(step, base)
	}

	return
}

func execStepSync(step *Step, base *Base) error {
	return execDo(step.do, step.options, base)
}

func execStepAsync(step *Step, wg *sync.WaitGroup, base *Base) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = execDo(step.do, step.options, base); nil != err {
			panic(err)
		}
	}()

	return
}

func execDo(do do, options *stepOptions, base *Base) (err error) {
	retry := options.retryable(base)
	fields := gox.Fields{
		field.String(`name`, options.name),
		field.Bool(`async`, options.async),
		field.Bool(`retry`, retry),
		field.Bool(`break`, options._break),
		field.Int(`counts`, base.Counts),
	}
	base.Info(`步骤执行开始`, fields...)

	undo := false
	rand.Seed(time.Now().UnixNano())
	for count := 0; count < base.Counts; count++ {
		if undo, err = do(); (nil == err) || (0 == count && !retry) || undo {
			break
		}

		backoff := base.backoff()
		base.Info(fmt.Sprintf(`步骤第%d次执行遇到错误`, count+1), fields.Connect(field.Error(err))...)
		base.Info(fmt.Sprintf(`休眠%s，继续执行步骤`, backoff), fields...)
		time.Sleep(backoff)
		base.Info(fmt.Sprintf(`步骤重试第%d次执行`, count+2), fields...)

		if count != base.Counts-1 {
			err = nil
		}
	}

	switch {
	case nil != err && retry:
		base.Error(`步骤执行尝试所有重试后出错`, fields.Connect(field.Error(err))...)
	case nil != err && !retry:
		base.Error(`步骤执行出错`, fields.Connect(field.Error(err))...)
	case nil == err && undo:
		base.Info(`步骤未执行`, fields...)
	case nil == err && !undo:
		base.Info(`步骤执行成功`, fields...)
	}

	return
}

func startCard(plugin Plugin, base *Base) {
	ticker := time.NewTimer(100 * time.Millisecond)
	defer func() {
		_ = ticker.Stop()
	}()

	for range ticker.C {
		if err := writeCard(plugin, base); nil != err {
			base.Warn(`写入卡片数据出错`, field.Error(err))
		}
		ticker.Reset(plugin.Interval())
	}
}

func writeCard(plugin Plugin, base *Base) (err error) {
	scheme := plugin.Scheme()
	if strings.HasPrefix(scheme, github) {
		scheme = fmt.Sprintf(`%s%s`, ghproxy, scheme)
	}

	if _card, ce := plugin.Card(); nil != ce {
		err = ce
	} else {
		err = base.writeCard(scheme, _card)
	}

	return
}
