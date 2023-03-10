package drone

import (
	"github.com/goexl/gox/field"
	"github.com/goexl/mengpo"
	"github.com/goexl/simaqian"
	"github.com/goexl/xiren"
)

func (b *bootstrap) setup() (err error) {
	// 处理别名
	if err = b.parseAliases(); nil != err {
		return
	}

	// 加载配置
	config := b.plugin.Config()
	err = mengpo.New().Getter(b.getter).Processor(b.processor).Build().Set(config)
	fields := config.Fields().Add(config.base().Fields()...)
	if nil != err {
		b.Error("加载配置出错", fields.Add(field.Error(err))...)
	} else {
		b.Info("加载配置成功", fields...)
	}
	if nil != err {
		return
	}

	b.Base = config.base()
	builder := simaqian.New()
	// 设置日志级别
	builder.Level(simaqian.ParseLevel(b.Level))
	// 向标准输出流输出日志
	zap := builder.Zap().Output(simaqian.Stdout())
	if logger, be := zap.Build(); nil != be {
		err = be
		b.Warn("配置日志失败", field.Error(be))
	} else {
		b.Logger = logger
	}
	if nil != err {
		return
	}

	// 设置配置信息
	if se := config.Setup(); nil != se {
		b.Error("设置配置信息出错", config.Fields().Add(field.Error(err))...)
		err = se
	} else {
		b.Info("设置配置信息完成，继续执行")
	}
	if nil != err {
		return
	}

	// 数据验证
	if err = xiren.Struct(config); nil != err {
		b.Error("配置验证未通过", config.Fields().Add(field.Error(err))...)
	} else {
		b.Info("配置验证通过，继续执行")
	}
	if nil != err {
		return
	}

	return
}
