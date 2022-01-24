package drone

var (
	_            = Args
	_ execOption = (*optionArgs)(nil)
)

type optionArgs struct {
	args []string
}

// Args 参数
func Args(args ...string) *optionArgs {
	return &optionArgs{
		args: args,
	}
}

func (a *optionArgs) applyExec(options *execOptions) {
	options.args = a.args
}
