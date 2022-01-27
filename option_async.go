package drone

var (
	_            = Async
	_            = Sync
	_ stepOption = (*optionAsync)(nil)
	_ execOption = (*optionAsync)(nil)
)

type optionAsync struct {
	async bool
}

// Async 配置异步执行
func Async() *optionAsync {
	return &optionAsync{
		async: true,
	}
}

// Sync 配置同步执行
func Sync() *optionAsync {
	return &optionAsync{
		async: false,
	}
}

func (a *optionAsync) applyStep(options *stepOptions) {
	options.async = a.async
}

func (a *optionAsync) applyExec(options *execOptions) {
	options.async = a.async
}
