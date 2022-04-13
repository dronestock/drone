package drone

import (
	"fmt"

	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (p *Plugin) Exec(command string, opts ...execOption) (err error) {
	_options := defaultExecOptions()
	for _, opt := range opts {
		opt.applyExec(_options)
	}

	fields := gox.Fields{
		field.String(`command`, command),
		field.Any(`args`, _options.args),
		field.Bool(`verbose`, p.Verbose),
		field.Bool(`debug`, p.Debug),
	}
	// 记录日志
	if p.Debug {
		p.Info(fmt.Sprintf(`开始执行%s命令`, _options.name), fields...)
	}

	gexOptions := gex.NewOptions(gex.Args(_options.args...))
	if `` != _options.dir {
		gexOptions = append(gexOptions, gex.Dir(_options.dir))
	}

	if 0 != len(_options.environments) {
		gexOptions = append(gexOptions, gex.StringEnvs(_options.environments...))
	}

	if _options.async {
		gexOptions = append(gexOptions, gex.Async())
	} else {
		gexOptions = append(gexOptions, gex.Sync())
	}

	// 增加检查
	for _, _checker := range _options.checkers {
		switch _checker.mode {
		case checkerModeContains:
			gexOptions = append(gexOptions, gex.ContainsChecker(_checker.args.(string)))
		case checkerModeEqual:
			gexOptions = append(gexOptions, gex.EqualChecker(_checker.args.(string)))
		}
	}

	// 增加输出
	for _, _collector := range _options.collectors {
		switch _collector.mode {
		case collectorModeString:
			gexOptions = append(gexOptions, gex.StringCollector(_collector.args.(*string)))
		}
	}

	if !p.Debug {
		gexOptions = append(gexOptions, gex.Quiet())
	} else {
		gexOptions = append(gexOptions, gex.Terminal())
	}

	// 执行命令
	if _, err = gex.Exec(command, gexOptions...); nil != err {
		p.Error(fmt.Sprintf(`执行%s命令出错`, _options.name), fields.Connect(field.Error(err))...)
	} else if p.Debug {
		p.Info(fmt.Sprintf(`执行%s命令成功`, _options.name), fields...)
	}

	return
}
