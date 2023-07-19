package stepper

import (
	"context"
)

type Stepper interface {
	Runnable() bool

	Run(ctx context.Context) (err error)
}
