package step

import (
	"fmt"
	"time"

	"github.com/dronestock/drone/internal/step/stepper"
	"github.com/goexl/gox"
)

var (
	_     = NewBuilder
	_     = NewDelay
	_     = NewDebug
	steps = 1
)

// Step 步骤
type Step struct {
	Stepper stepper.Stepper
	Options *Options

	_ gox.Pointerized
}

func New(stepper stepper.Stepper, options *Options) *Step {
	if "" == options.Name {
		options.Name = fmt.Sprintf("第%d步", steps)
		steps++
	}

	return &Step{
		Stepper: stepper,
		Options: options,
	}
}

// NewDelay 创建延迟步骤，调试使用
func NewDelay(delay time.Duration) *Step {
	return NewBuilder(stepper.NewDelay(delay)).Name("延迟步骤").Build()
}

// NewDebug 创建延迟步骤，调试使用
func NewDebug() *Step {
	return NewDelay(time.Hour)
}
