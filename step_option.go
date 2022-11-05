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
