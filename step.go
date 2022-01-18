package drone

import (
	`fmt`
)

var (
	_     = NewStep
	steps = 1
)

type step struct {
	do      do
	options *stepOptions
}

// NewStep 创建一个步骤
func NewStep(do do, opts ...stepOption) *step {
	_options := defaultStepOption()
	for _, opt := range opts {
		opt.applyStep(_options)
	}
	if `` == _options.name {
		_options.name = fmt.Sprintf(`第%d步`, steps)
		steps++
	}

	return &step{
		do:      do,
		options: _options,
	}
}
