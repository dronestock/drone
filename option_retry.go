package drone

var (
	_            = Retry
	_            = Interrupt
	_ stepOption = (*optionRetry)(nil)
)

type optionRetry struct {
	retry bool
}

// Retry 遇到错误重试
func Retry() *optionRetry {
	return &optionRetry{
		retry: true,
	}
}

// Interrupt 遇到错误中断执行
func Interrupt() *optionRetry {
	return &optionRetry{
		retry: false,
	}
}

func (r *optionRetry) applyStep(options *stepOptions) {
	options.retry = r.retry
}
