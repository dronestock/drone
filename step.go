package drone

import (
	`fmt`
)

var (
	_     = NewStep
	steps = 1
)

// Step 步骤
type Step struct {
	do      do
	options *stepOptions
}

// NewStep 创建一个步骤
func NewStep(do do, opts ...stepOption) *Step {
	_options := defaultStepOption()
	for _, opt := range opts {
		opt.applyStep(_options)
	}
	if `` == _options.name {
		_options.name = fmt.Sprintf(`第%d步`, steps)
		steps++
	}

	return &Step{
		do:      do,
		options: _options,
	}
}
