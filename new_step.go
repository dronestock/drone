package drone

import (
	"github.com/dronestock/drone/internal/step"
	"github.com/dronestock/drone/internal/step/stepper"
)

var _ = NewStep

func NewStep(stepper stepper.Stepper) *step.Builder {
	return step.NewBuilder(stepper)
}
