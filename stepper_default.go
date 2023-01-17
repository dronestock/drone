package drone

import (
	"context"
)

var _ stepper = (*defaultStepper)(nil)

type defaultStepper struct{}

func (ds *defaultStepper) Runnable() bool {
	return true
}

func (ds *defaultStepper) Run(_ context.Context) (err error) {
	return
}
