package step

import (
	"context"
	"time"
)

type Delay struct {
	delay time.Duration
}

func NewDelay(delay time.Duration) *Delay {
	return &Delay{
		delay: delay,
	}
}

func (d *Delay) Runnable() bool {
	return true
}

func (d *Delay) Run(_ context.Context) (err error) {
	return
}
