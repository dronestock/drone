package drone

type stepBuilder struct {
	stepper stepper
	options *stepOptions
}

// NewStep 创建一个步骤
func NewStep(stepper stepper) *stepBuilder {
	return &stepBuilder{
		stepper: stepper,
		options: &stepOptions{
			name:  "默认步骤",
			async: false,
			retry: true,
			br:    true,
		},
	}
}

func (sb *stepBuilder) Name(name string) *stepBuilder {
	sb.options.name = name

	return sb
}

func (sb *stepBuilder) Async() *stepBuilder {
	sb.options.async = true

	return sb
}

func (sb *stepBuilder) Continue() *stepBuilder {
	sb.options.br = false

	return sb
}

func (sb *stepBuilder) Break() *stepBuilder {
	sb.options.br = true

	return sb
}

func (sb *stepBuilder) Retry() *stepBuilder {
	sb.options.retry = true

	return sb
}

func (sb *stepBuilder) Interrupt() *stepBuilder {
	sb.options.retry = false

	return sb
}

func (sb *stepBuilder) Build() *Step {
	return newStep(sb.stepper, sb.options)
}
