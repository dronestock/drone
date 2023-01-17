package drone

import (
	"context"
)

type stepper interface {
	// Runnable 是否执行步骤
	Runnable() bool

	// Run 执行步骤
	Run(ctx context.Context) (err error)
}
