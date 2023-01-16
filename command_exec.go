package drone

import (
	"fmt"

	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (cb *commandBuilder) Exec() (err error) {
	fields := gox.Fields[any]{
		field.New("command", cb.command),
		field.New("args", cb.options.args),
		field.New("verbose", cb.base.Verbose),
		field.New("level", cb.base.Level),
	}
	// 记录日志
	if cb.base.Verbose {
		cb.base.Info(fmt.Sprintf("开始执行%s命令", cb.options.name), fields...)
	}

	gexOptions := gex.NewOptions(gex.Args(cb.options.args...))
	if "" != cb.options.dir {
		gexOptions = append(gexOptions, gex.Dir(cb.options.dir))
	}

	if 0 != len(cb.options.environments) {
		gexOptions = append(gexOptions, gex.StringEnvs(cb.options.environments...))
	}

	if cb.options.async {
		gexOptions = append(gexOptions, gex.Async())
	} else {
		gexOptions = append(gexOptions, gex.Sync())
	}

	// 增加检查
	for _, _checker := range cb.options.checkers {
		switch _checker.mode {
		case checkerModeContains:
			gexOptions = append(gexOptions, gex.ContainsChecker(_checker.args.(string)))
		case checkerModeEqual:
			gexOptions = append(gexOptions, gex.EqualChecker(_checker.args.(string)))
		}
	}

	// 增加输出
	for _, _collector := range cb.options.collectors {
		switch _collector.mode {
		case collectorModeString:
			gexOptions = append(gexOptions, gex.StringCollector(_collector.args.(*string)))
		}
	}

	// PWE处理
	if !cb.options.pwe {
		gexOptions = append(gexOptions, gex.DisablePwe())
	}

	if !cb.base.Verbose {
		gexOptions = append(gexOptions, gex.Quiet())
	} else {
		gexOptions = append(gexOptions, gex.Terminal())
		// 当开启了打印输出时，关闭遇到错误打印时
		gexOptions = append(gexOptions, gex.DisablePwe())
	}

	// 执行命令
	if _, err = gex.Exec(cb.command, gexOptions...); nil != err {
		cb.base.Error(fmt.Sprintf("执行%s命令出错", cb.options.name), fields.Connect(field.Error(err))...)
	} else if cb.base.Verbose {
		cb.base.Info(fmt.Sprintf("执行%s命令成功", cb.options.name), fields...)
	}

	return
}
