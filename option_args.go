package drone

var (
	_            = Args
	_ execOption = (*optionArgs)(nil)
)

type optionArgs struct {
	args []interface{}
}

// Args 参数
func Args(args ...interface{}) *optionArgs {
	return &optionArgs{
		args: args,
	}
}

func (a *optionArgs) applyExec(options *execOptions) {
	options.args = a.args
}
