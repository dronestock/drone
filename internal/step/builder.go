package step

type Builder struct {
	stepper Stepper
	options *Options
}

// NewBuilder 创建一个步骤
func NewBuilder(stepper Stepper) *Builder {
	return &Builder{
		stepper: stepper,
		options: &Options{
			Name:  "默认步骤",
			Async: false,
			Retry: true,
			Break: true,
		},
	}
}

func (b *Builder) Name(name string) *Builder {
	b.options.Name = name

	return b
}

func (b *Builder) Async() *Builder {
	b.options.Async = true

	return b
}

func (b *Builder) Continue() *Builder {
	b.options.Break = false

	return b
}

func (b *Builder) Break() *Builder {
	b.options.Break = true

	return b
}

func (b *Builder) Retry() *Builder {
	b.options.Retry = true

	return b
}

func (b *Builder) Interrupt() *Builder {
	b.options.Retry = false

	return b
}

func (b *Builder) Build() *Step {
	return New(b.stepper, b.options)
}
