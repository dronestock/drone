package drone

import (
	`fmt`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func (b *Base) Exec(command string, opts ...execOption) (err error) {
	_options := defaultExecOptions()
	for _, opt := range opts {
		opt.applyExec(_options)
	}

	fields := gox.Fields{
		field.String(`command`, command),
		field.Strings(`args`, _options.args...),
		field.Bool(`verbose`, b.Verbose),
		field.Bool(`debug`, b.Debug),
	}
	// 记录日志
	b.Info(fmt.Sprintf(`开始执行%s命令`, _options.name), fields...)

	gexOptions := gex.NewOptions(gex.Args(_options.args...))
	if `` != _options.dir {
		gexOptions = append(gexOptions, gex.Dir(_options.dir))
	}
	if !b.Debug {
		gexOptions = append(gexOptions, gex.Quiet())
	}
	if _, err = gex.Run(command, gexOptions...); nil != err {
		b.Error(fmt.Sprintf(`执行%s命令出错`, _options.name), fields.Connect(field.Error(err))...)
	} else {
		b.Info(fmt.Sprintf(`执行%s命令成功`, _options.name), fields...)
	}

	return
}
