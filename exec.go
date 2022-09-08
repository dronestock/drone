package drone

import (
	"fmt"

	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *Base) Exec(command string, opts ...execOption) error {
	_options := defaultExecOptions(b.Pwe)
	for _, opt := range opts {
		opt.applyExec(_options)
	}

	return b.exec(command, _options)
}

func (b *Base) exec(command string, options *execOptions) (err error) {
	fields := gox.Fields{
		field.String(`command`, command),
		field.Any(`args`, options.args),
		field.Bool(`verbose`, b.Verbose),
		field.Bool(`debug`, b.Debug),
	}
	// 记录日志
	if b.Debug {
		b.Info(fmt.Sprintf(`开始执行%s命令`, options.name), fields...)
	}

	gexOptions := gex.NewOptions(gex.Args(options.args...))
	if `` != options.dir {
		gexOptions = append(gexOptions, gex.Dir(options.dir))
	}

	if 0 != len(options.environments) {
		gexOptions = append(gexOptions, gex.StringEnvs(options.environments...))
	}

	if options.async {
		gexOptions = append(gexOptions, gex.Async())
	} else {
		gexOptions = append(gexOptions, gex.Sync())
	}

	// 增加检查
	for _, _checker := range options.checkers {
		switch _checker.mode {
		case checkerModeContains:
			gexOptions = append(gexOptions, gex.ContainsChecker(_checker.args.(string)))
		case checkerModeEqual:
			gexOptions = append(gexOptions, gex.EqualChecker(_checker.args.(string)))
		}
	}

	// 增加输出
	for _, _collector := range options.collectors {
		switch _collector.mode {
		case collectorModeString:
			gexOptions = append(gexOptions, gex.StringCollector(_collector.args.(*string)))
		}
	}

	// PWE处理
	if !options.pwe {
		gexOptions = append(gexOptions, gex.DisablePwe())
	}

	if !b.Debug {
		gexOptions = append(gexOptions, gex.Quiet())
	} else {
		gexOptions = append(gexOptions, gex.Terminal())
	}

	// 执行命令
	if _, err = gex.Exec(command, gexOptions...); nil != err {
		b.Error(fmt.Sprintf(`执行%s命令出错`, options.name), fields.Connect(field.Error(err))...)
	} else if b.Debug {
		b.Info(fmt.Sprintf(`执行%s命令成功`, options.name), fields...)
	}

	return
}
