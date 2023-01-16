package drone

import (
	"time"
)

type delayStepper struct {
	delay time.Duration
}

func newDelayStepper(delay time.Duration) *delayStepper {
	return &delayStepper{
		delay: delay,
	}
}

func (ds *delayStepper) Runnable() bool {
	return true
}

func (ds *delayStepper) Run() (err error) {
	return
}
