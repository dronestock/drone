package drone

type (
	stepOption interface {
		applyStep(options *stepOptions)
	}

	stepOptions struct {
		name        string
		parallelism bool
	}
)

func defaultStepOption() *stepOptions {
	return &stepOptions{
		parallelism: false,
	}
}
