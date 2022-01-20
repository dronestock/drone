package drone

type (
	stepOption interface {
		applyStep(options *stepOptions)
	}

	stepOptions struct {
		name   string
		async  bool
		_break bool
	}
)

func defaultStepOption() *stepOptions {
	return &stepOptions{
		async:  false,
		_break: true,
	}
}
