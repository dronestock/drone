package drone

type (
	stepOption interface {
		applyStep(options *stepOptions)
	}

	stepOptions struct {
		name      string
		async     bool
		interrupt bool
	}
)

func defaultStepOption() *stepOptions {
	return &stepOptions{
		async:     false,
		interrupt: true,
	}
}
