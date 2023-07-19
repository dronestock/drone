package step

import (
	"fmt"
	"time"

	"github.com/goexl/gox"
)

var (
	_     = NewBuilder
	_     = NewDelayStep
	_     = NewDebugStep
	steps = 1
)

// Step 步骤
type Step struct {
	gox.CannotCopy

	Stepper Stepper
	Options *Options
}

func New(stepper Stepper, options *Options) *Step {
	if "" == options.Name {
		options.Name = fmt.Sprintf("第%d步", steps)
		steps++
	}

	return &Step{
		Stepper: stepper,
		Options: options,
	}
}

// NewDelayStep 创建延迟步骤，调试使用
func NewDelayStep(delay time.Duration) *Step {
	return NewBuilder(NewDelay(delay)).Name("延迟步骤").Build()
}

// NewDebugStep 创建延迟步骤，调试使用
func NewDebugStep() *Step {
	return NewDelayStep(time.Hour)
}
