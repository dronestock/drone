package drone

type (
	stepOption interface {
		applyStep(options *stepOptions)
	}

	stepOptions struct {
		name   string
		async  bool
		retry  bool
		_break bool
	}
)

func defaultStepOption() *stepOptions {
	return &stepOptions{
		name:   `默认步骤`,
		async:  false,
		retry:  true,
		_break: true,
	}
}

func (o *stepOptions) retryable(base *Base) (retry bool) {
	retry = o.retry
	// 优先以本步骤的配置为准，如果本步骤配置为不重试，再以全局配置为准
	if retry {
		retry = base.Retry
	}

	return
}
