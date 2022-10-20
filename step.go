package drone

import (
	"fmt"
	"time"
)

var (
	_     = NewStep
	_     = NewDelayStep
	_     = NewDefaultDelayStep
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

// NewDelayStep 创建延迟步骤，调试使用
func NewDelayStep(delay time.Duration) *Step {
	return &Step{
		do: func() (undo bool, err error) {
			time.Sleep(delay)

			return
		},
		options: defaultStepOption(),
	}
}

// NewDefaultDelayStep 创建延迟步骤，调试使用
func NewDefaultDelayStep() *Step {
	return NewDelayStep(time.Hour)
}
