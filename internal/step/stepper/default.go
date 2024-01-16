package stepper

import (
	"context"
)

var _ Stepper = (*Default)(nil)

type Default struct{}

func (d *Default) Runnable() bool {
	return true
}

func (d *Default) Run(_ *context.Context) (err error) {
	return
}
