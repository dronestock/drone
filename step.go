package drone

import (
	"fmt"
	"time"

	"github.com/goexl/gox"
)

var (
	_     = NewStep
	_     = NewDelayStep
	_     = NewDebugStep
	steps = 1
)

// Step 步骤
type Step struct {
	gox.CannotCopy

	stepper stepper
	options *stepOptions
}

func newStep(stepper stepper, options *stepOptions) *Step {
	if "" == options.name {
		options.name = fmt.Sprintf("第%d步", steps)
		steps++
	}

	return &Step{
		stepper: stepper,
		options: options,
	}
}

// NewDelayStep 创建延迟步骤，调试使用
func NewDelayStep(delay time.Duration) *Step {
	return NewStep(newDelayStepper(delay)).Name("延迟步骤").Build()
}

// NewDebugStep 创建延迟步骤，调试使用
func NewDebugStep() *Step {
	return NewDelayStep(time.Hour)
}
