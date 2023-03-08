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
		field.New("args", cb.params.args),
		field.New("verbose", cb.base.Verbose),
		field.New("level", cb.base.Level),
	}
	// 记录日志
	if cb.base.Verbose {
		cb.base.Info(fmt.Sprintf("开始执行%s命令", cb.params.name), fields...)
	}

	gexOptions := gex.NewOptions(gex.Args(cb.params.args...))
	if "" != cb.params.dir {
		gexOptions = append(gexOptions, gex.Dir(cb.params.dir))
	}

	if 0 != len(cb.params.environments) {
		gexOptions = append(gexOptions, gex.StringEnvs(cb.params.environments...))
	}

	if cb.params.async {
		gexOptions = append(gexOptions, gex.Async())
	} else {
		gexOptions = append(gexOptions, gex.Sync())
	}

	// 增加检查
	for _, _checker := range cb.params.checkers {
		switch _checker.mode {
		case checkerModeContains:
			gexOptions = append(gexOptions, gex.ContainsChecker(_checker.args.(string)))
		case checkerModeEqual:
			gexOptions = append(gexOptions, gex.EqualChecker(_checker.args.(string)))
		}
	}

	// 增加输出
	for _, _collector := range cb.params.collectors {
		switch _collector.mode {
		case collectorModeString:
			gexOptions = append(gexOptions, gex.StringCollector(_collector.args.(*string)))
		}
	}

	// PWE处理
	if !cb.params.pwe {
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
		cb.base.Error(fmt.Sprintf("执行%s命令出错", cb.params.name), fields.Add(field.Error(err))...)
	} else if cb.base.Verbose {
		cb.base.Info(fmt.Sprintf("执行%s命令成功", cb.params.name), fields...)
	}

	return
}
