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
	logger := simaqian.Default()

	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	// 处理别名
	if err = parseAliases(_options.aliases...); nil != err {
		return
	}

	// 加载配置
	configuration := _plugin.Config()
	err = mengpo.Set(configuration, mengpo.EnvGetter(envGetter), mengpo.Processor(new(sliceProcessor)))
	fields := configuration.Fields().Connects(configuration.BaseConfig().Fields()...)
	if nil != err {
		logger.Error("加载配置出错", fields.Connect(field.Error(err))...)
	} else {
		logger.Info("加载配置成功", fields...)
	}
	if nil != err {
		return
	}

	// 数据验证
	if err = xiren.Struct(configuration); nil != err {
		logger.Error("配置验证未通过", configuration.Fields().Connect(field.Error(err))...)
	} else {
		logger.Info("配置验证通过，继续执行")
	}
	if nil != err {
		return
	}

	base := configuration.BaseConfig()
	builder := simaqian.New()
	// 设置日志级别
	builder.Level(simaqian.ParseLevel(base.Log.Level))
	// 向标准输出流输出日志
	zap := builder.Zap().Output(simaqian.Stdout())
	if base.Logger, err = zap.Build(); nil == err {
		logger = base.Logger
	}
	if nil != err {
		return
	}

	// 设置配置信息
	if unset, se := configuration.Setup(); nil != se {
		logger.Error("设置配置信息出错", configuration.Fields().Connect(field.Error(err))...)
		err = se
	} else if !unset {
		logger.Info("设置配置信息完成，继续执行")
	}
	if nil != err {
		return
	}

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
		logger.Error(fmt.Sprintf("%s插件执行出错，请检查", _options.name))
	} else {
		logger.Info(fmt.Sprintf("%s插件执行成功，恭喜", _options.name))

		// 退出程序，解决最外层panic报错的问题
		// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
		// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
		os.Exit(0)
	}

	return
}


