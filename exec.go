package drone

import (
	`fmt`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func (pb *PluginBase) Exec(command string, opts ...execOption) (err error) {
	_options := defaultExecOptions()
	for _, opt := range opts {
		opt.applyExec(_options)
	}

	fields := gox.Fields{
		field.String(`command`, command),
		field.Strings(`args`, _options.args...),
		field.Bool(`verbose`, pb.Verbose),
		field.Bool(`debug`, pb.Debug),
	}
	// 记录日志
	pb.Info(fmt.Sprintf(`开始执行%s命令`, _options.name), fields...)

	gexOptions := gex.NewOptions(gex.Args(_options.args...))
	if `` != _options.dir {
		gexOptions = append(gexOptions, gex.Dir(_options.dir))
	}

	if _options.async {
		gexOptions = append(gexOptions, gex.Async())
	} else {
		gexOptions = append(gexOptions, gex.Sync())
	}

	switch _options.checkerMode {
	case checkerModeContains:
		gexOptions = append(gexOptions, gex.ContainsChecker(_options.checkerArgs.(string)))
	case checkerModeEqual:
		gexOptions = append(gexOptions, gex.EqualChecker(_options.checkerArgs.(string)))
	}

	if !pb.Debug {
		gexOptions = append(gexOptions, gex.Quiet())
	}

	// 执行命令
	if _, err = gex.Run(command, gexOptions...); nil != err {
		pb.Error(fmt.Sprintf(`执行%s命令出错`, _options.name), fields.Connect(field.Error(err))...)
	} else {
		pb.Info(fmt.Sprintf(`执行%s命令成功`, _options.name), fields...)
	}

	return
}
